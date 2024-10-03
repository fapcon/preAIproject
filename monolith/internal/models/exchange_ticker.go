package models

import "github.com/shopspring/decimal"

//go:generate easytags $GOFILE json,mapper
type ExchangeTicker struct {
	ID         int             `json:"id" mapper:"id"`
	Pair       string          `json:"pair" mapper:"pair"`
	Price      decimal.Decimal `json:"price" mapper:"price"`
	ExchangeID int             `json:"exchange_id" mapper:"exchange_id"`
}
