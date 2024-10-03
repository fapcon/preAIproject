package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ptflp/gomapper"
	"github.com/shopspring/decimal"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"log"
	statservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/statistics/service"
	"sync"
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/storage"
	tservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/service"
	uservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/storages"
)

const (
	StatusUnknown = iota
	StatusProcessing
	StatusFinished
	StatusFailed

	RoundPlaces = 8
)

type Order struct {
	exchangeOrder          storage.ExchangeOrderer
	exchangeOrderLog       storage.ExchangeOrderLogger
	exchangeTickerService  tservice.ExchangeTicker
	exchangeUserKeyService uservice.ExchangeUserKeyer
	logger                 *zap.Logger
	errChan                chan models.ErrMessage
	statistic              statservice.Statisticer
}

func NewExchangeOrder(storages *storages.Storages, ticker tservice.ExchangeTicker, userKey uservice.ExchangeUserKeyer, statistic statservice.Statisticer, components *component.Components) ExchangeOrderer {
	return &Order{
		exchangeOrder:          storages.ExchangeOrder,
		exchangeOrderLog:       storages.ExchangeOrderLog,
		exchangeTickerService:  ticker,
		exchangeUserKeyService: userKey,
		logger:                 components.Logger,
		errChan:                components.ErrChan,
		statistic:              statistic,
	}
}

func (p *Order) CancelOrder(ctx context.Context, in OrderIn) (CancelOrderOut, error) {
	var condition utils.Condition

	exchangeClient := in.ExClient

	switch in.Signal.OrderType {
	case service.TypeCancelBuy:
		condition = utils.Condition{
			Equal: map[string]interface{}{"exchange_order_type": service.ExchangeOrderTypeLimit, "side": service.SideBuy, "status": service.OrderStatusNew, "bot_uuid": in.Bot.UUID},
		}
	case service.TypeAverage:
		condition = utils.Condition{
			Equal: map[string]interface{}{
				"exchange_order_type": service.ExchangeOrderTypeLimit,
				"side":                service.SideSell,
				"status":              service.OrderStatusNew,
				"bot_uuid":            in.Bot.UUID,
				"pair_id":             in.Signal.PairID,
			},
			NotEqual: map[string]interface{}{"buy_price": 0},
		}
	case service.TypeCancelSell:
		condition = utils.Condition{
			Equal: map[string]interface{}{"exchange_order_type": service.ExchangeOrderTypeLimit, "side": service.SideSell, "status": service.OrderStatusNew, "bot_uuid": in.Bot.UUID},
		}
	case service.TypeCancelAll:
		condition = utils.Condition{
			Equal: map[string]interface{}{"exchange_order_type": service.ExchangeOrderTypeLimit, "status": service.OrderStatusNew, "bot_uuid": in.Bot.UUID},
		}
	default:
		return CancelOrderOut{}, fmt.Errorf("cancel order: bad ordertype")
	}

	orders, err := p.exchangeOrder.GetList(ctx, condition)
	if err != nil {
		p.logger.Error("cancel orders: get orders", zap.Error(err))
		return CancelOrderOut{}, fmt.Errorf("cancel order: orders getlist error")
	}
	if len(orders) < 1 {
		return CancelOrderOut{}, fmt.Errorf("cancel order: not found")
	}

	exchangeOrders := make(chan service.OrderOut, 1024)
	var wg sync.WaitGroup
	for i := range orders {
		order := orders[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			exchangeOrder := exchangeClient.CancelOrder(ctx, service.CancelOrderIn{
				Pair:    order.Pair,
				OrderID: order.ExchangeOrderID,
			})
			if exchangeOrder.ErrorCode != errors.NoError {
				p.logger.Error("cancel orders: exchange cancel order", zap.Int("error code", exchangeOrder.ErrorCode), zap.Int("count", len(orders)))
				return
			}
			order.Status = exchangeOrder.Status
			err = p.exchangeOrder.Update(ctx, order)
			if err != nil {
				p.logger.Error("cancel orders: update order status", zap.Error(err))
				return
			}
			exchangeOrders <- exchangeOrder
		}()
	}
	wg.Wait()
	close(exchangeOrders)
	var exOrders []service.OrderOut
	platformOrders := make([]models.ExchangeOrderDTO, 0, len(orders))
	for exOrder := range exchangeOrders {
		for i := range orders {
			if orders[i].ExchangeOrderID == exOrder.OrderID {
				platformOrders = append(platformOrders, orders[i])
			}
		}
		exOrders = append(exOrders, exOrder)
	}

	return CancelOrderOut{
		ExchangeOrders: exOrders,
		PlatformOrders: platformOrders,
	}, nil
}

