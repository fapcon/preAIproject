package models

import (
	"time"

	"github.com/shopspring/decimal"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type ExchangeOrderLogDTO struct {
	ID         int64           `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id" db_ops:"id"`
	UUID       string          `json:"uuid" db:"uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index" mapper:"uuid"`
	OrderID    int64           `json:"order_id" db:"order_id" db_ops:"create" db_type:"bigint" db_default:"default 0" mapper:"order_id" db_index:"index"`
	UserID     int             `json:"user_id" db:"user_id" db_ops:"create" db_type:"int" db_default:"default 1" db_index:"index" mapper:"user_id"`
	ExchangeID int             `json:"exchange_id" db:"exchange_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"exchange_id"`
	Pair       string          `json:"pair" db:"pair" db_ops:"create" db_type:"varchar(21)" db_default:"not null" mapper:"pair"`
	Quantity   decimal.Decimal `json:"quantity" db:"quantity" db_ops:"create" db_type:"decimal(34,8)" db_default:"default 0" mapper:"quantity"`
	Amount     decimal.Decimal `json:"amount" db:"amount" db_ops:"create,update" db_type:"decimal(34,8)" db_default:"default 0" mapper:"amount"`
	Price      decimal.Decimal `json:"price" db:"price" db_ops:"create,update,upsert" db_type:"decimal(34,8)" db_default:"default 0"`
	Status     int             `json:"status" db:"status" db_ops:"create,update,upsert" db_type:"int" db_default:"default 0" mapper:"status"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt  types.NullTime  `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (e *ExchangeOrderLogDTO) TableName() string {
	return "exchange_order_log"
}

func (e *ExchangeOrderLogDTO) OnCreate() []string {
	return []string{}
}

func (e *ExchangeOrderLogDTO) SetUUID(uuid string) *ExchangeOrderLogDTO {
	e.UUID = uuid
	return e
}

func (e *ExchangeOrderLogDTO) GetUUID() string {
	return e.UUID
}

func (e *ExchangeOrderLogDTO) SetOrderID(orderID int64) *ExchangeOrderLogDTO {
	e.OrderID = orderID
	return e
}

func (e *ExchangeOrderLogDTO) GetOrderID() int64 {
	return e.OrderID
}

func (e *ExchangeOrderLogDTO) SetUserID(userId int) *ExchangeOrderLogDTO {
	e.UserID = userId
	return e
}

func (e *ExchangeOrderLogDTO) GetUserID() int {
	return e.UserID
}

func (e *ExchangeOrderLogDTO) SetExchangeID(exchangeID int) *ExchangeOrderLogDTO {
	e.ExchangeID = exchangeID
	return e
}

func (e *ExchangeOrderLogDTO) GetExchangeID() int {
	return e.ExchangeID
}

func (e *ExchangeOrderLogDTO) SetCreatedAt(createdAt time.Time) *ExchangeOrderLogDTO {
	e.CreatedAt = createdAt
	return e
}

func (e *ExchangeOrderLogDTO) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *ExchangeOrderLogDTO) SetUpdatedAt(updatedAt time.Time) *ExchangeOrderLogDTO {
	e.UpdatedAt = updatedAt
	return e
}

func (e *ExchangeOrderLogDTO) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (e *ExchangeOrderLogDTO) SetDeletedAt(deletedAt time.Time) *ExchangeOrderLogDTO {
	e.DeletedAt.Time.Time = deletedAt
	e.DeletedAt.Time.Valid = true
	return e
}

func (e *ExchangeOrderLogDTO) GetDeletedAt() time.Time {
	return e.DeletedAt.Time.Time
}

func (e *ExchangeOrderLogDTO) GetID() time.Time {
	return e.DeletedAt.Time.Time
}
