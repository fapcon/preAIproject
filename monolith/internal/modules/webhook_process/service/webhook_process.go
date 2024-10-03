package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ptflp/gomapper"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service"
	eservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	oservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
	uservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/webhook_process/storage"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/storages"
)

type WebhookProcess struct {
	exchangeOrder                oservice.ExchangeOrderer
	webhookProcessStorage        storage.WebhookProcesser
	webhookProcessHistoryStorage storage.WebhookProcessHistorer
	exchangeUserKeyService       uservice.ExchangeUserKeyer
	botService                   service.Boter
	logger                       *zap.Logger
	rateLimiter                  *eservice.RateLimiter
}

const (
	StatusUnknown = iota
	StatusProcessing
	StatusFinished
	StatusFailed
)

func NewWebhookProcess(storages *storages.Storages, order oservice.ExchangeOrderer, userKey uservice.ExchangeUserKeyer, bot service.Boter, components *component.Components) WebhookProcesser {
	return &WebhookProcess{
		webhookProcessStorage:        storages.WebhookProcess,
		webhookProcessHistoryStorage: storages.WebhookProcessHistory,
		exchangeOrder:                order,
		exchangeUserKeyService:       userKey,
		botService:                   bot,
		logger:                       components.Logger,
		rateLimiter:                  components.RateLimiter,
	}
}

func (p *WebhookProcess) extractSignal(slug string) (models.Signal, error) {
	signalRaw := strings.Split(slug, "_")
	if len(signalRaw) != 3 {
		return models.Signal{}, fmt.Errorf("bad signal format general")
	}

	botUUID, err := uuid.Parse(signalRaw[2])
	if err != nil {
		return models.Signal{}, fmt.Errorf("bad signal format in uuid")
	}

	if orderType, ok := eservice.OrderTypes[signalRaw[0]]; ok {
		return models.Signal{
			OrderType: orderType,
			Pair:      signalRaw[1],
			BotUUID:   botUUID.String(),
		}, nil
	}

	return models.Signal{}, fmt.Errorf("bad signal format in orderType")
}

func (p *WebhookProcess) WebhookProcess(ctx context.Context, in WebhookProcessIn) WebhookProcessOut {
	signal, err := p.extractSignal(in.Slug)
	if err != nil {
		return WebhookProcessOut{
			ErrorCode: errors.WebhookProcessError,
		}
	}

	botOut := p.botService.Get(ctx, service.BotGetIn{UUID: signal.BotUUID})
	if botOut.ErrorCode != errors.NoError {
		return WebhookProcessOut{
			ErrorCode: errors.WebhookProcessError,
		}
	}

	bot := botOut.Bot
	if !bot.Active {
		return WebhookProcessOut{
			ErrorCode: errors.WebhookBotInactive,
		}
	}
	webhook, err := p.CreateWebhookProcess(ctx, bot, in)
	if err != nil {
		return WebhookProcessOut{
			ErrorCode: errors.WebhookProcessError,
		}
	}

	keys, err := p.exchangeUserKeyService.ExchangeUserKeyListByIDs(ctx, bot.ExchangeID, bot.UserID)
	if err != nil {
		p.UpdateWebhookStatus(ctx, bot, webhook, "please add exchange API Keys", StatusFailed)
		return WebhookProcessOut{
			ErrorCode: errors.WebhookProcessError,
		}
	}
	orderIn := oservice.OrderIn{
		Webhook: webhook,
		Bot:     bot,
		Signal:  signal,
	}
	orderCtx := context.Background()

	for i := range keys {
		if !keys[i].MakeOrder {
			continue
		}
		oIn := orderIn
		oIn.Key = keys[i]
		oIn.ExClient = eservice.NewBinanceWithKey(bot.UserID, p.rateLimiter, keys[i].APIKey, keys[i].SecretKey)
		switch signal.OrderType {
		case eservice.TypeCancelBuy, eservice.TypeCancelAll, eservice.TypeCancelSell, eservice.TypeAverage:
			go p.exchangeOrder.CancelOrder(orderCtx, oIn)
		default:
			go func(ctx context.Context, in oservice.OrderIn) {
				out := p.exchangeOrder.PutOrder(orderCtx, in)
				if out.OrderID != 0 {
					in.Webhook.OrderID = out.OrderID
					in.Webhook.OrderUUID = out.OrderUUID
				}
				if out.ErrorCode != errors.NoError {
					if out.OrderID != 0 {
						p.UpdateWebhookStatus(ctx, in.Bot, in.Webhook, "", StatusFailed)
					}
					return
				}

				err := p.webhookProcessHistoryStorage.Create(ctx, models.WebhookProcessHistoryDTO{
					UserID:      in.Webhook.UserID,
					WebhookUUID: in.Webhook.UUID,
					WebhookID:   in.Webhook.ID,
					ExchangeID:  in.Bot.ExchangeID,
					Status:      StatusFinished,
				})
				if err != nil {
					return
				}

				in.Webhook.Status = StatusFinished
				in.Webhook.SetUpdatedAt(time.Now())
				_ = p.webhookProcessStorage.Update(ctx, in.Webhook)
			}(ctx, oIn)
		}
	}
	return WebhookProcessOut{
		Success: true,
	}
}

