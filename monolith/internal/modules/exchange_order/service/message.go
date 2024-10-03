package service

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
)

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

type StatisticOut struct {
	Keys []models.ExchangeUserKey `json:"keys"`
}

type PutOrderOut struct {
	ErrorCode int
	Success   bool
	OrderID   int64
	OrderUUID string
}

type CancelOrderOut struct {
	ExchangeOrders []service.OrderOut
	PlatformOrders []models.ExchangeOrderDTO
}
