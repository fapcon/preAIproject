package worker

import (
	"context"
	"github.com/shopspring/decimal"
	"time"

	"golang.org/x/sync/errgroup"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	tservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/service"
)

//const TELEGRAM_BOT_TOKEN = "6589663671:AAEzMKH0KWonpvJo5fv42MbDZnUkfKbLuEk"
//const TRUSTED_CHAN_ID = "-1001831256786"

type Ticker struct {
	ex               service.Exchanger
	ticker           tservice.ExchangeTicker
	ExchangeID       int
	notificationChan chan models.NotificationMessage
}

func NewTicker(ex service.Exchanger, services modules.Services, exchangeID int, notificationChan chan models.NotificationMessage) *Ticker {
	return &Ticker{ex: ex, ticker: services.Ticker, ExchangeID: exchangeID, notificationChan: notificationChan}
}

func (t *Ticker) Run() error {
	var errGroup errgroup.Group
	errGroup.Go(t.work)

	return errGroup.Wait()
}

func (t *Ticker) work() error {
	timeTicker := time.NewTicker(10 * time.Second)
	oldPrices := make(map[string]decimal.Decimal)

	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			out := t.ex.GetTicker(context.Background(), service.GetTickerIn{ExchangeID: 1})
			if out.ErrorCode != errors.NoError {
				continue
			}
			var tickers []models.ExchangeTicker
			for currency, price := range out.Data {
				if price.IsZero() {
					continue
				}
				ticker := models.ExchangeTicker{
					Pair:       currency,
					Price:      price,
					ExchangeID: t.ExchangeID,
				}
				err := t.NotificationTelegram(&oldPrices, ticker)
				if err != nil {
					return err
				}
				tickers = append(tickers, ticker)
			}
			err := t.ticker.Save(context.Background(), tickers)
			if err != nil {
				return err
			}
		}
	}
}
func (t *Ticker) NotificationTelegram(oldPrices *map[string]decimal.Decimal, ticker models.ExchangeTicker) error {
	if oldPrice, ok := (*oldPrices)[ticker.Pair]; ok {
		if !oldPrice.IsZero() {
			percent := oldPrice.Mul(decimal.NewFromFloat(0.05))
			difference := ticker.Price.Sub(oldPrice)
			res := percent.Cmp(difference.Abs())
			if res < 1 {
				t.notificationChan <- models.NotificationMessage{
					Pair:     ticker.Pair,
					Price:    ticker.Price,
					OldPrice: (*oldPrices)[ticker.Pair],
					Time:     time.Now(),
				}
				(*oldPrices)[ticker.Pair] = ticker.Price
			}
		} else {
			(*oldPrices)[ticker.Pair] = ticker.Price
		}

	}

	return nil
}