func (p *WebhookProcess) UpdateWebhookStatus(ctx context.Context, bot models.Bot, webhook models.WebhookProcessDTO, message string, status int) {
	p.WriteWebhookHistory(ctx, models.WebhookProcessHistoryDTO{
		UserID:      bot.UserID,
		WebhookUUID: webhook.UUID,
		WebhookID:   webhook.ID,
		ExchangeID:  bot.ExchangeID,
		Status:      status,
	})
	webhook.Status = status
	webhook.Message = message
	webhook.SetUpdatedAt(time.Now())
	_ = p.webhookProcessStorage.Update(ctx, webhook)
}

func (p *WebhookProcess) CreateWebhookProcess(ctx context.Context, bot models.Bot, in WebhookProcessIn) (models.WebhookProcessDTO, error) {
	webhookUUID := uuid.NewString()
	err := p.webhookProcessStorage.Create(ctx, models.WebhookProcessDTO{
		UUID:          webhookUUID,
		BotUUID:       bot.UUID,
		BotID:         bot.ID,
		UserID:        bot.UserID,
		Slug:          in.Slug,
		XForwardedFor: in.XForwardedFor,
		RemoteAddr:    in.RemoteAddr,
		Status:        StatusProcessing,
	})
	if err != nil {
		return models.WebhookProcessDTO{}, err
	}

	webhookProcessList, err := p.webhookProcessStorage.GetList(ctx, utils.Condition{
		Equal: map[string]interface{}{"uuid": webhookUUID},
	})
	if err != nil {
		return models.WebhookProcessDTO{}, err
	}
	if len(webhookProcessList) < 1 {
		return models.WebhookProcessDTO{}, fmt.Errorf("something went wrong, retrieving webhookprocess after create")
	}

	webhook := webhookProcessList[0]
	p.WriteWebhookHistory(ctx, models.WebhookProcessHistoryDTO{
		UserID:      bot.UserID,
		WebhookUUID: webhook.UUID,
		WebhookID:   webhook.ID,
		ExchangeID:  bot.ExchangeID,
		Status:      StatusProcessing,
	})
	return webhook, nil
}

func (p *WebhookProcess) WriteWebhookHistory(ctx context.Context, dto models.WebhookProcessHistoryDTO) {
	err := p.webhookProcessHistoryStorage.Create(ctx, dto)
	_ = err
}

