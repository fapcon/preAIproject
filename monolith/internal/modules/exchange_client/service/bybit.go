package service

import (
	"context"
	"github.com/hirokisan/bybit/v2"
	"github.com/shopspring/decimal"
	"go.uber.org/ratelimit"
	"log"
	"strings"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/helper"
)

const BybitRate = 10

func NewBybitWithKey(userID int, rateLimiter *RateLimiter, apiKey string, secretKey string) Exchanger {
	client := bybit.NewClient().WithAuth(apiKey, secretKey)

	limiter := rateLimiter.GetUserLimiter(userID, BybitRate)
	return NewBybitExchange(client, limiter)
}

type BybitExchange struct {
	limiter    ratelimit.Limiter
	bybit      *bybit.Client
	exchangeID int
}

func (b *BybitExchange) OrderMarket(in OrderMarketIn) OrderOut {
	var orderType int
	switch in.Side {
	case SideBuy:
		orderType = TypeBuyMarket
	case SideSell:
		orderType = TypeSellMarket
	default:
		return OrderOut{
			ErrorCode: errors.GeneralError,
		}
	}

	return b.CreateOrder(CreateOrderIn{
		Type:     orderType,
		Side:     in.Side,
		Pair:     in.Request.Pair,
		Quantity: in.Request.Quantity,
	})
}

func (b *BybitExchange) BuyMarket(in MarketIn) OrderOut {
	return b.OrderMarket(OrderMarketIn{
		Side:    SideBuy,
		Request: in,
	})
}

func (b *BybitExchange) SellMarket(in MarketIn) OrderOut {
	return b.OrderMarket(OrderMarketIn{
		Side:    SideSell,
		Request: in,
	})
}

func (b *BybitExchange) OrderLimit(in OrderLimitIn) OrderOut {
	var orderType int
	switch in.Side {
	case SideBuy:
		orderType = TypeBuyLimit
	case SideSell:
		orderType = TypeSellLimit
	default:
		return OrderOut{
			ErrorCode: errors.GeneralError,
		}
	}

	return b.CreateOrder(CreateOrderIn{
		Type:     orderType,
		Side:     in.Side,
		Pair:     in.Request.Pair,
		Quantity: in.Request.Quantity,
		Price:    in.Request.Price,
	})
}

func (b *BybitExchange) BuyLimit(in LimitIn) OrderOut {
	return b.OrderLimit(OrderLimitIn{
		Side:    SideBuy,
		Request: in,
	})
}

func (b *BybitExchange) SellLimit(in LimitIn) OrderOut {
	return b.OrderLimit(OrderLimitIn{
		Side:    SideSell,
		Request: in,
	})
}

func (b *BybitExchange) CreateOrder(in CreateOrderIn) OrderOut {
	var orderSide bybit.Side
	var positionIndex int
	switch in.Side {
	case SideBuy:
		orderSide = bybit.SideBuy
		positionIndex = 1
	case SideSell:
		orderSide = bybit.SideSell
		positionIndex = 2
	default:
		return OrderOut{
			ErrorCode: errors.GeneralError,
		}
	}

	var orderType bybit.OrderType
	switch in.Type {
	case TypeSellMarket:
		fallthrough
	case TypeBuyMarket:
		orderType = bybit.OrderTypeMarket
	case TypeBuyLimit:
		fallthrough
	case TypeSellLimit:
		orderType = bybit.OrderTypeLimit
	}

	quantity := GetQuantity(in.Quantity)
	price := GetPrice(in.Price)

	var timeInForce bybit.TimeInForce
	if in.Type == TypeBuyLimit || in.Type == TypeSellLimit {
		timeInForce = bybit.TimeInForceGoodTillCancel
	}

	orderParams := bybit.V5CreateOrderParam{
		Category:       bybit.CategoryV5Linear,
		Symbol:         bybit.SymbolV5(in.Pair),
		Side:           orderSide,
		OrderType:      orderType,
		Qty:            quantity,
		Price:          &price,
		TimeInForce:    &timeInForce,
		PositionIdx:    (*bybit.PositionIdx)(&positionIndex),
		ReduceOnly:     new(bool),
		CloseOnTrigger: new(bool),
	}

	b.limiter.Take()
	order, err := b.bybit.V5().Order().CreateOrder(orderParams)
	if err != nil {
		// refactor to recursion
		i := 1
		for err != nil {
			log.Println(err)
			switch {
			case strings.Contains(err.Error(), "recv_window"):
				b.limiter.Take()
				order, err = b.bybit.V5().Order().CreateOrder(orderParams)
				i++
			case strings.Contains(strings.ToLower(err.Error()), "insufficient"):
				in.Quantity = in.Quantity.Sub(in.Quantity.DivRound(decimal.NewFromInt(100), RoundPlaces))
				orderParams.Qty = GetQuantity(in.Quantity)
				b.limiter.Take()
				order, err = b.bybit.V5().Order().CreateOrder(orderParams)
				i++
			default:
				i = 9999
			}
			if i > 3 {
				break
			}
		}
	}

	if err != nil {
		return OrderOut{
			ErrorCode: order.RetCode,
			Message:   order.RetMsg,
		}
	}

	return b.composeOrder(order, in.Type)
}

