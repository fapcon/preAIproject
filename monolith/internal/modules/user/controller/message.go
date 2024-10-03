package controller

import "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

//go:generate easytags $GOFILE
type RegisterRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
	IdempotencyKey string `json:"idempotency_key"`
}

type ProfileResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
	Data      Data `json:"data"`
}

type Data struct {
	Message string      `json:"message,omitempty"`
	User    models.User `json:"user,omitempty"`
}
