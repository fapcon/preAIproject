package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type StrategySubscribersDTO struct {
	ID         int            `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id" db_ops:"id"`
	ExchangeID int            `json:"exchange_id" db:"exchange_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"exchange_id"`
	UserID     int            `json:"user_id" db:"user_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"user_id"`
	APIKey     string         `json:"api_key" db:"api_key" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"api_key"`
	SecretKey  string         `json:"secret_key" db:"secret_key" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"secret_key"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt  types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *StrategySubscribersDTO) TableName() string {
	return "strategy_subscribers"
}

func (s *StrategySubscribersDTO) OnCreate() []string {
	return []string{}
}

func (s *StrategySubscribersDTO) SetID(id int) *StrategySubscribersDTO {
	s.ID = id
	return s
}

func (s *StrategySubscribersDTO) GetID() int {
	return s.ID
}

func (s *StrategySubscribersDTO) SetExchangeID(id int) *StrategySubscribersDTO {
	s.ExchangeID = id
	return s
}

func (s *StrategySubscribersDTO) GetExchangeID() int {
	return s.ExchangeID
}

func (s *StrategySubscribersDTO) SetUserID(id int) *StrategySubscribersDTO {
	s.UserID = id
	return s
}

func (s *StrategySubscribersDTO) GetUserID() int {
	return s.UserID
}

func (s *StrategySubscribersDTO) SetAPIKey(apiKey string) *StrategySubscribersDTO {
	s.APIKey = apiKey
	return s
}

func (s *StrategySubscribersDTO) GetAPIKey() string {
	return s.APIKey
}

func (s *StrategySubscribersDTO) SetSecretKey(secretKey string) *StrategySubscribersDTO {
	s.SecretKey = secretKey
	return s
}

func (s *StrategySubscribersDTO) GetSecretKey() string {
	return s.SecretKey
}

func (s *StrategySubscribersDTO) SetCreatedAt(createdAt time.Time) *StrategySubscribersDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *StrategySubscribersDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *StrategySubscribersDTO) SetUpdatedAt(updatedAt time.Time) *StrategySubscribersDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *StrategySubscribersDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *StrategySubscribersDTO) SetDeletedAt(deletedAt time.Time) *StrategySubscribersDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *StrategySubscribersDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}
