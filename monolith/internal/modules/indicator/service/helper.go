package service

import (
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"time"
)

func (*IndicatorsService) candlesToTimeSeries(candles []service.CandlesData) *techan.TimeSeries {
	series := techan.NewTimeSeries()

	for _, candleData := range candles {
		startTime := candleData.OpenTime
		endTime := candleData.CloseTime

		timePeriod := techan.NewTimePeriod(startTime, endTime.Sub(startTime))
		candle := techan.NewCandle(timePeriod)

		candle.OpenPrice = big.NewDecimal(candleData.Open.InexactFloat64())
		candle.ClosePrice = big.NewDecimal(candleData.Close.InexactFloat64())
		candle.MaxPrice = big.NewDecimal(candleData.High.InexactFloat64())
		candle.MinPrice = big.NewDecimal(candleData.Low.InexactFloat64())
		candle.Volume = big.NewDecimal(candleData.QuoteAssetVolume.InexactFloat64())

		series.AddCandle(candle)
	}

	return series
}

// indicatorToSeries - helper to calculate stochastic RSI
func (*IndicatorsService) indicatorToSeries(rsi techan.Indicator, window int) *techan.TimeSeries {
	series := techan.NewTimeSeries()

	for i := 0; i < window; i++ {
		period := techan.NewTimePeriod(time.Unix(int64(i), 0), time.Second)
		v := rsi.Calculate(i)
		candle := techan.NewCandle(period)
		candle.OpenPrice = v
		candle.ClosePrice = v
		candle.MaxPrice = v
		candle.MinPrice = v
		candle.Volume = v
		series.AddCandle(candle)
	}

	return series
}
