package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type StrategyPairDTO struct {
	ID         int            `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id" db_ops:"id"`
	StrategyID int            `json:"strategy_id" db:"strategy_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"strategy_id"`
	PairID     int            `json:"pair_id" db:"pair_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"pair_id"`
	UserID     int            `json:"user_id" db:"user_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"user_id"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at" db_ops:"create,update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt  types.NullTime `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *StrategyPairDTO) TableName() string {
	return "strategy_pair"
}

func (s *StrategyPairDTO) OnCreate() []string {
	return []string{}
}

func (s *StrategyPairDTO) SetID(id int) *StrategyPairDTO {
	s.ID = id
	return s
}

func (s *StrategyPairDTO) GetID() int {
	return s.ID
}

func (s *StrategyPairDTO) SetStrategyID(strategyID int) *StrategyPairDTO {
	s.StrategyID = strategyID
	return s
}

func (s *StrategyPairDTO) GetStrategyID() int {
	return s.StrategyID
}

func (s *StrategyPairDTO) SetPairID(pairID int) *StrategyPairDTO {
	s.PairID = pairID
	return s
}

func (s *StrategyPairDTO) GetPairID() int {
	return s.PairID
}

func (s *StrategyPairDTO) SetCreatedAt(createdAt time.Time) *StrategyPairDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *StrategyPairDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *StrategyPairDTO) SetUpdatedAt(updatedAt time.Time) *StrategyPairDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *StrategyPairDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *StrategyPairDTO) SetDeletedAt(deletedAt time.Time) *StrategyPairDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *StrategyPairDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}
