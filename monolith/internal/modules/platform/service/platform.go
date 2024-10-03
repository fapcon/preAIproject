package service

import (
	"context"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service"
	oservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
	wservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/webhook_process/service"
)

type Platform struct {
	exchangeOrder         oservice.ExchangeOrderer
	webhookProcessService wservice.WebhookProcesser
	botService            service.Boter
	logger                *zap.Logger
}

func NewPlatform(webhook wservice.WebhookProcesser, order oservice.ExchangeOrderer, bot service.Boter, components *component.Components) *Platform {
	return &Platform{
		exchangeOrder:         order,
		webhookProcessService: webhook,
		botService:            bot,
		logger:                components.Logger,
	}
}

func (p *Platform) GetBotInfo(ctx context.Context, in GetBotInfoIn) GetBotInfoOut {
	out := p.botService.Get(ctx, service.BotGetIn{
		UUID: in.BotUUID,
	})
	if out.ErrorCode != errors.NoError {
		return GetBotInfoOut{
			ErrorCode: out.ErrorCode,
		}
	}

	orders, err := p.exchangeOrder.GetOrdersCondition(ctx, utils.Condition{Equal: map[string]interface{}{"bot_uuid": in.BotUUID, "user_id": in.UserID}})
	if err != nil {
		return GetBotInfoOut{
			ErrorCode: errors.InternalError,
		}
	}

	webhooksOut := p.webhookProcessService.GetBotWebhooks(ctx, wservice.GetBotRelationIn{BotUUID: in.BotUUID})
	if webhooksOut.ErrorCode != errors.NoError {
		return GetBotInfoOut{
			ErrorCode: webhooksOut.ErrorCode,
		}
	}

	return GetBotInfoOut{
		Data: GetBotInfoData{
			Bot:      out.Bot,
			Orders:   orders,
			Webhooks: webhooksOut.Data,
		},
	}
}