func (p *WebhookProcess) getWebhooksCondition(ctx context.Context, condition utils.Condition) ([]models.WebhookProcess, error) {
	webhooksDTO, err := p.webhookProcessStorage.GetList(ctx, condition)
	if err != nil {
		return nil, err
	}

	var webhooks []models.WebhookProcess
	err = gomapper.MapStructs(&webhooks, &webhooksDTO)
	if err != nil {
		return nil, err
	}

	err = p.addWebhookHistory(ctx, &webhooks)
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (p *WebhookProcess) addWebhookHistory(ctx context.Context, webhooks *[]models.WebhookProcess) error {
	m := make(map[int]*models.WebhookProcess, len(*webhooks))
	var ids []int
	for i := range *webhooks {
		id := (*webhooks)[i].ID
		m[id] = &(*webhooks)[i]
		ids = append(ids, id)
		switch (*webhooks)[i].Status {
		case StatusProcessing:
			(*webhooks)[i].StatusMsg = "Processing"
		case StatusFinished:
			(*webhooks)[i].StatusMsg = "Finished"
		case StatusFailed:
			(*webhooks)[i].StatusMsg = "Failed"
		}
	}

	webhookHistoryDTO, err := p.webhookProcessHistoryStorage.GetList(ctx, utils.Condition{Equal: map[string]interface{}{"webhook_id": ids}})
	if err != nil {
		return err
	}

	var webhooksHistory []models.WebhookProcessHistory

	err = gomapper.MapStructs(&webhooksHistory, &webhookHistoryDTO)
	if err != nil {
		return err
	}
	for i := range webhooksHistory {
		switch webhooksHistory[i].Status {
		case StatusProcessing:
			webhooksHistory[i].StatusMsg = "Processing"
		case StatusFinished:
			webhooksHistory[i].StatusMsg = "Finished"
		case StatusFailed:
			webhooksHistory[i].StatusMsg = "Failed"
		}
		m[webhooksHistory[i].WebhookID].History = append(m[webhooksHistory[i].WebhookID].History, webhooksHistory[i])
	}

	return nil
}

func (p *WebhookProcess) GetBotWebhooks(ctx context.Context, in GetBotRelationIn) GetWebhooksOut {
	webhooks, err := p.getWebhooksCondition(ctx, utils.Condition{Equal: map[string]interface{}{"bot_uuid": in.BotUUID}})
	if err != nil {
		return GetWebhooksOut{
			ErrorCode: errors.InternalError,
			Success:   false,
			Data:      []models.WebhookProcess{},
		}
	}

	return GetWebhooksOut{
		Success: true,
		Data:    webhooks,
	}
}

func (p *WebhookProcess) GetWebhookInfo(ctx context.Context, in GetWebhookInfoIn) GetWebhookInfoOut {
	webhooks, err := p.getWebhooksCondition(ctx, utils.Condition{Equal: map[string]interface{}{"user_id": in.UserID, "uuid": in.WebhookUUID}})
	if err != nil || len(webhooks) < 1 {
		return GetWebhookInfoOut{
			ErrorCode: errors.InternalError,
		}
	}

	orders, err := p.exchangeOrder.GetOrdersCondition(ctx, utils.Condition{Equal: map[string]interface{}{"webhook_uuid": in.WebhookUUID, "user_id": in.UserID}})
	if err != nil {
		return GetWebhookInfoOut{
			ErrorCode: errors.InternalError,
		}
	}

	return GetWebhookInfoOut{
		Data: GetWebhookInfoData{
			Webhook: webhooks[0],
			Orders:  orders,
		},
	}
}

func (p *WebhookProcess) GetUserWebhooks(ctx context.Context, in GetUserRelationIn) GetWebhooksOut {
	webhooks, err := p.getWebhooksCondition(ctx, utils.Condition{Equal: map[string]interface{}{"user_id": in.UserID}, Order: []*utils.Order{{Field: "id", Asc: false}}})
	if err != nil {
		return GetWebhooksOut{
			ErrorCode: errors.InternalError,
			Success:   false,
			Data:      []models.WebhookProcess{},
		}
	}

	return GetWebhooksOut{
		Success: true,
		Data:    webhooks,
	}
}

func (p *WebhookProcess) GetBotInfo(ctx context.Context, in GetBotInfoIn) GetBotInfoOut {
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

	webhooksOut := p.GetBotWebhooks(ctx, GetBotRelationIn{BotUUID: in.BotUUID})
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
