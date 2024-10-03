package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ptflp/godecoder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/posts/service"
)

type Poster interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	GetByIdUser(w http.ResponseWriter, r *http.Request)
	GetListTape(w http.ResponseWriter, r *http.Request)
}

type Post struct {
	service service.Poster
	responder.Responder
	godecoder.Decoder
}

func NewPost(service service.Poster, components *component.Components) Poster {
	return &Post{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Создание поста
// @Security ApiKeyAuth
// @Tags posts
// @ID postCreateRequest
// @Accept  json
// @Produce  json
// @Param object body PostUserCreateRequest true "UserSubDelRequest"
// @Router /api/1/posts/create [post]
func (u *Post) CreatePost(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	var req PostUserCreateRequest
	err = u.Decode(r.Body, &req)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	out := u.service.Create(r.Context(), service.PostCreateIn{
		Title:            req.Title,
		ShortDescription: req.ShortDescription,
		FullDescription:  req.FullDescription,
		Author:           claims.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, PostCreateResponse{
			ErrorCode: out.ErrorCode,
			Status:    out.Status,
		})
		return
	}
	u.OutputJSON(w, PostCreateResponse{
		Status: out.Status,
	})
}

// @Summary Удалить пост
// @Security ApiKeyAuth
// @Tags posts
// @ID postDeleteRequest
// @Accept  json
// @Produce  json
// @Param object body PostUserDeletedRequest true "PostUserDeletedRequest"
// @Success 200 {object} PostDeleteResponse
// @Router /api/1/posts/delete [post]
func (u *Post) DeletePost(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	var req PostUserDeletedRequest
	err = u.Decode(r.Body, &req)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	out := u.service.Delete(r.Context(), service.PostDeleteIn{
		Id:     req.Id,
		Author: claims.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, PostDeleteResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, PostDeleteResponse{Success: out.Success})
}

// @Summary Обновить пост
// @Security ApiKeyAuth
// @Tags posts
// @ID postUpdateRequest
// @Accept  json
// @Produce  json
// @Param object body PostUserUpdateRequest true "PostUserUpdateRequest"
// @Success 200 {object} PostUpdateResponse
// @Router /api/1/posts/update [post]
func (u *Post) UpdatePost(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	var req PostUserUpdateRequest
	err = u.Decode(r.Body, &req)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	out := u.service.UpdatePost(r.Context(), service.PostUpdateIn{
		Id:               req.Id,
		Title:            req.Title,
		ShortDescription: req.ShortDescription,
		FullDescription:  req.FullDescription,
		Author:           claims.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, PostUpdateResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, PostUpdateResponse{Success: out.Success})
}

// @Summary Поиск поста по id
// @Security ApiKeyAuth
// @Tags posts
// @ID postGetByIdRequest
// @Accept  json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} PostUpdateResponse
// @Router /api/1/posts/id [get]
func (u *Post) GetByIdUser(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	req, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		u.OutputJSON(w, PostUserGetByIdResponse{
			Success: false,
			Error:   fmt.Sprintf("invalid id: %v", req),
		})
		return
	}

	out := u.service.GetById(r.Context(), service.PostGetByIdIn{
		Id:     req,
		Author: claims.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, PostUserGetByIdResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, PostUserGetByIdResponse{
		Success: out.Success,
		Body: Data{
			Title:            out.Body.Title,
			ShortDescription: out.Body.ShortDescription,
			FullDescription:  out.Body.FullDescription,
		},
	})
}

// @Summary Вывод постов в ленте
// @Security ApiKeyAuth
// @Tags posts
// @ID postGetAllPostRequest
// @Accept  json
// @Produce json
// @Param Limit query integer true "Limit"
// @Param Offset query integer true "Offset"
// @Success 200 {object} PostListTapeResponse
// @Router /api/1/posts/tape [get]
func (u *Post) GetListTape(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		u.OutputJSON(w, PostUserGetByIdResponse{
			Success: false,
			Error:   fmt.Sprintf("invalid limit: %d", limit),
		})
		return
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		u.OutputJSON(w, PostUserGetByIdResponse{
			Success: false,
			Error:   fmt.Sprintf("invalid offset: %d", offset),
		})
		return
	}
	out := u.service.GetListTape(r.Context(), service.PostGetTapeIn{
		Limit:  limit,
		Offset: offset,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, PostListTapeResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, PostListTapeResponse{
		Success:   out.Success,
		ErrorCode: out.ErrorCode,
		Data:      out.Body,
	})
}
