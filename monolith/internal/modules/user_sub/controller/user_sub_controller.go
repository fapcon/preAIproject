package controller

import (
	"encoding/json"
	"github.com/ptflp/godecoder"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user_sub/service"
)

type UserSuber interface {
	Add(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type UserSub struct {
	service service.UserSuber
	responder.Responder
	godecoder.Decoder
}

func NewUserSubscription(service service.UserSuber, components *component.Components) UserSuber {
	return &UserSub{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Добавление подписки на пользователя
// @Security ApiKeyAuth
// @Tags user/subscription
// @ID userSubAddRequest
// @Accept  json
// @Produce  json
// @Param object body UserSubAddRequest true "UserSubAddRequest"
// @Success 200 {object} UserSubResponse
// @Router /api/1/user/subscription/add [post]
func (u *UserSub) Add(w http.ResponseWriter, r *http.Request) {
	user, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}

	var req UserSubAddRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.OutputJSON(w, UserSubResponse{
			Success:   false,
			ErrorCode: errors.UserSubAddError,
		})
		return
	}

	out := u.service.Add(r.Context(), service.UserSubAddIn{
		UserID:    user.ID,
		SubUserID: req.UserID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, UserSubResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
		})
		return
	}

	u.OutputJSON(w, UserSubResponse{
		Success:   true,
		ErrorCode: errors.NoError,
	})
}

// @Summary Список подписок пользователя
// @Security ApiKeyAuth
// @Tags user/subscription
// @ID userSubListRequest
// @Accept  json
// @Produce  json
// @Success 200 {object} UserSubListResponse
// @Router /api/1/user/subscription/list [get]
func (u *UserSub) List(w http.ResponseWriter, r *http.Request) {
	user, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}

	out := u.service.List(r.Context(), service.UserSubListIn{
		UserID: user.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, UserSubListResponse{
			SubUserIDs: nil,
			Success:    false,
			ErrorCode:  out.ErrorCode,
		})
		return
	}

	u.OutputJSON(w, UserSubListResponse{
		SubUserIDs: out.SubUserIDs,
		Success:    true,
		ErrorCode:  errors.NoError,
	})
}

// @Summary Удаление подписки на пользователя
// @Security ApiKeyAuth
// @Tags user/subscription
// @ID userSubDeleteRequest
// @Accept  json
// @Produce  json
// @Param object body UserSubDelRequest true "UserSubDelRequest"
// @Success 200 {object} UserSubResponse
// @Router /api/1/user/subscription/delete [delete]
func (u *UserSub) Delete(w http.ResponseWriter, r *http.Request) {
	user, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}

	var req UserSubDelRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.OutputJSON(w, UserSubResponse{
			Success:   false,
			ErrorCode: errors.UserSubAddError,
		})
		return
	}

	out := u.service.Delete(r.Context(), service.UserSubDelIn{
		UserID:    user.ID,
		SubUserID: req.UserID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, UserSubResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
		})
		return
	}

	u.OutputJSON(w, UserSubResponse{
		Success:   true,
		ErrorCode: errors.NoError,
	})
}
