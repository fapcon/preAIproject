package controller

import (
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules"
)

type ExchangeFacadeController struct {
	controllers modules.Controllers
}

func NewExchangeFacadeController(controllers modules.Controllers) ExchangeFacader {
	return ExchangeFacadeController{controllers: controllers}
}

// Ticker
func (c ExchangeFacadeController) Ticker(w http.ResponseWriter, r *http.Request) {
	c.controllers.Ticker.Ticker(w, r)
}

// ExchangeList
func (c ExchangeFacadeController) ListAdd(w http.ResponseWriter, r *http.Request) {
	c.controllers.List.ExchangeAdd(w, r)
}

func (c ExchangeFacadeController) ListDelete(w http.ResponseWriter, r *http.Request) {
	c.controllers.List.ExchangeDelete(w, r)
}

func (c ExchangeFacadeController) ListAll(w http.ResponseWriter, r *http.Request) {
	c.controllers.List.ExchangeList(w, r)
}

// User Key
func (c ExchangeFacadeController) UserKeyAdd(w http.ResponseWriter, r *http.Request) {
	c.controllers.UserKey.ExchangeUserKeyAdd(w, r)
}

func (c ExchangeFacadeController) UserKeyDelete(w http.ResponseWriter, r *http.Request) {
	c.controllers.UserKey.ExchangeUserKeyDelete(w, r)
}

func (c ExchangeFacadeController) UserKeyList(w http.ResponseWriter, r *http.Request) {
	c.controllers.UserKey.ExchangeUserKeyList(w, r)
}

// Order
func (c ExchangeFacadeController) OrderHistory(w http.ResponseWriter, r *http.Request) {
	c.controllers.Order.UserOrdersHistory(w, r)
}
