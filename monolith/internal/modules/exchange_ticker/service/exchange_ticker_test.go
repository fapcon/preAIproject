package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	smocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/storage/mocks"
	"testing"
)

func newTestLogger(conf config.AppConf) *zap.Logger {
	ws := bytes.NewBuffer(make([]byte, 0, 1000))
	logger := logs.NewLogger(conf, zapcore.AddSync(ws))
	return logger
}

type fields struct {
	storage *smocks.ExchangeTicker
	logger  *zap.Logger
}

func TestTicker_GetByID(t *testing.T) {
	type args struct {
		ctx      context.Context
		tickerID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ExchangeTicker
		wantErr bool
	}{
		{
			name: "Testing GetByID func, success",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:      context.Background(),
				tickerID: 1,
			},
			want:    models.ExchangeTicker{},
			wantErr: false,
		},
		{
			name: "Testing GetByID func, error",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:      context.Background(),
				tickerID: 1,
			},
			want:    models.ExchangeTicker{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Ticker{
				storage: tt.fields.storage,
				logger:  tt.fields.logger,
			}
			if !tt.wantErr {
				tt.fields.storage.On("GetByID", mock.Anything, mock.Anything).Return(models.ExchangeTicker{}, nil)
			} else {
				tt.fields.storage.On("GetByID", mock.Anything, mock.Anything).Return(models.ExchangeTicker{}, errors.New(""))
			}

			got, err := p.GetByID(tt.args.ctx, tt.args.tickerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicker_GetList(t *testing.T) {
	type args struct {
		ctx       context.Context
		condition utils.Condition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.ExchangeTicker
		wantErr bool
	}{
		{
			name: "Testing GetList func, success",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:       context.Background(),
				condition: utils.Condition{},
			},
			want:    []models.ExchangeTicker{{}, {}, {}},
			wantErr: false,
		},
		{
			name: "Testing GetList func, error",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:       context.Background(),
				condition: utils.Condition{},
			},
			want:    []models.ExchangeTicker{},
			wantErr: true,
		},
		{
			name: "Testing GetList func, empty list",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:       context.Background(),
				condition: utils.Condition{},
			},
			want:    []models.ExchangeTicker{{}, {}, {}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Ticker{
				storage: tt.fields.storage,
				logger:  tt.fields.logger,
			}

			if tt.name == "Testing GetList func, success" {
				tt.fields.storage.On("GetList", mock.Anything, mock.Anything).Return([]models.ExchangeTicker{{}, {}, {}}, nil)
			} else if tt.name == "Testing GetList func, error" {
				tt.fields.storage.On("GetList", mock.Anything, mock.Anything).Return([]models.ExchangeTicker{}, errors.New(""))
			} else {
				tt.fields.storage.On("GetList", mock.Anything, mock.Anything).Return([]models.ExchangeTicker{{}, {}, {}}, errors.New(""))
			}

			got, err := p.GetList(tt.args.ctx, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicker_GetTicker(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GetTickerOut
	}{
		{
			name: "Testing GetTicker func, success",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
			},
			want: GetTickerOut{Data: []models.ExchangeTicker{}, ErrorCode: 0},
		},
		{
			name: "Testing GetTicker func, error",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
			},
			want: GetTickerOut{ErrorCode: 2055},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Ticker{
				storage: tt.fields.storage,
				logger:  tt.fields.logger,
			}

			if tt.name == "Testing GetTicker func, success" {
				tt.fields.storage.On("GetTicker", mock.Anything).Return([]models.ExchangeTicker{}, nil)
			} else {
				tt.fields.storage.On("GetTicker", mock.Anything).Return([]models.ExchangeTicker{}, errors.New(""))
			}
			if got := p.GetTicker(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicker_Save(t *testing.T) {
	type args struct {
		ctx     context.Context
		tickers []models.ExchangeTicker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Testing Save func, success",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "Testing Save func, error",
			fields: fields{
				storage: smocks.NewExchangeTicker(t),
				logger:  newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:     context.Background(),
				tickers: []models.ExchangeTicker{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Ticker{
				storage: tt.fields.storage,
				logger:  tt.fields.logger,
			}

			if tt.wantErr {
				tt.fields.storage.On("Save", mock.Anything, mock.Anything).Return(errors.New(""))
			} else {
				tt.fields.storage.On("Save", mock.Anything, mock.Anything).Return(nil)
			}
			if err := p.Save(tt.args.ctx, tt.args.tickers); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
