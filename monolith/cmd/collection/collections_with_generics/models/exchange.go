package models

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"time"
)

type ExchangeOrderTest struct {
	ID           int64          `json:"id" mapper:"id"`
	UUID         string         `json:"uuid" mapper:"uuid"`
	OrderID      int64          `json:"order_id" mapper:"order_id"`
	UserID       int            `json:"user_id" mapper:"user_id"`
	ExchangeID   int            `json:"exchange_id" mapper:"exchange_id"`
	UnitedOrders int            `json:"united_orders" mapper:"united_orders"`
	OrderType    int            `json:"order_type" mapper:"order_type"`
	OrderTypeMsg string         `json:"order_type_msg" mapper:"order_type_msg"`
	Pair         string         `json:"pair" mapper:"pair"`
	Amount       float64        `json:"amount" mapper:"amount"`
	Quantity     float64        `json:"quantity" mapper:"quantity"`
	Price        float64        `json:"price" mapper:"price"`
	Side         int            `json:"side" mapper:"side"`
	SideMsg      string         `json:"side_msg" mapper:"side_msg"`
	Message      string         `json:"message" mapper:"message"`
	Status       int            `json:"status" mapper:"status"`
	StatusMsg    string         `json:"status_msg" mapper:"status_msg"`
	SumBuy       int            `json:"sumBuy" mapper:"sumBuy"`
	ApiKeyID     int            `json:"api_key_id" mapper:"api_key_id"`
	CreatedAt    time.Time      `json:"created_at" mapper:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" mapper:"updated_at"`
	DeletedAt    types.NullTime `json:"deleted_at" mapper:"deleted_at"`
}
