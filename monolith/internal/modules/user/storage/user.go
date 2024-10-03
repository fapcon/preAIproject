package storage

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/cache"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type SQLAdapter interface {
	Create(ctx context.Context, entity utils.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, tableName string, condition utils.Condition, opts ...interface{}) error
	Update(ctx context.Context, entity utils.Tabler, condition utils.Condition, operation string, opts ...interface{}) error
}

type UserStorage struct {
	adapter SQLAdapter
	cache   cache.Cache
}

const (
	userCacheKey     = "user:%d"
	userCacheTTL     = 15
	userCacheTimeout = 50
)

func NewUserStorage(sqlAdapter SQLAdapter, cache cache.Cache) *UserStorage {
	return &UserStorage{adapter: sqlAdapter, cache: cache}
}

func (s *UserStorage) Create(ctx context.Context, u models.UserDTO) (int, error) {
	err := s.adapter.Create(ctx, &u)

	return 0, err
}

func (s *UserStorage) Update(ctx context.Context, u models.UserDTO) error {
	err := s.adapter.Update(ctx, &u, utils.Condition{
		Equal: map[string]interface{}{
			"id": u.GetID(),
		},
	}, utils.Update)
	if err != nil {
		return err
	}
	_ = s.cache.Expire(ctx, fmt.Sprintf(userCacheKey, u.GetID()), 0)

	return nil
}

func (s *UserStorage) GetByID(ctx context.Context, userID int) (models.UserDTO, error) {
	var dto models.UserDTO
	var err error

	timeout, cancel := context.WithTimeout(context.Background(), userCacheTimeout*time.Millisecond)
	defer cancel()
	err = s.cache.Get(timeout, fmt.Sprintf(userCacheKey, userID), &dto)
	if err == nil {
		return dto, nil
	}

	var list []models.UserDTO
	err = s.adapter.List(ctx, &list, dto.TableName(), utils.Condition{
		Equal: map[string]interface{}{
			"id": userID,
		},
	})
	if err != nil {
		return models.UserDTO{}, err
	}
	if len(list) < 1 {
		return models.UserDTO{}, fmt.Errorf("user storage: GetByID not found")
	}

	go func() {
		timeout, cancel = context.WithTimeout(context.Background(), userCacheTimeout*time.Millisecond)
		defer cancel()
		s.cache.Set(timeout, fmt.Sprintf(userCacheKey, userID), list[0], userCacheTTL*time.Minute)
	}()

	return list[0], nil
}

func (s *UserStorage) GetByIDs(ctx context.Context) ([]models.UserDTO, error) {
	panic("implement me")
}

func (s *UserStorage) GetByFilter(ctx context.Context) ([]models.UserDTO, error) {
	panic("implement me")
}

func (s *UserStorage) GetByEmail(ctx context.Context, email string) (models.UserDTO, error) {
	var users []models.UserDTO
	err := s.adapter.List(ctx, &users, "users", utils.Condition{
		Equal: map[string]interface{}{
			"email": email,
		},
	})
	if err != nil {
		return models.UserDTO{}, err
	}
	if len(users) < 1 {
		return models.UserDTO{}, fmt.Errorf("user with email %s not found", email)
	}

	return users[0], nil
}
