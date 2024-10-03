package models

//go:generate easytags $GOFILE json,mapper
type ExchangeProvider struct {
	ID         int    `json:"id" mapper:"id"`
	UserID     int    `json:"user_id" mapper:"user_id"`
	ExchangeID int    `json:"exchange_id" mapper:"exchange_id"`
	APIKey     string `json:"api_key" mapper:"api_key"`
	SecretKey  string `json:"secret_key" mapper:"secret_key"`
}
