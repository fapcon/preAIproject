package models

import (
	"time"
)

//go:generate easytags $GOFILE json,mapper
type Bot struct {
	ID                   int            `json:"id" mapper:"id"`
	Kind                 int            `json:"kind" mapper:"kind"`
	UserID               int            `json:"user_id" mapper:"user_id"`
	Name                 string         `json:"name" mapper:"name"`
	Description          string         `json:"description" mapper:"description"`
	PairID               int            `json:"pair_id" mapper:"pair_id"`
	FixedAmount          float64        `json:"fixed_amount" mapper:"fixed_amount"`
	ExchangeType         int            `json:"exchange_type" mapper:"exchange_type"`
	ExchangeID           int            `json:"exchange_id" mapper:"exchange_id"`
	ExchangeUserKeyID    int            `json:"exchange_user_key_id" mapper:"exchange_user_key_id"`
	OrderType            int            `json:"order_type" mapper:"order_type"`
	SellPercent          float64        `json:"sell_percent" mapper:"sell_percent"`
	CommissionPercent    float64        `json:"commission_percent" mapper:"commission_percent"`
	AssetType            int            `json:"asset_type" mapper:"asset_type"`
	UUID                 string         `json:"uuid" mapper:"uuid"`
	Active               bool           `json:"active" mapper:"active"`
	LimitOrder           bool           `json:"limit_order" mapper:"limit_order"`
	LimitSellPercent     float64        `json:"limit_sell_percent" mapper:"limit_sell_percent"`
	LimitBuyPercent      float64        `json:"limit_buy_percent" mapper:"limit_buy_percent"`
	AutoSell             bool           `json:"auto_sell" mapper:"auto_sell"`
	AutoLimitSellPercent float64        `json:"auto_limit_sell_percent" mapper:"auto_limit_sell_percent"`
	OrderCountLimit      bool           `json:"order_count_limit" mapper:"order_count_limit"`
	OrderCount           int            `json:"order_count" mapper:"order_count"`
	Pairs                []StrategyPair `json:"pairs" mapper:"pairs"`
	CreatedAt            time.Time      `json:"created_at" mapper:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at" mapper:"updated_at"`
	DeletedAt            time.Time      `json:"deleted_at" mapper:"deleted_at"`
}
