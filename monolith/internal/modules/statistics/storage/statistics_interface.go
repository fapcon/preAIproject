package storage

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

type Statisticer interface {
	Create(ctx context.Context, botStatistics models.BotStatisticsDTO) error
	Update(ctx context.Context, botStatistics models.BotStatisticsDTO) error
	GetByID(ctx context.Context, id int) (models.BotStatisticsDTO, error)
	GetByUUID(ctx context.Context, botUUID string) (models.BotStatisticsDTO, error)
	GetByUserID(ctx context.Context, userID int) ([]models.BotStatisticsDTO, error)
	Delete(ctx context.Context, uuid string) error
}
