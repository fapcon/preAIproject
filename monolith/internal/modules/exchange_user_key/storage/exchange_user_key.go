package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/cache"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

var NotFound = fmt.Errorf("exchange user storage: GetByID not found")

const (
	exchangeUserKeyIDCacheKey     = "exchangeUserKey:ID:%d"
	exchangeUserKeyUserIDCacheKey = "exchangeUserKey:UserID:%d"
	apiKeyCacheTTL                = 24
	exchangeUserKeyTimeout        = 50
)

type ExchangeUserKey struct {
	adapter adapter.SQLAdapter
	cache   cache.Cache
}

func NewExchangeUserKey(adapter adapter.SQLAdapter, cache cache.Cache) *ExchangeUserKey {
	return &ExchangeUserKey{adapter: adapter, cache: cache}
}

func (e *ExchangeUserKey) Create(ctx context.Context, dto models.ExchangeUserKeyDTO) error {
	err := e.invalidateCache(ctx, dto)
	if err != nil {
		return err
	}

	return e.adapter.Create(ctx, &dto)
}

func (e *ExchangeUserKey) GetByUserID(ctx context.Context, userID int) ([]models.ExchangeUserKeyDTO, error) {
	var dtos []models.ExchangeUserKeyDTO
	var dto models.ExchangeUserKeyDTO
	var err error

	timeout, cancel := context.WithTimeout(context.Background(), exchangeUserKeyTimeout*time.Millisecond)
	defer cancel()
	err = e.cache.Get(timeout, fmt.Sprintf(exchangeUserKeyUserIDCacheKey, userID), &dtos)
	if err == nil {
		return dtos, nil
	}
	err = e.adapter.List(ctx, &dtos, dto.TableName(), utils.Condition{
		Equal: map[string]interface{}{"user_id": userID, "deleted_at": nil},
	})
	if err != nil {
		return nil, err
	}
	if len(dtos) < 1 {
		return nil, NotFound
	}

	timeout, cancel = context.WithTimeout(context.Background(), exchangeUserKeyTimeout*time.Millisecond)
	defer cancel()
	e.cache.Set(timeout, fmt.Sprintf(exchangeUserKeyUserIDCacheKey, userID), dtos, apiKeyCacheTTL*time.Minute)

	return dtos, nil
}

func (e *ExchangeUserKey) Update(ctx context.Context, dto models.ExchangeUserKeyDTO) error {
	err := e.adapter.Update(ctx, &dto, utils.Condition{
		Equal: map[string]interface{}{
			"id": dto.GetID(),
		},
	}, utils.Update)
	if err != nil {
		return err
	}

	err = e.invalidateCache(ctx, dto)
	if err != nil {
		return err
	}

	return nil
}

func (e *ExchangeUserKey) invalidateCache(ctx context.Context, dto models.ExchangeUserKeyDTO) error {
	err := e.cache.Expire(ctx, fmt.Sprintf(exchangeUserKeyIDCacheKey, dto.ID), 0)
	err = e.cache.Expire(ctx, fmt.Sprintf(exchangeUserKeyUserIDCacheKey, dto.UserID), 0)

	return err
}

func (e *ExchangeUserKey) GetList(ctx context.Context, condition utils.Condition) ([]models.ExchangeUserKeyDTO, error) {
	var list []models.ExchangeUserKeyDTO
	var dto models.ExchangeUserKeyDTO
	err := e.adapter.List(ctx, &list, dto.TableName(), condition)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (e *ExchangeUserKey) Delete(ctx context.Context, exchangeUserID int) error {
	dto, err := e.GetByID(ctx, exchangeUserID)
	if err != nil {
		return err
	}

	dto.SetDeletedAt(time.Now())

	err = e.adapter.Update(
		ctx,
		&dto,
		utils.Condition{
			Equal: map[string]interface{}{"id": dto.GetID()},
		},
		utils.Update,
	)
	if err != nil {
		return err
	}

	return e.invalidateCache(ctx, dto)
}

func (e *ExchangeUserKey) GetByID(ctx context.Context, exchangeUserKeyID int) (models.ExchangeUserKeyDTO, error) {
	var dto models.ExchangeUserKeyDTO
	var err error

	timeout, cancel := context.WithTimeout(context.Background(), exchangeUserKeyTimeout*time.Millisecond)
	defer cancel()
	err = e.cache.Get(timeout, fmt.Sprintf(exchangeUserKeyIDCacheKey, exchangeUserKeyID), &dto)
	if err == nil {
		return dto, nil
	}

	var list []models.ExchangeUserKeyDTO
	err = e.adapter.List(ctx, &list, dto.TableName(), utils.Condition{
		Equal: map[string]interface{}{"id": exchangeUserKeyID},
	})
	if err != nil {
		return models.ExchangeUserKeyDTO{}, err
	}
	if len(list) < 1 {
		return models.ExchangeUserKeyDTO{}, fmt.Errorf("exchange list storage: GetByID not found")
	}

	timeout, cancel = context.WithTimeout(context.Background(), exchangeUserKeyTimeout*time.Millisecond)
	defer cancel()
	e.cache.Set(timeout, fmt.Sprintf(exchangeUserKeyIDCacheKey, exchangeUserKeyID), list[0], apiKeyCacheTTL*time.Minute)

	return list[0], err
}
