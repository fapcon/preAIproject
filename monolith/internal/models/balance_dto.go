package models

import "github.com/shopspring/decimal"

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type BalanceDTO struct {
	Currency string          `json:"currency" mapper:"currency" db:"currency" db_ops:"currency" db_type:"currency" db_default:"currency"`
	Amount   decimal.Decimal `json:"amount" mapper:"amount" db:"amount" db_ops:"amount" db_type:"amount" db_default:"amount"`
	Locked   decimal.Decimal `json:"locked" mapper:"locked" db:"locked" db_ops:"locked" db_type:"locked" db_default:"locked"`
}

func (s *BalanceDTO) TableName() string {
	return "Balance"
}
