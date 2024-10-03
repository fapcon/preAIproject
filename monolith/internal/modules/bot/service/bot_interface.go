package service

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Boter
type Boter interface {
	Create(ctx context.Context, in BotCreateIn) BotOut
	Delete(ctx context.Context, in BotDeleteIn) BOut
	Update(ctx context.Context, in BotUpdateIn) BOut
	Get(ctx context.Context, in BotGetIn) BotOut
	Toggle(ctx context.Context, in BotToggleIn) BOut
	Subscribe(ctx context.Context, in BotSubscribeIn) BOut
	Unsubscribe(ctx context.Context, in BotSubscribeIn) BOut
	List(ctx context.Context, in BotListIn) BotListOut
	WebhookSignal(ctx context.Context, in WebhookSignalIn) WebhookSignalOut
}

type BotCreateIn struct {
	UserID int `json:"user_id"`
}

type BotSubscribeIn struct {
	UserIDs      []int  `json:"user_ids"`
	StrategyUUID string `json:"strategy_uuid"`
}

type BOut struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

type BotListIn struct {
	UserID int
}

type BotListOut struct {
	Data      []models.Bot `json:"data"`
	Success   bool         `json:"success"`
	ErrorCode int          `json:"error_code"`
}

type BotOut struct {
	ErrorCode int               `json:"error_code"`
	Bot       models.Bot        `json:"bot"`
	Hooks     map[string]string `json:"hooks"`
}

type Hooks struct {
	BuyMarket string `json:"buy_market"`
}

type BotDeleteIn struct {
	UUID   string `json:"uuid"`
	UserID int
}

type BotGetIn struct {
	UUID string `json:"uuid"`
	ID   int    `json:"id"`
}

type BotUpdateIn struct {
	Bot    models.Bot `json:"bot"`
	Fields []int      `json:"fields"`
	Pairs  []int
}

type BotPairAdd struct {
	BotID int
	Pairs []int
}

type BotToggleIn struct {
	UUID   string
	UserID int
	Active bool
}

type WebhookSignalIn struct {
	BotUUID string
	PairID  int `json:"pair_id"`
}

type WebhookSignalOut struct {
	Hook      string
	Signals   []string
	ErrorCode int
}