func (b *BybitExchange) CancelOrder(ctx context.Context, in CancelOrderIn) OrderOut {
	b.limiter.Take()
	order, err := b.bybit.V5().Order().CancelOrder(bybit.V5CancelOrderParam{OrderID: &in.OrderID})
	if err != nil {
		return OrderOut{ErrorCode: order.RetCode, Message: order.RetMsg}
	}

	return b.composeOrder(order, TypeGetOrder)
}

func (b *BybitExchange) GetOrder(ctx context.Context, in GetOrderIn) OrderOut {
	b.limiter.Take()
	getOrder, err := b.bybit.V5().Order().GetOpenOrders(bybit.V5GetOpenOrdersParam{
		Category: bybit.CategoryV5Linear,
		OrderID:  &in.OrderID,
	})
	if err != nil || len(getOrder.Result.List) == 0 {
		return OrderOut{
			ErrorCode: getOrder.RetCode,
			Message:   getOrder.RetMsg,
		}
	}

	return b.composeOrder(getOrder, TypeGetOrder)
}

func (b *BybitExchange) GetOrdersHistory(in EIn) EOut {
	//TODO implement me
	panic("implement me")
}

func (b *BybitExchange) GetOpenOrders(in EIn) EOut {
	//TODO implement me
	panic("implement me")
}

func getBybitBalance(bybitResponse *bybit.V5GetWalletBalanceResponse) []Balance {
	var balances []Balance
	for _, coin := range bybitResponse.Result.List[0].Coin {
		amount, _ := decimal.NewFromString(coin.WalletBalance)
		locked, _ := decimal.NewFromString(coin.Locked)
		if amount.Equal(decimal.NewFromInt(0)) && locked.Equal(decimal.NewFromInt(0)) {
			continue
		}
		balances = append(balances, Balance{
			Currency: string(coin.Coin),
			Amount:   amount,
			Locked:   locked,
		})
	}
	return balances
}

func getBybitFuturesWallet(futuresResponse *bybit.V5GetWalletBalanceResponse) []BalanceMargin {
	var balances []BalanceMargin
	for _, coin := range futuresResponse.Result.List[0].Coin {
		free, _ := decimal.NewFromString(coin.WalletBalance)
		locked, _ := decimal.NewFromString(coin.Locked)
		borrowed, _ := decimal.NewFromString(coin.BorrowAmount)
		interest, _ := decimal.NewFromString(coin.AccruedInterest)
		netAsset, _ := decimal.NewFromString(coin.Equity)
		if netAsset.Equal(decimal.NewFromInt(0)) {
			continue
		}
		balances = append(balances, BalanceMargin{
			Currency: string(coin.Coin),
			Free:     free,
			Locked:   locked,
			Borrowed: borrowed,
			Interest: interest,
			NetAsset: netAsset,
		})
	}
	return balances
}

