package service

import (
	"github.com/shopspring/decimal"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE json
type StatisticUpdateIn struct {
	BotUUID string          `json:"bot_uuid"`
	UserID  int             `json:"user_id"`
	Price   decimal.Decimal `json:"price"`
	Amount  decimal.Decimal `json:"amount"`
	Status  int             `json:"status"`
	Side    int             `json:"side"`
}

type StatisticUpdateOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type StatisticIn struct {
	BotUUID string `json:"bot_uuid"`
}

type StatisticOut struct {
	ErrorCode int                  `json:"error_code"`
	Statistic models.BotStatistics `json:"statistic"`
}

type UserStatisticIn struct {
	UserID int `json:"user_id"`
}

type UserStatisticOut struct {
	ErrorCode int                    `json:"error_code"`
	Statistic []models.BotStatistics `json:"statistic"`
}

type StatisticDeleteOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}
