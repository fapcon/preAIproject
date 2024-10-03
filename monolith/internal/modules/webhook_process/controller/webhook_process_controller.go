package controller

import (
	"bytes"
	"net/http"

	"github.com/ptflp/godecoder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/webhook_process/service"
)

type WebhookProcesser interface {
	WebhookProcess(w http.ResponseWriter, r *http.Request)
	UserWebhooksHistory(w http.ResponseWriter, r *http.Request)
	WebhookInfo(w http.ResponseWriter, r *http.Request)
	BotInfo(w http.ResponseWriter, r *http.Request)
}

type WebhookProcess struct {
	service service.WebhookProcesser
	responder.Responder
	godecoder.Decoder
}

func NewWebhookProcess(service service.WebhookProcesser, components *component.Components) WebhookProcesser {
	return &WebhookProcess{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Получение информации об вебхуке по uuid
// @Security ApiKeyAuth
// @Tags exchange/user/webhooks
// @ID exchangeUserWebhooksInfo
// @Accept  json
// @Produce  json
// @Param object body WebhookInfoRequest true "WebhookInfoRequest"
// @Success 200 {object} WebhookInfoResponse
// @Router /api/1/exchange/user/webhooks/info [post]
func (p *WebhookProcess) WebhookInfo(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	var req WebhookInfoRequest

	err = p.Decode(r.Body, &req)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.GetWebhookInfo(r.Context(), service.GetWebhookInfoIn{
		WebhookUUID: req.WebhookUUID,
		UserID:      userClaims.ID,
	})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, WebhookInfoResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, WebhookInfoResponse{
		Success: true,
		Data: WebhookInfoData{
			Webhook: out.Data.Webhook,
			Orders:  out.Data.Orders,
		},
	})
}

// @Summary Обработка вебхуков пользователя
// @Tags platform
// @ID platformHook
// @Accept  json
// @Produce  json
// @Param slug body string true "Slug"
// @Success 200 {object} WebhookProcessResponse
// @Router /platform/hook [post]
func (p *WebhookProcess) WebhookProcess(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	_, err := b.ReadFrom(r.Body)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	slug := string(b.Bytes())
	out := p.service.WebhookProcess(r.Context(), service.WebhookProcessIn{
		Slug:          slug,
		RemoteAddr:    r.RemoteAddr,
		XForwardedFor: r.Header.Get("X-Forwarded-For"),
	})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, WebhookProcessResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, WebhookProcessResponse{Success: true})
}

// @Summary Получение истории вебхуков пользователя
// @Security ApiKeyAuth
// @Tags exchange/user/webhooks
// @Accept  json
// @Produce  json
// @Success 200 {object} GetUserWebhooksResponse
// @Router /api/1/exchange/user/webhooks/history [get]
// @Router /api/1/exchange/user/webhooks/toggle [post]
func (p *WebhookProcess) UserWebhooksHistory(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.GetUserWebhooks(r.Context(), service.GetUserRelationIn{UserID: userClaims.ID})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, GetUserWebhooksResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, GetUserWebhooksResponse{
		Success: true,
		Data:    out.Data,
	})
}

// @Summary Получение информации об вебхуке по uuid
// @Security ApiKeyAuth
// @Tags exchange/user/bot
// @ID exchangeUserbotInfo
// @Accept  json
// @Produce  json
// @Param object body BotInfoRequest true "BotInfoRequest"
// @Success 200 {object} BotInfoResponse
// @Router /api/1/exchange/user/bot/info [post]
func (p *WebhookProcess) BotInfo(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	var req BotInfoRequest

	err = p.Decode(r.Body, &req)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.GetBotInfo(r.Context(), service.GetBotInfoIn{
		BotUUID: req.BotUUID,
		UserID:  userClaims.ID,
	})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, BotInfoResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, BotInfoResponse{
		Success: true,
		Data: BotInfoData{
			Bot:      out.Data.Bot,
			Orders:   out.Data.Orders,
			Webhooks: out.Data.Webhooks,
		},
	})
}
