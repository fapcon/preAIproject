package service

import (
	"context"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/strategy/storage"
)

type StrategyService struct {
	storage storage.Strateger
	logger  *zap.Logger
	bot     service.Boter
}

func NewStrategyService(storage storage.Strateger, logger *zap.Logger, bot service.Boter) *StrategyService {
	return &StrategyService{storage: storage, logger: logger, bot: bot}
}

func (s StrategyService) Create(ctx context.Context, in StrategyCreateIn) StrategyCreateOut {
	var dto models.StrategyDTO
	var err error

	dto.SetName(in.Name).
		SetDescription(in.Description).
		SetUUID(in.UUID).
		SetExchangeID(in.ExchangeID).
		SetBots(in.Bots) //Боты записаны по айди

	strategyID, err := s.storage.Create(ctx, dto)
	if err != nil {
		if v, ok := err.(*pq.Error); ok && v.Code == "23505" {
			s.logger.Error("strategy: Create err", zap.Error(err))
			return StrategyCreateOut{
				ErrorCode: errors.StrategyServiceStrategyAlreadyExists,
			}
		}
		s.logger.Error("strategy: Create err", zap.Error(err))
		return StrategyCreateOut{
			ErrorCode: errors.StrategyServiceCreateErr}
	}

	return StrategyCreateOut{StrategyID: strategyID}
}

func (s StrategyService) Update(ctx context.Context, in StrategyUpdateIn) StrategyUpdateOut {
	var dto models.StrategyDTO
	var err error

	dto.SetName(in.Strategy.Name).
		SetDescription(in.Strategy.Description).
		SetUUID(in.Strategy.UUID).
		SetExchangeID(in.Strategy.ExchangeID).
		SetBots(in.Strategy.Bots) //Боты записаны по айди

	err = s.storage.Update(ctx, dto)
	if err != nil {
		s.logger.Error("strategy: Update err", zap.Error(err))
		return StrategyUpdateOut{
			Success:   false,
			ErrorCode: errors.StrategyServiceUpdateErr,
		}
	}

	return StrategyUpdateOut{
		Success: true,
	}
}

func (s StrategyService) GetByID(ctx context.Context, in StrategyGetByIDIn) StrategyOut {
	var err error
	var dto models.StrategyDTO

	dto, err = s.storage.GetByID(ctx, in.ID)
	if err != nil {
		s.logger.Error("strategy: GetByID err", zap.Error(err))
		return StrategyOut{
			ErrorCode: errors.StrategyServiceRetrieveErr,
		}
	}

	var bots []models.Bot
	for _, botIDRaw := range dto.Bots {
		botID, err := strconv.Atoi(botIDRaw)
		if err != nil {
			s.logger.Error("strategy: BotRetrieve err", zap.Error(err))
			return StrategyOut{
				ErrorCode: errors.StrategyServiceRetrieveErr,
			}
		}
		botIn := service.BotGetIn{
			ID: botID,
		}
		bot := s.bot.Get(ctx, botIn)
		bots = append(bots, bot.Bot)
	}

	return StrategyOut{
		Strategy: &models.Strategy{
			ID:          dto.ID,
			Name:        dto.Name,
			UUID:        dto.UUID,
			Description: dto.Description,
			ExchangeID:  dto.ExchangeID,
			Bots:        bots,
		},
	}
}

func (s StrategyService) GetByName(ctx context.Context, in StrategyGetByNameIn) StrategyOut {
	var err error
	var dto models.StrategyDTO

	dto, err = s.storage.GetByName(ctx, in.Name)
	if err != nil {
		s.logger.Error("strategy: GetByName err", zap.Error(err))
		return StrategyOut{
			ErrorCode: errors.StrategyServiceRetrieveErr,
		}
	}

	var bots []models.Bot
	for _, botIDRaw := range dto.Bots {
		botID, err := strconv.Atoi(botIDRaw)
		if err != nil {
			s.logger.Error("strategy: BotRetrieve err", zap.Error(err))
			return StrategyOut{
				ErrorCode: errors.StrategyServiceRetrieveErr,
			}
		}
		botIn := service.BotGetIn{
			ID: botID,
		}
		bot := s.bot.Get(ctx, botIn)
		bots = append(bots, bot.Bot)
	}

	return StrategyOut{
		Strategy: &models.Strategy{
			ID:          dto.ID,
			Name:        dto.Name,
			UUID:        dto.UUID,
			Description: dto.Description,
			ExchangeID:  dto.ExchangeID,
			Bots:        bots,
		},
	}
}

func (s StrategyService) GetList(ctx context.Context) StrategiesOut {
	var strategies []models.Strategy
	var err error

	dto, err := s.storage.GetList(ctx) //TODO: refactor algorithm of exctraction
	if err != nil {
		s.logger.Error("strategy: GetList err", zap.Error(err))
		return StrategiesOut{
			ErrorCode: errors.StrategyServiceRetrieveErr,
		}
	}

	for _, strategyDTO := range dto {
		var bots []models.Bot

		for _, botIDRaw := range strategyDTO.Bots {
			botID, err := strconv.Atoi(botIDRaw)
			if err != nil {
				s.logger.Error("strategy: BotRetrieve err", zap.Error(err))
				return StrategiesOut{
					ErrorCode: errors.StrategyServiceRetrieveErr,
				}
			}
			botIn := service.BotGetIn{
				ID: botID,
			}
			bot := s.bot.Get(ctx, botIn)
			bots = append(bots, bot.Bot)
		}

		strategy := models.Strategy{
			ID:          strategyDTO.ID,
			Name:        strategyDTO.Name,
			UUID:        strategyDTO.UUID,
			Description: strategyDTO.Description,
			ExchangeID:  strategyDTO.ExchangeID,
			Bots:        bots,
		}
		strategies = append(strategies, strategy)
	}

	return StrategiesOut{
		Strategy: strategies,
	}
}

func (s StrategyService) Delete(ctx context.Context, in StrategyDeleteIn) StrategyDeleteOut {
	var err error

	err = s.storage.Delete(ctx, in.ID)
	if err != nil {
		if v, ok := err.(*pq.Error); ok && v.Code == "42703" {
			s.logger.Error("strategy: Delete err", zap.Error(err))
			return StrategyDeleteOut{
				ErrorCode: errors.StrategyServiceStrategyDoesntExist,
			}
		}
		s.logger.Error("strategy: Delete err", zap.Error(err))
		return StrategyDeleteOut{
			ErrorCode: errors.StrategyServiceDeleteErr,
		}
	}

	return StrategyDeleteOut{
		Success: true,
	}
}
