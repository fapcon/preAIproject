package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type WebhookProcessHistoryDTO struct {
	ID          int              `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" db_ops:"id" mapper:"id"`
	UserID      int              `json:"user_id" db:"user_id" db_ops:"create" db_type:"int" db_default:"default 1" db_index:"index" mapper:"user_id"`
	ExchangeID  int              `json:"exchange_id" db:"exchange_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"exchange_id"`
	WebhookUUID string           `json:"webhook_uuid" db:"webhook_uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" mapper:"webhook_uuid"`
	WebhookID   int              `json:"webhook_id" db:"webhook_id" db_ops:"create" db_type:"bigint" db_default:"not null" mapper:"webhook_id"`
	Status      int              `json:"status" db:"status" db_ops:"create,update,upsert" db_type:"int" db_default:"default 0" mapper:"status"`
	Message     types.NullString `json:"message" db:"message" db_ops:"create,update,upsert" db_type:"varchar(255)" db_default:"null" mapper:"message"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt   types.NullTime   `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (e *WebhookProcessHistoryDTO) TableName() string {
	return "webhook_process_history"
}

func (e *WebhookProcessHistoryDTO) OnCreate() []string {
	return []string{}
}

func (e *WebhookProcessHistoryDTO) SetUserID(userId int) *WebhookProcessHistoryDTO {
	e.UserID = userId
	return e
}

func (e *WebhookProcessHistoryDTO) GetUserID() int {
	return e.UserID
}

func (e *WebhookProcessHistoryDTO) SetExchangeID(exchangeID int) *WebhookProcessHistoryDTO {
	e.ExchangeID = exchangeID
	return e
}

func (e *WebhookProcessHistoryDTO) GetExchangeID() int {
	return e.ExchangeID
}

func (e *WebhookProcessHistoryDTO) SetCreatedAt(createdAt time.Time) *WebhookProcessHistoryDTO {
	e.CreatedAt = createdAt
	return e
}

func (e *WebhookProcessHistoryDTO) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *WebhookProcessHistoryDTO) SetUpdatedAt(updatedAt time.Time) *WebhookProcessHistoryDTO {
	e.UpdatedAt = updatedAt
	return e
}

func (e *WebhookProcessHistoryDTO) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (e *WebhookProcessHistoryDTO) SetDeletedAt(deletedAt time.Time) *WebhookProcessHistoryDTO {
	e.DeletedAt.Time.Time = deletedAt
	e.DeletedAt.Time.Valid = true
	return e
}

func (e *WebhookProcessHistoryDTO) GetDeletedAt() time.Time {
	return e.DeletedAt.Time.Time
}

func (e *WebhookProcessHistoryDTO) GetID() time.Time {
	return e.DeletedAt.Time.Time
}
