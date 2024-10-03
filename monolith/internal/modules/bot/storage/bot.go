package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/orm/db/adapter"
	"time"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/cache"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

const (
	botCacheKey = "bot:%s"
	botCacheTTL = 60
)

var NotFound = fmt.Errorf("bot storage: GetByID not found")
var NotFoundUUID = fmt.Errorf("bot storage: GetByUUID not found")

type BotStorage struct {
	adapter adapter.SQLAdapter
	cache   cache.Cache
}

func NewBotStorage(adapter adapter.SQLAdapter, cache cache.Cache) *BotStorage {
	return &BotStorage{adapter: adapter, cache: cache}
}

func (s *BotStorage) Create(ctx context.Context, bot models.BotDTO) (int, error) {
	err := s.adapter.Create(ctx, &bot)

	return 0, err
}

func (s *BotStorage) Update(ctx context.Context, bot models.BotDTO) error {
	updateTime(&bot)
	err := s.adapter.Update(ctx, &bot, utils.Condition{
		Equal: map[string]interface{}{
			"id": bot.GetID(),
		},
	}, utils.Update)
	if err != nil {
		return err
	}
	_ = s.cache.Expire(ctx, fmt.Sprintf(botCacheKey, bot.GetUUID()), 0)

	return nil
}

func (s *BotStorage) UpdateByUUID(ctx context.Context, bot models.BotDTO) error {
	updateTime(&bot)
	err := s.adapter.Update(ctx, &bot, utils.Condition{
		Equal: map[string]interface{}{
			"uuid": bot.GetUUID(),
		},
	}, utils.Update)
	if err != nil {
		return err
	}
	_ = s.cache.Expire(ctx, fmt.Sprintf(botCacheKey, bot.GetUUID()), 0)

	return nil
}

func (s *BotStorage) GetByID(ctx context.Context, botID int) (models.BotDTO, error) {
	var bots []models.BotDTO
	err := s.adapter.List(ctx, &bots, "bot", utils.Condition{
		Equal: map[string]interface{}{
			"id": botID,
		},
	})
	if err != nil {
		return models.BotDTO{}, err
	}
	if len(bots) < 1 {
		return models.BotDTO{}, NotFound
	}

	return bots[0], nil
}

func (s *BotStorage) GetByUUID(ctx context.Context, uuid string) (models.BotDTO, error) {
	var res models.BotDTO
	var err error
	timeout, cancel := context.WithTimeout(context.Background(), botCacheTTL*time.Millisecond)
	defer cancel()
	err = s.cache.Get(timeout, fmt.Sprintf(botCacheKey, uuid), &res)
	if err == nil {
		return res, nil
	}
	var bots []models.BotDTO
	err = s.adapter.List(ctx, &bots, "bot", utils.Condition{
		Equal: map[string]interface{}{
			"uuid": uuid,
		},
	})
	if err != nil {
		return models.BotDTO{}, err
	}
	if len(bots) < 1 {
		return models.BotDTO{}, NotFoundUUID
	}

	return bots[0], nil
}

func (s *BotStorage) GetDraft(ctx context.Context, userID int) (models.BotDTO, error) {
	var bots []models.BotDTO
	var err error
	err = s.adapter.List(ctx, &bots, "bot", utils.Condition{
		Equal: map[string]interface{}{
			"user_id":    userID,
			"active":     types.NullBool{},
			"deleted_at": types.NullTime{},
		},
		NotEqual: map[string]interface{}{},
		Order: []*utils.Order{
			{
				Field: "id",
				Asc:   false,
			},
		},
	})
	if err != nil {
		return models.BotDTO{}, err
	}
	if len(bots) < 1 {
		return models.BotDTO{}, NotFoundUUID
	}

	return bots[0], nil
}

func (s *BotStorage) GetList(ctx context.Context, condition utils.Condition) ([]models.BotDTO, error) {
	var bots []models.BotDTO
	var err error
	err = s.adapter.List(ctx, &bots, "bot", condition)
	if err != nil {
		return nil, err
	}
	if len(bots) < 1 {
		return nil, NotFoundUUID
	}

	return bots, nil
}

func (s *BotStorage) GetByIDs(ctx context.Context, ids []int) ([]models.BotDTO, error) {
	panic("implement me")
}

func (s *BotStorage) GetByFilter(ctx context.Context) ([]models.BotDTO, error) {
	panic("implement me")
}

func updateTime(bot *models.BotDTO) {
	bot.SetUpdatedAt(time.Now())
}
