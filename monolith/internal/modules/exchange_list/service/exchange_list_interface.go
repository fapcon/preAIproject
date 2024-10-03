package service

import (
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeLister

type ExchangeLister interface {
	ExchangeListList(ctx context.Context) ExchangeListOut
	ExchangeListDelete(ctx context.Context, exchangeListID int) error
	ExchangeListAdd(ctx context.Context, in ExchangeAddIn) ExchangeOut
}
