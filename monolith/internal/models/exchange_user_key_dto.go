package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type ExchangeUserKeyDTO struct {
	ID         int            `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id" db_ops:"id"`
	ExchangeID int            `json:"exchange_id" db:"exchange_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"exchange_id"`
	UserID     int            `json:"user_id" db:"user_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"user_id"`
	Label      string         `json:"label" db:"label" db_ops:"create,update" db_type:"varchar(55)" db_default:"default 0" mapper:"label"`
	MakeOrder  bool           `json:"make_order" db:"make_order" db_ops:"create,update" db_type:"boolean" db_default:"default true" mapper:"make_order"`
	APIKey     string         `json:"api_key" db:"api_key" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"api_key"`
	SecretKey  string         `json:"secret_key" db:"secret_key" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"secret_key"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt  types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *ExchangeUserKeyDTO) TableName() string {
	return "exchange_user_key"
}

func (s *ExchangeUserKeyDTO) OnCreate() []string {
	return []string{}
}

func (s *ExchangeUserKeyDTO) SetID(id int) *ExchangeUserKeyDTO {
	s.ID = id
	return s
}

func (s *ExchangeUserKeyDTO) GetID() int {
	return s.ID
}

func (s *ExchangeUserKeyDTO) SetExchangeID(id int) *ExchangeUserKeyDTO {
	s.ExchangeID = id
	return s
}

func (s *ExchangeUserKeyDTO) GetExchangeID() int {
	return s.ExchangeID
}

func (s *ExchangeUserKeyDTO) SetUserID(id int) *ExchangeUserKeyDTO {
	s.UserID = id
	return s
}

func (s *ExchangeUserKeyDTO) GetUserID() int {
	return s.UserID
}

func (s *ExchangeUserKeyDTO) SetAPIKey(apiKey string) *ExchangeUserKeyDTO {
	s.APIKey = apiKey
	return s
}

func (s *ExchangeUserKeyDTO) GetAPIKey() string {
	return s.APIKey
}

func (s *ExchangeUserKeyDTO) SetSecretKey(secretKey string) *ExchangeUserKeyDTO {
	s.SecretKey = secretKey
	return s
}

func (s *ExchangeUserKeyDTO) GetSecretKey() string {
	return s.SecretKey
}

func (s *ExchangeUserKeyDTO) SetCreatedAt(createdAt time.Time) *ExchangeUserKeyDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *ExchangeUserKeyDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *ExchangeUserKeyDTO) SetUpdatedAt(updatedAt time.Time) *ExchangeUserKeyDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *ExchangeUserKeyDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *ExchangeUserKeyDTO) SetDeletedAt(deletedAt time.Time) *ExchangeUserKeyDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *ExchangeUserKeyDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}

func (s *ExchangeUserKeyDTO) SetLabel(label string) *ExchangeUserKeyDTO {
	s.Label = label
	return s
}
func (s *ExchangeUserKeyDTO) GetLabel() string {
	return s.Label
}
