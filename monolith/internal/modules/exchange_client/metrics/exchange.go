package metrics

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/metrics"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"time"
)

type ExchangeAPIProxy struct {
	exchange service.Exchanger
	meter    metrics.MetricMeter
}

func (e *ExchangeAPIProxy) BuyMarket(in service.MarketIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.BuyMarket(in)
	e.meter.TimeCounting("Exchanger", "BuyMarket", startTime)
	e.meter.CountRequest("BuyMarket")
	return result
}

func (e *ExchangeAPIProxy) SellMarket(in service.MarketIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.SellMarket(in)
	e.meter.TimeCounting("Exchanger", "SellMarket", startTime)
	e.meter.CountRequest("SellMarket")
	return result
}

func (e *ExchangeAPIProxy) OrderLimit(in service.OrderLimitIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.OrderLimit(in)
	e.meter.TimeCounting("Exchanger", "OrderLimit", startTime)
	e.meter.CountRequest("OrderLimit")
	return result
}

func (e *ExchangeAPIProxy) BuyLimit(in service.LimitIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.BuyLimit(in)
	e.meter.TimeCounting("Exchanger", "BuyLimit", startTime)
	e.meter.CountRequest("BuyLimit")
	return result
}

func (e *ExchangeAPIProxy) SellLimit(in service.LimitIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.SellLimit(in)
	e.meter.TimeCounting("Exchanger", "SellLimit", startTime)
	e.meter.CountRequest("SellLimit")
	return result
}

func (e *ExchangeAPIProxy) CreateOrder(in service.CreateOrderIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.CreateOrder(in)
	e.meter.TimeCounting("Exchanger", "CreateOrder", startTime)
	e.meter.CountRequest("CreateOrder")
	return result
}

func (e *ExchangeAPIProxy) CancelOrder(ctx context.Context, in service.CancelOrderIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.CancelOrder(ctx, in)
	e.meter.TimeCounting("Exchanger", "CancelOrder", startTime)
	e.meter.CountRequest("CancelOrder")
	return result
}

func (e *ExchangeAPIProxy) GetOrder(ctx context.Context, in service.GetOrderIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.GetOrder(ctx, in)
	e.meter.TimeCounting("Exchanger", "GetOrder", startTime)
	e.meter.CountRequest("GetOrder")
	return result
}

func (e *ExchangeAPIProxy) GetOrdersHistory(in service.EIn) service.EOut {
	startTime := time.Now()
	result := e.exchange.GetOrdersHistory(in)
	e.meter.TimeCounting("Exchanger", "GetOrdersHistory", startTime)
	e.meter.CountRequest("GetOrdersHistory")
	return result
}

func (e *ExchangeAPIProxy) GetOpenOrders(in service.EIn) service.EOut {
	startTime := time.Now()
	result := e.exchange.GetOpenOrders(in)
	e.meter.TimeCounting("Exchanger", "GetOpenOrders", startTime)
	e.meter.CountRequest("GetOpenOrders")
	return result
}

func (e *ExchangeAPIProxy) GetAccount(ctx context.Context) service.GetAccountOut {
	startTime := time.Now()
	result := e.exchange.GetAccount(ctx)
	e.meter.TimeCounting("Exchanger", "GetAccount", startTime)
	e.meter.CountRequest("GetAccount")
	return result
}

func (e *ExchangeAPIProxy) GetBalances(ctx context.Context) service.GetAccountBalanceOut {
	startTime := time.Now()
	result := e.exchange.GetBalances(ctx)
	e.meter.TimeCounting("Exchanger", "GetBalances", startTime)
	e.meter.CountRequest("GetBalances")
	return result
}

func (e *ExchangeAPIProxy) GetTicker(ctx context.Context, in service.GetTickerIn) service.GetTickerOut {
	startTime := time.Now()
	result := e.exchange.GetTicker(ctx, in)
	e.meter.TimeCounting("Exchanger", "GetTicker", startTime)
	e.meter.CountRequest("GetTicker")
	return result
}

func (e *ExchangeAPIProxy) GetCandles(ctx context.Context, in service.GetCandlesIn) service.CandlesOut {
	startTime := time.Now()
	result := e.exchange.GetCandles(ctx, in)
	e.meter.TimeCounting("Exchanger", "GetCandles", startTime)
	e.meter.CountRequest("GetCandles")
	return result
}

func NewExchangeAPIProxy(exchanger service.Exchanger, metrics metrics.MetricMeter) *ExchangeAPIProxy {
	return &ExchangeAPIProxy{
		exchange: exchanger,
		meter:    metrics,
	}
}

func (e *ExchangeAPIProxy) OrderMarket(in service.OrderMarketIn) service.OrderOut {
	startTime := time.Now()
	result := e.exchange.OrderMarket(in)
	e.meter.TimeCounting("Exchanger", "OrderMarket", startTime)
	e.meter.CountRequest("OrderMarket")
	return result
}
