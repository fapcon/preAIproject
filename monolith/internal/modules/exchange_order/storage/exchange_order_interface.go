package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeOrderer
type ExchangeOrderer interface {
	Create(ctx context.Context, dto models.ExchangeOrderDTO) error
	Update(ctx context.Context, dto models.ExchangeOrderDTO) error
	GetByUUID(ctx context.Context, exchangeOrderUUID string) (models.ExchangeOrderDTO, error)
	GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error)
	Delete(ctx context.Context, exchangeOrderUUID string) error
	GetByBotUUID(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error)
}
