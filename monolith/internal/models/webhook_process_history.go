package models

//go:generate easytags $GOFILE json,mapper
type WebhookProcessHistory struct {
	OrderID    int    `json:"order_id" mapper:"order_id"`
	ExchangeID int    `json:"exchange_id" mapper:"exchange_id"`
	WebhookID  int    `json:"webhook_id" mapper:"webhook_id"`
	Pair       string `json:"pair" mapper:"pair"`
	Status     int    `json:"status" mapper:"status"`
	StatusMsg  string `json:"status_msg" mapper:"status_msg"`
}
