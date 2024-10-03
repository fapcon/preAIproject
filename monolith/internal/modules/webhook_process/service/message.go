package service

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type WebhookProcessIn struct {
	Slug          string
	XForwardedFor string
	RemoteAddr    string
}

type WebhookProcessOut struct {
	ErrorCode int
	Success   bool
}

type GetUserRelationIn struct {
	UserID int
}

type GetBotRelationIn struct {
	BotUUID string
}

type GetWebhooksOut struct {
	ErrorCode int
	Success   bool
	Data      []models.WebhookProcess
}

type GetWebhookInfoIn struct {
	WebhookUUID string
	UserID      int
}

type GetWebhookInfoOut struct {
	ErrorCode int
	Success   bool
	Data      GetWebhookInfoData
}

type GetWebhookInfoData struct {
	Webhook models.WebhookProcess
	Orders  []models.ExchangeOrder
}

type GetBotInfoIn struct {
	BotUUID string
	UserID  int
}

type GetBotInfoOut struct {
	ErrorCode int
	Success   bool
	Data      GetBotInfoData
}

type GetBotInfoData struct {
	Bot      models.Bot
	Orders   []models.ExchangeOrder
	Webhooks []models.WebhookProcess
}
