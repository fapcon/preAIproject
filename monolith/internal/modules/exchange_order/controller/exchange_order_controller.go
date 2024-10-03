package controller

import (
	"net/http"

	"github.com/ptflp/godecoder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
)

type ExchangeOrderer interface {
	UserOrdersHistory(w http.ResponseWriter, r *http.Request)
	GetBotOrders(w http.ResponseWriter, r *http.Request)
}

type ExchangeOrder struct {
	service service.ExchangeOrderer
	responder.Responder
	godecoder.Decoder
}

func NewExchangeOrder(service service.ExchangeOrderer, components *component.Components) ExchangeOrderer {
	return &ExchangeOrder{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Получение истории торгов пользователя
// @Security ApiKeyAuth
// @Tags exchange/user/orders
// @ID exchangeUserOrders
// @Accept  json
// @Produce  json
// @Success 200 {object} GetUserOrdersResponse
// @Router /api/1/exchange/user/orders/history [get]
func (p *ExchangeOrder) UserOrdersHistory(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.GetUserOrders(r.Context(), service.GetUserRelationIn{UserID: userClaims.ID})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, GetUserOrdersResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	statistic := p.service.GetAllOrdersStatistic(r.Context(), service.GetUserRelationIn{UserID: userClaims.ID})

	p.OutputJSON(w, GetUserOrdersResponse{
		Success:   true,
		Data:      out.Data,
		Statistic: statistic,
	})
}

// @Summary Получение истории торгов пользователя
// @Security ApiKeyAuth
// @Tags exchange/user/bot
// @ID exchangeBotOrders
// @Accept  json
// @Produce  json
// @Param object body BotInfoRequest true "BotInfoRequest"
// @Success 200 {object} GetUserOrdersResponse
// @Router /api/1/exchange/user/bot/list [post]
func (p *ExchangeOrder) GetBotOrders(w http.ResponseWriter, r *http.Request) {
	var req BotInfoRequest
	err := p.Decode(r.Body, &req)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	out := p.service.ExchangeOrderList(r.Context(), service.GetBotRelationIn{
		BotUUID: req.BotUUID,
	})

	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, GetUserOrdersResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	statistic := p.service.GetOrdersStatistic(r.Context(), service.GetBotRelationIn{
		BotUUID: req.BotUUID,
	})

	p.OutputJSON(w, GetUserOrdersResponse{
		Success:   true,
		Data:      out.Data,
		Statistic: statistic,
	})
}
