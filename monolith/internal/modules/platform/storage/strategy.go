package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type Strategy struct {
	adapter adapter.SQLAdapter
}

func NewStrategy(adapter adapter.SQLAdapter) *Strategy {
	return &Strategy{adapter: adapter}
}

func (e *Strategy) Create(ctx context.Context, dto models.StrategyDTO) error {
	return e.adapter.Create(ctx, &dto)
}

func (e *Strategy) Update(ctx context.Context, dto models.StrategyDTO) error {
	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (e *Strategy) GetByID(ctx context.Context, strategyID int) (models.StrategyDTO, error) {
	var list []models.StrategyDTO
	err := e.adapter.List(ctx, &list, "strategy", utils.Condition{
		Equal: map[string]interface{}{"id": strategyID},
	})
	if err != nil {
		return models.StrategyDTO{}, err
	}
	if len(list) < 1 {
		return models.StrategyDTO{}, fmt.Errorf("strategy storage: GetByID not found")
	}
	return list[0], err
}

func (e *Strategy) GetList(ctx context.Context) ([]models.StrategyDTO, error) {
	var list []models.StrategyDTO
	err := e.adapter.List(ctx, &list, "strategy", utils.Condition{Equal: map[string]interface{}{"deleted_at": nil}})
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (e *Strategy) Delete(ctx context.Context, strategyID int) error {
	dto, err := e.GetByID(ctx, strategyID)
	if err != nil {
		return err
	}

	dto.SetDeletedAt(time.Now())

	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}
