package service

import (
	"context"
	"github.com/shopspring/decimal"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules"
	cservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	iservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/service"
	lservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/service"
	oservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
	uservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service"
)

type ExchangeFacadeService struct {
	services modules.Services
	logger   *zap.Logger
}

func NewExchangeFacade(services modules.Services, components component.Components) ExchangeFacader {
	return ExchangeFacadeService{
		services: services,
		logger:   components.Logger,
	}
}

// Client
func (f ExchangeFacadeService) ClientBuyMarker(in cservice.MarketIn) cservice.OrderOut {
	return f.services.Client.BuyMarket(in)
}

func (f ExchangeFacadeService) ClientSellMarker(in cservice.MarketIn) cservice.OrderOut {
	return f.services.Client.SellMarket(in)
}

func (f ExchangeFacadeService) ClientOrderMarker(in cservice.OrderMarketIn) cservice.OrderOut {
	return f.services.Client.OrderMarket(in)
}

func (f ExchangeFacadeService) ClientOrderLimit(in cservice.OrderLimitIn) cservice.OrderOut {
	return f.services.Client.OrderLimit(in)
}

func (f ExchangeFacadeService) ClientBuyLimit(in cservice.LimitIn) cservice.OrderOut {
	return f.services.Client.BuyLimit(in)
}

func (f ExchangeFacadeService) ClientSellLimit(in cservice.LimitIn) cservice.OrderOut {
	return f.services.Client.SellLimit(in)
}

func (f ExchangeFacadeService) ClientCreateOrder(in cservice.CreateOrderIn) cservice.OrderOut {
	return f.services.Client.CreateOrder(in)
}

func (f ExchangeFacadeService) ClientCancelOrder(ctx context.Context, in cservice.CancelOrderIn) cservice.OrderOut {
	return f.services.Client.CancelOrder(ctx, in)
}

func (f ExchangeFacadeService) ClientGetOrder(ctx context.Context, in cservice.GetOrderIn) cservice.OrderOut {
	return f.services.Client.GetOrder(ctx, in)
}

func (f ExchangeFacadeService) ClientGetOrderHistory(in cservice.EIn) cservice.EOut {
	return f.services.Client.GetOrdersHistory(in)
}

func (f ExchangeFacadeService) ClientGetOpenOrders(in cservice.EIn) cservice.EOut {
	return f.services.Client.GetOpenOrders(in)
}

func (f ExchangeFacadeService) ClientGetBalances(ctx context.Context) cservice.GetAccountBalanceOut {
	return f.services.Client.GetBalances(ctx)
}

// Indicator
func (f ExchangeFacadeService) IndicatorEMA(ctx context.Context, symbol string, interval string, limit int, period1 int, period2 int, period3 int) iservice.EMA_out {
	return f.services.Indicator.EMA(symbol, interval, limit, period1, period2, period3, ctx)
}

func (f ExchangeFacadeService) IndicatorGetDynamicPairBinance(ctx context.Context) iservice.DynamicPairOur {
	return f.services.Indicator.GetDynamicPairBinance(ctx)
}

// ExchangeList
func (f ExchangeFacadeService) ListGetAll(ctx context.Context) lservice.ExchangeListOut {
	return f.services.ExchangeList.ExchangeListList(ctx)
}

func (f ExchangeFacadeService) ListDelete(ctx context.Context, exchangeListID int) error {
	return f.services.ExchangeList.ExchangeListDelete(ctx, exchangeListID)
}

func (f ExchangeFacadeService) ListAdd(ctx context.Context, in lservice.ExchangeAddIn) lservice.ExchangeOut {
	return f.services.ExchangeList.ExchangeListAdd(ctx, in)
}

// Order
func (f ExchangeFacadeService) OrderCancel(ctx context.Context, in oservice.OrderIn) (oservice.CancelOrderOut, error) {
	return f.services.Order.CancelOrder(ctx, in)
}

func (f ExchangeFacadeService) OrderWriteLog(ctx context.Context, in models.ExchangeOrderLogDTO) {
	f.services.Order.WriteOrderLog(ctx, in)
}

func (f ExchangeFacadeService) OrderWrite(ctx context.Context, in oservice.WriteOrderIn) error {
	return f.services.Order.WriteOrder(ctx, in)
}

