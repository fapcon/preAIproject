package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name WebhookProcessHistorer
type WebhookProcessHistorer interface {
	Create(ctx context.Context, dto models.WebhookProcessHistoryDTO) error
	Update(ctx context.Context, dto models.WebhookProcessHistoryDTO) error
	GetByID(ctx context.Context, webhookProcessHistoryID int) (models.WebhookProcessHistoryDTO, error)
	GetList(ctx context.Context, condition utils.Condition) ([]models.WebhookProcessHistoryDTO, error)
	Delete(ctx context.Context, WebhookProcessHistoryID int) error
}
