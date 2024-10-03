package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type ExchangeOrderLog struct {
	adapter adapter.SQLAdapter
}

func NewExchangeOrderLog(adapter adapter.SQLAdapter) ExchangeOrderLogger {
	return &ExchangeOrderLog{adapter: adapter}
}

func (e *ExchangeOrderLog) Create(ctx context.Context, dto models.ExchangeOrderLogDTO) error {
	return e.adapter.Create(ctx, &dto)
}

func (e *ExchangeOrderLog) Update(ctx context.Context, dto models.ExchangeOrderLogDTO) error {
	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (e *ExchangeOrderLog) GetByID(ctx context.Context, exchangeID int) (models.ExchangeOrderLogDTO, error) {
	var list []models.ExchangeOrderLogDTO
	err := e.adapter.List(ctx, &list, "exchange_order_log", utils.Condition{
		Equal: map[string]interface{}{"id": exchangeID},
	})
	if err != nil {
		return models.ExchangeOrderLogDTO{}, err
	}
	if len(list) < 1 {
		return models.ExchangeOrderLogDTO{}, fmt.Errorf("exchange orderlog storage: GetByID not found")
	}
	return list[0], err
}

func (e *ExchangeOrderLog) GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderLogDTO, error) {
	var list []models.ExchangeOrderLogDTO
	err := e.adapter.List(ctx, &list, "exchange_order_log", condition)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (e *ExchangeOrderLog) Delete(ctx context.Context, exchangeOrderLogID int) error {
	dto, err := e.GetByID(ctx, exchangeOrderLogID)
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
