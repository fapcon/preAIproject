package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/ptflp/godecoder"
	"net/http"
	"net/mail"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/service"
)

type Auther interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	Verify(w http.ResponseWriter, r *http.Request)
	SocialCallback(http.ResponseWriter, *http.Request)
	SocialRedirect(http.ResponseWriter, *http.Request)
}

type Auth struct {
	auth service.Auther
	responder.Responder
	godecoder.Decoder
}

func NewAuth(service service.Auther, components *component.Components) Auther {
	return &Auth{auth: service, Responder: components.Responder, Decoder: components.Decoder}
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// @Summary Регистрация пользователя
// @Tags auth
// @ID registerRequest
// @Accept  json
// @Produce  json
// @Param object body RegisterRequest true "RegisterRequest"
// @Success 200 {object} RegisterResponse
// @Router /api/1/auth/register [post]
func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := a.Decode(r.Body, &req)
	if err != nil {
		a.ErrorBadRequest(w, err)
		return
	}

	if !valid(req.Email) {
		a.OutputJSON(w, RegisterResponse{
			Success:   false,
			ErrorCode: http.StatusBadRequest,
			Data: Data{
				Message: "invalid email",
			},
		})
		return
	}

	if req.Password != req.RetypePassword {
		a.OutputJSON(w, RegisterResponse{
			Success:   false,
			ErrorCode: http.StatusBadRequest,
			Data: Data{
				Message: "passwords mismatch",
			},
		})
		return
	}

	out := a.auth.Register(r.Context(), service.RegisterIn{
		Email:          req.Email,
		Password:       req.Password,
		IdempotencyKey: req.IdempotencyKey,
	}, service.RegisterEmail)

	if out.ErrorCode != errors.NoError {
		msg := "register error"
		if out.ErrorCode == errors.UserServiceUserAlreadyExists {
			msg = "User already exists, please check your email"
		}
		a.OutputJSON(w, RegisterResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Data: Data{
				Message: msg,
			},
		})
		return
	}

	a.OutputJSON(w, RegisterResponse{
		Success: true,
		Data: Data{
			Message: "verification link sent to " + req.Email,
		},
	})
}

// @Summary Авторизация пользователя
// @Tags auth
// @ID loginRequest
// @Accept  json
// @Produce  json
// @Param object body LoginRequest true "LoginRequest"
// @Success 200 {object} AuthResponse
// @Router /api/1/auth/login [post]
func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := a.Decode(r.Body, &req)
	if err != nil {
		a.ErrorBadRequest(w, err)
		return
	}
	if len(req.Email) < 5 {
		a.OutputJSON(w, RegisterResponse{
			Success:   false,
			ErrorCode: http.StatusBadRequest,
			Data: Data{
				Message: "phone or email empty",
			},
		})
	}

	out := a.auth.AuthorizeEmail(r.Context(), service.AuthorizeEmailIn{
		Email:    req.Email,
		Password: req.Password,
	})
	if out.ErrorCode == errors.AuthServiceUserNotVerified {
		a.OutputJSON(w, AuthResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Data: LoginData{
				Message: "user email is not verified",
			},
		})
		return
	}

	if out.ErrorCode != errors.NoError {
		a.OutputJSON(w, AuthResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Data: LoginData{
				Message: "login or password mismatch",
			},
		})
		return
	}

	a.OutputJSON(w, AuthResponse{
		Success: true,
		Data: LoginData{
			Message:      "success login",
			AccessToken:  out.AccessToken,
			RefreshToken: out.RefreshToken,
		},
	})
}

// @Summary Обновление рефреш токена
// @Security ApiKeyAuth
// @Tags auth
// @ID refreshRequest
// @Accept  json
// @Produce  json
// @Success 200 {object} AuthResponse
// @Router /api/1/auth/refresh [post]
func (a *Auth) Refresh(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		a.ErrorBadRequest(w, err)
		return
	}
	out := a.auth.AuthorizeRefresh(r.Context(), service.AuthorizeRefreshIn{UserID: claims.ID})

	if out.ErrorCode != errors.NoError {
		a.OutputJSON(w, AuthResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Data: LoginData{
				Message: "login or password mismatch",
			},
		})
		return
	}

	a.OutputJSON(w, AuthResponse{
		Success: true,
		Data: LoginData{
			Message:      "success refresh",
			AccessToken:  out.AccessToken,
			RefreshToken: out.RefreshToken,
		},
	})
}

// @Summary Обновление рефреш токена
// @Tags auth
// @ID verifyRequest
// @Accept  json
// @Produce  json
// @Param object body VerifyRequest true "VerifyRequest"
// @Success 200 {object} AuthResponse
// @Router /api/1/auth/verify [post]
func (a *Auth) Verify(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	err := a.Decode(r.Body, &req)
	if err != nil {
		a.ErrorBadRequest(w, err)
		return
	}
	if len(req.Email) < 5 {
		a.OutputJSON(w, RegisterResponse{
			Success:   false,
			ErrorCode: http.StatusBadRequest,
			Data: Data{
				Message: "phone or email empty",
			},
		})
	}

	out := a.auth.VerifyEmail(r.Context(), service.VerifyEmailIn{
		Email: req.Email,
		Hash:  req.Hash,
	})

	if out.ErrorCode != errors.NoError {
		a.OutputJSON(w, AuthResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Data: LoginData{
				Message: "login or password mismatch",
			},
		})
		return
	}

	a.OutputJSON(w, AuthResponse{
		Success: true,
		Data: LoginData{
			Message: "email verification success",
		},
	})
}

// @Summary Обработка ответа сервисов авторизации
// @Tags auth
// @ID socialCallback
// @Accept  json
// @Produce  json
// @Param provider path string true "Provider"
// @Param code formData string true "Code"
// @Success 200 {object} AuthResponse
// @Router /api/1/auth/{provider}/callback [get]
func (a *Auth) SocialCallback(w http.ResponseWriter, r *http.Request) {
	providerName := chi.URLParam(r, "provider")
	code := r.FormValue("code")
	out := a.auth.SocialCallback(r.Context(), service.SocialCallbackIn{
		Code:     code,
		Provider: providerName,
	})

	if out.ErrorCode != errors.NoError {
		a.OutputJSON(w, AuthResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Data: LoginData{
				Message: "login or password mismatch",
			},
		})
		return
	}
	a.OutputJSON(w, AuthResponse{
		Success: true,
		Data: LoginData{
			Message:      "success login",
			AccessToken:  out.AccessToken,
			RefreshToken: out.RefreshToken,
		},
	})
}

// @Summary Редирект в сервисы авторизации
// @Tags auth
// @ID socialRedirect
// @Accept  json
// @Produce  json
// @Param provider path string true "Provider"
// @Router /api/1/auth/{provider}/login [get]
func (a *Auth) SocialRedirect(w http.ResponseWriter, r *http.Request) {
	providerName := chi.URLParam(r, "provider")
	urlOut := a.auth.SocialGetRedirectURL(r.Context(), service.SocialGetRedirectUrlIn{
		Provider: providerName,
	})
	redirectURL := urlOut.Url
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}
