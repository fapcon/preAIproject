package service

import (
	"context"

	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/storage"
)

type Ticker struct {
	storage storage.ExchangeTicker
	logger  *zap.Logger
}

func NewTicker(storage storage.ExchangeTicker, components *component.Components) ExchangeTicker {
	return &Ticker{
		storage: storage,
		logger:  components.Logger,
	}
}

func (p *Ticker) GetTicker(ctx context.Context) GetTickerOut {
	tickers, err := p.storage.GetTicker(ctx)
	if err != nil {
		return GetTickerOut{
			ErrorCode: errors.PlatformExchangeServiceGetTickerErr,
		}
	}

	return GetTickerOut{
		Data: tickers,
	}
}

func (p *Ticker) GetByID(ctx context.Context, tickerID int) (models.ExchangeTicker, error) {
	return p.storage.GetByID(ctx, tickerID)
}
func (p *Ticker) Save(ctx context.Context, tickers []models.ExchangeTicker) error {
	return p.storage.Save(ctx, tickers)
}

func (p *Ticker) GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeTicker, error) {
	return p.storage.GetList(ctx, condition)
}
