package service

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
)

type GetTickerOut struct {
	ErrorCode int
	Data      []models.ExchangeTicker
}

type ExchangeListOut struct {
	ErrorCode int
	Data      []models.ExchangeList
}
type ExchangeProviderListOut struct {
	ErrorCode int
	Data      []models.ExchangeProvider
}

type ExchangeAPIListOut struct {
	ErrorCode int
	Data      []models.ExchangeProvider
}

type ExchangeAddIn struct {
	UserID      int
	Name        string
	Description string
	Slug        string
}

type ExchangeUserKeyAddIn struct {
	ExchangeID int
	UserID     int
	APIKey     string
	SecretKey  string
}

type ExchangeOut struct {
	ErrorCode int
	Success   bool
}

type ExchangeUserListIn struct {
	UserID int
}

type ExchangeProviderAddIn struct {
	ExchangeID int
	APIKey     string
	SecretKey  string
}

type ExchangeUserListOut struct {
	ErrorCode int
	Data      []models.ExchangeUserKey
}

type WebhookProcessIn struct {
	Slug          string
	XForwardedFor string
	RemoteAddr    string
}

type WebhookProcessOut struct {
	ErrorCode int
	Success   bool
}

type OrderIn struct {
	ExClient service.Exchanger
	Webhook  models.WebhookProcessDTO
	Bot      models.Bot
	Signal   models.Signal
	Key      models.ExchangeUserKeyDTO
	BuyOrder models.ExchangeOrderDTO
}

type WriteOrderIn struct {
	ExchangeOrder service.OrderOut
	Webhook       models.WebhookProcessDTO
	Bot           models.Bot
	BuyOrder      models.ExchangeOrderDTO
	Side          int
	OrderType     int
	Signal        models.Signal
	Message       string
	UnitedOrders  int
	Key           models.ExchangeUserKeyDTO
}

type PutOrderOut struct {
	ErrorCode int
	Success   bool
}

type GetUserRelationIn struct {
	UserID int
}

type GetBotRelationIn struct {
	BotUUID string
}

type GetOrdersOut struct {
	ErrorCode int
	Success   bool
	Data      []models.ExchangeOrder
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

type ExchangeOrdersOut struct {
	ErrorCode int
	Success   bool
	Data      []models.ExchangeOrderNew
}

type StatisticOut struct {
	Keys []models.ExchangeUserKey `json:"keys"`
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
