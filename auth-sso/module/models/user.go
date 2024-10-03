package models

import (
	"database/sql"
)

type User struct {
	ID           int64        `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null"`
	Email        string       `json:"email" db:"email" db_type:"varchar(89)" db_default:"not null" db_index:"index,unique" db_ops:"create,update"`
	Password     []byte       `json:"-" db:"password" db_type:"varchar(144)" db_default:"not null" db_ops:"create,update"`
	IsAdmin      bool         `json:"is_admin" db:"is_admin" db_type:"boolean" db_default:"default false" db_ops:"create,update"`
	DeleteStatus bool         `json:"delete_status" db:"delete_status" db_type:"boolean" db_default:"default false" db_ops:"create,update"`
	CreatedAt    string       `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_ops:"created_at" db_index:"index" mapper:"created_at"`
	UpdatedAt    string       `json:"updated_at" db:"updated_at" db_type:"timestamp" db_default:"default (now()) not null" db_ops:"updated_at" db_index:"index" mapper:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) OnCreate() []string {
	return []string{}
}
