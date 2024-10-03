package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/storage/mocks"
	"testing"
)

func newTestLogger(conf config.AppConf) *zap.Logger {
	ws := bytes.NewBuffer(make([]byte, 0, 1000))
	logger := logs.NewLogger(conf, zapcore.AddSync(ws))
	return logger
}

type fields struct {
	exchangeListStorage *mocks.ExchangeLister
	logger              *zap.Logger
}

func TestExchangeList_ExchangeListAdd(t *testing.T) {
	type args struct {
		ctx context.Context
		in  ExchangeAddIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ExchangeOut
	}{
		{
			name: "testing ExchangeListAdd, success",
			fields: fields{
				exchangeListStorage: mocks.NewExchangeLister(t),
				logger:              newTestLogger(config.AppConf{})},
			args: args{
				ctx: context.Background(),
				in:  ExchangeAddIn{UserID: 1, Name: "Binance", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
			},
			want: ExchangeOut{
				Success:   true,
				ErrorCode: 0,
			},
		},
		{
			name: "testing ExchangeListAdd, error",
			fields: fields{
				exchangeListStorage: mocks.NewExchangeLister(t),
				logger:              newTestLogger(config.AppConf{})},
			args: args{
				ctx: context.Background(),
				in:  ExchangeAddIn{UserID: 1, Name: "Binance", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
			},
			want: ExchangeOut{
				Success:   false,
				ErrorCode: 2055,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ExchangeList{
				exchangeListStorage: tt.fields.exchangeListStorage,
				logger:              tt.fields.logger,
			}

			if tt.name != "testing ExchangeListAdd, error" {
				// Настройка ожидание вызова метода
				tt.fields.exchangeListStorage.On("Create", mock.Anything, mock.Anything).
					Return(nil)
				// возращаем ожидаемое значение
			} else {
				tt.fields.exchangeListStorage.On("Create", mock.Anything, mock.Anything).
					Return(errors.New(""))
			}

			if got := p.ExchangeListAdd(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExchangeListAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeList_ExchangeListDelete(t *testing.T) {
	type args struct {
		ctx            context.Context
		exchangeListID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Testing ExchangeListDelete, success",
			fields: fields{
				exchangeListStorage: mocks.NewExchangeLister(t),
				logger:              newTestLogger(config.AppConf{})},
			args:    args{ctx: context.Background(), exchangeListID: 1},
			wantErr: false,
		},
		{
			name: "Testing ExchangeListDelete, error",
			fields: fields{
				exchangeListStorage: mocks.NewExchangeLister(t),
				logger:              newTestLogger(config.AppConf{})},
			args:    args{ctx: context.Background(), exchangeListID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ExchangeList{
				exchangeListStorage: tt.fields.exchangeListStorage,
				logger:              tt.fields.logger,
			}

			if !tt.wantErr {
				tt.fields.exchangeListStorage.On("Delete", mock.Anything, mock.Anything).Return(nil)
			} else {
				tt.fields.exchangeListStorage.On("Delete", mock.Anything, mock.Anything).Return(errors.New(""))
			}
			if err := p.ExchangeListDelete(tt.args.ctx, tt.args.exchangeListID); (err != nil) != tt.wantErr {
				t.Errorf("ExchangeListDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExchangeList_ExchangeListList(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ExchangeListOut
	}{
		{
			name: "Testing ExchangeListList, success",
			fields: fields{
				exchangeListStorage: mocks.NewExchangeLister(t),
				logger:              newTestLogger(config.AppConf{})},
			args: args{ctx: context.Background()},
			want: ExchangeListOut{
				ErrorCode: 0,
				Data: []models.ExchangeList{
					{ID: 1, Name: "Binance", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
					{ID: 2, Name: "Haobi", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
					{ID: 3, Name: "Bybit", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
				},
			},
		},
		{
			name: "Testing ExchangeListList, error",
			fields: fields{
				exchangeListStorage: mocks.NewExchangeLister(t),
				logger:              newTestLogger(config.AppConf{})},
			args: args{ctx: context.Background()},
			want: ExchangeListOut{
				ErrorCode: 2056,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ExchangeList{
				exchangeListStorage: tt.fields.exchangeListStorage,
				logger:              tt.fields.logger,
			}

			if tt.name != "Testing ExchangeListList, error" {
				tt.fields.exchangeListStorage.On("GetList", mock.Anything).Return([]models.ExchangeListDTO{
					{ID: 1, Name: "Binance", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
					{ID: 2, Name: "Haobi", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
					{ID: 3, Name: "Bybit", Description: "Best Crypto Exchange", Slug: "best-crypto-exchange"},
				}, nil)
			} else {
				tt.fields.exchangeListStorage.On("GetList", mock.Anything).Return(nil, errors.New(""))
			}

			if got := p.ExchangeListList(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExchangeListList() = %v, want %v", got, tt.want)
			}
		})
	}
}
