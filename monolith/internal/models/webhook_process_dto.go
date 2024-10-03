package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type WebhookProcessDTO struct {
	ID            int            `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" db_ops:"id" mapper:"id"`
	UUID          string         `json:"uuid" db:"uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index,unique" mapper:"uuid"`
	BotUUID       string         `json:"bot_uuid" db:"bot_uuid" db_ops:"create,update" db_type:"char(36)" db_default:"not null" mapper:"bot_uuid"`
	OrderID       int64          `json:"order_id" db:"order_id" db_ops:"create,update" db_type:"bigint" db_default:"default -1" mapper:"order_id"`
	OrderUUID     string         `json:"order_uuid" db:"order_uuid" db_ops:"create,update" db_type:"char(36)" db_default:"not null" mapper:"order_uuid"`
	BotID         int            `json:"bot_id" db:"bot_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"bot_id"`
	UserID        int            `json:"user_id" db:"user_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"user_id"`
	Slug          string         `json:"slug" db:"slug" db_ops:"create,update" db_type:"varchar(100)" db_default:"not null" mapper:"slug"`
	Status        int            `json:"status" db:"status" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"status"`
	Message       string         `json:"message" db:"message" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"message"`
	XForwardedFor string         `json:"x_forwarded_for" db:"x_forwarded_for" db_ops:"create" db_type:"varchar(100)" db_default:"not null" mapper:"x_forwarded_for"`
	RemoteAddr    string         `json:"remote_addr" db:"remote_addr" db_ops:"create" db_type:"varchar(100)" db_default:"not null" mapper:"remote_addr"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at" db_ops:"create,update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt     types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *WebhookProcessDTO) TableName() string {
	return "webhook_process"
}

func (s *WebhookProcessDTO) OnCreate() []string {
	return []string{}
}

func (s *WebhookProcessDTO) SetID(id int) *WebhookProcessDTO {
	s.ID = id
	return s
}

func (s *WebhookProcessDTO) GetID() int {
	return s.ID
}

func (s *WebhookProcessDTO) SetUUID(uuid string) *WebhookProcessDTO {
	s.UUID = uuid
	return s
}

func (s *WebhookProcessDTO) GetUUID() string {
	return s.UUID
}

func (s *WebhookProcessDTO) SetCreatedAt(createdAt time.Time) *WebhookProcessDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *WebhookProcessDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *WebhookProcessDTO) SetUpdatedAt(updatedAt time.Time) *WebhookProcessDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *WebhookProcessDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *WebhookProcessDTO) SetDeletedAt(deletedAt time.Time) *WebhookProcessDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}
