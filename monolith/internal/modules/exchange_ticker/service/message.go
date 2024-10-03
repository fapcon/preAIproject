package service

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type GetTickerOut struct {
	ErrorCode int
	Data      []models.ExchangeTicker
}
