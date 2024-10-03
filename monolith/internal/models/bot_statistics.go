package models

import "github.com/shopspring/decimal"

//go:generate easytags $GOFILE json,mapper
type BotStatistics struct {
	ID            int             `json:"bot_id" mapper:"bot_id"`
	BotUUID       string          `json:"bot_uuid" mapper:"bot_uuid"`
	UserID        int             `json:"user_id" mapper:"user_id"`
	Profitability decimal.Decimal `json:"profitability" mapper:"profitability"` //прибыль
	OrderCount    int             `json:"order_count" mapper:"order_count"`
}
