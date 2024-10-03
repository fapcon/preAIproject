package service

import (
	"context"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name WebhookProcesser
type WebhookProcesser interface {
	WebhookProcess(ctx context.Context, in WebhookProcessIn) WebhookProcessOut
	UpdateWebhookStatus(ctx context.Context, bot models.Bot, webhook models.WebhookProcessDTO, message string, status int)
	CreateWebhookProcess(ctx context.Context, bot models.Bot, in WebhookProcessIn) (models.WebhookProcessDTO, error)
	WriteWebhookHistory(ctx context.Context, dto models.WebhookProcessHistoryDTO)
	GetWebhookInfo(ctx context.Context, in GetWebhookInfoIn) GetWebhookInfoOut
	GetUserWebhooks(ctx context.Context, in GetUserRelationIn) GetWebhooksOut

	GetBotWebhooks(ctx context.Context, in GetBotRelationIn) GetWebhooksOut
	GetBotInfo(ctx context.Context, in GetBotInfoIn) GetBotInfoOut
}
