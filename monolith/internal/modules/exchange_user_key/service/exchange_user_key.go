package service

import (
	"context"
	"fmt"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"

	"github.com/ptflp/gomapper"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/storage"
)

type UserKey struct {
	exchangeUserKeyStorage storage.ExchangeUserKeyer
	logger                 *zap.Logger
	rateLimiter            *service.RateLimiter
}

func NewUserKey(storage storage.ExchangeUserKeyer, components *component.Components) ExchangeUserKeyer {
	return &UserKey{
		exchangeUserKeyStorage: storage,
		logger:                 components.Logger,
	}
}

func (p *UserKey) ExchangeUserKeyAdd(ctx context.Context, in ExchangeUserKeyAddIn) ExchangeOut {
	var err error
	err = p.CheckKeys(ctx, in)
	if err != nil {
		return ExchangeOut{
			ErrorCode: errors.PlatformExchangeServiceAddExchangeUserErr,
		}
	}
	var dto models.ExchangeUserKeyDTO
	dto.SetUserID(in.UserID).
		SetLabel(in.Label).
		SetAPIKey(in.APIKey).
		SetSecretKey(in.SecretKey).
		SetExchangeID(in.ExchangeID)
	err = p.exchangeUserKeyStorage.Create(ctx, dto)
	if err != nil {
		return ExchangeOut{
			ErrorCode: errors.PlatformExchangeServiceAddExchangeUserErr,
		}
	}

	return ExchangeOut{
		Success: true,
	}
}

func (p *UserKey) CheckKeys(ctx context.Context, in ExchangeUserKeyAddIn) error {
	client := service.NewBinanceWithKey(in.UserID, p.rateLimiter, in.APIKey, in.SecretKey)
	account := client.GetAccount(ctx)
	if !account.Success {
		return fmt.Errorf("check account with keys error")
	}
	var spot bool
	for i := range account.DataSpot.Permissions {
		if account.DataSpot.Permissions[i] == service.PermissionSPOT {
			spot = true
		}
	}
	if !spot {
		return fmt.Errorf("check spot permission not passed")
	}

	return nil
}

func (p *UserKey) ExchangeUserKeyDelete(ctx context.Context, exchangeUserKeyID int, userID int) error {
	keys, err := p.exchangeUserKeyStorage.GetList(ctx, utils.Condition{Equal: map[string]interface{}{"id": exchangeUserKeyID, "user_id": userID}})
	if err != nil {
		return err
	}
	if len(keys) < 1 {
		return fmt.Errorf("exchange_user_key delete: permission denied")
	}
	return p.exchangeUserKeyStorage.Delete(ctx, exchangeUserKeyID)
}

func (p *UserKey) ExchangeUserKeyList(ctx context.Context, in ExchangeUserListIn) ExchangeUserListOut {
	dtos, err := p.exchangeUserKeyStorage.GetByUserID(ctx, in.UserID)
	if err != nil {
		return ExchangeUserListOut{
			ErrorCode: errors.PlatformExchangeServiceExchangeUserListErr,
		}
	}

	var exchangeUserList []models.ExchangeUserKey
	err = gomapper.MapStructs(&exchangeUserList, &dtos)
	if err != nil {
		return ExchangeUserListOut{
			ErrorCode: http.StatusInternalServerError,
			Data:      exchangeUserList,
		}
	}

	for i := range exchangeUserList {
		exchangeUserList[i].SecretKey = exchangeUserList[i].SecretKey[:len(exchangeUserList[i].SecretKey)/2]
		exchangeUserList[i].SecretKey = exchangeUserList[i].SecretKey + "*****************"
	}

	return ExchangeUserListOut{
		Data: exchangeUserList,
	}

}

func (p *UserKey) ExchangeUserKeyListByIDs(ctx context.Context, exchangeID, userID int) ([]models.ExchangeUserKeyDTO, error) {
	keysList, err := p.exchangeUserKeyStorage.GetList(ctx, utils.Condition{
		Equal: map[string]interface{}{"exchange_id": exchangeID, "user_id": userID, "deleted_at": nil},
	})
	if err != nil {
		return nil, err
	}

	if len(keysList) < 1 {
		return nil, fmt.Errorf("exchange user key not found")
	}

	return keysList, nil
}

func (p *UserKey) ExchangeUserKeyGetByUserID(ctx context.Context, userID int) ([]models.ExchangeUserKeyDTO, error) {
	return p.exchangeUserKeyStorage.GetByUserID(ctx, userID)
}

func (p *UserKey) ExchangeUserKeyGetByID(ctx context.Context, exchangeUserKeyID int) (models.ExchangeUserKeyDTO, error) {
	return p.exchangeUserKeyStorage.GetByID(ctx, exchangeUserKeyID)
}
