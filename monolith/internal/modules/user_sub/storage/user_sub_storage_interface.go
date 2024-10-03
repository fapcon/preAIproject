package storage

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type UserSuber interface {
	Create(ctx context.Context, in models.UserSubDTO) error
	Update(ctx context.Context, in models.UserSubDTO) error
	GetByID(ctx context.Context, in models.UserSubDTO) (models.UserSubDTO, error)
	GetList(ctx context.Context) ([]models.UserSubDTO, error)
	Delete(ctx context.Context, in models.UserSubDTO) error
}
