package service

import (
	"context"
	"github.com/ptflp/gomapper"
	"github.com/shopspring/decimal"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/statistics/storage"
	"time"
)

type Statistic struct {
	statistic storage.Statisticer
}

func NewStatistic(storage storage.Statisticer) *Statistic {
	return &Statistic{statistic: storage}
}

func (s *Statistic) UpdateStatistic(ctx context.Context, in StatisticUpdateIn) StatisticUpdateOut {
	var (
		botStatistic models.BotStatisticsDTO
		err          error
	)
	//получаем статистику текущего бота
	botStatistic, err = s.statistic.GetByUUID(ctx, in.BotUUID)
	//если статистики по текущему боту нет, то создаем ее
	if err == storage.NotFoundUUID {
		err = s.statistic.Create(ctx, models.BotStatisticsDTO{
			BotUUID:       in.BotUUID,
			UserID:        in.UserID,
			Profitability: decimal.Zero,
			OrderCount:    0,
			CreatedAt:     time.Now(),
		})
		if err != nil {
			return StatisticUpdateOut{
				ErrorCode: errors.StatisticServiceUpdateErr,
			}
		}
		return s.UpdateStatistic(ctx, in)
	}
	if err != nil {
		return StatisticUpdateOut{
			ErrorCode: errors.StatisticServiceUpdateErr,
		}
	}
	//считаем прибыльность только для тех ордеров, которые имеют статус filled
	//для остальных просто увеличиваем счетчик
	if in.Status == service.OrderStatusFilled {
		switch in.Side {
		case service.SideBuy:
			botStatistic.Profitability = botStatistic.Profitability.Sub(in.Amount)
		case service.SideSell, service.SideCancel:
			botStatistic.Profitability = botStatistic.Profitability.Add(in.Amount)
		}
	}

	botStatistic.OrderCount++

	err = s.statistic.Update(ctx, botStatistic)
	if err != nil {
		return StatisticUpdateOut{
			ErrorCode: errors.StatisticServiceUpdateErr,
		}
	}

	return StatisticUpdateOut{
		Success: true,
	}
}

func (s *Statistic) GetBotStatistic(ctx context.Context, in StatisticIn) StatisticOut {
	statistic, err := s.statistic.GetByUUID(ctx, in.BotUUID)
	if err != nil {
		return StatisticOut{
			ErrorCode: errors.StatisticServiceGetBotStatisticError,
			Statistic: models.BotStatistics{},
		}
	}
	return StatisticOut{
		Statistic: models.BotStatistics{
			ID:            statistic.ID,
			BotUUID:       statistic.BotUUID,
			UserID:        statistic.UserID,
			Profitability: statistic.Profitability,
			OrderCount:    statistic.OrderCount,
		},
	}
}

func (s *Statistic) GetUserBotStatistic(ctx context.Context, in UserStatisticIn) UserStatisticOut {
	statisticDTO, err := s.statistic.GetByUserID(ctx, in.UserID)
	var statistic []models.BotStatistics
	if err != nil {
		return UserStatisticOut{
			ErrorCode: errors.StatisticServiceGetUserStatisticError,
		}
	}
	err = gomapper.MapStructs(&statistic, &statisticDTO)
	if err != nil {
		return UserStatisticOut{
			ErrorCode: errors.StatisticServiceGetUserStatisticError,
		}
	}
	return UserStatisticOut{
		Statistic: statistic,
	}
}

func (s *Statistic) DeleteBotStatistic(ctx context.Context, in StatisticIn) StatisticDeleteOut {
	err := s.statistic.Delete(ctx, in.BotUUID)
	if err != nil {
		return StatisticDeleteOut{
			ErrorCode: errors.StatisticServiceDeleteErr,
		}
	}
	return StatisticDeleteOut{
		Success: true,
	}
}
