package controller

import (
	"net/http"

	"github.com/ptflp/godecoder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service"
)

type ExchangeUserKeyer interface {
	ExchangeUserKeyAdd(w http.ResponseWriter, r *http.Request)
	ExchangeUserKeyDelete(w http.ResponseWriter, r *http.Request)
	ExchangeUserKeyList(w http.ResponseWriter, r *http.Request)
	CheckKeys(w http.ResponseWriter, r *http.Request)
}

type UserKey struct {
	service service.ExchangeUserKeyer
	responder.Responder
	godecoder.Decoder
}

func NewUserKey(service service.ExchangeUserKeyer, components *component.Components) ExchangeUserKeyer {
	return &UserKey{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Добавление ключей пользователя от биржи торговли
// @Security ApiKeyAuth
// @Tags exchange/user/key
// @ID exchangeUserKeyAdd
// @Accept  json
// @Produce  json
// @Param object body ExchangeUserKeyAddRequest true "ExchangeUserKeyAddRequest"
// @Success 200 {object} ExchangeResponse
// @Router /api/1/exchange/user/key/add [post]
func (p *UserKey) ExchangeUserKeyAdd(w http.ResponseWriter, r *http.Request) {
	var req ExchangeUserKeyAddRequest
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

	out := p.service.ExchangeUserKeyAdd(r.Context(), service.ExchangeUserKeyAddIn{
		APIKey:     req.APIKey,
		SecretKey:  req.SecretKey,
		Label:      req.Label,
		UserID:     userClaims.ID,
		ExchangeID: req.ExchangeID,
	})

	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, ExchangeListResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, ExchangeResponse{
		Success: true,
	})
}

func (p *UserKey) CheckKeys(w http.ResponseWriter, r *http.Request) {
	var req ExchangeUserKeyAddRequest

	err := p.Decode(r.Body, &req)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	err = p.service.CheckKeys(r.Context(), service.ExchangeUserKeyAddIn{
		APIKey:     req.APIKey,
		SecretKey:  req.SecretKey,
		ExchangeID: req.ExchangeID,
	})

	if err != nil {
		p.OutputJSON(w, ExchangeListResponse{
			ErrorCode: errors.InternalError,
		})
		return
	}

	p.OutputJSON(w, ExchangeResponse{
		Success: true,
	})

}

// @Summary Получение списка ключей пользователя от биржи торговли
// @Security ApiKeyAuth
// @Tags exchange/user/key
// @ID exchangeUserKeyList
// @Accept  json
// @Produce  json
// @Success 200 {object} ExchangeUserKeyListResponse
// @Router /api/1/exchange/user/key/list [get]
func (p *UserKey) ExchangeUserKeyList(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	out := p.service.ExchangeUserKeyList(r.Context(), service.ExchangeUserListIn{UserID: userClaims.ID})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, ExchangeUserKeyListResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, ExchangeUserKeyListResponse{
		Success: true,
		Data:    out.Data,
	})
}

// @Summary Удаление ключей биржи по id
// @Security ApiKeyAuth
// @Tags exchange/user/key
// @ID exchangeUserKeyDelete
// @Accept  json
// @Produce  json
// @Param object body ExchangeUserKeyDeleteRequest true "ExchangeUserKeyDeleteRequest"
// @Success 200 {object} ExchangeResponse
// @Router /api/1/exchange/user/key/delete [post]
func (p *UserKey) ExchangeUserKeyDelete(w http.ResponseWriter, r *http.Request) {
	var req ExchangeUserKeyDeleteRequest

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

	err = p.service.ExchangeUserKeyDelete(r.Context(), req.ExchangeUserKeyID, userClaims.ID)
	if err != nil {
		p.ErrorForbidden(w, err)
		return
	}

	p.OutputJSON(w, ExchangeResponse{
		Success: true,
	})
}
