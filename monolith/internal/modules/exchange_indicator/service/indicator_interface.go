package service

import "context"

type Indicatorer interface {
	EMA(symbol string, interval string, limit int, period1 int, period2 int, period3 int, ctx context.Context) EMA_out
	GetDynamicPairBinance(ctx context.Context) DynamicPairOur
}

type EMA_out struct {
	ErrorCode int
	FirstEMA  []float64
	SecondEMA []float64
	ThirdEMA  []float64
}

type DynamicPairOur struct {
	ErrorCode int
	Symbols   []string
}
