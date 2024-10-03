package models

type Balance struct {
	Currency string  `json:"currency" mapper:"currency"'`
	Amount   float64 `json:"amount" mapper:"amount"`
	Locked   float64 `json:"locked" mapper:"locked"`
}
