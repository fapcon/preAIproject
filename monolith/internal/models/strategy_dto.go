package models

import (
	"strconv"
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type StrategyDTO struct {
	ID          int      `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id" db_ops:"id"`
	ExchangeID  int      `json:"exchange_id" db:"exchange_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"exchange_id"`
	Name        string   `json:"name" db:"name" db_ops:"create,update" db_type:"varchar(55)" db_default:"not null" mapper:"name"`
	Description string   `json:"description" db:"description" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"description"`
	UUID        string   `json:"uuid" db:"uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index,unique" mapper:"uuid"`
	Bots        []string `json:"bots" db:"bots" db_ops:"create,update" db_type:"text[]" db_default:"null" mapper:"bots"`
	//UserID     int            `json:"user_id" db:"user_id" db_ops:"create" db_type:"int" db_default:"default 1" mapper:"user_id"`
	//APIKey     string         `json:"api_key" db:"api_key" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"api_key"`
	//SecretKey  string         `json:"secret_key" db:"secret_key" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"secret_key"`
	CreatedAt time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *StrategyDTO) TableName() string {
	return "strategy"
}

func (s *StrategyDTO) OnCreate() []string {
	return []string{}
}

func (s *StrategyDTO) SetID(id int) *StrategyDTO {
	s.ID = id
	return s
}

func (s *StrategyDTO) SetBots(bots []Bot) *StrategyDTO {
	var botList []string

	for i := 0; i < len(bots); i++ {
		id := strconv.Itoa(bots[i].ID)
		botList = append(botList, id)
	}
	s.Bots = botList
	return s
}

func (s *StrategyDTO) GetID() int {
	return s.ID
}

func (s *StrategyDTO) SetExchangeID(id int) *StrategyDTO {
	s.ExchangeID = id
	return s
}

func (s *StrategyDTO) GetExchangeID() int {
	return s.ExchangeID
}

//func (s *StrategyDTO) SetUserID(id int) *StrategyDTO {
//	s.UserID = id
//	return s
//}

//func (s *StrategyDTO) GetUserID() int {
//	return s.UserID
//}
//
//func (s *StrategyDTO) SetAPIKey(apiKey string) *StrategyDTO {
//	s.APIKey = apiKey
//	return s
//}
//
//func (s *StrategyDTO) GetAPIKey() string {
//	return s.APIKey
//}

//func (s *StrategyDTO) SetSecretKey(secretKey string) *StrategyDTO {
//	s.SecretKey = secretKey
//	return s
//}
//
//func (s *StrategyDTO) GetSecretKey() string {
//	return s.SecretKey
//}

func (s *StrategyDTO) SetDescription(description string) *StrategyDTO {
	s.Description = description
	return s
}

func (s *StrategyDTO) SetName(name string) *StrategyDTO {
	s.Name = name
	return s
}

func (s *StrategyDTO) SetUUID(uuid string) *StrategyDTO {
	s.UUID = uuid
	return s
}

func (s *StrategyDTO) SetCreatedAt(createdAt time.Time) *StrategyDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *StrategyDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *StrategyDTO) SetUpdatedAt(updatedAt time.Time) *StrategyDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *StrategyDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *StrategyDTO) SetDeletedAt(deletedAt time.Time) *StrategyDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *StrategyDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}
