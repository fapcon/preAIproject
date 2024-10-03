package models

import (
	"time"

	"github.com/shopspring/decimal"
)

//go:generate easytags $GOFILE json,mapper
type ExchangeOrderLog struct {
	UUID       string          `json:"uuid" mapper:"uuid"`
	OrderID    int64           `json:"order_id" mapper:"order_id"`
	UserID     int             `json:"user_id" mapper:"user_id"`
	ExchangeID int             `json:"exchange_id" mapper:"exchange_id"`
	Pair       string          `json:"pair" mapper:"pair"`
	Quantity   decimal.Decimal `json:"quantity" mapper:"quantity"`
	Price      decimal.Decimal `json:"price" mapper:"price"`
	Status     int             `json:"status" mapper:"status"`
	StatusMsg  string          `json:"status_msg" mapper:"status_msg"`
	CreatedAt  time.Time       `json:"created_at" mapper:"created_at"`
}
