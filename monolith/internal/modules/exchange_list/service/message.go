package service

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type ExchangeListOut struct {
	ErrorCode int
	Data      []models.ExchangeList
}

type ExchangeAddIn struct {
	UserID      int
	Name        string
	Description string
	Slug        string
}

type ExchangeOut struct {
	ErrorCode int
	Success   bool
}
