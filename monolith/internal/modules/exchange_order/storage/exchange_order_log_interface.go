package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeOrderLogger
type ExchangeOrderLogger interface {
	Create(ctx context.Context, dto models.ExchangeOrderLogDTO) error
	Update(ctx context.Context, dto models.ExchangeOrderLogDTO) error
	GetByID(ctx context.Context, exchangeID int) (models.ExchangeOrderLogDTO, error)
	GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderLogDTO, error)
	Delete(ctx context.Context, exchangeOrderLogDTOID int) error
}