func (b *BybitExchange) GetAccount(ctx context.Context) GetAccountOut {
	b.limiter.Take()
	accSpot, err := b.bybit.V5().Account().GetWalletBalance(bybit.AccountType(bybit.AccountTypeV5SPOT), []bybit.Coin{})
	if err != nil {
		return GetAccountOut{
			ErrorCode: accSpot.RetCode,
			Message:   accSpot.RetMsg,
		}
	}
	balanceSpot := getBybitBalance(accSpot)
	b.limiter.Take()
	accContract, err := b.bybit.V5().Account().GetWalletBalance(bybit.AccountType(bybit.AccountTypeV5CONTRACT), []bybit.Coin{})
	if err != nil {
		return GetAccountOut{
			ErrorCode: accContract.RetCode,
			Message:   accContract.RetMsg,
		}
	}
	marginBalance := getBybitFuturesWallet(accContract)

	return GetAccountOut{
		DataSpot: AccountSpot{
			Balances: balanceSpot,
		},
		DataMargin: AccountMargin{
			Balances: marginBalance,
		},
		Success: true,
	}
}

func (b *BybitExchange) GetBalances(ctx context.Context) GetAccountBalanceOut {
	account := b.GetAccount(ctx)
	if account.ErrorCode != errors.NoError {
		return GetAccountBalanceOut{
			ErrorCode: account.ErrorCode,
			Success:   false,
		}
	}

	var res GetAccountBalanceOut
	dataSpotBalance := make([]models.BalanceDTO, len(account.DataSpot.Balances))
	for i := range account.DataSpot.Balances {
		dataSpotBalance[i].Currency = account.DataSpot.Balances[i].Currency
		dataSpotBalance[i].Locked = account.DataSpot.Balances[i].Locked
		dataSpotBalance[i].Amount = account.DataSpot.Balances[i].Amount
	}
	res.DataSpotBalance = dataSpotBalance
	dataMarginBalance := make([]models.BalanceDTO, len(account.DataMargin.Balances))
	for j := range account.DataMargin.Balances {
		dataMarginBalance[j].Currency = account.DataMargin.Balances[j].Currency
		dataMarginBalance[j].Locked = account.DataMargin.Balances[j].Locked
		dataMarginBalance[j].Amount = account.DataMargin.Balances[j].Free
	}
	res.DataMarginBalance = dataMarginBalance

	return res
}

func (b *BybitExchange) GetTicker(ctx context.Context, in GetTickerIn) GetTickerOut {
	b.limiter.Take()
	tickers, err := b.bybit.V5().Market().GetTickers(bybit.V5GetTickersParam{
		Category: bybit.CategoryV5Linear,
	})
	if err != nil {
		return GetTickerOut{
			ErrorCode: errors.ExchangeServiceGetTickerErr,
		}
	}
	res := make(map[string]decimal.Decimal, len(tickers.Result.LinearInverse.List))
	var price decimal.Decimal
	for _, ticker := range tickers.Result.LinearInverse.List {
		price, err = decimal.NewFromString(ticker.Ask1Price)
		if err != nil {
			return GetTickerOut{
				ErrorCode: errors.ExchangeServiceParsePriceErr,
			}
		}
		res[string(ticker.Symbol)] = price
	}

	return GetTickerOut{
		Data: res,
	}
}

