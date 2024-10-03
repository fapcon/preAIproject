package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type ExchangeOrder struct {
	adapter adapter.SQLAdapter
}

func NewExchangeOrder(adapter adapter.SQLAdapter) ExchangeOrderer {
	return &ExchangeOrder{adapter: adapter}
}

func (e *ExchangeOrder) Create(ctx context.Context, dto models.ExchangeOrderDTO) error {
	return e.adapter.Create(ctx, &dto)
}

func (e *ExchangeOrder) Update(ctx context.Context, dto models.ExchangeOrderDTO) error {
	dto.SetUpdatedAt(time.Now())
	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"uuid": dto.GetUUID()},
		},
		utils.Update,
	)
}

func (e *ExchangeOrder) GetByUUID(ctx context.Context, exchangeOrderUUID string) (models.ExchangeOrderDTO, error) {
	var list []models.ExchangeOrderDTO
	err := e.adapter.List(ctx, &list, "exchange_order", utils.Condition{
		Equal: map[string]interface{}{"uuid": exchangeOrderUUID},
	})
	if err != nil {
		return models.ExchangeOrderDTO{}, err
	}
	if len(list) < 1 {
		return models.ExchangeOrderDTO{}, fmt.Errorf("exchange order storage: GetByID not found")
	}
	return list[0], err
}

func (e *ExchangeOrder) GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error) {
	var list []models.ExchangeOrderDTO
	err := e.adapter.List(ctx, &list, "exchange_order", condition)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (e *ExchangeOrder) Delete(ctx context.Context, exchangeOrderUUID string) error {
	dto, err := e.GetByUUID(ctx, exchangeOrderUUID)
	if err != nil {
		return err
	}

	dto.SetDeletedAt(time.Now())

	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"uuid": dto.GetUUID()},
		},
		utils.Update,
	)
}

var NotFoundUUID = fmt.Errorf("bot storage: GetByBotUUID not found")

func (e *ExchangeOrder) GetByBotUUID(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error) {
	var exchangeOrders []models.ExchangeOrderDTO
	err := e.adapter.List(ctx, &exchangeOrders, "exchange_order", condition)
	if err != nil {
		return nil, err
	}
	if len(exchangeOrders) < 1 {
		return nil, NotFoundUUID
	}
	return exchangeOrders, nil
}
