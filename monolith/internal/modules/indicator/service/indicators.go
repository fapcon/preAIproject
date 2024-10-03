package service

import (
	"context"
	"fmt"
	"github.com/sdcoffey/techan"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
)

type IndicatorsService struct {
	exchange service.Exchanger
	logger   *zap.Logger
}

func (i *IndicatorsService) CalculateRSI(ctx context.Context, in RSIRequest) RSIResponse {
	candles := i.exchange.GetCandles(ctx, service.GetCandlesIn{
		Symbol:   in.Symbol,
		Interval: in.Timeframe,
		Limit:    in.Window + 1,
	})

	if candles.ErrorCode != 0 {
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrCandlesMsg, candles.ErrorMsg))
		return RSIResponse{
			ErrorCode: ErrCandlesCode,
			ErrorMsg:  ErrCandlesMsg,
		}
	}

	if len(candles.Candles) != in.Window+1 {
		i.logger.Warn(fmt.Sprintf("%s for %s: wanted: %d, got: %d", ErrCandlesCountMsg, in.Symbol, in.Window+1, len(candles.Candles)))
		return RSIResponse{
			ErrorCode: ErrCandlesCountCode,
			ErrorMsg:  ErrCandlesCountMsg,
		}
	}

	series := i.candlesToTimeSeries(candles.Candles)
	var pricesIndicator techan.Indicator
	switch in.PriceType {
	case ClosePrices:
		pricesIndicator = techan.NewClosePriceIndicator(series)
	case OpenPrices:
		pricesIndicator = techan.NewOpenPriceIndicator(series)
	case HighPrices:
		pricesIndicator = techan.NewHighPriceIndicator(series)
	case LowPrices:
		pricesIndicator = techan.NewLowPriceIndicator(series)
	case TypicalPrices:
		pricesIndicator = techan.NewTypicalPriceIndicator(series)
	default:
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType))
		return RSIResponse{
			ErrorCode: ErrWrongPricesTypeCode,
			ErrorMsg:  IndicatorError(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType)),
		}
	}

	rsiValue := techan.NewRelativeStrengthIndexIndicator(pricesIndicator, in.Window)

	return RSIResponse{
		ErrorCode: 0,
		ErrorMsg:  "OK",
		Value:     decimal.NewFromFloat(rsiValue.Calculate(series.LastIndex()).Float()),
	}
}

func (i *IndicatorsService) CalculateStochasticRSI(ctx context.Context, in StochasticRSIRequest) StochasticRSIResponse {
	candles := i.exchange.GetCandles(ctx, service.GetCandlesIn{
		Symbol:   in.Symbol,
		Interval: in.Timeframe,
		Limit:    in.RSIWindow + 1,
	})

	if candles.ErrorCode != 0 {
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrCandlesMsg, candles.ErrorMsg))
		return StochasticRSIResponse{
			ErrorCode: ErrCandlesCode,
			ErrorMsg:  ErrCandlesMsg,
		}
	}

	if len(candles.Candles) != in.RSIWindow+1 {
		i.logger.Warn(fmt.Sprintf("%s for %s: wanted: %d, got: %d", ErrCandlesCountMsg, in.Symbol, in.RSIWindow+1, len(candles.Candles)))
		return StochasticRSIResponse{
			ErrorCode: ErrCandlesCountCode,
			ErrorMsg:  ErrCandlesCountMsg,
		}
	}

	series := i.candlesToTimeSeries(candles.Candles)
	var pricesIndicator techan.Indicator
	switch in.PriceType {
	case ClosePrices:
		pricesIndicator = techan.NewClosePriceIndicator(series)
	case OpenPrices:
		pricesIndicator = techan.NewOpenPriceIndicator(series)
	case HighPrices:
		pricesIndicator = techan.NewHighPriceIndicator(series)
	case LowPrices:
		pricesIndicator = techan.NewLowPriceIndicator(series)
	case TypicalPrices:
		pricesIndicator = techan.NewTypicalPriceIndicator(series)
	default:
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType))
		return StochasticRSIResponse{
			ErrorCode: ErrWrongPricesTypeCode,
			ErrorMsg:  IndicatorError(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType)),
		}
	}

	rsiValue := techan.NewRelativeStrengthIndexIndicator(pricesIndicator, in.RSIWindow)
	a := i.indicatorToSeries(rsiValue, in.RSIWindow+1) // indicator type converts to techan.TimeSeries
	stochasticRSIKValue := techan.NewFastStochasticIndicator(a, in.KWindow)
	stochasticRSIDValue := techan.NewSlowStochasticIndicator(stochasticRSIKValue, in.DWindow)

	return StochasticRSIResponse{
		ErrorCode: 0,
		ErrorMsg:  "OK",
		KValue:    decimal.NewFromFloat(stochasticRSIKValue.Calculate(series.LastIndex()).Float()),
		DValue:    decimal.NewFromFloat(stochasticRSIDValue.Calculate(series.LastIndex()).Float()),
	}
}

