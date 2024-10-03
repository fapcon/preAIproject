package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type CommentDTO struct {
	ID        int            `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" db_ops:"id" mapper:"id"`
	AuthorID  int            `json:"author_id" db:"author_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"author_id"`
	Comment   string         `json:"comment" db:"comment" db_ops:"create,update" db_type:"text" db_default:"not null" mapper:"comment"`
	CreatedAt time.Time      `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"create" mapper:"created_at"`
	UpdatedAt types.NullTime `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"updated_at" mapper:"updated_at"`
	DeletedAt types.NullTime `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
	SourceID  int            `json:"source_id" db:"source_id" db_ops:"create,update" db_type:"varchar(255)" db_default:"not null" mapper:"source_id"`
	Source    int            `json:"source" db:"source" db_ops:"create,update" db_type:"varchar(255)" db_default:"not null" mapper:"source"`
}

func (c *CommentDTO) TableName() string {
	return "comments"
}

func (c *CommentDTO) OnCreate() []string {
	return []string{}
}

func (c *CommentDTO) SetID(id int) *CommentDTO {
	c.ID = id
	return c
}

func (c *CommentDTO) GetID() int {
	return c.ID
}

func (c *CommentDTO) SetAuthorID(authorID int) *CommentDTO {
	c.AuthorID = authorID
	return c
}

func (c *CommentDTO) GetAuthorID() int {
	return c.AuthorID
}

func (c *CommentDTO) SetSourceID(sourceId int) *CommentDTO {
	c.SourceID = sourceId
	return c
}

func (c *CommentDTO) GetSourceID() int {
	return c.SourceID
}

func (c *CommentDTO) SetSource(source int) *CommentDTO {
	c.Source = source
	return c
}

func (c *CommentDTO) GetSource() int {
	return c.Source
}

func (c *CommentDTO) SetComment(comment string) *CommentDTO {
	c.Comment = comment
	return c
}

func (c *CommentDTO) GetComment() string {
	return c.Comment
}

func (s *CommentDTO) SetCreatedAt(createdAt time.Time) *CommentDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *CommentDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *CommentDTO) SetUpdatedAt(updatedAt time.Time) *CommentDTO {
	s.UpdatedAt.Time.Time = updatedAt
	s.UpdatedAt.Time.Valid = true
	return s
}

func (s *CommentDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt.Time.Time
}

func (s *CommentDTO) SetDeletedAt(deletedAt time.Time) *CommentDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *CommentDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}
