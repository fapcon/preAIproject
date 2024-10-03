package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type StrategyPairStorage struct {
	adapter adapter.SQLAdapter
}

func NewStrategyPairStorage(adapter adapter.SQLAdapter) *StrategyPairStorage {
	return &StrategyPairStorage{adapter: adapter}
}

func (s *StrategyPairStorage) Create(ctx context.Context, dto models.StrategyPairDTO) error {
	return s.adapter.Create(ctx, &dto)
}

func (s *StrategyPairStorage) Update(ctx context.Context, dto models.StrategyPairDTO) error {
	return s.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (s *StrategyPairStorage) GetByStrategyID(ctx context.Context, exchangeID int) ([]models.StrategyPairDTO, error) {
	var list []models.StrategyPairDTO
	err := s.adapter.List(ctx, &list, "strategy_list", utils.Condition{
		Equal: map[string]interface{}{"id": exchangeID},
	})
	if err != nil {
		return []models.StrategyPairDTO{}, err
	}
	if len(list) < 1 {
		return []models.StrategyPairDTO{}, fmt.Errorf("strategy list storage: GetByID not found")
	}
	return nil, nil
}

func (s *StrategyPairStorage) GetList(ctx context.Context, condition utils.Condition) ([]models.StrategyPairDTO, error) {
	return nil, nil
}

func (s *StrategyPairStorage) Delete(ctx context.Context, strategyID int) error {
	return nil
}

func updateTimeStrategyPair(strategyPair *models.StrategyPairDTO) {
	strategyPair.SetUpdatedAt(time.Now())
}
