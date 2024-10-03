package controller

import (
	"encoding/json"
	"github.com/ptflp/godecoder"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/service"
)

type Userer interface {
	Profile(http.ResponseWriter, *http.Request)
	GetUsersInfo(http.ResponseWriter, *http.Request)
	ChangePassword(http.ResponseWriter, *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	SendResetCodeEmail(w http.ResponseWriter, r *http.Request)
}

type User struct {
	service service.Userer
	responder.Responder
	godecoder.Decoder
}

func NewUser(service service.Userer, components *component.Components) Userer {
	return &User{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Получение профиля пользователя
// @Security ApiKeyAuth
// @Tags user
// @ID profileRequest
// @Accept  json
// @Produce  json
// @Success 200 {object} ProfileResponse
// @Router /api/1/user/profile [get]
func (u *User) Profile(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	out := u.service.GetByID(r.Context(), service.GetByIDIn{UserID: claims.ID})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, ProfileResponse{
			ErrorCode: out.ErrorCode,
			Data: Data{
				Message: "retrieving user error",
			},
		})
		return
	}

	u.OutputJSON(w, ProfileResponse{
		Success:   true,
		ErrorCode: out.ErrorCode,
		Data: Data{
			User: *out.User,
		},
	})
}

func (u *User) GetUsersInfo(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// @Summary Изменение пароля пользователя
// @Security ApiKeyAuth
// @Tags user
// @ID changePasswordRequest
// @Accept json
// @Produce json
// @Param request body ChangePasswordRequest true "Запрос на изменение пароля"
// @Success 200 {object} ChangePasswordResponse
// @Router /api/1/user/changePassword [post]
func (u *User) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req service.ChangePasswordIn
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.OutputJSON(w, service.ChangePasswordOut{
			Success:   false,
			ErrorCode: errors.UserServiceChangePasswordErr,
		})
		return
	}

	out := u.service.ChangePassword(r.Context(), service.ChangePasswordIn{
		Email:              req.Email,
		OldPassword:        req.OldPassword,
		NewPassword:        req.NewPassword,
		ConfirmNewPassword: req.ConfirmNewPassword,
	})

	if out.Success {
		u.OutputJSON(w, service.ChangePasswordOut{
			Success: true,
		})
	} else {
		u.OutputJSON(w, service.ChangePasswordOut{
			Success:   false,
			ErrorCode: out.ErrorCode,
		})
	}
}

// @Summary Сброс пароля пользователя
// @Tags user
// @ID resetPasswordRequest
// @Accept json
// @Produce json
// @Param request body ResetPasswordRequest true "Запрос на сброс пароля"
// @Success 200 {object} ResetPasswordResponse
// @Router /api/1/user/resetPassword [post]
func (u *User) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req service.ResetPasswordIn
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.OutputJSON(w, service.ResetPasswordOut{
			Success:   false,
			ErrorCode: errors.UserServiceResetPasswordErr,
		})
		return
	}

	out := u.service.ResetPassword(r.Context(), req)

	if out.Success {
		u.OutputJSON(w, service.ResetPasswordOut{
			Success: true,
		})
	} else {
		u.OutputJSON(w, service.ResetPasswordOut{
			Success:   false,
			ErrorCode: out.ErrorCode,
		})
	}
}

// @Summary Отправка кода сброса по электронной почте
// @Tags user
// @ID sendCodeRequest
// @Accept json
// @Produce json
// @Param request body sendCodeRequest true "Запрос на отправку кода сброса по электронной почте"
// @Success 200 {object} sendCodeResponse
// @Router /api/1/user/sendCode [post]
func (u *User) SendResetCodeEmail(w http.ResponseWriter, r *http.Request) {
	var req service.SendResetCodeEmailIn
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.OutputJSON(w, service.SendResetCodeEmailOut{
			Success:   false,
			ErrorCode: errors.UserServiceSendResetCodeEmailErr,
		})
		return
	}

	out := u.service.SendResetCodeEmail(r.Context(), req)

	if out.Success {
		u.OutputJSON(w, service.SendResetCodeEmailOut{
			Success: true,
		})
	} else {
		u.OutputJSON(w, service.SendResetCodeEmailOut{
			Success:   false,
			ErrorCode: out.ErrorCode,
		})
	}
}

//
