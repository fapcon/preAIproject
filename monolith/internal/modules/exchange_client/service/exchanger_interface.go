package service

import (
	"context"
	"github.com/hirokisan/bybit/v2"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

const (
	Binance = iota + 1
	Bybit
)

const (
	OrderTypeCancel = -1
)

const (
	SideBuy = iota + 1
	SideSell
	SideCancel
)

var OrderSides = map[int]string{
	SideBuy:    "Buy",
	SideSell:   "Sell",
	SideCancel: "Cancel",
}

func GetOrderTypeRaw(orderType int) string {
	return OrderTypesRaw[orderType]
}

func GetOrderType(orderTypeRaw string) int {
	return OrderTypes[orderTypeRaw]
}

var OrderTypes = map[string]int{
	TypeBuyMarketRaw:     TypeBuyMarket,
	TypeSellMarketRaw:    TypeSellMarket,
	TypeBuyLimitRaw:      TypeBuyLimit,
	TypeSellLimitRaw:     TypeSellLimit,
	TypeCancelAllRaw:     TypeCancelAll,
	TypeCancelBuyRaw:     TypeCancelBuy,
	TypeCancelSellRaw:    TypeCancelSell,
	TypeAverageRaw:       TypeAverage,
	TypeAutoSellLimitRaw: TypeAutoSellLimit,
	TypeGetOrderRaw:      TypeGetOrder,
}

var OrderTypesRaw = map[int]string{
	TypeBuyMarket:     TypeBuyMarketRaw,
	TypeSellMarket:    TypeSellMarketRaw,
	TypeBuyLimit:      TypeBuyLimitRaw,
	TypeSellLimit:     TypeSellLimitRaw,
	TypeCancelAll:     TypeCancelAllRaw,
	TypeCancelBuy:     TypeCancelBuyRaw,
	TypeCancelSell:    TypeCancelSellRaw,
	TypeAverage:       TypeAverageRaw,
	TypeAutoSellLimit: TypeAutoSellLimitRaw,
	TypeGetOrder:      TypeGetOrderRaw,
}

const (
	TypeBuyMarketRaw     = "BUYMARKET"
	TypeSellMarketRaw    = "SELLMARKET"
	TypeBuyLimitRaw      = "BUYLIMIT"
	TypeSellLimitRaw     = "SELLLIMIT"
	TypeCancelAllRaw     = "CANCELALL"
	TypeCancelBuyRaw     = "CANCELBUY"
	TypeCancelSellRaw    = "CANCELSELL"
	TypeAverageRaw       = "AVERAGE"
	TypeAutoSellLimitRaw = "AUTOSELLLIMIT"
	TypeGetOrderRaw      = "GETORDER"
)

const (
	TypeBuyMarket = iota + 1
	TypeSellMarket
	TypeBuyLimit
	TypeSellLimit
	TypeCancelAll
	TypeCancelBuy
	TypeCancelSell
	TypeAverage
	TypeAutoSellLimit
	TypeGetOrder
)

const (
	ExchangeOrderTypeLimit = iota + 1
	ExchangeOrderTypeMarket
	ExchangeOrderTypeLimitMaker
	ExchangeOrderTypeStopLoss
	ExchangeOrderTypeStopLossLimit
	ExchangeOrderTypeTakeProfit
	ExchangeOrderTypeTakeProfitLimit
)

const (
	OrderStatusUnknown = iota
	OrderStatusProcessing
	OrderStatusFilled
	OrderStatusFailed
	OrderStatusNew
	OrderStatusPartiallyFilled
	OrderStatusCanceled
	OrderStatusPendingCancel
	OrderStatusRejected
	OrderStatusExpired
)

var numberStatus = map[int]string{
	OrderStatusProcessing:      "Processing",
	OrderStatusFailed:          "Failed",
	OrderStatusFilled:          string(binance.OrderStatusTypeFilled),
	OrderStatusNew:             string(binance.OrderStatusTypeNew),
	OrderStatusPartiallyFilled: string(binance.OrderStatusTypePartiallyFilled),
	OrderStatusCanceled:        string(binance.OrderStatusTypeCanceled),
	OrderStatusPendingCancel:   string(binance.OrderStatusTypePendingCancel),
	OrderStatusRejected:        string(binance.OrderStatusTypeRejected),
	OrderStatusExpired:         string(binance.OrderStatusTypeExpired),
}

var getStatuses = map[binance.OrderStatusType]int{
	binance.OrderStatusTypeFilled:          OrderStatusFilled,
	binance.OrderStatusTypeNew:             OrderStatusNew,
	binance.OrderStatusTypePartiallyFilled: OrderStatusPartiallyFilled,
	binance.OrderStatusTypeCanceled:        OrderStatusCanceled,
	binance.OrderStatusTypePendingCancel:   OrderStatusPendingCancel,
	binance.OrderStatusTypeRejected:        OrderStatusRejected,
	binance.OrderStatusTypeExpired:         OrderStatusExpired,
}

var getOrderType = map[binance.OrderType]int{
	binance.OrderTypeLimit:           ExchangeOrderTypeLimit,
	binance.OrderTypeMarket:          ExchangeOrderTypeMarket,
	binance.OrderTypeLimitMaker:      ExchangeOrderTypeLimitMaker,
	binance.OrderTypeStopLoss:        ExchangeOrderTypeStopLoss,
	binance.OrderTypeStopLossLimit:   ExchangeOrderTypeStopLossLimit,
	binance.OrderTypeTakeProfit:      ExchangeOrderTypeTakeProfit,
	binance.OrderTypeTakeProfitLimit: ExchangeOrderTypeTakeProfitLimit,
}

func GetStatus(status int) string {
	return numberStatus[status]
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Exchanger
type Exchanger interface {
	OrderMarket(in OrderMarketIn) OrderOut
	BuyMarket(in MarketIn) OrderOut
	SellMarket(in MarketIn) OrderOut
	OrderLimit(in OrderLimitIn) OrderOut
	BuyLimit(in LimitIn) OrderOut
	SellLimit(in LimitIn) OrderOut
	CreateOrder(in CreateOrderIn) OrderOut
	CancelOrder(ctx context.Context, in CancelOrderIn) OrderOut
	GetOrder(ctx context.Context, in GetOrderIn) OrderOut
	GetOrdersHistory(in EIn) EOut
	GetOpenOrders(in EIn) EOut
	GetAccount(ctx context.Context) GetAccountOut
	GetBalances(ctx context.Context) GetAccountBalanceOut
	GetTicker(ctx context.Context, in GetTickerIn) GetTickerOut
	GetCandles(ctx context.Context, in GetCandlesIn) CandlesOut
}

var getOrderTypeBybit = map[bybit.OrderType]int{
	bybit.OrderTypeLimit:  ExchangeOrderTypeLimit,
	bybit.OrderTypeMarket: ExchangeOrderTypeMarket,
}

var numberStatusBybit = map[int]string{
	OrderStatusFilled:          string(bybit.OrderStatusFilled),
	OrderStatusNew:             string(bybit.OrderStatusNew),
	OrderStatusPartiallyFilled: string(bybit.OrderStatusPartiallyFilled),
	OrderStatusCanceled:        string(bybit.OrderStatusCancelled),
	OrderStatusPendingCancel:   string(bybit.OrderStatusPendingCancel),
	OrderStatusRejected:        string(bybit.OrderStatusRejected),
}

var getStatusesBybit = map[bybit.OrderStatus]int{
	bybit.OrderStatusFilled:          OrderStatusFilled,
	bybit.OrderStatusNew:             OrderStatusNew,
	bybit.OrderStatusPartiallyFilled: OrderStatusPartiallyFilled,
	bybit.OrderStatusCancelled:       OrderStatusCanceled,
	bybit.OrderStatusPendingCancel:   OrderStatusPendingCancel,
	bybit.OrderStatusRejected:        OrderStatusRejected,
}

var getOrderSide = map[bybit.Side]int{
	bybit.SideBuy:  SideBuy,
	bybit.SideSell: SideSell,
}

type CreateOrderIn struct {
	Type     int
	Side     int
	Pair     string
	Quantity decimal.Decimal
	Price    decimal.Decimal
}

type GetTickerIn struct {
	ExchangeID int
}

type GetTickerOut struct {
	ErrorCode int
	Data      map[string]decimal.Decimal
}

type OrderMarketIn struct {
	Side    int
	Request MarketIn
}

type MarketIn struct {
	Pair     string
	Quantity decimal.Decimal
}

type OrderLimitIn struct {
	Side    int
	Request LimitIn
}

type LimitIn struct {
	Pair     string
	Quantity decimal.Decimal
	Price    decimal.Decimal
}

type OrderOut struct {
	ClientOrderID     string
	OrderID           string
	Price             decimal.Decimal
	Amount            decimal.Decimal
	Quantity          decimal.Decimal
	Pair              string
	Status            int
	Side              int
	Type              int
	ExchangeOrderType int
	ErrorCode         int
	Message           string
}

type EIn struct {
}

type GetOrderIn struct {
	Pair               string
	OrderID            string
	priceChangePercent int
}

type CancelOrderIn struct {
	Pair               string
	OrderID            string
	priceChangePercent int
}

type EOut struct {
	ErrorCode int
}

type GetAccountOut struct {
	ErrorCode  int
	DataSpot   AccountSpot
	DataMargin AccountMargin
	Message    string
	Success    bool
}

type GetCandlesIn struct {
	Symbol    string        // Название торговой пары
	Interval  KlineInterval // Торговый диапазон
	Limit     int           // Размер возвращаемого слайса с необходимыми данными
	StartTime int64         // Начало времени отчета / optional
	EndTime   int64         // Конец времени отчета / optional
} // Необходимые данные для получения свеч

type KlineInterval string

const (
	OneMinute      KlineInterval = "1m"
	ThreeMinutes   KlineInterval = "3m"
	FiveMinutes    KlineInterval = "5m"
	FifteenMinutes KlineInterval = "15m"
	ThirtyMinutes  KlineInterval = "30m"
	OneHour        KlineInterval = "1h"
	TwoHours       KlineInterval = "2h"
	FourHours      KlineInterval = "4h"
	SixHours       KlineInterval = "6h"
	EightHours     KlineInterval = "8h"
	TwelveHours    KlineInterval = "12h"
	OneDay         KlineInterval = "1d"
	ThreeDays      KlineInterval = "3d"
	OneWeek        KlineInterval = "1w"
	OneMonth       KlineInterval = "1M"
) // Временные интервалы Binance являются основными. Интервалы для других бирж отображаются на основании этих значений.

var bybitIntervals = map[KlineInterval]bybit.Interval{
	OneMinute:      "1",
	ThreeMinutes:   "3",
	FiveMinutes:    "5",
	FifteenMinutes: "15",
	ThirtyMinutes:  "30",
	OneHour:        "60",
	TwoHours:       "120",
	FourHours:      "240",
	SixHours:       "360",
	TwelveHours:    "720",
	OneDay:         "D",
	OneWeek:        "W",
	OneMonth:       "M",
}

type CandlesData struct {
	OpenTime                 time.Time       // Время открытия свечи
	Open                     decimal.Decimal // Цена открытия свечи
	High                     decimal.Decimal // Наивысшая цена, достигнутая свечей в течение периода
	Low                      decimal.Decimal // Наименьшая цена, достигнутая свечей в течение периода
	Close                    decimal.Decimal // Цена закрытия свечи
	Volume                   decimal.Decimal // Объем торгов в базовой валюте
	CloseTime                time.Time       // Время закрытия свечи
	QuoteAssetVolume         decimal.Decimal // Объем торгов в котируемой валюте
	TradeNum                 int64           // Количество сделок
	TakerBuyBaseAssetVolume  decimal.Decimal // Объем базовой валюты, купленной с помощью ордеров по рыночной цене во время свечи.
	TakerBuyQuoteAssetVolume decimal.Decimal // Объем котируемой валюты, потраченной на покупку базовой валюты
	// с помощью ордеров по рыночной цене во время свечи.
} // Данные которые возвращают свечи

type CandlesOut struct {
	ErrorCode int
	ErrorMsg  string
	Candles   []CandlesData
}

const (
	PermissionSPOTRaw = "SPOT"
)

const (
	PermissionUnknown = iota
	PermissionSPOT
)

type AccountSpot struct {
	CanTrade    bool
	CanDeposit  bool
	CanWithdraw bool
	Permissions []int
	Balances    []Balance
}

type AccountMargin struct {
	BorrowEnabled   bool //Флаг, указывающий, разрешено ли займы на маржу.
	TradeEnabled    bool //Флаг, указывающий, разрешена ли торговля на маржине.
	TransferEnabled bool //Флаг, указывающий, разрешены ли переводы между основным и маржинальным кошельками.
	Balances        []BalanceMargin
}

type Balance struct {
	Currency string
	Amount   decimal.Decimal
	Locked   decimal.Decimal
}

// Информация о активнах на маржинаольном кошельке
type BalanceMargin struct {
	Currency string          //Символ актива (например, BTC, ETH, BNB).
	Borrowed decimal.Decimal //Количество актива, взятого взаймы.
	Free     decimal.Decimal //Количество доступного актива.
	Interest decimal.Decimal //Количество начисленных процентов по активу.
	Locked   decimal.Decimal //Количество заблокированного актива.
	NetAsset decimal.Decimal //Общее значение актива (общий баланс).
}

type GetAccountBalanceOut struct {
	ErrorCode         int
	Success           bool
	DataSpotBalance   []models.BalanceDTO
	DataMarginBalance []models.BalanceDTO
}
