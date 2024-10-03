package controller

import "net/http"

type ExchangeFacader interface {
	Ticker(w http.ResponseWriter, r *http.Request)
	ListAdd(w http.ResponseWriter, r *http.Request)
	ListDelete(w http.ResponseWriter, r *http.Request)
	ListAll(w http.ResponseWriter, r *http.Request)
	UserKeyAdd(w http.ResponseWriter, r *http.Request)
	UserKeyDelete(w http.ResponseWriter, r *http.Request)
	UserKeyList(w http.ResponseWriter, r *http.Request)
	OrderHistory(w http.ResponseWriter, r *http.Request)
}