func (i *IndicatorsService) CalculateBollingerBands(ctx context.Context, in BollingerBandsRequest) BollingerBandsResponse {
	candles := i.exchange.GetCandles(ctx, service.GetCandlesIn{
		Symbol:   in.Symbol,
		Interval: in.Timeframe,
		Limit:    in.Window + 1,
	})

	if candles.ErrorCode != 0 {
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrCandlesMsg, candles.ErrorMsg))
		return BollingerBandsResponse{
			ErrorCode: ErrCandlesCode,
			ErrorMsg:  ErrCandlesMsg,
		}
	}

	if len(candles.Candles) != in.Window+1 {
		i.logger.Warn(fmt.Sprintf("%s for %s: wanted: %d, got: %d", ErrCandlesCountMsg, in.Symbol, in.Window+1, len(candles.Candles)))
		return BollingerBandsResponse{
			ErrorCode: ErrCandlesCountCode,
			ErrorMsg:  ErrCandlesCountMsg,
		}
	}

	series := i.candlesToTimeSeries(candles.Candles)
	var pricesIndicator techan.Indicator
	switch in.PriceType {
	case ClosePrices:
		pricesIndicator = techan.NewClosePriceIndicator(series)
	case OpenPrices:
		pricesIndicator = techan.NewOpenPriceIndicator(series)
	case HighPrices:
		pricesIndicator = techan.NewHighPriceIndicator(series)
	case LowPrices:
		pricesIndicator = techan.NewLowPriceIndicator(series)
	case TypicalPrices:
		pricesIndicator = techan.NewTypicalPriceIndicator(series)
	default:
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType))
		return BollingerBandsResponse{
			ErrorCode: ErrWrongPricesTypeCode,
			ErrorMsg:  IndicatorError(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType)),
		}
	}

	bollingerBandsUpperIndicator := techan.NewBollingerUpperBandIndicator(pricesIndicator, in.Window, in.Sigma)
	upperValue := decimal.NewFromFloat(bollingerBandsUpperIndicator.Calculate(series.LastIndex()).Float())
	bollingerBandsLowerIndicator := techan.NewBollingerLowerBandIndicator(pricesIndicator, in.Window, in.Sigma)
	lowerValue := decimal.NewFromFloat(bollingerBandsLowerIndicator.Calculate(series.LastIndex()).Float())

	return BollingerBandsResponse{
		ErrorCode: 0,
		ErrorMsg:  "OK",
		LowValue:  lowerValue,
		HighValue: upperValue,
	}
}

func (i *IndicatorsService) CalculateEMA(ctx context.Context, in EMARequest) EMAResponse {
	candles := i.exchange.GetCandles(ctx, service.GetCandlesIn{
		Symbol:   in.Symbol,
		Interval: in.Timeframe,
		Limit:    in.Window + 1,
	})

	if candles.ErrorCode != 0 {
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrCandlesMsg, candles.ErrorMsg))
		return EMAResponse{
			ErrorCode: ErrCandlesCode,
			ErrorMsg:  ErrCandlesMsg,
		}
	}

	if len(candles.Candles) != in.Window+1 {
		i.logger.Warn(fmt.Sprintf("%s for %s: wanted: %d, got: %d", ErrCandlesCountMsg, in.Symbol, in.Window+1, len(candles.Candles)))
		return EMAResponse{
			ErrorCode: ErrCandlesCountCode,
			ErrorMsg:  ErrCandlesCountMsg,
		}
	}

	series := i.candlesToTimeSeries(candles.Candles)
	var pricesIndicator techan.Indicator
	switch in.PriceType {
	case ClosePrices:
		pricesIndicator = techan.NewClosePriceIndicator(series)
	case OpenPrices:
		pricesIndicator = techan.NewOpenPriceIndicator(series)
	case HighPrices:
		pricesIndicator = techan.NewHighPriceIndicator(series)
	case LowPrices:
		pricesIndicator = techan.NewLowPriceIndicator(series)
	case TypicalPrices:
		pricesIndicator = techan.NewTypicalPriceIndicator(series)
	default:
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType))
		return EMAResponse{
			ErrorCode: ErrWrongPricesTypeCode,
			ErrorMsg:  IndicatorError(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType)),
		}
	}

	emaValue := techan.NewEMAIndicator(pricesIndicator, in.Window)

	return EMAResponse{
		ErrorCode: 0,
		ErrorMsg:  "OK",
		Value:     decimal.NewFromFloat(emaValue.Calculate(series.LastIndex()).Float()),
	}
}

