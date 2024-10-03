package models

import (
	"github.com/shopspring/decimal"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"time"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type BotStatisticsDTO struct {
	ID            int             `json:"id" db:"id" db_ops:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id"`
	BotUUID       string          `json:"bot_uuid" db:"bot_uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" mapper:"bot_uuid"`
	UserID        int             `json:"user_id" db:"user_id" db_ops:"create" db_type:"int" db_default:"not null" mapper:"user_id"`
	Profitability decimal.Decimal `json:"profitability" db:"profitability" db_ops:"create,update" db_type:"decimal(34,8)" db_default:"default 0" mapper:"profitability"`
	OrderCount    int             `json:"order_count" db:"order_count" db_ops:"create,update" db_type:"int" db_default:"default 0" mapper:"order_count"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt     types.NullTime  `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *BotStatisticsDTO) TableName() string {
	return "bot_statistics"
}

func (s *BotStatisticsDTO) OnCreate() []string {
	return []string{}
}

func (s *BotStatisticsDTO) SetID(id int) *BotStatisticsDTO {
	s.ID = id
	return s
}

func (s *BotStatisticsDTO) GetID() int {
	return s.ID
}

func (s *BotStatisticsDTO) SetUUID(uuid string) *BotStatisticsDTO {
	s.BotUUID = uuid
	return s
}

func (s *BotStatisticsDTO) GetUUID() string {
	return s.BotUUID
}

func (s *BotStatisticsDTO) SetUserID(userId int) *BotStatisticsDTO {
	s.UserID = userId
	return s
}

func (s *BotStatisticsDTO) GetUserID() int {
	return s.UserID
}

func (s *BotStatisticsDTO) SetProfitability(profitability decimal.Decimal) *BotStatisticsDTO {
	s.Profitability = profitability
	return s
}

func (s *BotStatisticsDTO) GetProfitability() decimal.Decimal {
	return s.Profitability
}

func (s *BotStatisticsDTO) SetOrderCount(orderCount int) *BotStatisticsDTO {
	s.OrderCount = orderCount
	return s
}

func (s *BotStatisticsDTO) GetOrderCount() int {
	return s.OrderCount
}

func (s *BotStatisticsDTO) SetCreatedAt(createdAt time.Time) *BotStatisticsDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *BotStatisticsDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *BotStatisticsDTO) SetUpdatedAt(updatedAt time.Time) *BotStatisticsDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *BotStatisticsDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *BotStatisticsDTO) SetDeletedAt(deletedAt time.Time) *BotStatisticsDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *BotStatisticsDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}
