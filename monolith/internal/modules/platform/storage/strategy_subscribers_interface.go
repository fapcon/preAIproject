package storage

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type StrategySubscriberer interface {
	Create(ctx context.Context, dto models.StrategySubscribersDTO) error
	Update(ctx context.Context, dto models.StrategySubscribersDTO) error
	GetByID(ctx context.Context, strategySubscribersID int) (models.StrategySubscribersDTO, error)
	GetList(ctx context.Context) ([]models.StrategySubscribersDTO, error)
	Delete(ctx context.Context, strategySubscribersID int) error
}
