package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"github.com/ptflp/gomapper"
	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type Ticker struct {
	adapter adapter.SQLAdapter
}

func NewTicker(adapter adapter.SQLAdapter) ExchangeTicker {
	return &Ticker{adapter: adapter}
}

func (t *Ticker) Save(ctx context.Context, tickers []models.ExchangeTicker) error {
	var dtos []utils.Tabler
	for i := range tickers {
		var dto models.ExchangeTickerDTO
		dto.SetCurrency(tickers[i].Pair).
			SetPrice(tickers[i].Price).
			SetUpdatedAt(time.Now()).
			SetExchangeID(tickers[i].ExchangeID).
			SetUpdatedAt(time.Now())
		dtos = append(dtos, utils.Tabler(&dto))
	}

	return t.adapter.Upsert(ctx, dtos)
}

func (t *Ticker) GetTicker(ctx context.Context) ([]models.ExchangeTicker, error) {
	var dtos []models.ExchangeTickerDTO
	err := t.adapter.List(ctx, &dtos, "exchange_ticker", utils.Condition{})
	if err != nil {
		return nil, err
	}
	var tickers []models.ExchangeTicker
	for i := range dtos {
		ticker := models.ExchangeTicker{
			ID:    dtos[i].GetID(),
			Pair:  dtos[i].GetCurrency(),
			Price: dtos[i].GetPrice(),
		}
		tickers = append(tickers, ticker)
	}

	return tickers, nil
}

func (t *Ticker) GetByID(ctx context.Context, tickerID int) (models.ExchangeTicker, error) {
	var tickers []models.ExchangeTickerDTO
	var ticker models.ExchangeTicker
	var dto models.ExchangeTickerDTO
	err := t.adapter.List(ctx, &tickers, dto.TableName(), utils.Condition{
		Equal: map[string]interface{}{
			"id": tickerID,
		},
	})
	if err != nil {
		return ticker, err
	}
	if len(tickers) < 1 {
		return ticker, fmt.Errorf("exchange ticker storage: GetTickerByID not found")
	}
	err = gomapper.MapStructs(&ticker, &tickers[0])
	if err != nil {
		return ticker, err
	}

	return ticker, nil
}

func (t *Ticker) GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeTicker, error) {
	var list []models.ExchangeTickerDTO
	var tickers []models.ExchangeTicker
	var dto models.ExchangeTickerDTO
	err := t.adapter.List(ctx, &list, dto.TableName(), condition)
	if err != nil {
		return nil, err
	}

	err = gomapper.MapStructs(&tickers, &list)
	if err != nil {
		return nil, err
	}
	return tickers, nil
}
