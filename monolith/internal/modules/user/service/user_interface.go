package service

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Userer
type Userer interface {
	Create(ctx context.Context, in UserCreateIn) UserCreateOut
	Update(ctx context.Context, in UserUpdateIn) UserUpdateOut
	VerifyEmail(ctx context.Context, in UserVerifyEmailIn) UserUpdateOut
	ChangePassword(ctx context.Context, in ChangePasswordIn) ChangePasswordOut
	GetByEmail(ctx context.Context, in GetByEmailIn) UserOut
	GetByPhone(ctx context.Context, in GetByPhoneIn) UserOut
	GetByID(ctx context.Context, in GetByIDIn) UserOut
	GetByIDs(ctx context.Context, in GetByIDsIn) UsersOut
	ResetPassword(ctx context.Context, in ResetPasswordIn) ResetPasswordOut
	SendResetCodeEmail(ctx context.Context, in SendResetCodeEmailIn) SendResetCodeEmailOut
}

const UserSuperMan = 99999

const (
	UserTypeDefault = iota + 1
)

type ChangePasswordIn struct {
	Email              string `json:"email"`
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type ChangePasswordOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type GetByIDIn struct {
	UserID int `json:"user_id"`
}

type GetByIDsIn struct {
	UserIDs []int `json:"user_i_ds"`
}

type UserOut struct {
	User      *models.User `json:"user"`
	ErrorCode int          `json:"error_code"`
}

type UsersOut struct {
	User      []models.User `json:"user"`
	ErrorCode int           `json:"error_code"`
}

type GetByEmailIn struct {
	Email string `json:"email"`
}

type GetByPhoneIn struct {
	Phone string `json:"phone"`
}

type UserCreateIn struct {
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           int    `json:"role"`
	IdempotencyKey string `json:"idempotency_key"`
}

type UserCreateOut struct {
	UserID    int `json:"user_id"`
	ErrorCode int `json:"error_code"`
}

type UserUpdateIn struct {
	User   models.User `json:"user"`
	Fields []int       `json:"fields"`
}

type UserUpdateOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type UOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type UserVerifyEmailIn struct {
	UserID int `json:"user_id"`
}

type ResetPasswordIn struct {
	Email              string `json:"email"`
	ResetCode          int    `json:"reset_code"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type ResetPasswordOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type SendResetCodeEmailIn struct {
	Email string `json:"email"`
}

type SendResetCodeEmailOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}
