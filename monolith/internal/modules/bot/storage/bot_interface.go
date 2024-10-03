package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type Boter interface {
	Create(ctx context.Context, strategy models.BotDTO) (int, error)
	Update(ctx context.Context, strategy models.BotDTO) error
	UpdateByUUID(ctx context.Context, bot models.BotDTO) error
	GetByID(ctx context.Context, strategyID int) (models.BotDTO, error)
	GetByUUID(ctx context.Context, uuid string) (models.BotDTO, error)
	GetDraft(ctx context.Context, userID int) (models.BotDTO, error)
	GetList(ctx context.Context, condition utils.Condition) ([]models.BotDTO, error)
	GetByIDs(ctx context.Context, ids []int) ([]models.BotDTO, error)
	GetByFilter(ctx context.Context) ([]models.BotDTO, error)
}
