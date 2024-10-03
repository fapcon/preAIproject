package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type StrategyPairer interface {
	Create(ctx context.Context, dto models.StrategyPairDTO) error
	Update(ctx context.Context, dto models.StrategyPairDTO) error
	GetByStrategyID(ctx context.Context, exchangeID int) ([]models.StrategyPairDTO, error)
	GetList(ctx context.Context, condition utils.Condition) ([]models.StrategyPairDTO, error)
	Delete(ctx context.Context, strategyID int) error
}
