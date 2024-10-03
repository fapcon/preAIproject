package service

import (
	"context"
	"go.uber.org/ratelimit"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/common"
	"github.com/shopspring/decimal"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

const RoundPlaces = 8
const BinanceRate = 10

var (
	apiKeyPlatform    = "5kOvWwvETSWX9fjxKC8vpjDWpUyn1vjzqtoefqljLatLvnGf4fMTe9ZHjqKSxRsz" // API key - предоставил Петр
	secretKeyPlatform = "et6TmjLwEkSKoo9RN5ajmhV4CGV0ccF47qqIZHXHPob9T3fwQE0sSv4LigK7vbFk" // Secret key - предоставил Петр
)

func NewPlatformBinance(rateLimiter *RateLimiter) Exchanger {
	client := binance.NewClient(apiKeyPlatform, secretKeyPlatform)
	limiter := rateLimiter.GetUserLimiter(0, BinanceRate)
	return NewExchange(client, limiter)
}

func NewBinanceWithKey(userID int, rateLimiter *RateLimiter, apiKey string, secretKey string) Exchanger {
	client := binance.NewClient(apiKey, secretKey)
	limiter := rateLimiter.GetUserLimiter(userID, BinanceRate)
	return NewExchange(client, limiter)
}

type Exchange struct {
	Limiter    ratelimit.Limiter
	Binance    *binance.Client
	ExchangeID int
}

func NewExchange(binance *binance.Client, rateLimiter ratelimit.Limiter) Exchanger {
	return &Exchange{Binance: binance, ExchangeID: Binance, Limiter: rateLimiter}
}

func (e *Exchange) BuyMarket(in MarketIn) OrderOut {
	return e.OrderMarket(OrderMarketIn{
		Side:    SideBuy,
		Request: in,
	})
}

func (e *Exchange) SellMarket(in MarketIn) OrderOut {
	return e.OrderMarket(OrderMarketIn{
		Side:    SideSell,
		Request: in,
	})
}

func (e *Exchange) OrderMarket(in OrderMarketIn) OrderOut {
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

	return e.CreateOrder(CreateOrderIn{
		Type:     orderType,
		Side:     in.Side,
		Pair:     in.Request.Pair,
		Quantity: in.Request.Quantity,
	})
}

func (e *Exchange) OrderLimit(in OrderLimitIn) OrderOut {
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

	return e.CreateOrder(CreateOrderIn{
		Type:     orderType,
		Side:     in.Side,
		Pair:     in.Request.Pair,
		Quantity: in.Request.Quantity,
		Price:    in.Request.Price,
	})
}

func (e *Exchange) BuyLimit(in LimitIn) OrderOut {
	return e.OrderLimit(OrderLimitIn{
		Side:    SideBuy,
		Request: in,
	})
}

func (e *Exchange) SellLimit(in LimitIn) OrderOut {
	return e.OrderLimit(OrderLimitIn{
		Side:    SideSell,
		Request: in,
	})
}

func (e *Exchange) CreateOrder(in CreateOrderIn) OrderOut {
	var binanceSide binance.SideType
	switch in.Side {
	case SideBuy:
		binanceSide = binance.SideTypeBuy
	case SideSell:
		binanceSide = binance.SideTypeSell
	default:
		return OrderOut{
			ErrorCode: errors.GeneralError,
		}
	}
	var orderType binance.OrderType
	switch in.Type {
	case TypeSellMarket:
		fallthrough
	case TypeBuyMarket:
		orderType = binance.OrderTypeMarket
	case TypeBuyLimit:
		fallthrough
	case TypeSellLimit:
		orderType = binance.OrderTypeLimit
	}
	quantity := GetQuantity(in.Quantity)
	createOrder := e.Binance.NewCreateOrderService().Symbol(in.Pair).
		Side(binanceSide).Type(orderType).Quantity(quantity)
	price := GetPrice(in.Price)
	if in.Type == TypeBuyLimit || in.Type == TypeSellLimit {
		createOrder.TimeInForce(binance.TimeInForceTypeGTC).Price(price)
	}
	log.Println("order started")
	e.Limiter.Take()
	order, err := createOrder.Do(context.Background())
	log.Println("order end")
	_ = quantity // bug fix
	if err != nil {
		// refactor to recursion
		i := 1
		for err != nil {
			log.Println(err)
			switch {
			case strings.Contains(err.Error(), "recvWindow"):
				e.Limiter.Take()
				order, err = createOrder.Do(context.Background())
				i++
			case strings.Contains(err.Error(), "LOT_SIZE"):
				quantity = quantity[:len(quantity)-1]
				if quantity[len(quantity)-1:] == "." {
					quantity = quantity[:len(quantity)-1]
				}
				createOrder = createOrder.Quantity(quantity)
				e.Limiter.Take()
				order, err = createOrder.Do(context.Background())
				i++
			case strings.Contains(err.Error(), "PRICE_FILTER"):
				price = price[:len(price)-1]
				if price[len(price)-1:] == "." {
					price = price[:len(price)-1]
				}
				createOrder = createOrder.Price(price)
				e.Limiter.Take()
				order, err = createOrder.Do(context.Background())
				i++
			case strings.Contains(err.Error(), "Account has insufficient balance"):
				in.Quantity = in.Quantity.Sub(in.Quantity.DivRound(decimal.NewFromInt(100), RoundPlaces))
				quantity = GetQuantity(in.Quantity)
				createOrder = createOrder.Quantity(quantity)
				e.Limiter.Take()
				order, err = createOrder.Do(context.Background())
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
		if apiError, ok := err.(*common.APIError); ok {
			return OrderOut{
				ErrorCode: int(apiError.Code),
				Message:   apiError.Message,
			}
		}
		return OrderOut{
			ErrorCode: errors.InternalError,
		}
	}

	return e.ComposeOrder(order, in.Type, in.Side)
}

func GetPrice(priceRaw decimal.Decimal) string {
	if priceRaw.GreaterThan(decimal.NewFromInt(10)) {
		return priceRaw.Round(1).String()
	}
	if priceRaw.GreaterThan(decimal.NewFromInt(1)) {
		return priceRaw.Round(2).String()
	}
	price := priceRaw.String()

	idx := 0
	passDot := false
	for i := range price {
		if price[i] == '.' {
			passDot = true
			continue
		}
		if passDot && price[i] != '0' && price[i] != '.' {
			break
		}
		idx++
	}
	if len(price) >= idx+4 {
		price = price[0 : idx+4]
	}

	return price
}

func GetQuantity(quantityRaw decimal.Decimal) string {
	if quantityRaw.GreaterThan(decimal.NewFromInt(10)) {
		return quantityRaw.Round(1).String()
	}
	if quantityRaw.GreaterThan(decimal.NewFromInt(1)) {
		return quantityRaw.Round(2).String()
	}

	quantity := quantityRaw.String()
	var idx int
	var passDot bool
	for i := range quantity {
		if quantity[i] == '.' {
			passDot = true
			continue
		}
		if passDot && quantity[i] != '0' && quantity[i] != '.' {
			break
		}
		idx++
	}
	if len(quantity) > idx+4 {
		quantity = quantity[0 : idx+4]
	}

	return quantity
}

func (e *Exchange) ComposeOrder(orderRaw interface{}, orderType, side int) OrderOut {
	switch orderType {
	case TypeSellMarket:
		fallthrough
	case TypeBuyMarket:
		fallthrough
	case TypeBuyLimit:
		fallthrough
	case TypeSellLimit:
		order := orderRaw.(*binance.CreateOrderResponse)

		quantity, _ := decimal.NewFromString(order.ExecutedQuantity)

		if quantity.Equal(decimal.NewFromInt(0)) {
			quantity, _ = decimal.NewFromString(order.OrigQuantity)
		}
		price, _ := decimal.NewFromString(order.Price)

		var fillPrices decimal.Decimal
		for i := range order.Fills {
			fillPrice, _ := decimal.NewFromString(order.Fills[i].Price)
			fillPrices = fillPrices.Add(fillPrice)
		}
		if fillPrices.GreaterThan(decimal.NewFromInt(0)) {
			price = fillPrices.DivRound(decimal.NewFromInt(int64(len(order.Fills))), RoundPlaces)
		}
		amount, _ := decimal.NewFromString(order.CummulativeQuoteQuantity)
		if amount.Equal(decimal.NewFromInt(0)) {
			amount = price.Mul(quantity).Round(RoundPlaces)
		}
		status, ok := getStatuses[order.Status]
		if !ok {
			return OrderOut{
				ErrorCode: errors.InternalError,
			}
		}
		return OrderOut{
			ClientOrderID:     order.ClientOrderID,
			OrderID:           strconv.FormatInt(order.OrderID, 10),
			Amount:            amount,
			Price:             price,
			Status:            status,
			Type:              orderType,
			ExchangeOrderType: getOrderType[order.Type],
			Pair:              order.Symbol,
			Quantity:          quantity,
			Side:              side,
		}
	case TypeCancelAll,
		TypeCancelBuy,
		TypeCancelSell,
		OrderTypeCancel:
		order := orderRaw.(*binance.CancelOrderResponse)

		quantity, _ := decimal.NewFromString(order.ExecutedQuantity)

		if quantity.Equal(decimal.NewFromInt(0)) {
			quantity, _ = decimal.NewFromString(order.OrigQuantity)
		}
		price, _ := decimal.NewFromString(order.Price)
		amount, _ := decimal.NewFromString(order.CummulativeQuoteQuantity)

		if amount.Equal(decimal.NewFromInt(0)) {
			amount = price.Mul(quantity)
		}

		status, ok := getStatuses[order.Status]
		if !ok {
			return OrderOut{
				ErrorCode: errors.InternalError,
			}
		}
		return OrderOut{
			ClientOrderID:     order.ClientOrderID,
			OrderID:           strconv.FormatInt(order.OrderID, 10),
			Amount:            amount,
			Price:             price,
			Status:            status,
			Type:              orderType,
			Pair:              order.Symbol,
			ExchangeOrderType: getOrderType[order.Type],
			Quantity:          quantity,
			Side:              side,
		}
	case TypeGetOrder:
		order := orderRaw.(*binance.Order)

		quantityExecuted, _ := decimal.NewFromString(order.ExecutedQuantity)

		price, _ := decimal.NewFromString(order.Price)
		amount, _ := decimal.NewFromString(order.CummulativeQuoteQuantity)

		status, ok := getStatuses[order.Status]
		if !ok {
			return OrderOut{
				ErrorCode: errors.InternalError,
			}
		}
		return OrderOut{
			ClientOrderID:     order.ClientOrderID,
			OrderID:           strconv.FormatInt(order.OrderID, 10),
			Amount:            amount,
			Price:             price,
			Status:            status,
			Type:              orderType,
			ExchangeOrderType: getOrderType[order.Type],
			Pair:              order.Symbol,
			Quantity:          quantityExecuted,
			Side:              side,
		}
	}

	return OrderOut{
		ErrorCode: errors.InternalError,
	}
}

func (e *Exchange) CancelOrder(ctx context.Context, in CancelOrderIn) OrderOut {
	orderID, err := strconv.Atoi(in.OrderID)
	if err != nil {
		return OrderOut{ErrorCode: errors.InternalError}
	}
	e.Limiter.Take()
	cancelOrder, err := e.Binance.NewCancelOrderService().OrderID(int64(orderID)).Symbol(in.Pair).Do(ctx)
	if err != nil {
		return OrderOut{ErrorCode: errors.InternalError}
	}

	return e.ComposeOrder(cancelOrder, OrderTypeCancel, SideCancel)
}

func (e *Exchange) GetOrder(ctx context.Context, in GetOrderIn) OrderOut {
	orderID, err := strconv.Atoi(in.OrderID)
	if err != nil {
		return OrderOut{ErrorCode: errors.InternalError}
	}
	e.Limiter.Take()
	order, err := e.Binance.NewGetOrderService().Symbol(in.Pair).
		OrderID(int64(orderID)).Do(ctx)
	if err != nil {
		return OrderOut{
			ErrorCode: 3000,
		}
	}

	return e.ComposeOrder(order, TypeGetOrder, 0)
}

func (e *Exchange) GetOrdersHistory(in EIn) EOut {
	panic("implement me")
}

func (e *Exchange) GetOpenOrders(in EIn) EOut {
	panic("implement me")
}

func (e *Exchange) GetBalances(ctx context.Context) GetAccountBalanceOut {
	account := e.GetAccount(ctx)
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

func (e *Exchange) GetTicker(ctx context.Context, in GetTickerIn) GetTickerOut {
	resRaw, err := e.Binance.NewListBookTickersService().Do(ctx)
	if err != nil {
		return GetTickerOut{
			ErrorCode: errors.ExchangeServiceGetTickerErr,
		}
	}
	res := make(map[string]decimal.Decimal, len(resRaw))
	var price decimal.Decimal
	for _, ticker := range resRaw {
		price, err = decimal.NewFromString(ticker.AskPrice)
		if err != nil {
			return GetTickerOut{
				ErrorCode: errors.ExchangeServiceParsePriceErr,
			}
		}
		res[ticker.Symbol] = price
	}

	return GetTickerOut{
		Data: res,
	}
}

// Разобраться с ErrorCode
func (e *Exchange) GetCandles(ctx context.Context, in GetCandlesIn) CandlesOut {
	// Необходимо ли реализация функции если StartTime и EndTime не передаются и в каком формате данных их принимать
	// Сейчас это int
	res, err := e.Binance.NewKlinesService().Symbol(in.Symbol).Limit(in.Limit).Interval(string(in.Interval)).
		StartTime(in.StartTime).EndTime(in.EndTime).Do(ctx)
	if err != nil {
		return CandlesOut{
			ErrorCode: errors.GeneralError,
		}
	}
	klines := make([]CandlesData, len(res))
	for i, lines := range res {
		klines[i].OpenTime = time.Unix(0, lines.OpenTime*int64(time.Millisecond))
		klines[i].Open, _ = decimal.NewFromString(lines.Open)
		klines[i].Low, _ = decimal.NewFromString(lines.Low)
		klines[i].High, _ = decimal.NewFromString(lines.High)
		klines[i].Close, _ = decimal.NewFromString(lines.Close)
		klines[i].Open, _ = decimal.NewFromString(lines.Open)
		klines[i].CloseTime = time.Unix(0, lines.CloseTime*int64(time.Millisecond))
		// Возможно не нужные значения
		klines[i].QuoteAssetVolume, _ = decimal.NewFromString(lines.QuoteAssetVolume)
		klines[i].TradeNum = lines.TradeNum
		klines[i].TakerBuyBaseAssetVolume, _ = decimal.NewFromString(lines.TakerBuyQuoteAssetVolume)
		klines[i].TakerBuyBaseAssetVolume, _ = decimal.NewFromString(lines.TakerBuyBaseAssetVolume)
	}
	return CandlesOut{Candles: klines}
}

func (e *Exchange) GetAccount(ctx context.Context) GetAccountOut {
	//апи со спотового аккаунта
	e.Limiter.Take()
	spotAccount, err := e.Binance.NewGetAccountService().Do(ctx)
	if err != nil {
		if binanceErr, ok := err.(*common.APIError); ok {
			return GetAccountOut{
				ErrorCode: int(binanceErr.Code),
				Message:   binanceErr.Message,
			}
		}

		return GetAccountOut{ErrorCode: errors.InternalError}
	}
	balanceSpot := GetSpotBalance(spotAccount)

	var permissions []int
	for i := range spotAccount.Permissions {
		if spotAccount.Permissions[i] == PermissionSPOTRaw {
			permissions = append(permissions, PermissionSPOT)
		}
	}

	//апи с маржинального аккаунта
	e.Limiter.Take()
	marginAccount, err := e.Binance.NewGetMarginAccountService().Do(ctx)
	if err != nil {
		if binanceErr, ok := err.(*common.APIError); ok {
			return GetAccountOut{
				ErrorCode: int(binanceErr.Code),
				Message:   binanceErr.Message,
			}
		}

		return GetAccountOut{ErrorCode: errors.InternalError}
	}
	marginBalance := GetMarginBalance(marginAccount)

	return GetAccountOut{
		DataSpot: AccountSpot{
			CanTrade:    spotAccount.CanTrade,
			CanDeposit:  spotAccount.CanDeposit,
			CanWithdraw: spotAccount.CanWithdraw,
			Permissions: permissions,
			Balances:    balanceSpot,
		},
		DataMargin: AccountMargin{
			BorrowEnabled:   marginAccount.BorrowEnabled,
			TradeEnabled:    marginAccount.TradeEnabled,
			TransferEnabled: marginAccount.TransferEnabled,
			Balances:        marginBalance,
		},
		Success: true,
	}
}

func GetSpotBalance(spotAccount *binance.Account) []Balance {
	var balances []Balance
	var amount decimal.Decimal
	var locked decimal.Decimal
	for i := range spotAccount.Balances {
		amount, _ = decimal.NewFromString(spotAccount.Balances[i].Free)
		locked, _ = decimal.NewFromString(spotAccount.Balances[i].Locked)
		if amount.Equal(decimal.NewFromInt(0)) && locked.Equal(decimal.NewFromInt(0)) {
			continue
		}
		balances = append(balances, Balance{
			Currency: spotAccount.Balances[i].Asset,
			Amount:   amount,
			Locked:   locked,
		})
	}
	return balances
}

func GetMarginBalance(marginBalance *binance.MarginAccount) []BalanceMargin {
	var balances []BalanceMargin
	var free decimal.Decimal
	var locked decimal.Decimal
	var borrowed decimal.Decimal
	var interest decimal.Decimal
	var netAsset decimal.Decimal
	for i := range marginBalance.UserAssets {
		borrowed, _ = decimal.NewFromString(marginBalance.UserAssets[i].Borrowed)
		free, _ = decimal.NewFromString(marginBalance.UserAssets[i].Free)
		locked, _ = decimal.NewFromString(marginBalance.UserAssets[i].Locked)
		interest, _ = decimal.NewFromString(marginBalance.UserAssets[i].Interest)
		netAsset, _ = decimal.NewFromString(marginBalance.UserAssets[i].NetAsset)
		if netAsset.Equal(decimal.NewFromInt(0)) {
			continue
		}
		balances = append(balances, BalanceMargin{
			Currency: marginBalance.UserAssets[i].Asset,
			Free:     free,
			Locked:   locked,
			Borrowed: borrowed,
			Interest: interest,
			NetAsset: netAsset,
		})
	}
	return balances
}
