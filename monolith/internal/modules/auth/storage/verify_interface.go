package storage

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Verifier
type Verifier interface {
	GetByEmail(ctx context.Context, email, hash string) (models.EmailVerifyDTO, error)
	GetByUserID(ctx context.Context, userID int) (models.EmailVerifyDTO, error)
	Verify(ctx context.Context, userID int) error
	VerifyEmail(ctx context.Context, email, hash string) error
	Create(ctx context.Context, email, hash string, userID int) error
}
