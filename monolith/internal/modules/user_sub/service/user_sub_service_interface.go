package service

import (
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name UserSuber
type UserSuber interface {
	Add(ctx context.Context, in UserSubAddIn) UserSubOut
	List(ctx context.Context, in UserSubListIn) UserSubListOut
	Delete(ctx context.Context, in UserSubDelIn) UserSubOut
}

type UserSubAddIn struct {
	UserID    int
	SubUserID int
}

type UserSubDelIn struct {
	UserID    int
	SubUserID int
}

type UserSubOut struct {
	Success   bool
	ErrorCode int
}

type UserSubListIn struct {
	UserID int
}

type UserSubListOut struct {
	SubUserIDs []int
	Success    bool
	ErrorCode  int
}
