package service

import (
	"context"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name CommentService
type CommentService interface {
	CreateComment(ctx context.Context, in CommentCreateIn) CommentCreateOut
	UpdateComment(ctx context.Context, in CommentUpdateIn) CommentUpdateOut
	GetCommentByID(ctx context.Context, in CommentGetByIdIn) CommentGetByIdOut
	DeleteComment(ctx context.Context, in CommentDeleteIn) CommentDeleteOut
	GetCommentList(ctx context.Context) CommentGetTapeOut
}

type CommentCreateIn struct {
	Comment  string `json:"comment" mapper:"comment"`
	AuthorID int    `json:"author" mapper:"author"`
	Source   int    `json:"source" mapper:"source"`
	SourceID int    `json:"source_id" mapper:"source_id"`
}

type CommentCreateOut struct {
	Status    int
	ErrorCode int
}

type CommentUpdateIn struct {
	Id       int    `json:"id"`
	Comment  string `json:"comment"`
	AuthorID int    `json:"author"`
}

type CommentUpdateOut struct {
	Success   bool
	ErrorCode int
}

type CommentGetByIdIn struct {
	Id       int `json:"id"`
	AuthorID int `json:"author"`
}

type CommentGetByIdOut struct {
	Success   bool
	ErrorCode int
	Body      Data
}

type Data struct {
	Comment  string `json:"comment"`
	Source   int    `json:"source" `
	SourceID int    `json:"source_id" `
}

type CommentDeleteIn struct {
	Id       int `json:"id"`
	AuthorID int `json:"author"`
}

type CommentDeleteOut struct {
	Success   bool
	ErrorCode int
}

type CommentGetTapeOut struct {
	Body      []models.CommentDTO
	ErrorCode int
	Success   bool
}
