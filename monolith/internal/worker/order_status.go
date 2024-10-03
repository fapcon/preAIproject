package worker

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	oservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
	tservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/service"
	kservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/storages"
)

type OrderStatus struct {
	exchangeOrderService oservice.ExchangeOrderer
	userKeyService       kservice.ExchangeUserKeyer
	ticker               tservice.ExchangeTicker
	logger               *zap.Logger
	rateLimiter          *service.RateLimiter
}

var userApiKeys = make(map[int]models.ExchangeUserKeyDTO)

func NewOrderStatus(storage *storages.Storages, services modules.Services, components *component.Components) *OrderStatus {
	return &OrderStatus{exchangeOrderService: services.Order, ticker: services.Ticker, userKeyService: services.UserKey, logger: components.Logger, rateLimiter: components.RateLimiter}
}

func (o *OrderStatus) Run() error {
	var errGroup errgroup.Group
	errGroup.Go(o.work)

	return errGroup.Wait()
}

func (o *OrderStatus) work() error {
	timeTicker := time.NewTicker(21 * time.Second)
	timeTicker3min := time.NewTicker(3 * time.Minute)
	ctx := context.Background()
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			dtos, err := o.exchangeOrderService.GetOrderList(ctx, utils.Condition{
				NotEqual: map[string]interface{}{"status": []interface{}{service.OrderStatusFilled, service.OrderStatusCanceled, service.OrderStatusFailed, service.OrderStatusRejected, service.OrderStatusExpired}},
			})
			if err != nil {
				continue
			}
			if len(dtos) < 1 {
				continue
			}
			ordersByID := make(map[int][]models.ExchangeOrderDTO, len(dtos))
			var pairIDs []interface{}
			for i := range dtos {
				ordersByID[dtos[i].PairID] = append(ordersByID[dtos[i].PairID], dtos[i])
				pairIDs = append(pairIDs, dtos[i].PairID)
			}
			pairs, err := o.ticker.GetList(ctx, utils.Condition{Equal: map[string]interface{}{"id": pairIDs}})
			if err != nil {
				o.logger.Error("order status worker retrieving ticker list", zap.Error(err))
			}
			for _, pair := range pairs {
				for _, order := range ordersByID[pair.ID] {
					if order.Side == service.SideSell {
						if order.Price.GreaterThan(order.Price) {
							o.UpdateStatus(pair, order)
						}
					}
					if order.Side == service.SideBuy {
						if order.Price.LessThan(order.Price) {
							o.UpdateStatus(pair, order)
						}
					}
				}
			}
		case <-timeTicker3min.C:
			dtos, err := o.exchangeOrderService.GetOrderList(ctx, utils.Condition{
				NotEqual: map[string]interface{}{"status": []interface{}{service.OrderStatusFilled, service.OrderStatusCanceled, service.OrderStatusFailed, service.OrderStatusRejected, service.OrderStatusExpired}},
			})
			if err != nil {
				continue
			}
			if len(dtos) < 1 {
				continue
			}
			ordersByID := make(map[int][]models.ExchangeOrderDTO, len(dtos))
			var pairIDs []interface{}
			for i := range dtos {
				ordersByID[dtos[i].PairID] = append(ordersByID[dtos[i].PairID], dtos[i])
				pairIDs = append(pairIDs, dtos[i].PairID)
			}
			pairs, err := o.ticker.GetList(ctx, utils.Condition{Equal: map[string]interface{}{"id": pairIDs}})
			if err != nil {
				o.logger.Error("order status worker retrieving ticker list", zap.Error(err))
			}
			for _, pair := range pairs {
				for _, order := range ordersByID[pair.ID] {
					o.UpdateStatus(pair, order)
				}
			}
		}
	}
}

func (o *OrderStatus) UpdateStatus(pair models.ExchangeTicker, order models.ExchangeOrderDTO) {
	ctx := context.Background()
	var err error
	var keys []models.ExchangeUserKeyDTO
	var apiKey, secretKey string

	if v, ok := userApiKeys[order.ApiKeyID]; ok {
		apiKey = v.APIKey
		secretKey = v.SecretKey
	} else {
		keys, err = o.userKeyService.ExchangeUserKeyGetByUserID(ctx, order.UserID)
		if err != nil {
			o.logger.Error(fmt.Sprintf("order status worker retrieve user key: %d", order.UserID), zap.Error(err))
			return
		}

		for i := range keys {
			userApiKeys[keys[i].ID] = keys[i]
		}
	}
	if v, ok := userApiKeys[order.ApiKeyID]; ok {
		apiKey = v.APIKey
		secretKey = v.SecretKey
	}

	exchangeClient := service.NewBinanceWithKey(order.UserID, o.rateLimiter, apiKey, secretKey)
	exchangeOrder := exchangeClient.GetOrder(ctx, service.GetOrderIn{
		Pair:    pair.Pair,
		OrderID: order.ExchangeOrderID,
	})
	if exchangeOrder.ErrorCode != errors.NoError {
		return
	}

	if order.Status == exchangeOrder.Status {
		return
	}
	order.Status = exchangeOrder.Status
	err = o.exchangeOrderService.UpdateOrder(ctx, order)
	if err != nil {
		o.logger.Error("order status worker update order status", zap.Error(err))
	}
}
