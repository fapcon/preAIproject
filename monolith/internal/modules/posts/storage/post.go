package storage

import (
	"context"
	"fmt"

	"gitlab.com/golight/orm/utils"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/cache"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type SQLAdapter interface {
	Create(ctx context.Context, entity utils.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, tableName string, condition utils.Condition, opts ...interface{}) error
	Update(ctx context.Context, entity utils.Tabler, condition utils.Condition, operation string, opts ...interface{}) error
}

var NotFound = fmt.Errorf("bot storage: GetByID not found")

type PostStorage struct {
	adapter SQLAdapter
	cache   cache.Cache
}

func NewPostStorage(sqlAdapter SQLAdapter, cache cache.Cache) *PostStorage {
	return &PostStorage{adapter: sqlAdapter, cache: cache}
}

func (u *PostStorage) Create(ctx context.Context, dto models.PostDTO) error {
	err := u.adapter.Create(ctx, &dto)
	return err
}

func (u *PostStorage) Update(ctx context.Context, dto models.PostDTO) error {
	err := u.adapter.Update(ctx, &dto, utils.Condition{
		Equal: map[string]interface{}{
			"id": dto.ID,
		},
	}, utils.Update)
	return err
}

func (u *PostStorage) GetById(ctx context.Context, id int) (models.PostDTO, error) {
	var post []models.PostDTO
	err := u.adapter.List(ctx, &post, "post", utils.Condition{
		Equal: map[string]interface{}{
			"id":         id,
			"deleted_at": types.NullTime{},
		},
	})
	if err != nil {
		return models.PostDTO{}, err
	}
	if len(post) < 1 {
		return models.PostDTO{}, NotFound
	}
	return post[0], nil
}

func (u *PostStorage) GetList(ctx context.Context) ([]models.PostDTO, error) {
	var post []models.PostDTO
	err := u.adapter.List(ctx, &post, "post", utils.Condition{
		Equal: map[string]interface{}{
			"deleted_at": types.NullTime{},
		},
	})
	if err != nil {
		return []models.PostDTO{}, err
	}
	if len(post) < 1 {
		return []models.PostDTO{}, NotFound
	}
	return post, nil
}