func (p *Order) PutOrder(ctx context.Context, in OrderIn) PutOrderOut {
	exchangeClient := in.ExClient
	currencyTicker, err := p.exchangeTickerService.GetByID(ctx, in.Bot.PairID)
	if err != nil || currencyTicker.Pair != in.Signal.Pair {
		p.errChan <- models.ErrMessage{
			Message: "Get pair error",
			ErrNum:  errors.PlatformExchangePutOrderGetPairError,
			Time:    time.Now(),
		}
		return PutOrderOut{
			ErrorCode: errors.PlatformExchangePutOrderGetPairError,
		}
	}

	if currencyTicker.Price.IsZero() || currencyTicker.Price.IsNegative() {
		p.errChan <- models.ErrMessage{
			Message: "Ticker price error",
			ErrNum:  errors.PlatformExchangePutOrderGetPairError,
			Time:    time.Now(),
		}
		return PutOrderOut{
			ErrorCode: errors.PlatformExchangePutOrderGetPairError,
		}
	}

	// Выбор актива, в котором будет выражаться цена ордера
	// BaseAsset - Базовый актив (например, BTC) - in.Bot.AssetType == 0
	// QuoteAsset - Котируемый актив (например, USD) - in.Bot.AssetType == 1

	if in.Bot.AssetType == 1 {
		quoteAsset := currencyTicker.Price
		in.Signal.PairPrice = quoteAsset
	} else if in.Bot.AssetType == 0 {
		baseAsset := decimal.NewFromInt(int64(1)).Div(currencyTicker.Price)
		in.Signal.PairPrice = baseAsset
	}

	in.Signal.PairID = currencyTicker.ID
	var side int
	var price decimal.Decimal
	switch in.Signal.OrderType {
	case service.TypeBuyMarket, service.TypeAverage:
		side = service.SideBuy
		price = currencyTicker.Price
	case service.TypeBuyLimit:
		price = currencyTicker.Price.Mul(decimal.NewFromFloat(1 + in.Bot.LimitBuyPercent/100)).Round(RoundPlaces)
		side = service.SideBuy
	case service.TypeSellMarket:
		side = service.SideSell
		price = currencyTicker.Price
	case service.TypeSellLimit:
		side = service.SideSell
		price = currencyTicker.Price.Mul(decimal.NewFromFloat(1 + in.Bot.LimitSellPercent/100)).Round(RoundPlaces)
	}
	quantity := decimal.NewFromFloat(in.Bot.FixedAmount).DivRound(price, RoundPlaces)

	orderUUID := uuid.NewString()
	err = p.exchangeOrder.Create(ctx, models.ExchangeOrderDTO{
		BotUUID:      in.Bot.UUID,
		UUID:         orderUUID,
		UserID:       in.Bot.UserID,
		ExchangeID:   in.Bot.ExchangeID,
		PairID:       in.Signal.PairID,
		Pair:         in.Signal.Pair,
		OrderType:    in.Signal.OrderType,
		ApiKeyID:     in.Key.ID,
		UnitedOrders: 1,
		Status:       service.OrderStatusProcessing,
		Side:         side,
	})
	if err != nil {
		p.errChan <- models.ErrMessage{
			Message: "Create order error",
			ErrNum:  errors.PlatformExchangeCreateOrderError,
			Time:    time.Now(),
		}
		return PutOrderOut{
			ErrorCode: errors.PlatformExchangeCreateOrderError,
		}
	}

	orderDTO, err := p.exchangeOrder.GetByUUID(ctx, orderUUID)
	if err != nil {
		p.errChan <- models.ErrMessage{
			Message: "Update order error",
			ErrNum:  errors.PlatformExchangeUpdateOrderError,
			Time:    time.Now(),
		}
		return PutOrderOut{
			ErrorCode: errors.PlatformExchangeUpdateOrderError,
		}
	}

	p.WriteOrderLog(ctx, models.ExchangeOrderLogDTO{
		UUID:       orderDTO.UUID,
		UserID:     in.Bot.UserID,
		ExchangeID: in.Bot.ExchangeID,
		Pair:       in.Signal.Pair,
		OrderID:    orderDTO.ID,
		Status:     service.OrderStatusProcessing,
	})

	var exchangeOrder service.OrderOut
	var openOrdersCount int
	if in.Bot.OrderCountLimit {
		openOrders, err := p.GetOrderList(ctx, utils.Condition{
			Equal: map[string]interface{}{"exchange_order_type": service.ExchangeOrderTypeLimit, "side": service.SideSell, "status": service.OrderStatusNew, "bot_uuid": in.Bot.UUID},
		})
		if err != nil {
			exchangeOrder.ErrorCode = 30001
			exchangeOrder.Message = "get order count limit err"
		}
		for i := range openOrders {
			openOrdersCount += openOrders[i].UnitedOrders
		}
	}

	if exchangeOrder.ErrorCode == errors.NoError {

		switch in.Signal.OrderType {
		case service.TypeBuyMarket, service.TypeAverage:
			if in.Signal.OrderType == service.TypeBuyMarket && in.Bot.AutoSell && in.Bot.OrderCountLimit && openOrdersCount >= in.Bot.OrderCount {
				exchangeOrder.ErrorCode = 3000
				exchangeOrder.Message = fmt.Sprintf("order count limit exceeded, limit: %d, open orders: %d", in.Bot.OrderCount, openOrdersCount)
				break
			}
			p.logger.Info("buy market")
			exchangeOrder = exchangeClient.BuyMarket(service.MarketIn{
				Pair:     in.Signal.Pair,
				Quantity: quantity,
			})
			p.logger.Info("buy market end")
			if exchangeOrder.ErrorCode == errors.NoError {
				quantity = exchangeOrder.Quantity
			}

			if in.Bot.AutoSell && exchangeOrder.ErrorCode == errors.NoError && in.Signal.OrderType != service.TypeAverage {
				autoSellPrice := exchangeOrder.Price.Mul(decimal.NewFromFloat(1 + in.Bot.AutoLimitSellPercent/100)).Round(RoundPlaces)
				in.BuyOrder = orderDTO
				in.BuyOrder.Price = exchangeOrder.Price
				in.Signal.OrderType = service.TypeAutoSellLimit
				go p.OrderSellLimit(ctx, in, quantity, autoSellPrice, 1)
			}
			if exchangeOrder.ErrorCode == errors.NoError && in.Signal.OrderType == service.TypeAverage {
				p.logger.Info("average started")
				log.Println(exchangeOrder)
				if in.Bot.OrderCountLimit && openOrdersCount >= in.Bot.OrderCount {
					exchangeOrder.ErrorCode = 3000
					exchangeOrder.Message = fmt.Sprintf("order count limit exceeded, limit: %d, open orders: %d", in.Bot.OrderCount, openOrdersCount)
					break
				}
				in.BuyOrder = orderDTO
				_ = func() error {
					var canceledQuantitySumm, averageSellPrice, averageBuyPrice decimal.Decimal
					// cancel all sell limit orders
					var cancelOrderOut CancelOrderOut
					cancelOrderOut, err = p.CancelOrder(ctx, in)
					var count float64
					var unitedOrders int
					for _, platformOrder := range cancelOrderOut.PlatformOrders {
						if platformOrder.BuyPrice.GreaterThan(decimal.NewFromInt(0)) {
							averageBuyPrice = averageBuyPrice.Add(platformOrder.BuyPrice)
							unitedOrders += platformOrder.UnitedOrders
							count++
						}
						canceledQuantitySumm = canceledQuantitySumm.Add(platformOrder.Quantity)
					}

					// add current order
					averageBuyPrice = averageBuyPrice.Add(exchangeOrder.Price)
					canceledQuantitySumm = canceledQuantitySumm.Add(quantity)
					count++
					unitedOrders++

					// calculate average
					averageBuyPrice = averageBuyPrice.DivRound(decimal.NewFromInt(int64(count)), RoundPlaces)

					averageSellPrice = averageBuyPrice.Mul(decimal.NewFromFloat(1 + in.Bot.AutoLimitSellPercent/100)).Round(RoundPlaces)
					in.BuyOrder.Price = averageBuyPrice
					p.logger.Info("order sell limit started")
					p.OrderSellLimit(ctx, in, canceledQuantitySumm, averageSellPrice, unitedOrders)
					p.logger.Info("order sell limit end")

					return nil
				}()
				p.logger.Info("average end")
			}
		case service.TypeSellMarket:
			exchangeOrder = exchangeClient.SellMarket(service.MarketIn{
				Pair:     in.Signal.Pair,
				Quantity: quantity,
			})
		case service.TypeBuyLimit:
			if in.Bot.OrderCountLimit && openOrdersCount >= in.Bot.OrderCount {
				exchangeOrder.ErrorCode = 3000
				exchangeOrder.Message = fmt.Sprintf("order count limit exceeded, limit: %d, open orders: %d", in.Bot.OrderCount, openOrdersCount)
				break
			}
			exchangeOrder = exchangeClient.BuyLimit(service.LimitIn{
				Pair:     in.Signal.Pair,
				Quantity: quantity,
				Price:    price,
			})
		case service.TypeSellLimit:
			if in.Bot.OrderCountLimit && openOrdersCount >= in.Bot.OrderCount {
				exchangeOrder.ErrorCode = 3000
				exchangeOrder.Message = fmt.Sprintf("order count limit exceeded, limit: %d, open orders: %d", in.Bot.OrderCount, openOrdersCount)
				break
			}
			exchangeOrder = exchangeClient.SellLimit(service.LimitIn{
				Pair:     in.Signal.Pair,
				Quantity: quantity,
				Price:    price,
			})
		}
	}

	if exchangeOrder.ErrorCode != errors.NoError {
		p.WriteOrderLog(ctx, models.ExchangeOrderLogDTO{
			UUID:       orderUUID,
			OrderID:    orderDTO.ID,
			Amount:     exchangeOrder.Amount,
			UserID:     in.Bot.UserID,
			ExchangeID: in.Bot.ExchangeID,
			Pair:       in.Signal.Pair,
			Quantity:   quantity,
			Price:      in.Signal.PairPrice,
			Status:     StatusFailed,
		})
		orderDTO.Message = exchangeOrder.Message
		orderDTO.SetStatus(service.OrderStatusFailed).
			SetMessage(exchangeOrder.Message).
			SetPrice(in.Signal.PairPrice).
			SetQuantity(quantity).
			SetAmount(quantity.Mul(in.Signal.PairPrice).Round(RoundPlaces))
		_ = p.UpdateOrder(ctx, orderDTO)
		return PutOrderOut{
			ErrorCode: exchangeOrder.ErrorCode,
			OrderID:   orderDTO.ID,
			OrderUUID: orderDTO.UUID,
		}
	}

	orderDTO.SetStatus(exchangeOrder.Status).
		SetExchangeOrderID(exchangeOrder.OrderID).
		SetAmount(exchangeOrder.Amount).
		SetPrice(exchangeOrder.Price).
		SetQuantity(quantity)
	orderDTO.ExchangeOrderType = exchangeOrder.ExchangeOrderType
	err = p.UpdateOrder(ctx, orderDTO)
	if err != nil {
		p.errChan <- models.ErrMessage{
			Message: "Update order error",
			ErrNum:  errors.PlatformExchangeUpdateOrderError,
			Time:    time.Now(),
		}
		return PutOrderOut{
			ErrorCode: errors.PlatformExchangeUpdateOrderError,
		}
	}

	statOut := p.statistic.UpdateStatistic(ctx, statservice.StatisticUpdateIn{
		BotUUID: orderDTO.BotUUID,
		UserID:  orderDTO.UserID,
		Price:   orderDTO.Price,
		Amount:  orderDTO.Amount,
		Status:  orderDTO.Status,
		Side:    orderDTO.Side,
	})
	if statOut.ErrorCode != errors.NoError {
		p.errChan <- models.ErrMessage{
			Message: "Update statistic error",
			ErrNum:  errors.StatisticServiceUpdateErr,
			Time:    time.Now(),
		}
		return PutOrderOut{
			ErrorCode: errors.StatisticServiceUpdateErr,
		}
	}

	p.WriteOrderLog(ctx, models.ExchangeOrderLogDTO{
		UUID:       orderUUID,
		OrderID:    orderDTO.ID,
		Amount:     exchangeOrder.Amount,
		UserID:     in.Bot.UserID,
		ExchangeID: in.Bot.ExchangeID,
		Pair:       in.Signal.Pair,
		Quantity:   quantity,
		Price:      exchangeOrder.Price,
		Status:     exchangeOrder.Status,
	})
	if err != nil {
		p.errChan <- models.ErrMessage{
			Message: "Create order error",
			ErrNum:  errors.PlatformExchangeCreateOrderLogError,
			Time:    time.Now(),
		}
		return PutOrderOut{
			ErrorCode: errors.PlatformExchangeCreateOrderLogError,
		}
	}

	return PutOrderOut{
		Success:   true,
		OrderID:   orderDTO.ID,
		OrderUUID: orderDTO.UUID,
	}
}

