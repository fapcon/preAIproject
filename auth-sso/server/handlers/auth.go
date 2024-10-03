package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/clients/auth"

	"net/http"
	"time"
)

type Auther interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	SocialCallback(w http.ResponseWriter, r *http.Request)
	SocialGetRedirectURL(w http.ResponseWriter, r *http.Request)
}

type Auth struct {
	service auth.ClientService
}

func NewAuth(service auth.ClientService) *Auth {
	return &Auth{service: service}
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.service.Login(context.Background(), request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := a.service.Profile(context.Background(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Expires:  time.Now().Add(8760 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := a.service.RegisterNewUser(context.Background(), request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(RegisterResponse{UserID: userID})
}

func (a *Auth) Profile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Error(w, fmt.Errorf("user is not authorized").Error(), http.StatusForbidden)
		}
		return
	}

	token := cookie.Value

	user, err := a.service.Profile(context.Background(), token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt_token",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
}

func (a *Auth) SocialCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	code := r.FormValue("code")

	token, err := a.service.SocialCallback(context.Background(), provider, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := a.service.Profile(context.Background(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Expires:  time.Now().Add(8760 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *Auth) SocialGetRedirectURL(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	url, err := a.service.SocialGetRedirectURL(context.Background(), provider)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)
}
