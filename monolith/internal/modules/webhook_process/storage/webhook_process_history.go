package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type WebhookProcessHistory struct {
	adapter adapter.SQLAdapter
}

func NewWebhookProcessHistory(adapter adapter.SQLAdapter) WebhookProcessHistorer {
	return &WebhookProcessHistory{adapter: adapter}
}

func (e *WebhookProcessHistory) Create(ctx context.Context, dto models.WebhookProcessHistoryDTO) error {
	return e.adapter.Create(ctx, &dto)
}

func (e *WebhookProcessHistory) Update(ctx context.Context, dto models.WebhookProcessHistoryDTO) error {
	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (e *WebhookProcessHistory) GetByID(ctx context.Context, exchangeID int) (models.WebhookProcessHistoryDTO, error) {
	var list []models.WebhookProcessHistoryDTO
	err := e.adapter.List(ctx, &list, "webhook_process_history", utils.Condition{
		Equal: map[string]interface{}{"id": exchangeID},
	})
	if err != nil {
		return models.WebhookProcessHistoryDTO{}, err
	}
	if len(list) < 1 {
		return models.WebhookProcessHistoryDTO{}, fmt.Errorf("exchange orderlog storage: GetByID not found")
	}
	return list[0], err
}

func (e *WebhookProcessHistory) GetList(ctx context.Context, condition utils.Condition) ([]models.WebhookProcessHistoryDTO, error) {
	var list []models.WebhookProcessHistoryDTO
	err := e.adapter.List(ctx, &list, "webhook_process_history", condition)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (e *WebhookProcessHistory) Delete(ctx context.Context, webhookProcessHistoryID int) error {
	dto, err := e.GetByID(ctx, webhookProcessHistoryID)
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
