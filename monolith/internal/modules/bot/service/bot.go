package service

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"time"

	"github.com/google/uuid"
	"github.com/ptflp/gomapper"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	sstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/storage"
	tiservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/storages"
)

type Bot struct {
	bot          sstorage.Boter
	strategyPair sstorage.StrategyPairer
	ticker       tiservice.ExchangeTicker
	logger       *zap.Logger
	conf         config.AppConf
}

func NewBot(storages *storages.Storages, ticker tiservice.ExchangeTicker, logger *zap.Logger, conf config.AppConf) *Bot {
	return &Bot{bot: storages.Bot, ticker: ticker, logger: logger, conf: conf, strategyPair: storages.StrategyPair}
}

func (b *Bot) Create(ctx context.Context, in BotCreateIn) BotOut {
	var dto models.BotDTO
	var err error
	dto, err = b.bot.GetDraft(ctx, in.UserID)
	if err == nil && err != sstorage.NotFoundUUID {
		return BotOut{
			Bot: models.Bot{
				ID:           dto.GetID(),
				Kind:         dto.GetKind(),
				UserID:       dto.GetUserID(),
				Name:         dto.GetName(),
				Description:  dto.GetDescription(),
				PairID:       dto.GetPairID(),
				ExchangeType: dto.GetExchangeType(),
				ExchangeID:   dto.GetExchangeID(),
				OrderType:    dto.GetOrderType(),
				SellPercent:  dto.GetSellPercent(),
				AssetType:    dto.GetAssetType(),
				UUID:         dto.GetUUID(),
				Active:       dto.GetActive(),
				CreatedAt:    dto.GetCreatedAt(),
				UpdatedAt:    dto.GetUpdatedAt(),
				DeletedAt:    dto.GetDeletedAt(),
			},
			Hooks: b.createHook(ctx, dto),
		}
	}
	var UUID uuid.UUID
	UUID, err = uuid.NewUUID()
	if err != nil {
		return BotOut{
			ErrorCode: errors.BotServiceUUIDGenerateErr,
			Bot:       models.Bot{},
		}
	}
	dto.
		SetUserID(in.UserID).
		SetUUID(UUID.String())
	_, err = b.bot.Create(ctx, dto)
	if err != nil {
		return BotOut{
			ErrorCode: errors.BotServiceCreateErr,
			Bot:       models.Bot{},
		}
	}
	dto, err = b.bot.GetByUUID(ctx, dto.GetUUID())
	if err != nil {
		return BotOut{
			ErrorCode: errors.BotServiceCreateErr,
			Bot:       models.Bot{},
		}
	}
	var strategy models.Bot
	err = gomapper.MapStructs(&strategy, &dto)
	if err != nil {
		return BotOut{
			ErrorCode: errors.BotServiceCreateErr,
			Bot:       models.Bot{},
		}
	}

	return BotOut{
		Bot:   strategy,
		Hooks: b.createHook(ctx, dto),
	}
}

func (b *Bot) Get(ctx context.Context, in BotGetIn) BotOut {
	var err error
	var dto models.BotDTO
	var strategy models.Bot
	switch {
	case in.ID > 0:
		dto, err = b.bot.GetByID(ctx, in.ID)
	case len(in.UUID) > 0:
		dto, err = b.bot.GetByUUID(ctx, in.UUID)
	}
	if err != nil {
		return BotOut{
			ErrorCode: errors.BotServiceGetErr,
			Bot:       strategy,
		}
	}

	err = gomapper.MapStructs(&strategy, &dto)
	if err != nil {
		return BotOut{
			ErrorCode: http.StatusInternalServerError,
			Bot:       strategy,
		}
	}
	return BotOut{
		Bot:   strategy,
		Hooks: b.createHook(ctx, dto),
	}
}

func (b *Bot) Toggle(ctx context.Context, in BotToggleIn) BOut {
	bot, err := b.bot.GetByUUID(ctx, in.UUID)
	if err != nil {
		return BOut{
			ErrorCode: errors.InternalError,
		}
	}

	if bot.UserID != in.UserID {
		return BOut{
			ErrorCode: errors.GeneralError,
		}
	}

	bot.SetActive(in.Active)
	err = b.bot.UpdateByUUID(ctx, bot)
	if err != nil {
		return BOut{
			ErrorCode: errors.InternalError,
		}
	}

	return BOut{
		Success: true,
	}
}

func (b *Bot) Delete(ctx context.Context, in BotDeleteIn) BOut {
	bot, err := b.bot.GetByUUID(ctx, in.UUID)
	if err != nil {
		return BOut{
			ErrorCode: errors.InternalError,
		}
	}

	bot.SetDeletedAt(time.Now())
	err = b.bot.Update(ctx, bot)
	if err != nil {
		return BOut{
			ErrorCode: errors.InternalError,
		}
	}

	return BOut{
		Success: true,
	}
}

