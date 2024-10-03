package storage

import (
	"context"
	"fmt"
	"gitlab.com/golight/boilerplate/infrastructure/cache"
	"gitlab.com/golight/orm/db/adapter"
	"gitlab.com/golight/orm/utils"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"time"
)

const (
	botCacheKey = "botStatistic:%s"
)

var NotFoundID = fmt.Errorf("bot statistics storage: GetByID not found")
var NotFoundUUID = fmt.Errorf("bot statistics storage: GetByUUID not found ")
var NotFoundUserID = fmt.Errorf("bot statistics storage: GetByUserID not found")

type StatisticStorage struct {
	adapter adapter.SQLAdapter
	cache   cache.Cache
}

func NewStatisticStorage(adapter adapter.SQLAdapter, cache cache.Cache) *StatisticStorage {
	return &StatisticStorage{adapter: adapter, cache: cache}
}

func (s *StatisticStorage) Create(ctx context.Context, botStatistics models.BotStatisticsDTO) error {
	return s.adapter.Create(ctx, &botStatistics)
}

func (s *StatisticStorage) Update(ctx context.Context, botStatistics models.BotStatisticsDTO) error {
	botStatistics.SetUpdatedAt(time.Now())
	err := s.adapter.Update(ctx, &botStatistics, utils.Condition{
		Equal: map[string]interface{}{
			"id": botStatistics.GetID(),
		},
	}, utils.Update)
	if err != nil {
		return err
	}

	_ = s.cache.Expire(ctx, fmt.Sprintf(botCacheKey, botStatistics.GetUUID()), 0)

	return nil
}

func (s *StatisticStorage) GetByID(ctx context.Context, id int) (models.BotStatisticsDTO, error) {
	var botStatistics []models.BotStatisticsDTO
	err := s.adapter.List(ctx, &botStatistics, "bot_statistics", utils.Condition{
		Equal: map[string]interface{}{
			"id":         id,
			"deleted_at": nil,
		},
	})
	if err != nil {
		return models.BotStatisticsDTO{}, err
	}
	if len(botStatistics) < 1 {
		return models.BotStatisticsDTO{}, NotFoundID
	}
	return botStatistics[0], nil
}

func (s *StatisticStorage) GetByUUID(ctx context.Context, botUUID string) (models.BotStatisticsDTO, error) {
	var botStatistics []models.BotStatisticsDTO
	err := s.adapter.List(ctx, &botStatistics, "bot_statistics", utils.Condition{
		Equal: map[string]interface{}{
			"bot_uuid":   botUUID,
			"deleted_at": nil,
		},
	})
	if err != nil {
		return models.BotStatisticsDTO{}, err
	}
	if len(botStatistics) < 1 {
		return models.BotStatisticsDTO{}, NotFoundUUID
	}
	return botStatistics[0], nil
}

func (s *StatisticStorage) GetByUserID(ctx context.Context, userID int) ([]models.BotStatisticsDTO, error) {
	var botStatistics []models.BotStatisticsDTO
	err := s.adapter.List(ctx, &botStatistics, "bot_statistics", utils.Condition{
		Equal: map[string]interface{}{
			"user_id":    userID,
			"deleted_at": nil,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(botStatistics) < 1 {
		return nil, NotFoundUserID
	}
	return botStatistics, nil
}

func (s *StatisticStorage) Delete(ctx context.Context, uuid string) error {
	botStatistic, err := s.GetByUUID(ctx, uuid)
	if err != nil {
		return err
	}

	botStatistic.SetDeletedAt(time.Now())

	return s.adapter.Update(
		ctx, &botStatistic, utils.Condition{
			Equal: map[string]interface{}{"id": botStatistic.GetID()},
		}, utils.Update)
}
