package storage

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type SQLAdapter interface {
	Create(ctx context.Context, entity utils.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, tableName string, condition utils.Condition, opts ...interface{}) error
	Update(ctx context.Context, entity utils.Tabler, condition utils.Condition, operation string, opts ...interface{}) error
}

type Comment struct {
	adapter SQLAdapter
}

func NewComment(sqlAdapter SQLAdapter) Commenter {
	return &Comment{adapter: sqlAdapter}
}

func (c *Comment) Create(ctx context.Context, dto models.CommentDTO) error {
	return c.adapter.Create(ctx, &dto)
}

func (c *Comment) Update(ctx context.Context, dto models.CommentDTO) error {
	dto.SetUpdatedAt(time.Now())
	return c.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (c *Comment) GetByID(ctx context.Context, commentID int) (models.CommentDTO, error) {
	var list []models.CommentDTO
	err := c.adapter.List(ctx, &list, "comments", utils.Condition{
		Equal: map[string]interface{}{"id": commentID},
	})
	if err != nil {
		return models.CommentDTO{}, err
	}
	if len(list) < 1 {
		return models.CommentDTO{}, fmt.Errorf("comment storage: GetByID not found")
	}
	return list[0], err
}

func (c *Comment) GetList(ctx context.Context) ([]models.CommentDTO, error) {
	var list []models.CommentDTO
	err := c.adapter.List(ctx, &list, "comments", utils.Condition{
		Equal: map[string]interface{}{
			"deleted_at": types.NullTime{},
		},
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c *Comment) Delete(ctx context.Context, commentID int) error {
	dto, err := c.GetByID(ctx, commentID)
	if err != nil {
		return err
	}

	dto.SetDeletedAt(time.Now())

	return c.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}
