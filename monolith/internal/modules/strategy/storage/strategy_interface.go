package storage

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type Strateger interface {
	Create(ctx context.Context, dto models.StrategyDTO) (int, error)
	Update(ctx context.Context, dto models.StrategyDTO) error
	GetByID(ctx context.Context, strategyID int) (models.StrategyDTO, error)
	GetByName(ctx context.Context, name string) (models.StrategyDTO, error)
	GetList(ctx context.Context) ([]models.StrategyDTO, error)
	Delete(ctx context.Context, strategyID int) error
}
