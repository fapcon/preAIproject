package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name WebhookProcesser
type WebhookProcesser interface {
	Create(ctx context.Context, dto models.WebhookProcessDTO) error
	Update(ctx context.Context, dto models.WebhookProcessDTO) error
	GetByID(ctx context.Context, exchangeID int) (models.WebhookProcessDTO, error)
	GetList(ctx context.Context, condition utils.Condition) ([]models.WebhookProcessDTO, error)
	Delete(ctx context.Context, WebhookProcessID int) error
}
