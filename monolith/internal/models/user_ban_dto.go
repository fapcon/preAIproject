package models

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"time"
)

//go:generate easytags user_ban.go json,db,db_ops,db_type,db_default,db_index

const (
	StatusDefault = iota
	StatusBanned
)

type UsersBanedDTO struct {
	ID            int              `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null"`
	UserID        int              `json:"user_id" db:"user_id" db_type:"int" db_default:"default 0" db_ops:"create,update"`
	Reason        types.NullString `json:"reason" db:"reason" db_type:"varchar(200)" db_default:"default null" db_ops:"create,update"`
	BanStartedAt  time.Time        `json:"ban_started_at" db:"ban_started_at" db_type:"timestamp" db_default:"default (now()) not null" db_ops:"create,update" `
	BanFinishedAt time.Time        `json:"ban_finished_at" db:"ban_finished_at" db_type:"timestamp" db_default:"default (now()) not null" db_ops:"create,update"`
}

func (u *UsersBanedDTO) TableName() string {
	return "users_banned"
}

func (u *UsersBanedDTO) OnCreate() []string {
	return []string{}
}

func (u *UsersBanedDTO) SetID(id int) *UsersBanedDTO {
	u.ID = id
	return u
}

func (u *UsersBanedDTO) GetID() int {
	return u.ID
}

func (u *UsersBanedDTO) SetUserID(id int) *UsersBanedDTO {
	u.UserID = id
	return u
}

func (u *UsersBanedDTO) GetUserID() int {
	return u.UserID
}

func (u *UsersBanedDTO) SetReason(reason string) *UsersBanedDTO {
	u.Reason = types.NewNullString(reason)
	return u
}

func (u *UsersBanedDTO) GetReason() types.NullString {
	return u.Reason
}

func (u *UsersBanedDTO) SetBanStartedAt(t time.Time) *UsersBanedDTO {
	u.BanStartedAt = t
	return u
}

func (u *UsersBanedDTO) GetBanStartedAt() time.Time {
	return u.BanStartedAt
}

func (u *UsersBanedDTO) SetBanFinishedAt(t time.Time) *UsersBanedDTO {
	u.BanFinishedAt = t
	return u
}

func (u *UsersBanedDTO) GetBanFinishedAt() time.Time {
	return u.BanFinishedAt
}
