package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type StrategyCreateRequest struct {
	ID          int          `json:"id" mapper:"id"`
	Name        string       `json:"name" mapper:"name"`
	UUID        string       `json:"uuid" mapper:"uuid"`
	Description string       `json:"description" mapper:"description"`
	ExchangeID  int          `json:"exchange_id" mapper:"exchange_id"`
	Bots        []models.Bot `json:"bots" mapper:"bots"`
}

type StrategyCreateResponse struct {
	StrategyID int `json:"strategy_id"`
	ErrorCode  int `json:"error_code"`
}

type StrategyUpdateRequest struct {
	Strategy models.Strategy `json:"strategy"`
	Fields   []int           `json:"fields"`
}

type StrategyDefaultResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type StrategyResponse struct {
	Strategy  *models.Strategy `json:"strategy"`
	ErrorCode int              `json:"error_code"`
}

type StrategiesResponse struct {
	Strategy  []models.Strategy `json:"strategy"`
	ErrorCode int               `json:"error_code"`
}
