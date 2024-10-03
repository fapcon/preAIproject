package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
)

//go:generate easytags $GOFILE json

type GetUserOrdersResponse struct {
	Success   bool                   `json:"success"`
	ErrorCode int                    `json:"error_code"`
	Data      []models.ExchangeOrder `json:"data"`
	Statistic service.StatisticOut   `swaggerignore:"true"`
}

type BotInfoRequest struct {
	BotUUID string `json:"bot_uuid"`
}
