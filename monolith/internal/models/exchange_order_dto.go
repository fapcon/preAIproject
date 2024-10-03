package models

import (
	"time"

	"github.com/shopspring/decimal"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type ExchangeOrderDTO struct {
	ID                int64           `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id" db_ops:"id"`
	UUID              string          `json:"uuid" db:"uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index,unique" mapper:"uuid"`
	BotUUID           string          `json:"bot_uuid" db:"bot_uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index" mapper:"bot_uuid"`
	OrderType         int             `json:"order_type" db:"order_type" db_ops:"create" db_type:"int" db_default:"default 0" mapper:"order_type"`
	ExchangeOrderType int             `json:"exchange_order_type" db:"exchange_order_type" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"exchange_order_type"`
	Side              int             `json:"side" db:"side" db_ops:"create" db_type:"int" db_default:"default 0" mapper:"side"`
	ExchangeOrderID   string          `json:"exchange_order_id" db:"exchange_order_id" db_ops:"create,update" db_type:"char(36)" db_default:"" db_index:"index" mapper:"exchange_order_id"`
	UnitedOrders      int             `json:"united_orders" db:"united_orders" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"united_orders"`
	UserID            int             `json:"user_id" db:"user_id" db_ops:"create" db_type:"bigint" db_default:"default 1" db_index:"index" mapper:"user_id"`
	ExchangeID        int             `json:"exchange_id" db:"exchange_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"exchange_id"`
	PairID            int             `json:"pair_id" db:"pair_id" db_ops:"create" db_type:"int" db_default:"not null" mapper:"pair_id"`
	BuyOrderID        int64           `json:"order_id" db:"order_id" db_ops:"create" db_type:"int" db_default:"default 0" mapper:"order_id"`
	ApiKeyID          int             `json:"api_key_id" db:"api_key_id" db_ops:"create" db_type:"int" db_default:"default 0" mapper:"api_key_id"`
	Pair              string          `json:"pair" db:"pair" db_ops:"create" db_type:"varchar(21)" db_default:"not null" mapper:"pair"`
	Amount            decimal.Decimal `json:"amount" db:"amount" db_ops:"create,update" db_type:"decimal(34,8)" db_default:"default 0" mapper:"amount"`
	Quantity          decimal.Decimal `json:"quantity" db:"quantity" db_ops:"create,update" db_type:"decimal(34,8)" db_default:"default 0" mapper:"quantity"`
	Price             decimal.Decimal `json:"price" db:"price" db_ops:"create,update,upsert" db_type:"decimal(34,8)" db_default:"default 0" mapper:"price"`
	BuyPrice          decimal.Decimal `json:"buy_price" db:"buy_price" db_ops:"create,update,upsert" db_type:"decimal(34,8)" db_default:"default 0" mapper:"buy_price"`
	Status            int             `json:"status" db:"status" db_ops:"create,update,upsert" db_type:"int" db_default:"default 0" mapper:"status"`
	WebhookUUID       string          `json:"webhook_uuid" db:"webhook_uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" mapper:"webhook_uuid"`
	Message           string          `json:"message" db:"message" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"message"`
	CreatedAt         time.Time       `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt         types.NullTime  `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (e *ExchangeOrderDTO) TableName() string {
	return "exchange_order"
}

func (e *ExchangeOrderDTO) OnCreate() []string {
	return []string{}
}

func (e *ExchangeOrderDTO) SetUUID(uuid string) *ExchangeOrderDTO {
	e.UUID = uuid
	return e
}

func (e *ExchangeOrderDTO) GetUUID() string {
	return e.UUID
}

func (e *ExchangeOrderDTO) SetExchangeOrderID(exchangeOrderID string) *ExchangeOrderDTO {
	e.ExchangeOrderID = exchangeOrderID
	return e
}

func (e *ExchangeOrderDTO) GetExchangeOrderID() string {
	return e.ExchangeOrderID
}

func (e *ExchangeOrderDTO) SetUserID(userId int) *ExchangeOrderDTO {
	e.UserID = userId
	return e
}

func (e *ExchangeOrderDTO) GetUserID() int {
	return e.UserID
}

func (e *ExchangeOrderDTO) SetStatus(status int) *ExchangeOrderDTO {
	e.Status = status
	return e
}

func (e *ExchangeOrderDTO) GetStatus() int {
	return e.Status
}

func (e *ExchangeOrderDTO) SetAmount(amount decimal.Decimal) *ExchangeOrderDTO {
	e.Amount = amount
	return e
}

func (e *ExchangeOrderDTO) GetAmount() decimal.Decimal {
	return e.Amount
}

func (e *ExchangeOrderDTO) SetPrice(price decimal.Decimal) *ExchangeOrderDTO {
	e.Price = price
	return e
}

func (e *ExchangeOrderDTO) GetPrice() decimal.Decimal {
	return e.Price
}

func (e *ExchangeOrderDTO) SetQuantity(quantity decimal.Decimal) *ExchangeOrderDTO {
	e.Quantity = quantity
	return e
}

func (e *ExchangeOrderDTO) GetQuantity() decimal.Decimal {
	return e.Quantity
}

func (e *ExchangeOrderDTO) SetSide(side int) *ExchangeOrderDTO {
	e.Side = side
	return e
}

func (e *ExchangeOrderDTO) GetSide() int {
	return e.Side
}

func (e *ExchangeOrderDTO) SetWebhookUUID(webhookUUID string) *ExchangeOrderDTO {
	e.WebhookUUID = webhookUUID
	return e
}

func (e *ExchangeOrderDTO) GetWebhookUUID() string {
	return e.WebhookUUID
}

func (e *ExchangeOrderDTO) SetMessage(message string) *ExchangeOrderDTO {
	e.Message = message
	return e
}

func (e *ExchangeOrderDTO) GetMessage() string {
	return e.Message
}

func (e *ExchangeOrderDTO) SetExchangeID(exchangeID int) *ExchangeOrderDTO {
	e.ExchangeID = exchangeID
	return e
}

func (e *ExchangeOrderDTO) GetExchangeID() int {
	return e.ExchangeID
}

func (e *ExchangeOrderDTO) SetCreatedAt(createdAt time.Time) *ExchangeOrderDTO {
	e.CreatedAt = createdAt
	return e
}

func (e *ExchangeOrderDTO) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *ExchangeOrderDTO) SetUpdatedAt(updatedAt time.Time) *ExchangeOrderDTO {
	e.UpdatedAt = updatedAt
	return e
}

func (e *ExchangeOrderDTO) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (e *ExchangeOrderDTO) SetDeletedAt(deletedAt time.Time) *ExchangeOrderDTO {
	e.DeletedAt.Time.Time = deletedAt
	e.DeletedAt.Time.Valid = true
	return e
}

func (e *ExchangeOrderDTO) GetDeletedAt() time.Time {
	return e.DeletedAt.Time.Time
}

func (e *ExchangeOrderDTO) GetID() time.Time {
	return e.DeletedAt.Time.Time
}
