package service

import (
	"context"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name ExchangeUserKeyer
type ExchangeUserKeyer interface {
	ExchangeUserKeyAdd(ctx context.Context, in ExchangeUserKeyAddIn) ExchangeOut
	ExchangeUserKeyDelete(ctx context.Context, exchangeUserKeyID int, userID int) error

	ExchangeUserKeyList(ctx context.Context, in ExchangeUserListIn) ExchangeUserListOut
	ExchangeUserKeyListByIDs(ctx context.Context, exchangeID, userID int) ([]models.ExchangeUserKeyDTO, error)

	ExchangeUserKeyGetByID(ctx context.Context, exchangeUserKeyID int) (models.ExchangeUserKeyDTO, error)
	ExchangeUserKeyGetByUserID(ctx context.Context, userID int) ([]models.ExchangeUserKeyDTO, error)

	CheckKeys(ctx context.Context, in ExchangeUserKeyAddIn) error
}
