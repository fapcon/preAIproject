package models

import (
	"time"

	"github.com/shopspring/decimal"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default

type ExchangeTickerDTO struct {
	ID         int             `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" db_ops:"id"`
	Pair       string          `json:"pair" db:"pair" db_ops:"create,update,conflict" db_type:"varchar(21)" db_default:"not null"`
	ExchangeID int             `json:"exchange_id" db:"exchange_id" db_ops:"create,update,conflict" db_type:"int" db_default:"default 1"`
	Price      decimal.Decimal `json:"price" db:"price" db_ops:"create,update,upsert" db_type:"decimal(34,8)" db_default:"default 0"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" db:"updated_at" db_ops:"update,upsert" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
	DeletedAt  types.NullTime  `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at"`
}

func (s *ExchangeTickerDTO) TableName() string {
	return "exchange_ticker"
}

func (s *ExchangeTickerDTO) OnCreate() []string {
	return []string{
		"create unique index exchange_ticker_pair_exchange_id_index on exchange_ticker (pair, exchange_id);",
	}
}

func (s *ExchangeTickerDTO) SetID(id int) *ExchangeTickerDTO {
	s.ID = id
	return s
}

func (s *ExchangeTickerDTO) GetCurrency() string {
	return s.Pair
}

func (s *ExchangeTickerDTO) SetCurrency(currency string) *ExchangeTickerDTO {
	s.Pair = currency
	return s
}

func (s *ExchangeTickerDTO) GetPrice() decimal.Decimal {
	return s.Price
}

func (s *ExchangeTickerDTO) SetPrice(price decimal.Decimal) *ExchangeTickerDTO {
	s.Price = price
	return s
}

func (s *ExchangeTickerDTO) GetExchangeID() int {
	return s.ExchangeID
}

func (s *ExchangeTickerDTO) SetExchangeID(id int) *ExchangeTickerDTO {
	s.ExchangeID = id
	return s
}

func (s *ExchangeTickerDTO) GetID() int {
	return s.ID
}

func (s *ExchangeTickerDTO) SetCreatedAt(createdAt time.Time) *ExchangeTickerDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *ExchangeTickerDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *ExchangeTickerDTO) SetUpdatedAt(updatedAt time.Time) *ExchangeTickerDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *ExchangeTickerDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *ExchangeTickerDTO) SetDeletedAt(deletedAt time.Time) *ExchangeTickerDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *ExchangeTickerDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}