func (b *Bot) Update(ctx context.Context, in BotUpdateIn) BOut {
	var dto models.BotDTO
	strategy := in.Bot
	if strategy.ExchangeID < 1 {
		strategy.ExchangeID = 1
	}
	// TODO: validate data
	dto.
		SetUUID(strategy.UUID).
		SetExchangeID(strategy.ExchangeID).
		SetExchangeType(strategy.ExchangeType).
		SetKind(strategy.Kind).
		SetName(strategy.Name).
		SetActive(strategy.Active).
		SetDescription(strategy.Description).
		SetOrderType(strategy.OrderType).
		SetPairID(strategy.PairID).
		SetSellPercent(strategy.SellPercent).
		SetCommissionPercent(strategy.CommissionPercent).
		SetFixedAmount(strategy.FixedAmount).
		SetLimitOrder(strategy.LimitOrder).
		SetLimitSellPercent(strategy.LimitSellPercent).
		SetLimitBuyPercent(strategy.LimitBuyPercent).
		SetAssetType(strategy.AssetType).
		SetAutoSell(strategy.AutoSell).
		SetAutoLimitSellPercent(strategy.AutoLimitSellPercent).
		SetOrderCountLimit(strategy.OrderCountLimit).
		SetOrderCount(strategy.OrderCount)
	err := b.bot.UpdateByUUID(ctx, dto)
	if err != nil {
		return BOut{
			Success:   false,
			ErrorCode: errors.BotServiceUpdateErr,
		}
	}

	pairsOut := b.addPairs(ctx, BotPairAdd{
		BotID: in.Bot.ID,
		Pairs: in.Pairs,
	})

	if pairsOut.ErrorCode != errors.NoError {
		return BOut{
			Success:   false,
			ErrorCode: errors.BotServiceUpdateErr,
		}
	}

	return BOut{
		Success: true,
	}
}

func (b *Bot) List(ctx context.Context, in BotListIn) BotListOut {
	dtos, err := b.bot.GetList(ctx, utils.Condition{
		Equal:    map[string]interface{}{"user_id": in.UserID, "deleted_at": nil},
		NotEqual: map[string]interface{}{"active": nil},
		Order: []*utils.Order{{
			Field: "id",
		}},
	})
	if err != nil {
		return BotListOut{
			ErrorCode: 2000,
		}
	}

	var strategies []models.Bot
	err = gomapper.MapStructs(&strategies, &dtos)
	if err != nil {
		return BotListOut{
			ErrorCode: 3001,
		}
	}

	return BotListOut{
		Success: true,
		Data:    strategies,
	}
}

func (b *Bot) Subscribe(ctx context.Context, in BotSubscribeIn) BOut {
	panic("implement me")
}

func (b *Bot) Unsubscribe(ctx context.Context, in BotSubscribeIn) BOut {
	panic("implement me")
}

func (b *Bot) createHook(ctx context.Context, dto models.BotDTO) map[string]string {
	var pair string
	ticker, err := b.ticker.GetByID(ctx, dto.PairID)
	if err == nil {
		pair = ticker.Pair
	} else {
		pair = "N"
	}

	path := fmt.Sprintf("BUYMARKET_%s_%s", pair, dto.GetUUID())

	buyMarket := fmt.Sprintf("%s/%s", b.conf.APIUrl, path)

	return map[string]string{
		models.OrderTypes[models.BuyMarket]: buyMarket,
	}
}

func (b *Bot) addPairs(ctx context.Context, in BotPairAdd) BOut {
	var err error
	for i := range in.Pairs {
		var dto models.StrategyPairDTO
		dto.SetPairID(in.Pairs[i]).SetStrategyID(in.BotID)
		err = b.strategyPair.Create(ctx, dto)
		if err != nil {
			return BOut{
				ErrorCode: 2000,
			}
		}
	}

	return BOut{
		Success: true,
	}
}

func (b *Bot) WebhookSignal(ctx context.Context, in WebhookSignalIn) WebhookSignalOut {
	bot, err := b.bot.GetByUUID(ctx, in.BotUUID)
	if err != nil {
		return WebhookSignalOut{
			ErrorCode: errors.InternalError,
		}
	}

	pair, err := b.ticker.GetByID(ctx, in.PairID)
	if err != nil {
		return WebhookSignalOut{
			ErrorCode: errors.InternalError,
		}
	}

	signals := WebhookSignalOut{
		Hook: b.conf.APIUrl,
	}

	for _, signal := range service.OrderTypesRaw {
		if signal == service.TypeGetOrderRaw ||
			signal == service.TypeAutoSellLimitRaw {
			continue
		}
		signals.Signals = append(signals.Signals, fmt.Sprintf("%s_%s_%s", signal, pair.Pair, bot.UUID))
	}
	SortNameAscend(signals.Signals)

	return signals
}

func SortNameAscend(files []string) {
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})
}
