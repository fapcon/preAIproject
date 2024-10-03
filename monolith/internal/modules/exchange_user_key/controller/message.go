package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE json

type ExchangeResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
}

type ExchangeListResponse struct {
	Success   bool                  `json:"success"`
	ErrorCode int                   `json:"error_code,omitempty"`
	Data      []models.ExchangeList `json:"data"`
}

type ExchangeUserKeyAddRequest struct {
	ExchangeID int    `json:"exchange_id"`
	Label      string `json:"label"`
	APIKey     string `json:"api_key"`
	SecretKey  string `json:"secret_key"`
}

type ExchangeUserKeyDeleteRequest struct {
	ExchangeUserKeyID int `json:"exchange_user_key_id"`
}

type ExchangeUserKeyListResponse struct {
	Success   bool                     `json:"success"`
	ErrorCode int                      `json:"error_code,omitempty"`
	Data      []models.ExchangeUserKey `json:"data"`
}
