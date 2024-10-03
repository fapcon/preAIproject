package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type ExchangeList struct {
	adapter adapter.SQLAdapter
}

func NewExchangeList(adapter adapter.SQLAdapter) *ExchangeList {
	return &ExchangeList{adapter: adapter}
}

func (e ExchangeList) Create(ctx context.Context, dto models.ExchangeListDTO) error {
	return e.adapter.Create(ctx, &dto)
}

func (e ExchangeList) Update(ctx context.Context, dto models.ExchangeListDTO) error {
	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (e ExchangeList) GetByID(ctx context.Context, exchangeID int) (models.ExchangeListDTO, error) {
	var list []models.ExchangeListDTO
	err := e.adapter.List(ctx, &list, "exchange_list", utils.Condition{
		Equal: map[string]interface{}{"id": exchangeID},
	})
	if err != nil {
		return models.ExchangeListDTO{}, err
	}
	if len(list) < 1 {
		return models.ExchangeListDTO{}, fmt.Errorf("exchange list storage: GetByID not found")
	}
	return list[0], err
}

func (e ExchangeList) GetList(ctx context.Context) ([]models.ExchangeListDTO, error) {
	var list []models.ExchangeListDTO
	err := e.adapter.List(ctx, &list, "exchange_list", utils.Condition{Equal: map[string]interface{}{"deleted_at": nil}})
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (e ExchangeList) Delete(ctx context.Context, exchangeID int) error {
	dto, err := e.GetByID(ctx, exchangeID)
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
