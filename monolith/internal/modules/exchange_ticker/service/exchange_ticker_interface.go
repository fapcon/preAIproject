package service

import (
	"context"

	"gitlab.com/golight/orm/utils"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeTicker
type ExchangeTicker interface {
	Save(ctx context.Context, tickers []models.ExchangeTicker) error
	GetTicker(ctx context.Context) GetTickerOut
	GetByID(ctx context.Context, tickerID int) (models.ExchangeTicker, error)
	GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeTicker, error)
}
