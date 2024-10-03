package storage

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Poster
type Poster interface {
	Create(ctx context.Context, dto models.PostDTO) error
	GetById(ctx context.Context, id int) (models.PostDTO, error)
	Update(ctx context.Context, dto models.PostDTO) error
	GetList(ctx context.Context) ([]models.PostDTO, error)
}
