package models

import "github.com/shopspring/decimal"

//go:generate easytags $GOFILE json,mapper
type ExchangeUserKey struct {
	ID            int           `json:"id" mapper:"id"`
	ExchangeID    int           `json:"exchange_id" mapper:"exchange_id"`
	UserID        int           `json:"user_id" mapper:"user_id"`
	Label         string        `json:"label" mapper:"label"`
	MakeOrder     bool          `json:"make_order" mapper:"make_order"`
	APIKey        string        `json:"api_key" mapper:"api_key"`
	SecretKey     string        `json:"secret_key" mapper:"secret_key"`
	StatisticData StatisticData `json:"statistic_data" mapper:"statistic_data"`
}

type StatisticData struct {
	SumSell decimal.Decimal `json:"sum_sell" mapper:"sum_sell"`
	SumBuy  decimal.Decimal `json:"sum_buy" mapper:"sum_buy"`
	Profit  decimal.Decimal `json:"profit" mapper:"profit"`
	ToSell  decimal.Decimal `json:"to_sell" mapper:"to_sell"`
	ToEarn  decimal.Decimal `json:"to_earn" mapper:"to_earn"`
	Earned  decimal.Decimal `json:"earned" mapper:"earned"`
}