func (p *Order) WriteOrderLog(ctx context.Context, orderLogDTO models.ExchangeOrderLogDTO) {
	err := p.exchangeOrderLog.Create(ctx, orderLogDTO)
	if err != nil {
		_ = err
	}
}

func (p *Order) WriteOrder(ctx context.Context, in WriteOrderIn) error {
	if in.OrderType == 0 || in.Side == 0 {
		return fmt.Errorf("bad input data, ordertype, order side")
	}
	orderUUID := uuid.NewString()
	orderDTO := models.ExchangeOrderDTO{
		BotUUID:           in.Bot.UUID,
		UUID:              orderUUID,
		UserID:            in.Bot.UserID,
		ExchangeID:        in.Bot.ExchangeID,
		Pair:              in.Signal.Pair,
		PairID:            in.Signal.PairID,
		OrderType:         in.OrderType,
		Message:           in.Message,
		ExchangeOrderType: in.ExchangeOrder.ExchangeOrderType,
		Side:              in.Side,
		UnitedOrders:      in.UnitedOrders,
		ApiKeyID:          in.Key.ID,
	}
	orderDTO.SetStatus(in.ExchangeOrder.Status).
		SetExchangeOrderID(in.ExchangeOrder.OrderID).
		SetAmount(in.ExchangeOrder.Amount).
		SetPrice(in.ExchangeOrder.Price).
		SetQuantity(in.ExchangeOrder.Quantity)

	orderDTO.BuyPrice = in.BuyOrder.Price
	orderDTO.BuyOrderID = in.BuyOrder.ID

	return p.exchangeOrder.Create(ctx, orderDTO)
}

