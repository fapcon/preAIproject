package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type WebhookProcess struct {
	adapter adapter.SQLAdapter
}

func NewWebhookProcess(adapter adapter.SQLAdapter) WebhookProcesser {
	return &WebhookProcess{adapter: adapter}
}

func (e *WebhookProcess) Create(ctx context.Context, dto models.WebhookProcessDTO) error {
	return e.adapter.Create(ctx, &dto)
}

func (e *WebhookProcess) Update(ctx context.Context, dto models.WebhookProcessDTO) error {
	return e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"uuid": dto.GetUUID()},
		},
		utils.Update,
	)
}

func (e *WebhookProcess) GetByID(ctx context.Context, WebhookProcessID int) (models.WebhookProcessDTO, error) {
	var list []models.WebhookProcessDTO
	err := e.adapter.List(ctx, &list, "webhook_process", utils.Condition{
		Equal: map[string]interface{}{"id": WebhookProcessID},
	})
	if err != nil {
		return models.WebhookProcessDTO{}, err
	}
	if len(list) < 1 {
		return models.WebhookProcessDTO{}, fmt.Errorf("WebhookProcess storage: GetByID not found")
	}
	return list[0], err
}

func (e *WebhookProcess) GetList(ctx context.Context, condition utils.Condition) ([]models.WebhookProcessDTO, error) {
	var list []models.WebhookProcessDTO
	err := e.adapter.List(ctx, &list, "webhook_process", condition)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (e *WebhookProcess) Delete(ctx context.Context, WebhookProcessID int) error {
	dto, err := e.GetByID(ctx, WebhookProcessID)
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
