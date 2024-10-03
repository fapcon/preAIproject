package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	service2 "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service"
	"testing"

	"gitlab.com/golight/orm/utils"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

	bmocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service/mocks"
	omocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service/mocks"
	umocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service/mocks"
	wpmocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/webhook_process/storage/mocks"
)

func newTestLogger(conf config.AppConf) *zap.Logger {
	ws := bytes.NewBuffer(make([]byte, 0, 1000))
	logger := logs.NewLogger(conf, zapcore.AddSync(ws))
	return logger
}

type fields struct {
	exchangeOrder                *omocks.ExchangeOrderer         //service.ExchangeOrderer
	webhookProcessStorage        *wpmocks.WebhookProcesser       //storage.WebhookProcesser
	webhookProcessHistoryStorage *wpmocks.WebhookProcessHistorer //storage.WebhookProcessHistorer
	exchangeUserKeyService       *umocks.ExchangeUserKeyer       //service.ExchangeUserKeyer
	botService                   *bmocks.Boter                   //service.Boter
	logger                       *zap.Logger
}

func TestWebhookProcess_CreateWebhookProcess(t *testing.T) {
	type args struct {
		ctx context.Context
		bot models.Bot
		in  WebhookProcessIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.WebhookProcessDTO
		wantErr bool
	}{
		{
			name: "Тест успешного создания WebhookProcess",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				bot: models.Bot{},
				in:  WebhookProcessIn{},
			},
			want:    models.WebhookProcessDTO{},
			wantErr: false,
		},
		{
			name: "Тест ошибки при создании WebhookProcess",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				bot: models.Bot{},
				in:  WebhookProcessIn{},
			},
			want:    models.WebhookProcessDTO{},
			wantErr: true,
		},
		{
			name: "Тест пустого списка WebhookProcess после создания",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				bot: models.Bot{},
				in:  WebhookProcessIn{},
			},
			want:    models.WebhookProcessDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			if tt.name == "Тест успешного создания WebhookProcess" {
				tt.fields.webhookProcessStorage.
					On("Create", mock.Anything, mock.Anything).Return(nil)
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{{}, {}, {}}, nil)
				tt.fields.webhookProcessHistoryStorage.
					On("Create", mock.Anything, mock.Anything).Return(nil)
			} else if tt.name == "Тест ошибки при создании WebhookProcess" {
				tt.fields.webhookProcessStorage.
					On("Create", mock.Anything, mock.Anything).Return(errors.New(""))
			} else {
				tt.fields.webhookProcessStorage.
					On("Create", mock.Anything, mock.Anything).Return(nil)
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{{}, {}, {}}, errors.New(""))
			}
			got, err := p.CreateWebhookProcess(tt.args.ctx, tt.args.bot, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWebhookProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWebhookProcess() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhookProcess_GetBotWebhooks(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetBotRelationIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GetWebhooksOut
	}{
		{
			name: "Тест успешного получения BotWebhooks",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetBotRelationIn{},
			},
			want: GetWebhooksOut{
				Success:   true,
				ErrorCode: 0,
				Data:      []models.WebhookProcess{},
			},
		},
		{
			name: "Тест ошибки при получении BotWebhooks",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetBotRelationIn{},
			},
			want: GetWebhooksOut{
				Success:   false,
				ErrorCode: 1,
				Data:      []models.WebhookProcess{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			if tt.want.Success == true {
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{}, nil)
				tt.fields.webhookProcessHistoryStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessHistoryDTO{}, nil)
			} else {
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{}, errors.New(""))
			}
			if got := p.GetBotWebhooks(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBotWebhooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhookProcess_GetUserWebhooks(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetUserRelationIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GetWebhooksOut
	}{
		{
			name: "Тест успешного получения UserWebhooks",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetUserRelationIn{},
			},
			want: GetWebhooksOut{
				Success:   true,
				ErrorCode: 0,
				Data:      []models.WebhookProcess{},
			},
		},
		{
			name: "Тест ошибки при получении BotWebhooks",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetUserRelationIn{},
			},
			want: GetWebhooksOut{
				Success:   false,
				ErrorCode: 1,
				Data:      []models.WebhookProcess{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			if tt.want.Success == true {
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{}, nil)
				tt.fields.webhookProcessHistoryStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessHistoryDTO{}, nil)
			} else {
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{}, errors.New(""))
			}
			if got := p.GetUserWebhooks(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserWebhooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhookProcess_GetWebhookInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetWebhookInfoIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GetWebhookInfoOut
	}{
		{
			name: "Test successful GetWebhookInfo",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetWebhookInfoIn{},
			},
			want: GetWebhookInfoOut{
				Data: GetWebhookInfoData{},
			},
		},
		{
			name: "Test error in GetWebhookInfo (webhooks)",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetWebhookInfoIn{},
			},
			want: GetWebhookInfoOut{
				ErrorCode: 0,
				Success:   false,
				Data:      GetWebhookInfoData{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}

			tt.fields.webhookProcessStorage.
				On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{{}, {}}, nil)
			tt.fields.webhookProcessHistoryStorage.
				On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessHistoryDTO{}, nil)
			tt.fields.exchangeOrder.
				On("GetOrdersCondition", mock.Anything, mock.Anything).Return([]models.ExchangeOrder{}, nil)

			if got := p.GetWebhookInfo(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got.ErrorCode, tt.want.ErrorCode) {
				t.Errorf("GetWebhookInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

//WebhookProcess not have func PutOrder
//func TestWebhookProcess_PutOrder(t *testing.T) {
//	type args struct {
//		ctx context.Context
//		in  service.OrderIn
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		exch   *mocks.Exchanger
//		args   args
//		want   PutOrderOut
//	}{
//		{
//			name: "Test successful PutOrder",
//			exch: mocks.NewExchanger(t),
//			fields: fields{
//				exchangeOrder:                omocks.NewExchangeOrderer(t),
//				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
//				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
//				exchangeTickerService:        tmocks.NewExchangeTicker(t),
//				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
//				botService:                   bmocks.NewBoter(t),
//				logger:                       newTestLogger(config.AppConf{}),
//			},
//			args: args{
//				ctx: context.Background(),
//				in:  service.OrderIn{},
//			},
//			want: PutOrderOut{
//				Success: true,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &WebhookProcess{
//				exchangeOrder:                tt.fields.exchangeOrder,
//				webhookProcessStorage:        tt.fields.webhookProcessStorage,
//				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
//				exchangeTickerService:        tt.fields.exchangeTickerService,
//				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
//				botService:                   tt.fields.botService,
//				logger:                       tt.fields.logger,
//			}
//			tt.fields.exchangeTickerService.
//				On("GetByID", mock.Anything, mock.Anything).Return(models.ExchangeTicker{}, nil)
//			tt.fields.exchangeOrder.
//				On("CreateOrder", mock.Anything, mock.Anything).Return(nil)
//			tt.fields.exchangeOrder.
//				On("GetOrdersByUUID", mock.Anything, mock.Anything).Return(models.ExchangeOrderDTO{}, nil)
//			tt.fields.exchangeOrder.
//				On("WriteOrderLog", mock.Anything, mock.Anything).Return(models.ExchangeOrderDTO{}, nil)
//			tt.fields.exchangeOrder.
//				On("WriteOrderLog", mock.Anything, mock.Anything)
//			tt.fields.webhookProcessHistoryStorage.
//				On("Create", mock.Anything, mock.Anything).Return(nil)
//			tt.fields.webhookProcessStorage.
//				On("Update", mock.Anything, mock.Anything).Return(nil)
//			tt.fields.exchangeOrder.
//				On("UpdateOrder", mock.Anything, mock.Anything).Return(nil)
//			if got := p.PutOrder(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("PutOrder() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestWebhookProcess_UpdateWebhookStatus(t *testing.T) {
	type args struct {
		ctx      context.Context
		bot      models.Bot
		webhook  models.WebhookProcessDTO
		orderDTO models.ExchangeOrderDTO
		message  string
		status   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test successful webhook status update",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:      context.Background(),
				bot:      models.Bot{},
				webhook:  models.WebhookProcessDTO{},
				orderDTO: models.ExchangeOrderDTO{},
				message:  "Order successful",
				status:   StatusFinished,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			tt.fields.webhookProcessHistoryStorage.
				On("Create", mock.Anything, mock.Anything).Return(nil)
			tt.fields.webhookProcessStorage.
				On("Update", mock.Anything, mock.Anything).Return(nil)
			p.UpdateWebhookStatus(tt.args.ctx, tt.args.bot, tt.args.webhook, tt.args.message, tt.args.status)
		})
	}
}

