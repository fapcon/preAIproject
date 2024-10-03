package service

type IndicatorError string

const (
	ErrCandlesMsg         IndicatorError = "error fetching candles"
	ErrWrongPricesTypeMsg IndicatorError = "wrong prices type used"
	ErrCandlesCountMsg    IndicatorError = "candles count mismatch"
)

const (
	ErrCandlesCode = iota + 1
	ErrWrongPricesTypeCode
	ErrCandlesCountCode
)
