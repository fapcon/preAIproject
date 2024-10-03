package service

import (
	"bytes"
	"context"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"testing"
	"time"

	eomocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/storage/mocks"
	omocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/storage/mocks"
	eumocks "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service/mocks"
)

type fields struct {
	exchangeOrder          *omocks.ExchangeOrderer
	exchangeOrderLog       *eomocks.ExchangeOrderLogger
	exchangeUserKeyService *eumocks.ExchangeUserKeyer
	logger                 *zap.Logger
}

func newTestLogger(conf config.AppConf) *zap.Logger {
	ws := bytes.NewBuffer(make([]byte, 0, 1000))
	logger := logs.NewLogger(conf, zapcore.AddSync(ws))
	return logger
}

func TestOrder_AddOrdersStatistic(t *testing.T) {
	type args struct {
		ctx    context.Context
		orders *[]models.ExchangeOrder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   StatisticOut
	}{
		{
			name: "Test add orders statistic",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				orders: &[]models.ExchangeOrder{
					{
						ID:           4,
						UUID:         "uuid4",
						OrderID:      103,
						UserID:       4,
						ExchangeID:   4,
						UnitedOrders: 1,
						OrderType:    4,
						OrderTypeMsg: "Stop-Market Order",
						Pair:         "SOL/USDT",
						Amount:       decimal.NewFromInt(10),
						Quantity:     decimal.NewFromInt(10),
						Price:        decimal.NewFromInt(4),
						Side:         1,
						SideMsg:      "Buy",
						Message:      "",
						Status:       2,
						StatusMsg:    "New",
						History:      nil,
						SumBuy:       10000,
						ApiKeyID:     1,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
				},
			},
			want: StatisticOut{
				Keys: []models.ExchangeUserKey{
					{
						ID:         4,
						UserID:     4,
						ExchangeID: 4,
						MakeOrder:  true,
						APIKey:     strconv.Itoa(1),
						StatisticData: models.StatisticData{
							Profit: decimal.NewFromInt(0),
							ToEarn: decimal.NewFromInt(0),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeUserKeyService.
				On("ExchangeUserKeyGetByID", mock.Anything, mock.Anything).
				Return(models.ExchangeUserKeyDTO{
					ID:         4,
					UserID:     4,
					ExchangeID: 4,
					MakeOrder:  true,
					APIKey:     strconv.Itoa(1),
					SecretKey:  "",
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			if got := p.AddOrdersStatistic(tt.args.ctx, tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOrdersStatistic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_ExchangeOrderList(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetBotRelationIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GetOrdersOut
	}{
		{
			name: "Test exchange order list",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetBotRelationIn{},
			},
			want: GetOrdersOut{
				ErrorCode: 0,
				Success:   true,
				Data:      []models.ExchangeOrder{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("GetList", mock.Anything, mock.Anything).Return([]models.ExchangeOrderDTO{}, nil)
			tt.fields.exchangeOrderLog.
				On("GetList", mock.Anything, mock.Anything).Return([]models.ExchangeOrderLogDTO{}, nil)
			if got := p.ExchangeOrderList(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExchangeOrderList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_GetAllOrdersStatistic(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetUserRelationIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   StatisticOut
	}{
		{
			name: "Test get all orders statistic",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{},
			want: StatisticOut{
				Keys: []models.ExchangeUserKey{
					{
						ID:         4,
						UserID:     4,
						ExchangeID: 4,
						Label:      "",
						MakeOrder:  true,
						APIKey:     strconv.Itoa(1),
						SecretKey:  "",
						StatisticData: models.StatisticData{
							Profit: decimal.NewFromInt(0),
							ToEarn: decimal.NewFromInt(0),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("GetList", mock.Anything, mock.Anything).Return([]models.ExchangeOrderDTO{
				{
					ID:           4,
					UUID:         "uuid4",
					UserID:       4,
					ExchangeID:   4,
					UnitedOrders: 1,
					OrderType:    4,
					Pair:         "SOL/USDT",
					Amount:       decimal.NewFromInt(10),
					Quantity:     decimal.NewFromInt(10),
					Price:        decimal.NewFromInt(4),
					Side:         1,
					Message:      "",
					Status:       2,
					ApiKeyID:     1,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				},
			}, nil)
			tt.fields.exchangeUserKeyService.
				On("ExchangeUserKeyGetByID", mock.Anything, mock.Anything).
				Return(models.ExchangeUserKeyDTO{
					ID:         4,
					UserID:     4,
					ExchangeID: 4,
					MakeOrder:  true,
					APIKey:     strconv.Itoa(1),
					SecretKey:  "",
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			if got := p.GetAllOrdersStatistic(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllOrdersStatistic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_GetBotOrders(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetBotRelationIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GetOrdersOut
	}{
		{
			name: "Test get bot orders",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetBotRelationIn{},
			},
			want: GetOrdersOut{
				ErrorCode: 0,
				Success:   true,
				Data:      []models.ExchangeOrder{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderDTO{}, nil)
			tt.fields.exchangeOrderLog.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderLogDTO{}, nil)
			if got := p.GetBotOrders(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBotOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_GetOrdersCondition(t *testing.T) {
	type args struct {
		ctx       context.Context
		condition utils.Condition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.ExchangeOrder
		wantErr bool
	}{
		{
			name: "Test get bot orders",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:       context.Background(),
				condition: utils.Condition{},
			},
			want:    []models.ExchangeOrder{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderDTO{}, nil)
			tt.fields.exchangeOrderLog.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderLogDTO{}, nil)
			got, err := p.GetOrdersCondition(tt.args.ctx, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrdersCondition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrdersCondition() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_addOrdersHistory(t *testing.T) {
	type args struct {
		ctx    context.Context
		orders *[]models.ExchangeOrder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test add orders history",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx:    context.Background(),
				orders: &[]models.ExchangeOrder{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrderLog.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderLogDTO{}, nil)
			if err := p.addOrdersHistory(tt.args.ctx, tt.args.orders); (err != nil) != tt.wantErr {
				t.Errorf("addOrdersHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_GetOrdersStatistic(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetBotRelationIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   StatisticOut
	}{
		{
			name: "Test get orders statistic",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetBotRelationIn{},
			},
			want: StatisticOut{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("GetList", mock.Anything, mock.Anything).Return([]models.ExchangeOrderDTO{}, nil)
			if got := p.GetOrdersStatistic(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrdersStatistic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_GetUserOrders(t *testing.T) {
	type args struct {
		ctx context.Context
		in  GetUserRelationIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   GetOrdersOut
	}{
		{
			name: "Test get user orders",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in:  GetUserRelationIn{},
			},
			want: GetOrdersOut{
				ErrorCode: 0,
				Success:   true,
				Data:      []models.ExchangeOrder{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderDTO{}, nil)
			tt.fields.exchangeOrderLog.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderLogDTO{}, nil)
			if got := p.GetUserOrders(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_WriteOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		in  WriteOrderIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test write order",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 logs.NewLogger(config.AppConf{}, zapcore.AddSync(bytes.NewBuffer(make([]byte, 0, 1000)))),
			},
			args: args{
				ctx: context.Background(),
				in: WriteOrderIn{
					OrderType: 1,
					Side:      1,
					Webhook:   models.WebhookProcessDTO{},
					Bot:       models.Bot{},
					Signal:    models.Signal{},
					Key:       models.ExchangeUserKeyDTO{},
					BuyOrder:  models.ExchangeOrderDTO{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("Create", mock.Anything, mock.Anything).
				Return(nil)
			if err := p.WriteOrder(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("WriteOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_WriteOrderLog(t *testing.T) {
	type args struct {
		ctx         context.Context
		orderLogDTO models.ExchangeOrderLogDTO
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test write order log",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 logs.NewLogger(config.AppConf{}, zapcore.AddSync(bytes.NewBuffer(make([]byte, 0, 1000)))),
			},
			args: args{
				ctx:         context.Background(),
				orderLogDTO: models.ExchangeOrderLogDTO{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrderLog.
				On("Create", mock.Anything, mock.Anything).
				Return(nil)
			p.WriteOrderLog(tt.args.ctx, tt.args.orderLogDTO)
		})
	}
}

func TestOrder_OrderSellLimit(t *testing.T) {
	mockRateLimiter := service.NewRateLimiter()
	type args struct {
		ctx          context.Context
		in           OrderIn
		quantity     decimal.Decimal
		price        decimal.Decimal
		unitedOrders int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test order sell limit",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 logs.NewLogger(config.AppConf{}, zapcore.AddSync(bytes.NewBuffer(make([]byte, 0, 1000)))),
			},
			args: args{
				ctx: context.Background(),
				in: OrderIn{
					ExClient: service.NewPlatformBinance(mockRateLimiter),
					Webhook:  models.WebhookProcessDTO{},
					Bot:      models.Bot{},
					Signal: models.Signal{
						Pair:      "BTC/USDT",
						OrderType: 1,
					},
					Key:      models.ExchangeUserKeyDTO{},
					BuyOrder: models.ExchangeOrderDTO{},
				},
				quantity: decimal.NewFromInt(10),
				price:    decimal.NewFromInt(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("Create", mock.Anything, mock.Anything).
				Return(nil)
			//ctx context.Context,        in OrderIn, quantity,         price decimal.Decimal, unitedOrders int
			p.OrderSellLimit(tt.args.ctx, tt.args.in, tt.args.quantity, tt.args.price, tt.args.unitedOrders)
		})
	}
}

func TestOrder_CancelOrder(t *testing.T) {
	mockRateLimiter := service.NewRateLimiter()

	type args struct {
		ctx context.Context
		in  OrderIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    CancelOrderOut
		wantErr bool
	}{
		{
			name: "Test cancel order",
			fields: fields{
				exchangeOrder:          omocks.NewExchangeOrderer(t),
				exchangeOrderLog:       eomocks.NewExchangeOrderLogger(t),
				exchangeUserKeyService: eumocks.NewExchangeUserKeyer(t),
				logger:                 newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: OrderIn{
					BuyOrder: models.ExchangeOrderDTO{},
					Signal: models.Signal{
						OrderType: 6,
					},
					ExClient: service.NewPlatformBinance(mockRateLimiter),
				},
			},
			want: CancelOrderOut{
				ExchangeOrders: nil,
				PlatformOrders: make([]models.ExchangeOrderDTO, 0, 2),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Order{
				exchangeOrder:          tt.fields.exchangeOrder,
				exchangeOrderLog:       tt.fields.exchangeOrderLog,
				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
				logger:                 tt.fields.logger,
			}
			tt.fields.exchangeOrder.
				On("GetList", mock.Anything, mock.Anything).
				Return([]models.ExchangeOrderDTO{
					{
						ID:                1,
						UUID:              "uuid1",
						BotUUID:           "bot_uuid1",
						OrderType:         1,
						ExchangeOrderType: 1,
						Side:              1,
						ExchangeOrderID:   "1",
						UnitedOrders:      1,
						UserID:            1,
						ExchangeID:        1,
						PairID:            1,
						BuyOrderID:        1,
						ApiKeyID:          1,
						Pair:              "BTC",
						Amount:            decimal.NewFromInt(10),
						Quantity:          decimal.NewFromInt(10),
						Price:             decimal.NewFromInt(10000),
						BuyPrice:          decimal.NewFromInt(10000),
						Status:            1,
						WebhookUUID:       "webhook_uuid1",
						Message:           "This is a test message",
						CreatedAt:         time.Now(),
						UpdatedAt:         time.Now(),
					},
					{
						ID:                2,
						UUID:              "uuid2",
						BotUUID:           "bot_uuid2",
						OrderType:         2,
						ExchangeOrderType: 2,
						Side:              2,
						ExchangeOrderID:   "2",
						UnitedOrders:      2,
						UserID:            2,
						ExchangeID:        2,
						PairID:            2,
						BuyOrderID:        2,
						ApiKeyID:          2,
						Pair:              "ETH",
						Amount:            decimal.NewFromInt(20),
						Quantity:          decimal.NewFromInt(20),
						Price:             decimal.NewFromInt(20000),
						BuyPrice:          decimal.NewFromInt(20000),
						Status:            2,
						WebhookUUID:       "webhook_uuid2",
						Message:           "This is a second test message",
						CreatedAt:         time.Now(),
						UpdatedAt:         time.Now(),
					},
				},
					nil)
			//tt.fields.exchangeOrder.
			//	On("CancelOrder", mock.Anything, mock.Anything).
			//	Return(exchange_client.OrderOut{})
			//tt.fields.exchangeOrder.
			//	On("Update", mock.Anything, mock.Anything).
			//	Return(nil)
			got, err := p.CancelOrder(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ExchangeOrders, tt.want.ExchangeOrders) {
				t.Errorf("CancelOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: Сделать тесты ниже, когда будет сделана реализация функций.

//func TestOrder_CreateOrder(t *testing.T) {
//	type args struct {
//		ctx      context.Context
//		orderDTO models.ExchangeOrderDTO
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &Order{
//				exchangeOrder:          tt.fields.exchangeOrder,
//				exchangeOrderLog:       tt.fields.exchangeOrderLog,
//				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
//				logger:                 tt.fields.logger,
//			}
//			if err := p.CreateOrder(tt.args.ctx, tt.args.orderDTO); (err != nil) != tt.wantErr {
//				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//
//func TestOrder_UpdateOrder(t *testing.T) {
//	type args struct {
//		ctx context.Context
//		dto models.ExchangeOrderDTO
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &Order{
//				exchangeOrder:          tt.fields.exchangeOrder,
//				exchangeOrderLog:       tt.fields.exchangeOrderLog,
//				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
//				logger:                 tt.fields.logger,
//			}
//			if err := p.UpdateOrder(tt.args.ctx, tt.args.dto); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestOrder_GetOrdersByUUID(t *testing.T) {
//	type args struct {
//		ctx  context.Context
//		uuid string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    models.ExchangeOrderDTO
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &Order{
//				exchangeOrder:          tt.fields.exchangeOrder,
//				exchangeOrderLog:       tt.fields.exchangeOrderLog,
//				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
//				logger:                 tt.fields.logger,
//			}
//			got, err := p.GetOrdersByUUID(tt.args.ctx, tt.args.uuid)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetOrdersByUUID() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetOrdersByUUID() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOrder_GetOrderList(t *testing.T) {
//	type args struct {
//		ctx       context.Context
//		condition utils.Condition
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []models.ExchangeOrderDTO
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &Order{
//				exchangeOrder:          tt.fields.exchangeOrder,
//				exchangeOrderLog:       tt.fields.exchangeOrderLog,
//				exchangeUserKeyService: tt.fields.exchangeUserKeyService,
//				logger:                 tt.fields.logger,
//			}
//			got, err := p.GetOrderList(tt.args.ctx, tt.args.condition)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetOrderList() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetOrderList() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
