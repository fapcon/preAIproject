package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE json
type BotStatisticRequest struct {
	BotUUID string `json:"bot_uuid"`
}

type BotStatisticResponse struct {
	Success   bool                 `json:"success"`
	ErrorCode int                  `json:"error_code"`
	Data      models.BotStatistics `json:"data"`
}

type UserStatisticRequest struct {
	UserID int `json:"user_id"`
}

type UserStatisticResponse struct {
	Success   bool                   `json:"success"`
	ErrorCode int                    `json:"error_code"`
	Data      []models.BotStatistics `json:"data"`
}

type BotStatisticDeleteRequest struct {
	BotUUID string `json:"bot_uuid"`
}

type BotStatisticDeleteResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}
