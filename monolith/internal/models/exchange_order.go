package models

import (
	"time"

	"github.com/shopspring/decimal"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,mapper
type ExchangeOrder struct {
	ID           int64              `json:"id" mapper:"id"`
	UUID         string             `json:"uuid" mapper:"uuid"`
	OrderID      int64              `json:"order_id" mapper:"order_id"`
	UserID       int                `json:"user_id" mapper:"user_id"`
	ExchangeID   int                `json:"exchange_id" mapper:"exchange_id"`
	UnitedOrders int                `json:"united_orders" mapper:"united_orders"`
	OrderType    int                `json:"order_type" mapper:"order_type"`
	OrderTypeMsg string             `json:"order_type_msg" mapper:"order_type_msg"`
	Pair         string             `json:"pair" mapper:"pair"`
	Amount       decimal.Decimal    `json:"amount" mapper:"amount"`
	Quantity     decimal.Decimal    `json:"quantity" mapper:"quantity"`
	Price        decimal.Decimal    `json:"price" mapper:"price"`
	Side         int                `json:"side" mapper:"side"`
	SideMsg      string             `json:"side_msg" mapper:"side_msg"`
	Message      string             `json:"message" mapper:"message"`
	Status       int                `json:"status" mapper:"status"`
	StatusMsg    string             `json:"status_msg" mapper:"status_msg"`
	History      []ExchangeOrderLog `json:"history" mapper:"history"`
	SumBuy       int                `json:"sumBuy" mapper:"sumBuy"`
	ApiKeyID     int                `json:"api_key_id" mapper:"api_key_id"`
	CreatedAt    time.Time          `json:"created_at" mapper:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" mapper:"updated_at"`
	DeletedAt    types.NullTime     `json:"deleted_at" mapper:"deleted_at"`
}

type ExchangeOrderNew struct {
	ID              int64          `json:"id" mapper:"id"`
	UUID            string         `json:"uuid" mapper:"uuid"`
	BotUUID         string         `json:"bot_uuid" mapper:"bot_uuid"`
	OrderType       int            `json:"order_type" mapper:"order_type"`
	Side            int            `json:"side" mapper:"side"`
	ExchangeOrderID int64          `json:"exchange_order_id" mapper:"exchange_order_id"`
	UserID          int            `json:"user_id" mapper:"user_id"`
	ExchangeID      int            `json:"exchange_id" mapper:"exchange_id"`
	Pair            string         `json:"pair" mapper:"pair"`
	Amount          float64        `json:"amount" mapper:"amount"`
	Quantity        float64        `json:"quantity" mapper:"quantity"`
	Price           float64        `json:"price" mapper:"price"`
	Status          int            `json:"status" mapper:"status"`
	WebhookUUID     string         `json:"webhook_uuid" mapper:"webhook_uuid"`
	Message         string         `json:"message" mapper:"message"`
	CreatedAt       time.Time      `json:"created_at" mapper:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" mapper:"updated_at"`
	DeletedAt       types.NullTime `json:"deleted_at" mapper:"deleted_at"`
}
