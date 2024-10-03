package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE json
type TickerResponse struct {
	Success   bool                    `json:"success"`
	ErrorCode int                     `json:"error_code,omitempty"`
	Data      []models.ExchangeTicker `json:"data"`
}
