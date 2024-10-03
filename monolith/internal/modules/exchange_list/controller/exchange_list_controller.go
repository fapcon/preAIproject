package controller

import (
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"

	"github.com/ptflp/godecoder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/service"
)

type ExchangeLister interface {
	ExchangeAdd(w http.ResponseWriter, r *http.Request)
	ExchangeDelete(w http.ResponseWriter, r *http.Request)
	ExchangeList(w http.ResponseWriter, r *http.Request)
}

type ExchangeList struct {
	service service.ExchangeLister
	responder.Responder
	godecoder.Decoder
}

func NewExchangeList(service service.ExchangeLister, components *component.Components) ExchangeLister {
	return &ExchangeList{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Добавление биржи торговли
// @Security ApiKeyAuth
// @Tags exchange
// @ID exchangeAdd
// @Accept  json
// @Produce  json
// @Param object body ExchangeAddRequest true "ExchangeAddRequest"
// @Success 200 {object} ExchangeResponse
// @Router /api/1/exchange/add [post]
func (p *ExchangeList) ExchangeAdd(w http.ResponseWriter, r *http.Request) {
	var req ExchangeAddRequest
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	err = p.Decode(r.Body, &req)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.ExchangeListAdd(r.Context(), service.ExchangeAddIn{
		UserID:      userClaims.ID,
		Name:        req.Name,
		Description: req.Description,
		Slug:        req.Slug,
	})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, ExchangeResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, ExchangeResponse{
		Success: true,
	})
}

// @Summary Получение списка бирж торговли
// @Security ApiKeyAuth
// @Tags exchange
// @ID exchangeList
// @Accept  json
// @Produce  json
// @Success 200 {object} ExchangeListResponse
// @Router /api/1/exchange/list [get]
func (p *ExchangeList) ExchangeList(w http.ResponseWriter, r *http.Request) {
	out := p.service.ExchangeListList(r.Context())
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, ExchangeListResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, ExchangeListResponse{
		Success: true,
		Data:    out.Data,
	})
}

// @Summary Удаление биржи торговли
// @Security ApiKeyAuth
// @Tags exchange
// @ID exchangeDelete
// @Accept  json
// @Produce  json
// @Param object body ExchangeDeleteRequest true "ExchangeDeleteRequest"
// @Success 200 {object} ExchangeResponse
// @Router /api/1/exchange/delete [post]
func (p *ExchangeList) ExchangeDelete(w http.ResponseWriter, r *http.Request) {
	var req ExchangeDeleteRequest

	err := p.Decode(r.Body, &req)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	err = p.service.ExchangeListDelete(r.Context(), req.ID)
	if err != nil {
		p.OutputJSON(w, ExchangeResponse{
			ErrorCode: errors.GeneralError,
		})
		return
	}

	p.OutputJSON(w, ExchangeResponse{
		Success: true,
	})

}
