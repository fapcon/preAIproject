package service

import (
	"context"
	"net/http"

	"github.com/ptflp/gomapper"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/storage"
)

type ExchangeList struct {
	exchangeListStorage storage.ExchangeLister
	logger              *zap.Logger
}

func NewExchangeList(storage storage.ExchangeLister, components *component.Components) ExchangeLister {
	return &ExchangeList{
		exchangeListStorage: storage,
		logger:              components.Logger,
	}
}

func (p *ExchangeList) ExchangeListList(ctx context.Context) ExchangeListOut {
	dto, err := p.exchangeListStorage.GetList(ctx)
	if err != nil {
		return ExchangeListOut{
			ErrorCode: errors.PlatformExchangeServiceGetExchangeListErr,
		}
	}

	var exchangeList []models.ExchangeList
	err = gomapper.MapStructs(&exchangeList, &dto)
	if err != nil {
		return ExchangeListOut{
			ErrorCode: http.StatusInternalServerError,
			Data:      exchangeList,
		}
	}

	return ExchangeListOut{
		Data: exchangeList,
	}
}

func (p *ExchangeList) ExchangeListDelete(ctx context.Context, exchangeListID int) error {
	return p.exchangeListStorage.Delete(ctx, exchangeListID)
}

func (p *ExchangeList) ExchangeListAdd(ctx context.Context, in ExchangeAddIn) ExchangeOut {
	err := p.exchangeListStorage.Create(ctx, models.ExchangeListDTO{
		Name:        in.Name,
		Description: in.Description,
	})

	if err != nil {
		return ExchangeOut{
			ErrorCode: errors.PlatformExchangeServiceGetTickerErr,
		}
	}

	return ExchangeOut{
		Success: true,
	}
}
