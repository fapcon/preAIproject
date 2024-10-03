package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	service2 "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/platform/service"
)

//go:generate easytags $GOFILE json
type TickerResponse struct {
	Success   bool                    `json:"success"`
	ErrorCode int                     `json:"error_code,omitempty"`
	Data      []models.ExchangeTicker `json:"data"`
}

type ExchangeAddRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

type ExchangeResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
}

type ExchangeListResponse struct {
	Success   bool                  `json:"success"`
	ErrorCode int                   `json:"error_code,omitempty"`
	Data      []models.ExchangeList `json:"data"`
}

type ExchangeDeleteRequest struct {
	ID int `json:"id"`
}

type ExchangeUserKeyAddRequest struct {
	ExchangeID int    `json:"exchange_id"`
	Label      string `json:"label"`
	APIKey     string `json:"api_key"`
	SecretKey  string `json:"secret_key"`
}

type ExchangeUserKeyDeleteRequest struct {
	ExchangeUserKeyID int `json:"exchange_user_key_id"`
}

type ExchangeUserKeyListResponse struct {
	Success   bool                     `json:"success"`
	ErrorCode int                      `json:"error_code,omitempty"`
	Data      []models.ExchangeUserKey `json:"data"`
}

type WebhookProcessResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
}

type WebhookHistoryRequest struct {
	WebhookUUID string `json:"webhook_uuid"`
}

type GetUserOrdersResponse struct {
	Success   bool                   `json:"success"`
	ErrorCode int                    `json:"error_code"`
	Data      []models.ExchangeOrder `json:"data"`
	Statistic service.StatisticOut
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

type ExchangeOrderResponse struct {
	Success   bool                      `json:"success"`
	ErrorCode int                       `json:"error_code"`
	Data      []models.ExchangeOrderNew `json:"data"`
}

type GetAccountBalanceOut struct {
	Success    bool
	ErrorCode  int
	DataSpot   BalanceInfo
	DataMargin BalanceInfo
}

type BalanceInfo struct {
	Message  string
	Balances []models.BalanceDTO
}

type GetCandlesRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	Limit     int    `json:"limit"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

type CandlesResponse struct {
	ErrorCode int
	Candles   []service2.CandlesData // Заменить в дальнейшем на модель DTO
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
