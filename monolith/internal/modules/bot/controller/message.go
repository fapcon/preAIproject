package controller

import "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

//go:generate easytags $GOFILE json
type BotResponse struct {
	Success   bool    `json:"success"`
	ErrorCode int     `json:"error_code,omitempty"`
	Data      BotData `json:"data"`
}

type DefaultResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type BotData struct {
	Message string            `json:"message,omitempty"`
	Bot     models.Bot        `json:"strategy,omitempty"`
	Hooks   map[string]string `json:"hooks"`
}

type GetBotRequest struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
}

type UpdateRequest struct {
	Strategy models.Bot `json:"strategy,omitempty"`
	Pairs    []int      `json:"pairs"`
}

type UpdateResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
}

type BotListResponse struct {
	Success   bool         `json:"success"`
	ErrorCode int          `json:"error_code,omitempty"`
	Data      []models.Bot `json:"data"`
}

type BotToggleRequest struct {
	BotUUID string `json:"bot_uuid"`
	Active  bool   `json:"active"`
}

type BotUUIDRequest struct {
	BotUUID string `json:"bot_uuid"`
}

type WebhookSignalRequest struct {
	BotUUID string `json:"bot_uuid"`
	PairID  int    `json:"pair_id"`
}

type WebhookSignalResponse struct {
	ErrorCode int
	Success   bool
	Hook      string   `json:"hook"`
	Signals   []string `json:"signals"`
}
