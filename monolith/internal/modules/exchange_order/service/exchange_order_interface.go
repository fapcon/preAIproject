package service

import (
	"context"

	"github.com/shopspring/decimal"
	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeOrderer
type ExchangeOrderer interface {
	CancelOrder(ctx context.Context, in OrderIn) (CancelOrderOut, error)
	WriteOrderLog(ctx context.Context, orderLogDTO models.ExchangeOrderLogDTO)
	WriteOrder(ctx context.Context, in WriteOrderIn) error
	OrderSellLimit(ctx context.Context, in OrderIn, quantity, price decimal.Decimal, unitedOrders int)
	GetOrdersStatistic(ctx context.Context, in GetBotRelationIn) StatisticOut
	AddOrdersStatistic(ctx context.Context, orders *[]models.ExchangeOrder) StatisticOut
	GetBotOrders(ctx context.Context, in GetBotRelationIn) GetOrdersOut
	GetUserOrders(ctx context.Context, in GetUserRelationIn) GetOrdersOut
	ExchangeOrderList(ctx context.Context, in GetBotRelationIn) GetOrdersOut
	GetAllOrdersStatistic(ctx context.Context, in GetUserRelationIn) StatisticOut
	GetOrdersCondition(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrder, error)

	GetOrderList(ctx context.Context, condition utils.Condition) ([]models.ExchangeOrderDTO, error)
	UpdateOrder(ctx context.Context, dto models.ExchangeOrderDTO) error

	PutOrder(ctx context.Context, in OrderIn) PutOrderOut
}
