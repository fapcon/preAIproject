package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default
type ExchangeListDTO struct {
	ID          int            `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" db_ops:"id"`
	Name        string         `json:"name" db:"name" db_ops:"create,update" db_type:"varchar(21)" db_default:"not null" db_index:"index,unique"`
	Description string         `json:"description" db:"description" db_ops:"create,update" db_type:"varchar(55)" db_default:"not null"`
	Slug        string         `json:"slug" db:"slug" db_ops:"create,update" db_type:"varchar(21)" db_default:"not null"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at" db_ops:"create,update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
	DeletedAt   types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at"`
}

func (s *ExchangeListDTO) TableName() string {
	return "exchange_list"
}

func (s *ExchangeListDTO) OnCreate() []string {
	return []string{}
}

func (s *ExchangeListDTO) SetID(id int) *ExchangeListDTO {
	s.ID = id
	return s
}

func (s *ExchangeListDTO) GetID() int {
	return s.ID
}

func (s *ExchangeListDTO) SetName(name string) *ExchangeListDTO {
	s.Name = name
	return s
}

func (s *ExchangeListDTO) GetName() string {
	return s.Name
}

func (s *ExchangeListDTO) SetDescription(description string) *ExchangeListDTO {
	s.Description = description
	return s
}

func (s *ExchangeListDTO) GetDescription() string {
	return s.Description
}

func (s *ExchangeListDTO) SetCreatedAt(createdAt time.Time) *ExchangeListDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *ExchangeListDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *ExchangeListDTO) SetUpdatedAt(updatedAt time.Time) *ExchangeListDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *ExchangeListDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *ExchangeListDTO) SetDeletedAt(deletedAt time.Time) *ExchangeListDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}
