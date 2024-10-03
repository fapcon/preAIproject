package models

type StrategyTest struct {
	ID         int    `json:"id" mapper:"id"`
	ExchangeID int    `json:"exchange_id" mapper:"exchange_id"`
	UserID     int    `json:"user_id" mapper:"user_id"`
	APIKey     string `json:"api_key" mapper:"api_key"`
	SecretKey  string `json:"secret_key" mapper:"secret_key"`
}
