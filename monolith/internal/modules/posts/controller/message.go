package controller

import "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

type PostUserCreateRequest struct {
	Title            string `json:"title" `
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
}

type PostCreateResponse struct {
	Status    int
	ErrorCode int
}

type PostUserDeletedRequest struct {
	Id int `json:"id"`
}

type PostDeleteResponse struct {
	Success   bool
	ErrorCode int
}

type PostUserUpdateRequest struct {
	Id               int    `json:"id"`
	Title            string `json:"title" `
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
}

type PostUpdateResponse struct {
	Success   bool
	ErrorCode int
}

type PostUserGetByIdRequest struct {
	Id int `json:"id"`
}

type PostUserGetByIdResponse struct {
	Success   bool
	ErrorCode int
	Error     string
	Body      Data
}

type Data struct {
	Title            string `json:"title" `
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
}

type PostListTapeResponse struct {
	Success   bool
	ErrorCode int
	Data      []models.PostDTO
}
