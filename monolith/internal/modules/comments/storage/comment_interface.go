package storage

import (
	"context"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Commenter
type Commenter interface {
	Create(ctx context.Context, dto models.CommentDTO) error
	Update(ctx context.Context, dto models.CommentDTO) error
	GetByID(ctx context.Context, commentID int) (models.CommentDTO, error)
	GetList(ctx context.Context) ([]models.CommentDTO, error)
	Delete(ctx context.Context, commentID int) error
}
