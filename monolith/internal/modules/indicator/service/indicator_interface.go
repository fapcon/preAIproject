package service

import (
	"context"
	"github.com/shopspring/decimal"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
)

type IndicatorName string

type PriceType string

const (
	ClosePrices   PriceType = "close"
	OpenPrices    PriceType = "open"
	HighPrices    PriceType = "high"
	LowPrices     PriceType = "low"
	TypicalPrices PriceType = "typical"
)

type IndicatorServicer interface {
	CalculateRSI(ctx context.Context, in RSIRequest) RSIResponse
	CalculateStochasticRSI(ctx context.Context, in StochasticRSIRequest) StochasticRSIResponse
	CalculateBollingerBands(ctx context.Context, in BollingerBandsRequest) BollingerBandsResponse
	CalculateEMA(ctx context.Context, in EMARequest) EMAResponse
	CalculateSMA(ctx context.Context, in SMARequest) SMAResponse
	CalculateCCI(ctx context.Context, in CCIRequest) CCIResponse
	CalculateMACD(ctx context.Context, in MACDRequest) MACDResponse
}

type MarketType string

const (
	SpotMarket     MarketType = "SPOT"
	ContractMarket MarketType = "CONTRACT"
)

type RSIRequest struct {
	Symbol     string // ex: "BTCUSDT"
	MarketType MarketType
	ExchangeID int
	Window     int
	Timeframe  service.KlineInterval
	PriceType  PriceType
}

type RSIResponse struct {
	ErrorCode int
	ErrorMsg  IndicatorError
	Value     decimal.Decimal
}

type StochasticRSIRequest struct {
	Symbol     string // ex: "BTCUSDT"
	MarketType MarketType
	ExchangeID int
	RSIWindow  int
	KWindow    int
	DWindow    int
	Timeframe  service.KlineInterval
	PriceType  PriceType
}

type StochasticRSIResponse struct {
	ErrorCode int
	ErrorMsg  IndicatorError
	KValue    decimal.Decimal
	DValue    decimal.Decimal
}

type BollingerBandsRequest struct {
	Symbol     string // ex: "BTCUSDT"
	MarketType MarketType
	ExchangeID int
	Window     int
	Sigma      float64
	Timeframe  service.KlineInterval
	PriceType  PriceType
}

type BollingerBandsResponse struct {
	ErrorCode int
	ErrorMsg  IndicatorError
	LowValue  decimal.Decimal
	HighValue decimal.Decimal
}

type EMARequest struct {
	Symbol     string // ex: "BTCUSDT"
	MarketType MarketType
	ExchangeID int
	Window     int
	Timeframe  service.KlineInterval
	PriceType  PriceType
}

type EMAResponse struct {
	ErrorCode int
	ErrorMsg  IndicatorError
	Value     decimal.Decimal
}

type SMARequest struct {
	Symbol     string // ex: "BTCUSDT"
	MarketType MarketType
	ExchangeID int
	Window     int
	Timeframe  service.KlineInterval
	PricesType PriceType
}

type SMAResponse struct {
	ErrorCode int
	ErrorMsg  IndicatorError
	Value     decimal.Decimal
}

type CCIRequest struct {
	Symbol     string // ex: "BTCUSDT"
	MarketType MarketType
	ExchangeID int
	Window     int
	Timeframe  service.KlineInterval
}

type CCIResponse struct {
	ErrorCode int
	ErrorMsg  IndicatorError
	Value     decimal.Decimal
}

type MACDRequest struct {
	Symbol           string // ex: "BTCUSDT"
	MarketType       MarketType
	ExchangeID       int
	LongWindow       int
	SignalLineWindow int
	ShortWindow      int
	Timeframe        service.KlineInterval
	PriceType        PriceType
}

type MACDResponse struct {
	ErrorCode int
	ErrorMsg  IndicatorError
	Value     decimal.Decimal
	Histogram decimal.Decimal
}
