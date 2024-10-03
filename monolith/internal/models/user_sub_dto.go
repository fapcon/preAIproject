package models

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"time"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type UserSubDTO struct {
	ID        int            `json:"id" db:"id" db_ops:"" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id"`
	UserID    int            `json:"user_id" db:"user_id" db_ops:"create,update" db_type:"bigint" db_default:"not null" db_index:"index" mapper:"user_id"`
	SubUserID int            `json:"sub_user_id" db:"sub_user_id" db_ops:"create,update" db_type:"bigint" db_default:"not null" db_index:"index" mapper:"sub_user_id"`
	CreatedAt time.Time      `json:"created_at" db:"created_at" db_ops:"" db_type:"timestamp" db_default:"default (now()) not null" mapper:"created_at"`
	DeletedAt types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" mapper:"deleted_at"`
}

func (s *UserSubDTO) TableName() string {
	return "users_sub"
}

func (s *UserSubDTO) OnCreate() []string {
	return []string{}
}

func (s *UserSubDTO) SetID(id int) *UserSubDTO {
	s.ID = id
	return s
}

func (s *UserSubDTO) GetID() int {
	return s.ID
}

func (s *UserSubDTO) SetUserID(userID int) *UserSubDTO {
	s.UserID = userID
	return s
}

func (s *UserSubDTO) GetUserID() int {
	return s.UserID
}

func (s *UserSubDTO) SetSubscriberID(id int) *UserSubDTO {
	s.SubUserID = id
	return s
}

func (s *UserSubDTO) GetSubscriberID() int {
	return s.SubUserID
}

func (s *UserSubDTO) SetCreatedAt(createdAt time.Time) *UserSubDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *UserSubDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *UserSubDTO) SetDeletedAt(deletedAt time.Time) *UserSubDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}
