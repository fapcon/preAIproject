package service

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Poster
type Poster interface {
	Create(ctx context.Context, in PostCreateIn) PostCreateOut
	UpdatePost(ctx context.Context, in PostUpdateIn) PostUpdateOut
	Delete(ctx context.Context, in PostDeleteIn) PostDeleteOut
	GetById(ctx context.Context, in PostGetByIdIn) PostGetByIdOut
	GetListTape(ctx context.Context, in PostGetTapeIn) PostGetTapeOut
}

type PostCreateIn struct {
	Title            string `json:"title" mapper:"title"`
	ShortDescription string `json:"short_description" mapper:"short_description"`
	FullDescription  string `json:"full_description" mapper:"full_description"`
	Author           int    `json:"author" mapper:"author"`
}

type PostCreateOut struct {
	Status    int
	ErrorCode int
}

type PostDeleteIn struct {
	Id     int `json:"id"`
	Author int `json:"author"`
}

type PostDeleteOut struct {
	Success   bool
	ErrorCode int
}

type PostUpdateIn struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
	Author           int    `json:"author"`
}

type PostUpdateOut struct {
	Success   bool
	ErrorCode int
}

type PostGetByIdIn struct {
	Id     int `json:"id"`
	Author int `json:"author"`
}

type PostGetByIdOut struct {
	Success   bool
	ErrorCode int
	Body      Data
}

type Data struct {
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
}

type PostGetTapeIn struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type PostGetTapeOut struct {
	Body      []models.PostDTO
	ErrorCode int
	Success   bool
}
