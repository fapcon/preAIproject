package service

import (
	"context"
	"go.uber.org/zap"
	"math"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/posts/storage"
	"time"
)

type PostService struct {
	storage storage.Poster
	logger  *zap.Logger
}

func NewPostService(storage storage.Poster, logger *zap.Logger) *PostService {
	return &PostService{storage: storage, logger: logger}
}

func (p *PostService) Create(ctx context.Context, in PostCreateIn) PostCreateOut {
	var dto models.PostDTO
	dto.SetTittle(in.Title).
		SetShortDescription(in.ShortDescription).
		SetFullDescription(in.FullDescription).
		SetAuthor(in.Author).
		SetCreatedAt(time.Now())
	err := p.storage.Create(ctx, dto)
	if err != nil {
		return PostCreateOut{
			Status:    http.StatusInternalServerError,
			ErrorCode: errors.PostServiceCreatePostErr,
		}
	}
	return PostCreateOut{Status: http.StatusOK, ErrorCode: errors.NoError}
}

func (p *PostService) Delete(ctx context.Context, in PostDeleteIn) PostDeleteOut {
	post, err := p.storage.GetById(ctx, in.Id)
	if err != nil {
		return PostDeleteOut{
			Success:   false,
			ErrorCode: errors.PostServiceGetByIdErr,
		}
	}
	if post.Author != in.Author {
		return PostDeleteOut{
			Success:   false,
			ErrorCode: errors.PostServiceCreatedOtherUser,
		}
	}
	post.SetDeletedAt(time.Now())
	err = p.storage.Update(ctx, post)
	if err != nil {
		return PostDeleteOut{
			Success:   false,
			ErrorCode: errors.PostServiceDeletePostErr,
		}
	}
	return PostDeleteOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}

func (p *PostService) UpdatePost(ctx context.Context, in PostUpdateIn) PostUpdateOut {
	post, err := p.storage.GetById(ctx, in.Id)
	if err != nil {
		return PostUpdateOut{
			Success:   false,
			ErrorCode: errors.PostServiceGetByIdErr,
		}
	}
	if post.Author != in.Author {
		return PostUpdateOut{
			Success:   false,
			ErrorCode: errors.PostServiceCreatedOtherUser,
		}
	}
	post.SetTittle(in.Title).
		SetShortDescription(in.ShortDescription).
		SetFullDescription(in.FullDescription).
		SetUpdatedAt(time.Now())

	err = p.storage.Update(ctx, post)
	if err != nil {
		return PostUpdateOut{
			Success:   false,
			ErrorCode: errors.PostServiceUpdatePostErr,
		}
	}

	return PostUpdateOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}

func (p *PostService) GetById(ctx context.Context, in PostGetByIdIn) PostGetByIdOut {
	post, err := p.storage.GetById(ctx, in.Id)
	if err != nil {
		return PostGetByIdOut{
			Success:   false,
			ErrorCode: errors.PostServiceGetByIdErr,
		}
	}
	if post.Author != in.Author {
		return PostGetByIdOut{
			Success:   false,
			ErrorCode: errors.PostServiceCreatedOtherUser,
		}
	}
	return PostGetByIdOut{
		Success:   true,
		ErrorCode: errors.NoError,
		Body: Data{
			Title:            post.GetTittle(),
			ShortDescription: post.GetShortDescription(),
			FullDescription:  post.GetFullDescription(),
		},
	}
}

func (p *PostService) GetListTape(ctx context.Context, in PostGetTapeIn) PostGetTapeOut {
	out, err := p.storage.GetList(ctx)
	if err != nil {
		return PostGetTapeOut{
			Success:   false,
			ErrorCode: errors.PostServiceGetListErr,
		}
	}
	return paginate(out, in.Limit, in.Offset)
}

func paginate(data []models.PostDTO, limit int, offset int) PostGetTapeOut {
	if len(data) == 0 || offset <= 0 || limit <= 0 {
		return PostGetTapeOut{
			Success:   false,
			ErrorCode: errors.PostServiceLimitOffsetLenErr,
			Body:      []models.PostDTO{},
		}
	}
	totalPages := int(math.Ceil(float64(len(data)) / float64(limit)))
	if offset > totalPages {
		return PostGetTapeOut{
			Success:   false,
			ErrorCode: errors.PostServiceOffsetErr,
			Body:      []models.PostDTO{},
		}
	}
	startIdx := (offset - 1) * limit
	endIdx := int(math.Min(float64(startIdx+limit), float64(len(data))))
	return PostGetTapeOut{
		Success:   true,
		ErrorCode: errors.NoError,
		Body:      data[startIdx:endIdx],
	}
}
