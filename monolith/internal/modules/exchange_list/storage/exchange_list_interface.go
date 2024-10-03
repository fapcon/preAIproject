package storage

import (
	"context"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeLister

type ExchangeLister interface {
	Create(ctx context.Context, dto models.ExchangeListDTO) error
	Update(ctx context.Context, dto models.ExchangeListDTO) error
	GetByID(ctx context.Context, exchangeID int) (models.ExchangeListDTO, error)
	GetList(ctx context.Context) ([]models.ExchangeListDTO, error)
	Delete(ctx context.Context, exchangeID int) error
}
