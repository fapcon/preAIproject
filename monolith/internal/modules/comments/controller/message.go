package controller

import "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

type CommentUserCreateRequest struct {
	Comment  string `json:"comment" `
	Source   int    `json:"source" `
	SourceID int    `json:"source_id" `
}

type CommentCreateResponse struct {
	Status    int
	ErrorCode int
}

type CommentUserUpdateRequest struct {
	Id      int    `json:"id"`
	Comment string `json:"comment" `
}

type CommentUpdateResponse struct {
	Success   bool
	ErrorCode int
}

type CommentGetByIdRequest struct {
	Id int `json:"id"`
}

type CommentGetByIdResponse struct {
	Success   bool
	ErrorCode int
	Error     string
	Body      Data
}

type Data struct {
	Comment string `json:"comment" `
	Source  int    `json:"source" `
}

type CommentUserGetByIdResponse struct {
	Success   bool
	ErrorCode int
	Error     string
	Body      Data
}

type CommentUserDeletedRequest struct {
	Id int `json:"id"`
}

type CommentDeleteResponse struct {
	Success   bool
	ErrorCode int
}

type CommentListTapeResponse struct {
	Success   bool
	ErrorCode int
	Data      []models.CommentDTO
}
