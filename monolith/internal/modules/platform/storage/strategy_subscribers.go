package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type StrategySubscribers struct {
	adapter adapter.SQLAdapter
}

func NewStrategySubscribers(adapter adapter.SQLAdapter) StrategySubscriberer {
	return &StrategySubscribers{adapter: adapter}
}

func (s *StrategySubscribers) Create(ctx context.Context, dto models.StrategySubscribersDTO) error {
	return s.adapter.Create(ctx, &dto)
}

func (s *StrategySubscribers) Update(ctx context.Context, dto models.StrategySubscribersDTO) error {
	return s.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (s *StrategySubscribers) GetByID(ctx context.Context, strategyID int) (models.StrategySubscribersDTO, error) {
	var list []models.StrategySubscribersDTO
	err := s.adapter.List(ctx, &list, "strategy_subscribers", utils.Condition{
		Equal: map[string]interface{}{"id": strategyID},
	})
	if err != nil {
		return models.StrategySubscribersDTO{}, err
	}
	if len(list) < 1 {
		return models.StrategySubscribersDTO{}, fmt.Errorf("strategy_subscribers storage: GetByID not found")
	}
	return list[0], err
}

func (s *StrategySubscribers) GetList(ctx context.Context) ([]models.StrategySubscribersDTO, error) {
	var list []models.StrategySubscribersDTO
	err := s.adapter.List(ctx, &list, "strategy_subscribers", utils.Condition{Equal: map[string]interface{}{"deleted_at": nil}})
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *StrategySubscribers) Delete(ctx context.Context, strategyID int) error {
	dto, err := s.GetByID(ctx, strategyID)
	if err != nil {
		return err
	}

	dto.SetDeletedAt(time.Now())

	return s.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}