func (i *IndicatorsService) CalculateSMA(ctx context.Context, in SMARequest) SMAResponse {
	candles := i.exchange.GetCandles(ctx, service.GetCandlesIn{
		Symbol:   in.Symbol,
		Interval: in.Timeframe,
		Limit:    in.Window + 1,
	})

	if candles.ErrorCode != 0 {
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrCandlesMsg, candles.ErrorMsg))
		return SMAResponse{
			ErrorCode: ErrCandlesCode,
			ErrorMsg:  ErrCandlesMsg,
		}
	}

	if len(candles.Candles) != in.Window+1 {
		i.logger.Warn(fmt.Sprintf("%s for %s: wanted: %d, got: %d", ErrCandlesCountMsg, in.Symbol, in.Window+1, len(candles.Candles)))
		return SMAResponse{
			ErrorCode: ErrCandlesCountCode,
			ErrorMsg:  ErrCandlesCountMsg,
		}
	}

	series := i.candlesToTimeSeries(candles.Candles)
	var pricesIndicator techan.Indicator
	switch in.PricesType {
	case ClosePrices:
		pricesIndicator = techan.NewClosePriceIndicator(series)
	case OpenPrices:
		pricesIndicator = techan.NewOpenPriceIndicator(series)
	case HighPrices:
		pricesIndicator = techan.NewHighPriceIndicator(series)
	case LowPrices:
		pricesIndicator = techan.NewLowPriceIndicator(series)
	case TypicalPrices:
		pricesIndicator = techan.NewTypicalPriceIndicator(series)
	default:
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PricesType))
		return SMAResponse{
			ErrorCode: ErrWrongPricesTypeCode,
			ErrorMsg:  IndicatorError(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PricesType)),
		}
	}

	smaValue := techan.NewSimpleMovingAverage(pricesIndicator, in.Window)

	return SMAResponse{
		ErrorCode: 0,
		ErrorMsg:  "OK",
		Value:     decimal.NewFromFloat(smaValue.Calculate(series.LastIndex()).Float()),
	}
}

func (i *IndicatorsService) CalculateCCI(ctx context.Context, in CCIRequest) CCIResponse {
	candles := i.exchange.GetCandles(ctx, service.GetCandlesIn{
		Symbol:   in.Symbol,
		Interval: in.Timeframe,
		Limit:    in.Window + 1,
	})

	if candles.ErrorCode != 0 {
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrCandlesMsg, candles.ErrorMsg))
		return CCIResponse{
			ErrorCode: ErrCandlesCode,
			ErrorMsg:  ErrCandlesMsg,
		}
	}

	if len(candles.Candles) != in.Window+1 {
		i.logger.Warn(fmt.Sprintf("%s for %s: wanted: %d, got: %d", ErrCandlesCountMsg, in.Symbol, in.Window+1, len(candles.Candles)))
		return CCIResponse{
			ErrorCode: ErrCandlesCountCode,
			ErrorMsg:  ErrCandlesCountMsg,
		}
	}

	series := i.candlesToTimeSeries(candles.Candles)
	cciValue := techan.NewCCIIndicator(series, in.Window)

	return CCIResponse{
		ErrorCode: 0,
		ErrorMsg:  "OK",
		Value:     decimal.NewFromFloat(cciValue.Calculate(series.LastIndex()).Float()),
	}
}

func (i *IndicatorsService) CalculateMACD(ctx context.Context, in MACDRequest) MACDResponse {
	candles := i.exchange.GetCandles(ctx, service.GetCandlesIn{
		Symbol:   in.Symbol,
		Interval: in.Timeframe,
		Limit:    in.LongWindow*2 + 1,
	})

	if candles.ErrorCode != 0 {
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrCandlesMsg, candles.ErrorMsg))
		return MACDResponse{
			ErrorCode: ErrCandlesCode,
			ErrorMsg:  ErrCandlesMsg,
		}
	}

	if len(candles.Candles) != in.LongWindow+1 {
		i.logger.Warn(fmt.Sprintf("%s for %s: wanted: %d, got: %d", ErrCandlesCountMsg, in.Symbol, in.LongWindow+1, len(candles.Candles)))
		return MACDResponse{
			ErrorCode: ErrCandlesCountCode,
			ErrorMsg:  ErrCandlesCountMsg,
		}
	}

	series := i.candlesToTimeSeries(candles.Candles)
	var pricesIndicator techan.Indicator
	switch in.PriceType {
	case ClosePrices:
		pricesIndicator = techan.NewClosePriceIndicator(series)
	case OpenPrices:
		pricesIndicator = techan.NewOpenPriceIndicator(series)
	case HighPrices:
		pricesIndicator = techan.NewHighPriceIndicator(series)
	case LowPrices:
		pricesIndicator = techan.NewLowPriceIndicator(series)
	case TypicalPrices:
		pricesIndicator = techan.NewTypicalPriceIndicator(series)
	default:
		i.logger.Warn(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType))
		return MACDResponse{
			ErrorCode: ErrWrongPricesTypeCode,
			ErrorMsg:  IndicatorError(fmt.Sprintf("%s: %s", ErrWrongPricesTypeMsg, in.PriceType)),
		}
	}

	macdIndicator := techan.NewMACDIndicator(pricesIndicator, in.ShortWindow, in.LongWindow)
	macdHistIndicator := techan.NewMACDHistogramIndicator(macdIndicator, in.SignalLineWindow)

	return MACDResponse{
		ErrorCode: 0,
		ErrorMsg:  "OK",
		Value:     decimal.NewFromFloat(macdHistIndicator.Calculate(series.LastIndex()).Float()),
		Histogram: decimal.NewFromFloat(macdHistIndicator.Calculate(series.LastIndex()).Float()),
	}
}
