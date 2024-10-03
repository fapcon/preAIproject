package storage

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type SQLAdapter interface {
	Create(ctx context.Context, entity utils.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, tableName string, condition utils.Condition, opts ...interface{}) error
	Update(ctx context.Context, entity utils.Tabler, condition utils.Condition, operation string, opts ...interface{}) error
}

type UserSubStorage struct {
	adapter SQLAdapter
}

func NewUserSubStorage(sqlAdapter SQLAdapter) *UserSubStorage {
	return &UserSubStorage{adapter: sqlAdapter}
}

func (u *UserSubStorage) Create(ctx context.Context, in models.UserSubDTO) error {
	return u.adapter.Create(ctx, &in)
}

func (u *UserSubStorage) Update(ctx context.Context, dto models.UserSubDTO) error {
	return u.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}

func (u *UserSubStorage) GetByID(ctx context.Context, in models.UserSubDTO) (models.UserSubDTO, error) {
	var list []models.UserSubDTO
	err := u.adapter.List(ctx, &list, "users_sub", utils.Condition{
		Equal: map[string]interface{}{"user_id": in.UserID, "sub_user_id": in.SubUserID, "deleted_at": nil},
	})
	if err != nil {
		return models.UserSubDTO{}, err
	}
	if len(list) < 1 {
		return models.UserSubDTO{}, fmt.Errorf("user_sub storage: GetByID not found")
	}
	return list[0], err
}

func (u *UserSubStorage) GetList(ctx context.Context) ([]models.UserSubDTO, error) {
	var list []models.UserSubDTO
	err := u.adapter.List(ctx, &list, "users_sub", utils.Condition{Equal: map[string]interface{}{"deleted_at": nil}})
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (u *UserSubStorage) Delete(ctx context.Context, in models.UserSubDTO) error {
	dto, err := u.GetByID(ctx, in)
	if err != nil {
		return err
	}

	dto.SetDeletedAt(time.Now())

	return u.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
}