func (p *Order) OrderSellLimit(ctx context.Context, in OrderIn, quantity, price decimal.Decimal, unitedOrders int) {
	var err error
	p.logger.Info("exchange sell limit started")
	autoSellOrder := in.ExClient.SellLimit(service.LimitIn{
		Pair:     in.Signal.Pair,
		Quantity: quantity,
		Price:    price,
	})
	p.logger.Info("exchange sell limit end")
	if autoSellOrder.ErrorCode != errors.NoError {
		autoSellOrder.Quantity = quantity
		autoSellOrder.Price = price
	}
	err = p.WriteOrder(ctx, WriteOrderIn{
		ExchangeOrder: autoSellOrder,
		Webhook:       in.Webhook,
		Bot:           in.Bot,
		Signal:        in.Signal,
		BuyOrder:      in.BuyOrder,
		Message:       autoSellOrder.Message,
		Side:          service.SideSell,
		OrderType:     in.Signal.OrderType,
		UnitedOrders:  unitedOrders,
		Key:           in.Key,
	})
	if err != nil {
		p.logger.Error("error auto sell write order", zap.Error(err))
	}
}

func (p *Order) GetOrdersCondition(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrder, error) {
	ordersDTO, err := p.exchangeOrder.GetList(ctx, condition)
	if err != nil {
		return nil, err
	}

	var orders []models.ExchangeOrder
	err = gomapper.MapStructs(&orders, &ordersDTO)
	if err != nil {
		return nil, err
	}

	err = p.addOrdersHistory(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (p *Order) addOrdersHistory(ctx context.Context, orders *[]models.ExchangeOrder) error {
	m := make(map[int64]*models.ExchangeOrder, len(*orders))
	var ids []int64
	for i := range *orders {
		id := (*orders)[i].ID
		m[id] = &(*orders)[i]
		ids = append(ids, id)
		(*orders)[i].StatusMsg = service.GetStatus((*orders)[i].Status)
		(*orders)[i].OrderTypeMsg = service.GetOrderTypeRaw((*orders)[i].OrderType)

		switch (*orders)[i].Side {
		case service.SideSell:
			(*orders)[i].SideMsg = "Sell"
		case service.SideBuy:
			(*orders)[i].SideMsg = "Buy"
		}
	}

	orderLogsDTO, err := p.exchangeOrderLog.GetList(ctx, utils.Condition{Equal: map[string]interface{}{"order_id": ids}})
	if err != nil {
		return err
	}

	var orderLogs []models.ExchangeOrderLog

	err = gomapper.MapStructs(&orderLogs, &orderLogsDTO)
	if err != nil {
		return err
	}
	for i := range orderLogs {
		orderLogs[i].StatusMsg = service.GetStatus(orderLogs[i].Status)
		m[orderLogs[i].OrderID].History = append(m[orderLogs[i].OrderID].History, orderLogs[i])
	}

	return nil
}

func (p *Order) GetOrdersStatistic(ctx context.Context, in GetBotRelationIn) StatisticOut {
	ordersDTO, err := p.exchangeOrder.GetList(ctx, utils.Condition{
		Equal: map[string]interface{}{"bot_uuid": in.BotUUID, "deleted_at": nil},
		Order: []*utils.Order{{
			Field: "id",
		}},
	})
	if err != nil {
		return StatisticOut{}
	}

	var orders []models.ExchangeOrder
	err = gomapper.MapStructs(&orders, &ordersDTO)
	if err != nil {
		return StatisticOut{}
	}

	return p.AddOrdersStatistic(ctx, &orders)
}

func (p *Order) AddOrdersStatistic(ctx context.Context, orders *[]models.ExchangeOrder) StatisticOut {
	var ids []int64
	buyOrders := make(map[int]map[int64]*models.ExchangeOrder, len(*orders))
	sellOrders := make(map[int][]*models.ExchangeOrder)
	keys := make(map[int]*models.ExchangeUserKey)
	for i := range *orders {
		id := (*orders)[i].ID
		ids = append(ids, id)
		(*orders)[i].StatusMsg = service.GetStatus((*orders)[i].Status)
		(*orders)[i].OrderTypeMsg = service.GetOrderTypeRaw((*orders)[i].OrderType)

		if _, ok := keys[(*orders)[i].ApiKeyID]; !ok {
			exchangeUserKeyDTO, err := p.exchangeUserKeyService.ExchangeUserKeyGetByID(ctx, (*orders)[i].ApiKeyID)
			if err != nil {
				continue
			}
			var exchangeUserKey models.ExchangeUserKey
			err = gomapper.MapStructs(&exchangeUserKey, &exchangeUserKeyDTO)
			if err != nil {
				continue
			}

			keys[exchangeUserKey.ID] = &exchangeUserKey
		}

		switch (*orders)[i].Side {
		case service.SideSell:
			(*orders)[i].SideMsg = service.OrderSides[service.SideSell]
			key := keys[(*orders)[i].ApiKeyID]
			if key == nil {
				continue
			}
			if (*orders)[i].Status == service.OrderStatusFilled {
				key.StatisticData.SumSell = key.StatisticData.SumSell.Add((*orders)[i].Amount)
				sellOrders[(*orders)[i].ApiKeyID] = append(sellOrders[(*orders)[i].ApiKeyID], &(*orders)[i])
			}
			if (*orders)[i].Status == service.OrderStatusNew || (*orders)[i].Status == service.OrderStatusPartiallyFilled {
				key.StatisticData.ToSell = key.StatisticData.ToSell.Add((*orders)[i].Amount)
			}
		case service.SideBuy:
			(*orders)[i].SideMsg = service.OrderSides[service.SideBuy]
			if (*orders)[i].Status == service.OrderStatusFilled {
				key := keys[(*orders)[i].ApiKeyID]
				if key == nil {
					continue
				}
				key.StatisticData.SumBuy = key.StatisticData.SumBuy.Sub((*orders)[i].Amount)
				if buyOrders[(*orders)[i].ApiKeyID] == nil {
					buyOrders[(*orders)[i].ApiKeyID] = make(map[int64]*models.ExchangeOrder)
				}
				buyOrders[(*orders)[i].ApiKeyID][(*orders)[i].ID] = &(*orders)[i]
			}
		}
	}

	// calc earned
	for keyID := range sellOrders {
		for i := range sellOrders[keyID] {
			if v, ok := buyOrders[keyID][sellOrders[keyID][i].OrderID]; ok {
				earned := sellOrders[keyID][i].Amount.Sub(v.Amount)
				keys[keyID].StatisticData.Earned = keys[keyID].StatisticData.Earned.Add(earned)
			}
		}
	}

	// calc Profit
	var keysData []models.ExchangeUserKey
	for keyID := range keys {
		keys[keyID].StatisticData.Profit = keys[keyID].StatisticData.SumSell.Sub(keys[keyID].StatisticData.SumBuy)
		keys[keyID].StatisticData.ToEarn = keys[keyID].StatisticData.Profit.Sub(keys[keyID].StatisticData.ToSell)
		keysData = append(keysData, *keys[keyID])
	}

	return StatisticOut{
		Keys: keysData,
	}
}

func (p *Order) GetBotOrders(ctx context.Context, in GetBotRelationIn) GetOrdersOut {
	orders, err := p.GetOrdersCondition(ctx, utils.Condition{Equal: map[string]interface{}{"bot_uuid": in.BotUUID}})
	if err != nil {
		return GetOrdersOut{
			ErrorCode: errors.InternalError,
		}
	}

	return GetOrdersOut{
		Success: true,
		Data:    orders,
	}
}

func (p *Order) GetUserOrders(ctx context.Context, in GetUserRelationIn) GetOrdersOut {
	orders, err := p.GetOrdersCondition(ctx, utils.Condition{Equal: map[string]interface{}{"user_id": in.UserID, "deleted_at": nil}, Order: []*utils.Order{{Field: "id", Asc: false}}})
	if err != nil {
		return GetOrdersOut{
			ErrorCode: errors.InternalError,
		}
	}

	return GetOrdersOut{
		Success: true,
		Data:    orders,
	}
}

func (p *Order) ExchangeOrderList(ctx context.Context, in GetBotRelationIn) GetOrdersOut {
	exchangeOrders, err := p.GetOrdersCondition(ctx, utils.Condition{
		Equal: map[string]interface{}{"bot_uuid": in.BotUUID, "deleted_at": nil},
		Order: []*utils.Order{{
			Field: "id",
		}},
	})
	if err != nil {
		return GetOrdersOut{
			ErrorCode: 2000,
		}
	}

	return GetOrdersOut{
		Success: true,
		Data:    exchangeOrders,
	}
}

func (p *Order) GetAllOrdersStatistic(ctx context.Context, in GetUserRelationIn) StatisticOut {
	ordersDTO, err := p.exchangeOrder.GetList(ctx, utils.Condition{
		Equal: map[string]interface{}{"user_id": in.UserID, "deleted_at": nil},
		Order: []*utils.Order{{
			Field: "id",
		}},
	})
	if err != nil {
		return StatisticOut{}
	}

	var orders []models.ExchangeOrder
	err = gomapper.MapStructs(&orders, &ordersDTO)
	if err != nil {
		return StatisticOut{}
	}

	return p.AddOrdersStatistic(ctx, &orders)
}

func (p *Order) GetOrderList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error) {
	return p.exchangeOrder.GetList(ctx, condition)
}

func (p *Order) UpdateOrder(ctx context.Context, dto models.ExchangeOrderDTO) error {
	return p.exchangeOrder.Update(ctx, dto)
}
