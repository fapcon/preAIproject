package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type ExchangeUserKeyer interface {
	Create(ctx context.Context, dto models.ExchangeUserKeyDTO) error
	GetByUserID(ctx context.Context, userID int) ([]models.ExchangeUserKeyDTO, error)
	GetByID(ctx context.Context, exchangeUserKeyID int) (models.ExchangeUserKeyDTO, error)
	Update(ctx context.Context, dto models.ExchangeUserKeyDTO) error
	GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeUserKeyDTO, error)
	Delete(ctx context.Context, exchangeUserID int) error
}
