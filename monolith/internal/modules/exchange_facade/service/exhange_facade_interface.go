package service

import (
	"context"
	"github.com/shopspring/decimal"
	"gitlab.com/golight/orm/utils"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	iservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/service"
	lservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/service"
	oservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
)

type ExchangeFacader interface {
	ClientBuyMarker(in service.MarketIn) service.OrderOut
	ClientSellMarker(in service.MarketIn) service.OrderOut
	ClientOrderMarker(in service.OrderMarketIn) service.OrderOut
	ClientOrderLimit(in service.OrderLimitIn) service.OrderOut
	ClientBuyLimit(in service.LimitIn) service.OrderOut
	ClientSellLimit(in service.LimitIn) service.OrderOut
	ClientCreateOrder(in service.CreateOrderIn) service.OrderOut
	ClientCancelOrder(ctx context.Context, in service.CancelOrderIn) service.OrderOut
	ClientGetOrder(ctx context.Context, in service.GetOrderIn) service.OrderOut
	ClientGetOrderHistory(in service.EIn) service.EOut
	ClientGetOpenOrders(in service.EIn) service.EOut
	ClientGetBalances(ctx context.Context) service.GetAccountBalanceOut

	IndicatorEMA(ctx context.Context, symbol string, interval string, limit int, period1 int, period2 int, period3 int) iservice.EMA_out
	IndicatorGetDynamicPairBinance(ctx context.Context) iservice.DynamicPairOur

	ListGetAll(ctx context.Context) lservice.ExchangeListOut
	ListDelete(ctx context.Context, exchangeListID int) error
	ListAdd(ctx context.Context, in lservice.ExchangeAddIn) lservice.ExchangeOut

	OrderCancel(ctx context.Context, in oservice.OrderIn) (oservice.CancelOrderOut, error)
	OrderWriteLog(ctx context.Context, in models.ExchangeOrderLogDTO)
	OrderWrite(ctx context.Context, in oservice.WriteOrderIn) error
	OrderSellLimit(ctx context.Context, in oservice.OrderIn, quantity decimal.Decimal, price decimal.Decimal, unitedOrders int)
	OrderGetCondition(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrder, error)
	OrderGetStatistic(ctx context.Context, in oservice.GetBotRelationIn) oservice.StatisticOut
	OrderAddOrdersStatistic(ctx context.Context, orders *[]models.ExchangeOrder) oservice.StatisticOut
	OrderGetBotOrders(ctx context.Context, in oservice.GetBotRelationIn) oservice.GetOrdersOut
	OrderGetUserOrders(ctx context.Context, in oservice.GetUserRelationIn) oservice.GetOrdersOut
	OrderExchangeList(ctx context.Context, in oservice.GetBotRelationIn) oservice.GetOrdersOut
	OrderGetAllStatistic(ctx context.Context, in oservice.GetUserRelationIn) oservice.StatisticOut
	OrderCreate(ctx context.Context, in models.ExchangeOrderDTO) error
	OrderGetByUUID(ctx context.Context, uuid string) (models.ExchangeOrderDTO, error)
	OrderGetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error)
	OrderUpdate(ctx context.Context, dto models.ExchangeOrderDTO) error

	TickerGetByID(ctx context.Context, id int) (models.ExchangeTicker, error)
	TickerSave(ctx context.Context, tickers []models.ExchangeTicker) error
	TickerGetlist(ctx context.Context, condition utils.Condition) ([]models.ExchangeTicker, error)
}
