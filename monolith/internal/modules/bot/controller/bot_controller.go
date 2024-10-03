package controller

import (
	"github.com/ptflp/godecoder"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service"
)

type Boter interface {
	Create(http.ResponseWriter, *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Toggle(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	WebhookSignal(w http.ResponseWriter, r *http.Request)
}

type Bot struct {
	bot service.Boter
	responder.Responder
	godecoder.Decoder
}

func NewBot(strategy service.Boter, components *component.Components) *Bot {
	return &Bot{bot: strategy, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Поиск стратегии ID или UUID
// @Security ApiKeyAuth
// @Tags strategy
// @ID strategyGet
// @Accept  json
// @Produce  json
// @Param object body GetBotRequest true "GetBotRequest"
// @Success 200 {object} BotResponse
// @Router /api/1/strategy/get [post]
func (b *Bot) Get(w http.ResponseWriter, r *http.Request) {
	var req GetBotRequest
	var res BotResponse
	err := b.Decode(r.Body, &req)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	out := b.bot.Get(r.Context(), service.BotGetIn{
		UUID: req.UUID,
		ID:   req.ID,
	})

	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		b.OutputJSON(w, res)
		return
	}

	b.OutputJSON(w, BotResponse{
		Success: true,
		Data: BotData{
			Bot:   out.Bot,
			Hooks: out.Hooks,
		},
	})
}

// @Summary Создание стратегии
// @Security ApiKeyAuth
// @Tags strategy
// @ID strategyCreate
// @Accept  json
// @Produce  json
// @Success 200 {object} BotResponse
// @Router /api/1/strategy/create [get]
func (b *Bot) Create(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}
	out := b.bot.Create(r.Context(), service.BotCreateIn{UserID: userClaims.ID})

	if out.ErrorCode != errors.NoError {
		b.OutputJSON(w, BotResponse{
			ErrorCode: out.ErrorCode,
			Data: BotData{
				Message: "bot create error",
			},
		})
		return
	}

	b.OutputJSON(w, BotResponse{
		Success: true,
		Data: BotData{
			Bot:   out.Bot,
			Hooks: out.Hooks,
		},
	})
}

// @Summary Редактирование стратегии
// @Security ApiKeyAuth
// @Tags strategy
// @ID strategyUpdate
// @Accept  json
// @Produce  json
// @Param object body UpdateRequest true "UpdateRequest"
// @Success 200 {object} UpdateResponse
// @Router /api/1/strategy/update [post]
func (b *Bot) Update(w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest
	var res UpdateResponse
	err := b.Decode(r.Body, &req)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}
	req.Strategy.UserID = userClaims.ID
	out := b.bot.Update(r.Context(), service.BotUpdateIn{
		Bot:   req.Strategy,
		Pairs: req.Pairs,
	})
	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		b.OutputJSON(w, res)
		return
	}

	res.Success = true
	b.OutputJSON(w, res)
}

// @Summary Получение списка стратегий
// @Security ApiKeyAuth
// @Tags strategy
// @ID strategyGetList
// @Accept  json
// @Produce  json
// @Success 200 {object} BotListResponse
// @Router /api/1/strategy/list [get]
func (b *Bot) List(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}
	var res BotListResponse
	out := b.bot.List(r.Context(), service.BotListIn{UserID: userClaims.ID})

	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		b.OutputJSON(w, res)
		return
	}

	res.Success = true
	res.Data = out.Data
	b.OutputJSON(w, res)
}

// @Summary Переключение активности стратегии
// @Security ApiKeyAuth
// @Tags strategy
// @ID strategyToggle
// @Accept  json
// @Produce  json
// @Param object body BotToggleRequest true "BotToggleRequest"
// @Success 200 {object} DefaultResponse
// @Router /api/1/strategy/toggle [post]
func (b *Bot) Toggle(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	var req BotToggleRequest
	err = b.Decode(r.Body, &req)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	out := b.bot.Toggle(r.Context(), service.BotToggleIn{
		UUID:   req.BotUUID,
		UserID: userClaims.ID,
		Active: req.Active,
	})

	if out.ErrorCode != errors.NoError {
		b.OutputJSON(w, DefaultResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	b.OutputJSON(w, DefaultResponse{
		Success: true,
	})
}

// @Summary Удаление стратегии
// @Security ApiKeyAuth
// @Tags strategy
// @ID strategyDelete
// @Accept  json
// @Produce  json
// @Param object body BotUUIDRequest true "BotUUIDRequest"
// @Success 200 {object} DefaultResponse
// @Router /api/1/strategy/delete [post]
func (b *Bot) Delete(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	var req BotUUIDRequest
	err = b.Decode(r.Body, &req)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	out := b.bot.Delete(r.Context(), service.BotDeleteIn{UUID: req.BotUUID, UserID: userClaims.ID})

	if out.ErrorCode != errors.NoError {
		b.OutputJSON(w, DefaultResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	b.OutputJSON(w, DefaultResponse{
		Success: true,
	})
}

// @Summary Обработка сигналов для вебхуков
// @Security ApiKeyAuth
// @Tags strategy
// @ID strategySignal
// @Accept  json
// @Produce  json
// @Param object body WebhookSignalRequest true "WebhookSignalRequest"
// @Success 200 {object} WebhookSignalResponse
// @Router /api/1/strategy/signal [post]
func (b *Bot) WebhookSignal(w http.ResponseWriter, r *http.Request) {
	var err error
	var req WebhookSignalRequest
	err = b.Decode(r.Body, &req)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	out := b.bot.WebhookSignal(r.Context(), service.WebhookSignalIn{
		BotUUID: req.BotUUID,
		PairID:  req.PairID,
	})

	if out.ErrorCode != errors.NoError {
		b.OutputJSON(w, WebhookSignalResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	b.OutputJSON(w, WebhookSignalResponse{
		Hook:    out.Hook,
		Signals: out.Signals,
		Success: true,
	})
}
