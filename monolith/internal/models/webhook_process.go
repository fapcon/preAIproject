package models

//go:generate easytags $GOFILE json,mapper
type WebhookProcess struct {
	ID        int                     `json:"id" mapper:"id"`
	OrderID   int64                   `json:"order_id" mapper:"order_id"`
	BotID     int                     `json:"bot_id" mapper:"bot_id"`
	Status    int                     `json:"status" mapper:"status"`
	StatusMsg string                  `json:"status_msg" mapper:"status_msg"`
	History   []WebhookProcessHistory `json:"history" mapper:"history"`
}
