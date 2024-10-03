package service

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Strateger
type Strateger interface {
	Create(ctx context.Context, in StrategyCreateIn) StrategyCreateOut
	Update(ctx context.Context, in StrategyUpdateIn) StrategyUpdateOut
	GetByID(ctx context.Context, in StrategyGetByIDIn) StrategyOut
	GetByName(ctx context.Context, in StrategyGetByNameIn) StrategyOut
	GetList(ctx context.Context) StrategiesOut
	Delete(ctx context.Context, in StrategyDeleteIn) StrategyDeleteOut
}

type StrategyCreateIn struct {
	ID          int          `json:"id" mapper:"id"`
	Name        string       `json:"name" mapper:"name"`
	UUID        string       `json:"uuid" mapper:"uuid"`
	Description string       `json:"description" mapper:"description"`
	ExchangeID  int          `json:"exchange_id" mapper:"exchange_id"`
	Bots        []models.Bot `json:"bots" mapper:"bots"`
}

type StrategyCreateOut struct {
	StrategyID int `json:"strategy_id"`
	ErrorCode  int `json:"error_code"`
}

type StrategyUpdateIn struct {
	Strategy models.Strategy `json:"strategy"`
	Fields   []int           `json:"fields"`
}

type StrategyUpdateOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type StrategyGetByIDIn struct {
	ID int `json:"id"`
}

type StrategyGetByNameIn struct {
	Name string `json:"name"`
}

type StrategyDeleteIn struct {
	ID int `json:"id"`
}

type StrategyDeleteOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type StrategyOut struct {
	Strategy  *models.Strategy `json:"strategy"`
	ErrorCode int              `json:"error_code"`
}

type StrategiesOut struct {
	Strategy  []models.Strategy `json:"strategy"`
	ErrorCode int               `json:"error_code"`
}
