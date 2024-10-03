package service

import (
	"context"
)

type Statisticer interface {
	UpdateStatistic(ctx context.Context, in StatisticUpdateIn) StatisticUpdateOut
	GetBotStatistic(ctx context.Context, in StatisticIn) StatisticOut
	GetUserBotStatistic(ctx context.Context, in UserStatisticIn) UserStatisticOut
	DeleteBotStatistic(ctx context.Context, in StatisticIn) StatisticDeleteOut
}