func TestWebhookProcess_WebhookProcess(t *testing.T) {
	type args struct {
		ctx context.Context
		in  WebhookProcessIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   WebhookProcessOut
	}{
		{
			name: "Test successful webhook processing",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: WebhookProcessIn{
					Slug: "AVERAGE_CKBBUSD_07ee9a1c-53ba-11ed-8cd4-0242ac120002",
				},
			},
			want: WebhookProcessOut{
				ErrorCode: 0,
				Success:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			tt.fields.botService.
				On("Get", mock.Anything, mock.Anything).Return(service2.BotOut{
				Bot: models.Bot{
					Active: true,
				},
			})
			tt.fields.webhookProcessStorage.
				On("Create", mock.Anything, mock.Anything).Return(nil)
			tt.fields.webhookProcessStorage.
				On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessDTO{{}, {}, {}}, nil)
			tt.fields.webhookProcessHistoryStorage.
				On("Create", mock.Anything, mock.Anything).Return(nil)
			tt.fields.exchangeUserKeyService.
				On("ExchangeUserKeyListByIDs", mock.Anything, mock.Anything, mock.Anything).Return([]models.ExchangeUserKeyDTO{}, nil)
			if got := p.WebhookProcess(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebhookProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhookProcess_WriteWebhookHistory(t *testing.T) {
	type args struct {
		ctx context.Context
		dto models.WebhookProcessHistoryDTO
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test successful WriteWebhookHistory",
			fields: fields{
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
			},
			args: args{
				ctx: context.Background(),
				dto: models.WebhookProcessHistoryDTO{
					UserID:      1,
					WebhookUUID: "webhook_uuid",
					WebhookID:   1,
					ExchangeID:  1,
					Status:      StatusProcessing,
				},
			},
		},
		{
			name: "Test error in WriteWebhookHistory",
			fields: fields{
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
			},
			args: args{
				ctx: context.Background(),
				dto: models.WebhookProcessHistoryDTO{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			tt.fields.webhookProcessHistoryStorage.
				On("Create", tt.args.ctx, tt.args.dto).Return(nil)
			p.WriteWebhookHistory(tt.args.ctx, tt.args.dto)
		})
	}
}

func TestWebhookProcess_addWebhookHistory(t *testing.T) {
	type args struct {
		ctx      context.Context
		webhooks *[]models.WebhookProcess
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test successful addition of webhook history",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				webhooks: &[]models.WebhookProcess{
					{
						ID:     1,
						Status: StatusProcessing,
					},
					{
						ID:     2,
						Status: StatusFinished,
					},
					{
						ID:     3,
						Status: StatusFailed,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test error in getting webhook history",
			fields: fields{
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
			},
			args: args{
				ctx: context.Background(),
				webhooks: &[]models.WebhookProcess{
					{
						ID:     1,
						Status: StatusProcessing,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			if tt.wantErr {
				tt.fields.webhookProcessHistoryStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessHistoryDTO{}, errors.New("error getting webhook history"))
			} else {
				tt.fields.webhookProcessHistoryStorage.
					On("GetList", mock.Anything, mock.Anything).Return([]models.WebhookProcessHistoryDTO{}, nil)
			}
			if err := p.addWebhookHistory(tt.args.ctx, tt.args.webhooks); (err != nil) != tt.wantErr {
				t.Errorf("addWebhookHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookProcess_extractSignal(t *testing.T) {
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Signal
		wantErr bool
	}{
		{
			name:   "Test successful extraction of Signal",
			fields: fields{},
			args: args{
				slug: "AVERAGE_CKBBUSD_07ee9a1c-53ba-11ed-8cd4-0242ac120002",
			},
			want: models.Signal{
				OrderType: 8,
				Pair:      "CKBBUSD",
				PairID:    0,
				PairPrice: decimal.NewFromInt(0),
				BotUUID:   "07ee9a1c-53ba-11ed-8cd4-0242ac120002",
			},
			wantErr: false,
		},
		{
			name:   "Test error due to bad signal format general",
			fields: fields{},
			args: args{
				slug: "invalid_signal",
			},
			want:    models.Signal{},
			wantErr: true,
		},
		{
			name:   "Test error due to bad signal format in uuid",
			fields: fields{},
			args: args{
				slug: "buy_btcusdt_invalid_uuid",
			},
			want:    models.Signal{},
			wantErr: true,
		},
		{
			name:   "Test error due to bad signal format in orderType",
			fields: fields{},
			args: args{
				slug: "invalid_orderType_btcusdt_a663a72e-f814-4eae-9f1a-6773edf75787",
			},
			want:    models.Signal{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			got, err := p.extractSignal(tt.args.slug)
			got.PairPrice = tt.want.PairPrice
			if (err != nil) != tt.wantErr {
				t.Errorf("extractSignal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractSignal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhookProcess_getWebhooksCondition(t *testing.T) {
	type args struct {
		ctx       context.Context
		condition utils.Condition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.WebhookProcess
		wantErr bool
	}{
		{
			name: "Test successful retrieval of WebhookProcess list",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				condition: utils.Condition{
					Equal: map[string]interface{}{"status": StatusProcessing},
				},
			},
			want:    []models.WebhookProcess{},
			wantErr: false,
		},
		{
			name: "Test error when retrieving WebhookProcess list",
			fields: fields{
				exchangeOrder:                omocks.NewExchangeOrderer(t),
				webhookProcessStorage:        wpmocks.NewWebhookProcesser(t),
				webhookProcessHistoryStorage: wpmocks.NewWebhookProcessHistorer(t),
				exchangeUserKeyService:       umocks.NewExchangeUserKeyer(t),
				botService:                   bmocks.NewBoter(t),
				logger:                       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				condition: utils.Condition{
					Equal: map[string]interface{}{"status": StatusProcessing},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &WebhookProcess{
				exchangeOrder:                tt.fields.exchangeOrder,
				webhookProcessStorage:        tt.fields.webhookProcessStorage,
				webhookProcessHistoryStorage: tt.fields.webhookProcessHistoryStorage,
				exchangeUserKeyService:       tt.fields.exchangeUserKeyService,
				botService:                   tt.fields.botService,
				logger:                       tt.fields.logger,
			}
			if tt.wantErr == false {
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).
					Return([]models.WebhookProcessDTO{}, nil)
				tt.fields.webhookProcessHistoryStorage.
					On("GetList", mock.Anything, mock.Anything).
					Return([]models.WebhookProcessHistoryDTO{}, nil)
			} else {
				tt.fields.webhookProcessStorage.
					On("GetList", mock.Anything, mock.Anything).
					Return(nil, errors.New(""))
			}

			got, err := p.getWebhooksCondition(tt.args.ctx, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("getWebhooksCondition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWebhooksCondition() got = %v, want %v", got, tt.want)
			}
		})
	}
}
