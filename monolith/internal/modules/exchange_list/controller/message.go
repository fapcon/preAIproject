package controller

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE json
type ExchangeAddRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

type ExchangeResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
}

type ExchangeListResponse struct {
	Success   bool                  `json:"success"`
	ErrorCode int                   `json:"error_code,omitempty"`
	Data      []models.ExchangeList `json:"data"`
}

type ExchangeDeleteRequest struct {
	ID int `json:"id"`
}
