package service

import (
	"context"
	"net/http"
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/comments/storage"
)

type commentService struct {
	storage storage.Commenter
}

func NewCommentService(storage storage.Commenter) CommentService {
	return &commentService{
		storage: storage,
	}
}

func (p *commentService) CreateComment(ctx context.Context, in CommentCreateIn) CommentCreateOut {
	var dto models.CommentDTO
	dto.SetComment(in.Comment).
		SetSource(in.Source).
		SetAuthorID(in.AuthorID).
		SetSourceID(in.SourceID).
		SetCreatedAt(time.Now())
	err := p.storage.Create(ctx, dto)
	if err != nil {
		return CommentCreateOut{
			Status:    http.StatusInternalServerError,
			ErrorCode: errors.CommentServiceCreateCommentErr,
		}
	}
	return CommentCreateOut{Status: http.StatusOK, ErrorCode: errors.NoError}
}

func (p *commentService) UpdateComment(ctx context.Context, in CommentUpdateIn) CommentUpdateOut {
	comment, err := p.storage.GetByID(ctx, in.Id)
	if err != nil {
		return CommentUpdateOut{
			Success:   false,
			ErrorCode: errors.CommentServiceGetByIdErr,
		}
	}
	if comment.AuthorID != in.AuthorID {
		return CommentUpdateOut{
			Success:   false,
			ErrorCode: errors.CommentServiceCreatedOtherUser,
		}
	}
	comment.SetComment(in.Comment).
		SetUpdatedAt(time.Now())

	err = p.storage.Update(ctx, comment)
	if err != nil {
		return CommentUpdateOut{
			Success:   false,
			ErrorCode: errors.CommentServiceUpdateCommentErr,
		}
	}

	return CommentUpdateOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}

func (p *commentService) GetCommentByID(ctx context.Context, in CommentGetByIdIn) CommentGetByIdOut {
	comment, err := p.storage.GetByID(ctx, in.Id)
	if err != nil {
		return CommentGetByIdOut{
			Success:   false,
			ErrorCode: errors.CommentServiceGetByIdErr,
		}
	}
	if comment.AuthorID != in.AuthorID {
		return CommentGetByIdOut{
			Success:   false,
			ErrorCode: errors.CommentServiceCreatedOtherUser,
		}
	}
	return CommentGetByIdOut{
		Success:   true,
		ErrorCode: errors.NoError,
		Body: Data{
			Comment: comment.GetComment(),
			Source:  comment.GetSource(),
		},
	}
}

func (p *commentService) DeleteComment(ctx context.Context, in CommentDeleteIn) CommentDeleteOut {
	comment, err := p.storage.GetByID(ctx, in.Id)
	if err != nil {
		return CommentDeleteOut{
			Success:   false,
			ErrorCode: errors.CommentServiceGetByIdErr,
		}
	}
	if comment.AuthorID != in.AuthorID {
		return CommentDeleteOut{
			Success:   false,
			ErrorCode: errors.CommentServiceCreatedOtherUser,
		}
	}
	comment.SetDeletedAt(time.Now())
	err = p.storage.Update(ctx, comment)
	if err != nil {
		return CommentDeleteOut{
			Success:   false,
			ErrorCode: errors.CommentServiceDeleteCommentErr,
		}
	}
	return CommentDeleteOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}

func (p *commentService) GetCommentList(ctx context.Context) CommentGetTapeOut {
	out, err := p.storage.GetList(ctx)
	if err != nil {
		return CommentGetTapeOut{
			Success:   false,
			ErrorCode: errors.CommentServiceGetListErr,
		}
	}
	return CommentGetTapeOut{
		Success:   true,
		ErrorCode: errors.NoError,
		Body:      out,
	}
}
