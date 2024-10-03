package service

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type ExchangeUserKeyAddIn struct {
	ExchangeID int
	Label      string
	UserID     int
	APIKey     string
	SecretKey  string
}

type ExchangeOut struct {
	ErrorCode int
	Success   bool
}

type ExchangeUserListIn struct {
	UserID int
}

type ExchangeUserListOut struct {
	ErrorCode int
	Data      []models.ExchangeUserKey
}
