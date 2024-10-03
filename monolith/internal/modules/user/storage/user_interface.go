package storage

import (
	"context"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate mockgen -source user_interface.go -destination=mocks/user_mock.go

type Userer interface {
	Create(ctx context.Context, u models.UserDTO) (int, error)
	Update(ctx context.Context, u models.UserDTO) error
	GetByID(ctx context.Context, userID int) (models.UserDTO, error)
	GetByIDs(ctx context.Context) ([]models.UserDTO, error)
	GetByEmail(ctx context.Context, email string) (models.UserDTO, error)
	GetByFilter(ctx context.Context) ([]models.UserDTO, error)
}
