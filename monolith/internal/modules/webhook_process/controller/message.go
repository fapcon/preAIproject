package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE json

type WebhookProcessResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
}

type GetUserWebhooksResponse struct {
	Success   bool                    `json:"success"`
	ErrorCode int                     `json:"error_code"`
	Data      []models.WebhookProcess `json:"data"`
}

type WebhookInfoRequest struct {
	WebhookUUID string `json:"webhook_uuid"`
}

type WebhookInfoResponse struct {
	Success   bool            `json:"success"`
	ErrorCode int             `json:"error_code"`
	Data      WebhookInfoData `json:"data"`
}

type WebhookInfoData struct {
	Webhook models.WebhookProcess  `json:"bot"`
	Orders  []models.ExchangeOrder `json:"orders"`
}

type BotInfoRequest struct {
	BotUUID string `json:"bot_uuid"`
}

type BotInfoResponse struct {
	Success   bool        `json:"success"`
	ErrorCode int         `json:"error_code"`
	Data      BotInfoData `json:"data"`
}

type BotInfoData struct {
	Bot      models.Bot              `json:"bot"`
	Orders   []models.ExchangeOrder  `json:"orders"`
	Webhooks []models.WebhookProcess `json:"webhooks"`
}
