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
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/comments/service"
)

type Comenter interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
	UpdateComment(w http.ResponseWriter, r *http.Request)
	GetCommentByID(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
	GetCommentsList(w http.ResponseWriter, r *http.Request)
}

type CommentController struct {
	service service.CommentService
	responder.Responder
	godecoder.Decoder
}

func NewCommentController(service service.CommentService, components *component.Components) Comenter {
	return &CommentController{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Создание комментария
// @Security ApiKeyAuth
// @Tags comments
// @ID createComment
// @Accept  json
// @Produce  json
// @Param object body CommentUserCreateRequest true "CommentUserCreateRequest"
// @Success 200 {object} CommentCreateResponse
// @Router /api/1/comments/create [post]
func (u *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	var req CommentUserCreateRequest
	err = u.Decode(r.Body, &req)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	out := u.service.CreateComment(r.Context(), service.CommentCreateIn{
		Comment:  req.Comment,
		Source:   req.Source,
		SourceID: req.SourceID,
		AuthorID: claims.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, CommentCreateResponse{
			ErrorCode: out.ErrorCode,
			Status:    out.Status,
		})
		return
	}
	u.OutputJSON(w, CommentCreateResponse{
		Status: out.Status,
	})
}

// @Summary Измениение комментария
// @Security ApiKeyAuth
// @Tags comments
// @ID updateComment
// @Accept  json
// @Produce  json
// @Param object body CommentUserUpdateRequest true "CommentUserUpdateRequest"
// @Success 200 {object} CommentUpdateResponse
// @Router /api/1/comments/update [post]
func (u *CommentController) UpdateComment(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	var req CommentUserUpdateRequest
	err = u.Decode(r.Body, &req)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	out := u.service.UpdateComment(r.Context(), service.CommentUpdateIn{
		Id:       req.Id,
		AuthorID: claims.ID,
		Comment:  req.Comment,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, CommentUpdateResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, CommentUpdateResponse{Success: out.Success})
}

// @Summary Поиск комментприй по ID
// @Security ApiKeyAuth
// @Tags comments
// @ID commentGet
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {object} CommentUserGetByIdResponse
// @Router /api/1/comments/id [get]
func (u *CommentController) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	req, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		u.OutputJSON(w, CommentGetByIdResponse{
			Success: false,
			Error:   fmt.Sprintf("invalid id: %d", req),
		})
		return
	}

	out := u.service.GetCommentByID(r.Context(), service.CommentGetByIdIn{
		Id:       req,
		AuthorID: claims.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, CommentUserGetByIdResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, CommentUserGetByIdResponse{
		Success: out.Success,
		Body: Data{
			Comment: out.Body.Comment,
			Source:  out.Body.Source,
		},
	})
}

// @Summary Удаление комментария
// @Security ApiKeyAuth
// @Tags comments
// @ID вeleteComment
// @Accept  json
// @Produce  json
// @Param object body CommentUserDeletedRequest true "CommentUserDeletedRequest"
// @Success 200 {object} CommentDeleteResponse
// @Router /api/1/comments/delete [post]
func (u *CommentController) DeleteComment(w http.ResponseWriter, r *http.Request) {
	claims, err := handlers.ExtractUser(r)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	var req CommentUserDeletedRequest
	err = u.Decode(r.Body, &req)
	if err != nil {
		u.ErrorBadRequest(w, err)
		return
	}
	out := u.service.DeleteComment(r.Context(), service.CommentDeleteIn{
		Id:       req.Id,
		AuthorID: claims.ID,
	})
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, CommentDeleteResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, CommentDeleteResponse{Success: out.Success})
}

// @Summary Получение списка комментариев
// @Security ApiKeyAuth
// @Tags comments
// @ID сommentsList
// @Accept  json
// @Produce  json
// @Success 200 {object} CommentListTapeResponse
// @Router /api/1/comments/list [get]
func (u *CommentController) GetCommentsList(w http.ResponseWriter, r *http.Request) {
	out := u.service.GetCommentList(r.Context())
	if out.ErrorCode != errors.NoError {
		u.OutputJSON(w, CommentListTapeResponse{
			Success:   out.Success,
			ErrorCode: out.ErrorCode,
		})
		return
	}
	u.OutputJSON(w, CommentListTapeResponse{
		Success:   out.Success,
		ErrorCode: out.ErrorCode,
		Data:      out.Body,
	})
}