func (f ExchangeFacadeService) OrderSellLimit(ctx context.Context, in oservice.OrderIn, quantity decimal.Decimal, price decimal.Decimal, unitedOrders int) {
	f.services.Order.OrderSellLimit(ctx, in, quantity, price, unitedOrders)
}

func (f ExchangeFacadeService) OrderGetCondition(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrder, error) {
	return f.services.Order.GetOrdersCondition(ctx, condition)
}

func (f ExchangeFacadeService) OrderGetStatistic(ctx context.Context, in oservice.GetBotRelationIn) oservice.StatisticOut {
	return f.services.Order.GetOrdersStatistic(ctx, in)
}

func (f ExchangeFacadeService) OrderAddOrdersStatistic(ctx context.Context, orders *[]models.ExchangeOrder) oservice.StatisticOut {
	return f.services.Order.AddOrdersStatistic(ctx, orders)
}

func (f ExchangeFacadeService) OrderGetBotOrders(ctx context.Context, in oservice.GetBotRelationIn) oservice.GetOrdersOut {
	return f.services.Order.GetBotOrders(ctx, in)
}

func (f ExchangeFacadeService) OrderGetUserOrders(ctx context.Context, in oservice.GetUserRelationIn) oservice.GetOrdersOut {
	return f.services.Order.GetUserOrders(ctx, in)
}

func (f ExchangeFacadeService) OrderExchangeList(ctx context.Context, in oservice.GetBotRelationIn) oservice.GetOrdersOut {
	return f.services.Order.ExchangeOrderList(ctx, in)
}

func (f ExchangeFacadeService) OrderGetAllStatistic(ctx context.Context, in oservice.GetUserRelationIn) oservice.StatisticOut {
	return f.services.Order.GetAllOrdersStatistic(ctx, in)
}

func (f ExchangeFacadeService) OrderCreate(ctx context.Context, in models.ExchangeOrderDTO) error {
	//f.services.Order.CreateOrder(ctx, in)
	return nil
}

func (f ExchangeFacadeService) OrderGetByUUID(ctx context.Context, uuid string) (models.ExchangeOrderDTO, error) {
	//f.services.Order.GetOrdersByUUID(ctx, uuid)
	return models.ExchangeOrderDTO{}, nil
}

func (f ExchangeFacadeService) OrderGetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error) {
	return f.services.Order.GetOrderList(ctx, condition)
}

func (f ExchangeFacadeService) OrderUpdate(ctx context.Context, dto models.ExchangeOrderDTO) error {
	return f.services.Order.UpdateOrder(ctx, dto)
}

// Ticker
func (f ExchangeFacadeService) TickerGetByID(ctx context.Context, id int) (models.ExchangeTicker, error) {
	return f.services.Ticker.GetByID(ctx, id)
}

func (f ExchangeFacadeService) TickerSave(ctx context.Context, tickers []models.ExchangeTicker) error {
	return f.services.Ticker.Save(ctx, tickers)
}

func (f ExchangeFacadeService) TickerGetlist(ctx context.Context, condition utils.Condition) ([]models.ExchangeTicker, error) {
	return f.services.Ticker.GetList(ctx, condition)
}

// User Key
func (f ExchangeFacadeService) UserKeyAdd(ctx context.Context, in uservice.ExchangeUserKeyAddIn) uservice.ExchangeOut {
	return f.services.UserKey.ExchangeUserKeyAdd(ctx, in)
}

func (f ExchangeFacadeService) UserKeyCheck(ctx context.Context, in uservice.ExchangeUserKeyAddIn) error {
	return f.services.UserKey.CheckKeys(ctx, in)
}

func (f ExchangeFacadeService) UserKeyDelete(ctx context.Context, exchangeUserKeyID int, userID int) error {
	return f.services.UserKey.ExchangeUserKeyDelete(ctx, exchangeUserKeyID, userID)
}

func (f ExchangeFacadeService) UserKeyList(ctx context.Context, in uservice.ExchangeUserListIn) uservice.ExchangeUserListOut {
	return f.services.UserKey.ExchangeUserKeyList(ctx, in)
}