func (b *BybitExchange) GetCandles(ctx context.Context, in GetCandlesIn) CandlesOut {
	interval, ok := bybitIntervals[in.Interval]
	if !ok {
		return CandlesOut{ErrorCode: errors.InternalError}
	}
	startTime := int(in.StartTime)
	endTime := int(in.EndTime)
	b.limiter.Take()
	candles, err := b.bybit.V5().Market().GetKline(bybit.V5GetKlineParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5(in.Symbol),
		Interval: interval,
		Start:    &startTime,
		End:      &endTime,
		Limit:    &in.Limit,
	})
	if err != nil {
		return CandlesOut{
			ErrorCode: errors.GeneralError,
		}
	}
	klines := make([]CandlesData, len(candles.Result.List))
	for i, lines := range candles.Result.List {
		klines[i].OpenTime, _ = helper.MsToTime(lines.StartTime)
		klines[i].Open, _ = decimal.NewFromString(lines.Open)
		klines[i].Low, _ = decimal.NewFromString(lines.Low)
		klines[i].High, _ = decimal.NewFromString(lines.High)
		klines[i].Close, _ = decimal.NewFromString(lines.Close)
		klines[i].QuoteAssetVolume, _ = decimal.NewFromString(lines.Volume)
	}
	return CandlesOut{Candles: klines}
}

func NewBybitExchange(bybit *bybit.Client, rateLimiter ratelimit.Limiter) Exchanger {
	return &BybitExchange{bybit: bybit, exchangeID: Bybit, limiter: rateLimiter}
}

func (b *BybitExchange) composeOrder(orderRaw interface{}, orderType int) OrderOut {
	switch orderType {
	case TypeGetOrder:
		getOrder := orderRaw.(*bybit.V5GetOrdersResponse)
		if len(getOrder.Result.List) == 0 {
			return OrderOut{
				ErrorCode: errors.InternalError,
			}
		}
		quantity, _ := decimal.NewFromString(getOrder.Result.List[0].CumExecQty)

		if quantity.Equal(decimal.NewFromInt(0)) {
			quantity, _ = decimal.NewFromString(getOrder.Result.List[0].Qty)
		}
		price, _ := decimal.NewFromString(getOrder.Result.List[0].AvgPrice)
		if price.Equal(decimal.NewFromInt(0)) {
			price, _ = decimal.NewFromString(getOrder.Result.List[0].Price)
		}
		amount, _ := decimal.NewFromString(getOrder.Result.List[0].CumExecQty)
		if amount.Equal(decimal.NewFromInt(0)) {
			amount = price.Mul(quantity).Round(RoundPlaces)
		}

		status, ok := getStatusesBybit[getOrder.Result.List[0].OrderStatus]
		if !ok {
			return OrderOut{
				ErrorCode: errors.InternalError,
			}
		}

		side, ok := getOrderSide[getOrder.Result.List[0].Side]
		if !ok {
			return OrderOut{
				ErrorCode: errors.InternalError,
			}
		}

		return OrderOut{
			ClientOrderID:     getOrder.Result.List[0].OrderLinkID,
			OrderID:           getOrder.Result.List[0].OrderID,
			Price:             price,
			Amount:            amount,
			Quantity:          quantity,
			Pair:              string(getOrder.Result.List[0].Symbol),
			Status:            status,
			Side:              side,
			Type:              orderType,
			ExchangeOrderType: getOrderTypeBybit[getOrder.Result.List[0].OrderType],
			ErrorCode:         0,
			Message:           "OK",
		}

	case TypeSellMarket:
		fallthrough
	case TypeBuyMarket:
		fallthrough
	case TypeBuyLimit:
		fallthrough
	case TypeSellLimit:
		order := orderRaw.(*bybit.V5CreateOrderResponse)
		// API V5 возвращает только orderID после размещения ордера, поэтому проверяем ордер для получения параметров
		getOrder := b.GetOrder(context.Background(), GetOrderIn{OrderID: order.Result.OrderID})
		getOrder.Type = orderType

		return getOrder
	case TypeCancelAll,
		TypeCancelBuy,
		TypeCancelSell,
		OrderTypeCancel:
		order := orderRaw.(*bybit.V5CancelOrderResponse)
		getOrder := b.GetOrder(context.Background(), GetOrderIn{OrderID: order.Result.OrderID})
		getOrder.Type = orderType

		return getOrder
	}

	return OrderOut{
		ErrorCode: errors.InternalError,
	}
}
