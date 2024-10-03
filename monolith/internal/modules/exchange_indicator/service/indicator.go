package service

import (
	"context"
	"github.com/adshao/go-binance/v2"
	"go.uber.org/zap"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/storage"
)

type Indicator struct {
	logger  *zap.Logger
	conf    config.AppConf
	storage storage.Indicatorer
	client  *binance.Client
}

var (
	apiKeyPlatform    = "5kOvWwvETSWX9fjxKC8vpjDWpUyn1vjzqtoefqljLatLvnGf4fMTe9ZHjqKSxRsz" // API key - предоставил Петр
	secretKeyPlatform = "et6TmjLwEkSKoo9RN5ajmhV4CGV0ccF47qqIZHXHPob9T3fwQE0sSv4LigK7vbFk" // Secret key - предоставил Петр
)

func NewIndicator(logger *zap.Logger, conf config.AppConf, storage storage.Indicatorer) Indicatorer {
	client := binance.NewClient(apiKeyPlatform, secretKeyPlatform)

	return &Indicator{logger: logger, conf: conf, storage: storage, client: client}
}

func (i *Indicator) EMA(symbol string, interval string, limit int, period1 int, period2 int, period3 int, ctx context.Context) EMA_out {
	klines, err := i.client.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit).Do(ctx)
	if err != nil {
		return EMA_out{ErrorCode: errors.GeneralError}
	}
	close := make([]float64, len(klines))
	for i := range klines {
		k, _ := strconv.ParseFloat(klines[i].Close, 64)
		close[i] = k
	}
	firstEMA := calculateEMA(close, period1)
	SecondEMA := calculateEMA(close, period2)
	ThirdEma := calculateEMA(close, period3)
	return EMA_out{
		FirstEMA:  firstEMA,
		SecondEMA: SecondEMA,
		ThirdEMA:  ThirdEma,
	}
}

func calculateEMA(close []float64, period int) []float64 {
	ema := make([]float64, len(close))

	// Рассчитываем первое значение EMA как простое среднее цены
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += close[i]
	}
	ema[period-1] = sum / float64(period)

	// Рассчитываем остальные значения EMA с использованием формулы
	multiplier := 2.0 / (float64(period) + 1)
	for i := period; i < len(close); i++ {
		ema[i] = (close[i]-ema[i-1])*multiplier + ema[i-1]
	}
	return ema
}

func (i *Indicator) GetDynamicPairBinance(ctx context.Context) DynamicPairOur {
	exchangeInfo, err := i.client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return DynamicPairOur{ErrorCode: errors.GeneralError}
	}
	symbols := make([]string, 0, len(exchangeInfo.Symbols))
	for _, symbol := range exchangeInfo.Symbols {
		symbols = append(symbols, symbol.Symbol)
	}
	return DynamicPairOur{Symbols: symbols}
}
