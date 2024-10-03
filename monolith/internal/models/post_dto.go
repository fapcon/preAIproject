package models

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"time"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type PostDTO struct {
	ID               int            `json:"id" db:"id" db_ops:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id"`
	Title            string         `json:"title" db:"title" db_ops:"create,update" db_type:"varchar(55)" db_default:"not null" mapper:"title"`
	ShortDescription string         `json:"short_description" db:"short_description" db_ops:"create,update" db_type:"varchar(55)" db_default:"null" mapper:"short_description"`
	FullDescription  string         `json:"full_description" db:"full_description" db_ops:"create,update" db_type:"varchar(55)" db_default:"null" mapper:"full_description"`
	Author           int            `json:"author" db:"author" db_ops:"create,update" db_type:"varchar(55)" db_default:"not null" mapper:"author"`
	CreatedAt        time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"create" mapper:"created_at"`
	UpdatedAt        types.NullTime `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"updated_at" mapper:"updated_at"`
	DeletedAt        types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *PostDTO) TableName() string {
	return "post"
}

func (s *PostDTO) OnCreate() []string {
	return []string{}
}

func (s *PostDTO) SetID(id int) *PostDTO {
	s.ID = id
	return s
}

func (s *PostDTO) GetID() int {
	return s.ID
}

func (s *PostDTO) SetTittle(tittle string) *PostDTO {
	s.Title = tittle
	return s
}

func (s *PostDTO) GetTittle() string {
	return s.Title
}

func (s *PostDTO) SetShortDescription(ShortDescription string) *PostDTO {
	s.ShortDescription = ShortDescription
	return s
}

func (s *PostDTO) GetShortDescription() string {
	return s.ShortDescription
}

func (s *PostDTO) SetFullDescription(FullDescription string) *PostDTO {
	s.FullDescription = FullDescription
	return s
}

func (s *PostDTO) GetFullDescription() string {
	return s.FullDescription
}

func (s *PostDTO) SetAuthor(Author int) *PostDTO {
	s.Author = Author
	return s
}

func (s *PostDTO) GetAuthor() int {
	return s.Author
}

func (s *PostDTO) SetCreatedAt(createdAt time.Time) *PostDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *PostDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *PostDTO) SetUpdatedAt(updatedAt time.Time) *PostDTO {
	s.UpdatedAt.Time.Time = updatedAt
	s.UpdatedAt.Time.Valid = true
	return s
}

func (s *PostDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt.Time.Time
}

func (s *PostDTO) SetDeletedAt(deletedAt time.Time) *PostDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *PostDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}

//type PostDTO struct {
//	ID               int            `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" db_ops:"id" mapper:"id"`
//	Title            string         `json:"title" mapper:"title" db:"title" db_ops:"create,update" db_type:"title" db_default:"title"`
//	ShortDescription string         `json:"short_description" mapper:"short_description" db:"short_description" db_ops:"create,update" db_type:"short_description" db_default:"short_description"`
//	FullDescription  string         `json:"full_description" mapper:"full_description" db:"full_description" db_ops:"create,update" db_type:"full_description" db_default:"full_description"`
//	Author           int            `json:"author" mapper:"author" db:"author" db_ops:"create,update" db_type:"author" db_default:"author"`
//	CreatedAt        time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
//	UpdatedAt        time.Time      `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
//	DeletedAt        types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
//}
