package test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.uber.org/ratelimit"
	"reflect"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service/mocks"
	"sync"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"

	"bou.ke/monkey"
	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/assert"
)

func patch_binance_lib() {
	//var cl *binance.Client
	//monkey.PatchInstanceMethod(reflect.TypeOf(cl), "NewCreateOrderService", func(_ *binance.Client) *binance.CreateOrderService {
	//	return &binance.CreateOrderService{}
	//})

	var cos *binance.CreateOrderService
	monkey.PatchInstanceMethod(reflect.TypeOf(cos), "Do", func(c *binance.CreateOrderService, _ context.Context, _ ...binance.RequestOption) (res *binance.CreateOrderResponse, err error) {
		//err = c.Test(context.Background())
		//if err != nil {
		//	return nil, err
		//}

		return &binance.CreateOrderResponse{
			Status: binance.OrderStatusTypeNew,
		}, nil
	})
	//monkey.PatchInstanceMethod(reflect.TypeOf(cos), "Symbol", func(c *binance.CreateOrderService, _ string) *binance.CreateOrderService {
	//	return &binance.CreateOrderService{}
	//})
	//monkey.PatchInstanceMethod(reflect.TypeOf(cos), "Side", func(c *binance.CreateOrderService, _ binance.SideType) *binance.CreateOrderService {
	//	return &binance.CreateOrderService{}
	//})
	//monkey.PatchInstanceMethod(reflect.TypeOf(cos), "Type", func(c *binance.CreateOrderService, _ binance.OrderType) *binance.CreateOrderService {
	//	return &binance.CreateOrderService{}
	//})

	var cnos *binance.CancelOrderService
	monkey.PatchInstanceMethod(reflect.TypeOf(cnos), "Do", func(_ *binance.CancelOrderService, _ context.Context, _ ...binance.RequestOption) (res *binance.CancelOrderResponse, err error) {
		return &binance.CancelOrderResponse{
			Status: binance.OrderStatusTypeCanceled,
		}, nil
	})

	var gos *binance.GetOrderService
	monkey.PatchInstanceMethod(reflect.TypeOf(gos), "Do", func(_ *binance.GetOrderService, _ context.Context, _ ...binance.RequestOption) (res *binance.Order, err error) {
		return &binance.Order{}, nil
	})

	var lbts *binance.ListBookTickersService
	monkey.PatchInstanceMethod(reflect.TypeOf(lbts), "Do", func(_ *binance.ListBookTickersService, _ context.Context, _ ...binance.RequestOption) (res []*binance.BookTicker, err error) {
		return []*binance.BookTicker{&binance.BookTicker{}}, nil
	})

	var ks *binance.KlinesService
	monkey.PatchInstanceMethod(reflect.TypeOf(ks), "Do", func(_ *binance.KlinesService, _ context.Context, _ ...binance.RequestOption) (res []*binance.Kline, err error) {
		return []*binance.Kline{&binance.Kline{}}, nil
	})

	var gas *binance.GetAccountService
	monkey.PatchInstanceMethod(reflect.TypeOf(gas), "Do", func(_ *binance.GetAccountService, _ context.Context, _ ...binance.RequestOption) (res *binance.Account, err error) {
		return &binance.Account{}, nil
	})

	var gmas *binance.GetMarginAccountService
	monkey.PatchInstanceMethod(reflect.TypeOf(gmas), "Do", func(_ *binance.GetMarginAccountService, _ context.Context, _ ...binance.RequestOption) (res *binance.MarginAccount, err error) {
		return &binance.MarginAccount{}, nil
	})
}

func TestMain(m *testing.M) {
	//setup
	patch_binance_lib()

	//run
	_ = m.Run()

	//teardown
	monkey.UnpatchAll()

}

func TestNewPlatformBinance(t *testing.T) {
	tests := []struct {
		name string
		want service.Exchanger
	}{
		{
			name: "not_nil_result",
		},
	}
	for _, tt := range tests {
		rl := &service.RateLimiter{
			UserLimits: map[int]ratelimit.Limiter{},
			Mu:         sync.Mutex{},
		}
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, service.NewPlatformBinance(rl), "NewPlatformBinance()")
		})
	}
}

func TestNewBinanceWithKey(t *testing.T) {
	type args struct {
		apiKey    string
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want service.Exchanger
	}{
		{
			name: "not_nit_result",
			args: args{
				apiKey:    "",
				secretKey: "",
			},
		},
	}
	for _, tt := range tests {
		rl := &service.RateLimiter{
			UserLimits: map[int]ratelimit.Limiter{},
			Mu:         sync.Mutex{},
		}
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, service.NewBinanceWithKey(1, rl, tt.args.apiKey, tt.args.secretKey), "NewBinanceWithKey(%v, %v)", tt.args.apiKey, tt.args.secretKey)
		})
	}
}

func TestNewExchange(t *testing.T) {
	type args struct {
		binance *binance.Client
	}
	tests := []struct {
		name string
		args args
		want service.Exchanger
	}{
		{
			name: "not_nil_result",
			args: args{
				binance: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//assert.Equalf(t, tt.want, NewExchange(tt.args.binance), "NewExchange(%v)", tt.args.binance)
			assert.NotEmpty(t, service.NewExchange(tt.args.binance, nil), "NewExchange(%v)", tt.args.binance)
		})
	}
}

func TestExchange_BuyLimit(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedInput := service.LimitIn{Price: decimal.NewFromFloat(1.0), Quantity: decimal.NewFromFloat(1.0), Pair: "BTCUSDT"}
	expectedOutput := service.OrderOut{
		OrderID: "1",
		Status:  service.OrderStatusNew,
		Side:    service.SideBuy,
		Type:    service.TypeBuyLimit,
		Amount:  decimal.NewFromFloat(1),
	}

	mockExchanger.Mock.On("BuyLimit", expectedInput).Return(expectedOutput)

	result := mockExchanger.BuyLimit(expectedInput)

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("BuyLimit() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_BuyMarket(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedInput := service.MarketIn{Quantity: decimal.NewFromFloat(1.0), Pair: "BTCUSDT"}
	expectedOutput := service.OrderOut{
		OrderID: "2",
		Status:  service.OrderStatusNew,
		Side:    service.SideBuy,
		Type:    service.TypeBuyMarket,
		Amount:  decimal.NewFromFloat(1),
	}

	mockExchanger.Mock.On("BuyMarket", expectedInput).Return(expectedOutput)
	result := mockExchanger.BuyMarket(expectedInput)
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("BuyMarket() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_CancelOrder(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedInput := service.CancelOrderIn{
		OrderID: "3",
	}
	expectedOutput := service.OrderOut{
		OrderID: "3",
		Status:  service.OrderStatusCanceled,
		Side:    service.SideSell,
		Type:    service.TypeSellMarket,
		Amount:  decimal.NewFromFloat(1),
	}

	mockExchanger.Mock.On("CancelOrder", mock.Anything, expectedInput).Return(expectedOutput)

	result := mockExchanger.CancelOrder(context.Background(), expectedInput)

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("CancelOrder() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_CreateOrder(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}
	expectedInput := service.CreateOrderIn{
		Pair:     "BTCUSDT",
		Quantity: decimal.NewFromFloat(1.0),
		Price:    decimal.NewFromFloat(10000.0),
		Side:     service.SideBuy,
		Type:     service.TypeBuyLimit,
	}
	expectedOutput := service.OrderOut{
		OrderID: "4",
		Status:  service.OrderStatusNew,
		Side:    service.SideBuy,
		Type:    service.TypeSellLimit,
		Amount:  decimal.NewFromFloat(1),
	}

	mockExchanger.Mock.On("CreateOrder", expectedInput).Return(expectedOutput)

	result := mockExchanger.CreateOrder(expectedInput)

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("CreateOrder() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_GetAccount(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}
	expectedOutput := service.GetAccountOut{
		ErrorCode: 0,
		DataSpot: service.AccountSpot{
			CanTrade:    true,
			CanDeposit:  true,
			CanWithdraw: true,
			Permissions: []int{1, 2, 3},
			Balances: []service.Balance{
				{
					Currency: "BTC",
					Amount:   decimal.NewFromFloat(1.0),
					Locked:   decimal.NewFromFloat(1.0),
				},
				{
					Currency: "USDT",
					Amount:   decimal.NewFromFloat(1.0),
					Locked:   decimal.NewFromFloat(1.0),
				},
			},
		},
		DataMargin: service.AccountMargin{
			BorrowEnabled:   true,
			TradeEnabled:    true,
			TransferEnabled: true,
			Balances: []service.BalanceMargin{
				{
					Currency: "BTC",
					Free:     decimal.NewFromFloat(1.0),
					Locked:   decimal.NewFromFloat(1.0),
				},
				{
					Currency: "USDT",
					Free:     decimal.NewFromFloat(1.0),
					Locked:   decimal.NewFromFloat(1.0),
				},
			},
		},
		Message: "Success",
		Success: true,
	}

	mockExchanger.Mock.On("GetAccount", mock.Anything).Return(expectedOutput)

	result := mockExchanger.GetAccount(context.Background())

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("GetAccount() = %v, want %v", result, expectedOutput)
	}
}

// not implement
func TestExchange_GetBalances(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.GetAccountBalanceOut{
		ErrorCode: 0,
		DataSpotBalance: []models.BalanceDTO{
			{
				Currency: "BTC",
				Amount:   decimal.NewFromFloat(10.0),
				Locked:   decimal.NewFromFloat(5.0),
			},
			{
				Currency: "ETH",
				Amount:   decimal.NewFromFloat(20.0),
				Locked:   decimal.NewFromFloat(0.0),
			},
		},
		DataMarginBalance: []models.BalanceDTO{
			{
				Currency: "BTC",
				Amount:   decimal.NewFromFloat(10.0),
				Locked:   decimal.NewFromFloat(5.0),
			},
			{
				Currency: "ETH",
				Amount:   decimal.NewFromFloat(20.0),
				Locked:   decimal.NewFromFloat(0.0),
			},
		},
		Success: true,
	}

	mockExchanger.Mock.On("GetBalances", mock.Anything).Return(expectedOutput)
	result := mockExchanger.GetBalances(context.Background())

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("GetBalances() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_GetCandles(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.CandlesOut{
		ErrorCode: 0,
		Candles: []service.CandlesData{
			{
				OpenTime:  time.Unix(0, 0),
				CloseTime: time.Unix(0, 0),
			},
		},
	}

	mockExchanger.Mock.On("GetCandles", mock.Anything, mock.Anything).Return(expectedOutput)

	result := mockExchanger.GetCandles(context.Background(), service.GetCandlesIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("GetCandles() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_GetOpenOrders(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.EOut{
		ErrorCode: 0,
	}

	mockExchanger.Mock.On("GetOpenOrders", mock.Anything).Return(expectedOutput)

	result := mockExchanger.GetOpenOrders(service.EIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("GetOpenOrders() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_GetOrder(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.OrderOut{
		OrderID: "123",
		Status:  service.OrderStatusFilled,
		Side:    service.SideSell,
		Type:    service.ExchangeOrderTypeMarket,
		Amount:  decimal.NewFromFloat(1.5),
	}

	mockExchanger.Mock.On("GetOrder", mock.Anything, mock.Anything).Return(expectedOutput)

	result := mockExchanger.GetOrder(context.Background(), service.GetOrderIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("GetOrder() = %v, want %v", result, expectedOutput)
	}
}

// not implement
func TestExchange_GetOrdersHistory(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.EOut{
		ErrorCode: 0,
	}

	mockExchanger.Mock.On("GetOrdersHistory", mock.Anything).Return(expectedOutput)

	result := mockExchanger.GetOrdersHistory(service.EIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("GetOrdersHistory() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_GetTicker(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.GetTickerOut{
		ErrorCode: 0,
		Data: map[string]decimal.Decimal{
			"BTCUSDT": decimal.NewFromFloat(1.0),
			"ETHUSDT": decimal.NewFromFloat(2.0),
		},
	}

	mockExchanger.Mock.On("GetTicker", mock.Anything, mock.Anything).Return(expectedOutput)

	result := mockExchanger.GetTicker(context.Background(), service.GetTickerIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("GetTicker() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_OrderLimit(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.OrderOut{
		OrderID:  "123",
		Status:   service.OrderStatusCanceled,
		Side:     service.SideBuy,
		Type:     service.TypeGetOrder,
		Price:    decimal.NewFromFloat(100.0),
		Quantity: decimal.NewFromFloat(1.0),
	}

	mockExchanger.Mock.On("OrderLimit", mock.Anything).Return(expectedOutput)

	result := mockExchanger.OrderLimit(service.OrderLimitIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("OrderLimit() = %v, want %v", result, expectedOutput)
	}
}
func TestExchange_OrderMarket(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.OrderOut{
		OrderID:  "456",
		Status:   service.OrderStatusNew,
		Side:     service.SideSell,
		Type:     service.ExchangeOrderTypeMarket,
		Price:    decimal.NewFromFloat(200.0),
		Quantity: decimal.NewFromFloat(2.0),
	}

	mockExchanger.Mock.On("OrderMarket", mock.Anything).Return(expectedOutput)

	result := mockExchanger.OrderMarket(service.OrderMarketIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("OrderMarket() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_SellLimit(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	expectedOutput := service.OrderOut{
		OrderID:  "789",
		Status:   service.OrderStatusNew,
		Side:     service.SideSell,
		Type:     service.TypeSellLimit,
		Price:    decimal.NewFromFloat(300.0),
		Quantity: decimal.NewFromFloat(3.0),
	}

	mockExchanger.Mock.On("SellLimit", mock.Anything).Return(expectedOutput)

	result := mockExchanger.SellLimit(service.LimitIn{})

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("SellLimit() = %v, want %v", result, expectedOutput)
	}
}

func TestExchange_SellMarket(t *testing.T) {
	mockExchanger := &mocks.Exchanger{
		Mock: mock.Mock{},
	}

	// Ожидаемый вывод для успешной продажи на рынке
	expectedOutputSuccess := service.OrderOut{
		OrderID:  "123",
		Status:   service.OrderStatusNew,
		Side:     service.SideSell,
		Type:     service.TypeSellMarket,
		Quantity: decimal.NewFromFloat(2.0),
		Price:    decimal.NewFromFloat(200.0),
	}

	// Настройка мока для успешной продажи на рынке
	mockExchanger.Mock.On("SellMarket", mock.Anything).Return(expectedOutputSuccess)

	// Тест успешной продажи на рынке
	resultSuccess := mockExchanger.SellMarket(service.MarketIn{})
	if !reflect.DeepEqual(resultSuccess, expectedOutputSuccess) {
		t.Errorf("SellMarket() = %v, want %v", resultSuccess, expectedOutputSuccess)
	}
}

func Test_getSpotBalance(t *testing.T) {
	type args struct {
		spotAccount *binance.Account
	}
	tests := []struct {
		name string
		args args
		want []service.Balance
	}{
		{
			name: "test_with_balance",
			args: args{
				spotAccount: &binance.Account{
					Balances: []binance.Balance{
						binance.Balance{
							Asset:  "1",
							Free:   "1",
							Locked: "1",
						}},
				},
			},
			want: []service.Balance{
				service.Balance{
					Currency: "1",
					Amount:   decimal.NewFromInt(1),
					Locked:   decimal.NewFromInt(1),
				},
			},
		},
		{
			name: "test_without_balance",
			args: args{
				spotAccount: &binance.Account{
					Balances: []binance.Balance{
						binance.Balance{
							Asset:  "0",
							Free:   "0",
							Locked: "0",
						}},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, service.GetSpotBalance(tt.args.spotAccount), "getSpotBalance(%v)", tt.args.spotAccount)
		})
	}
}

func Test_getMarginBalance(t *testing.T) {
	type args struct {
		marginBalance *binance.MarginAccount
	}
	tests := []struct {
		name string
		args args
		want []service.BalanceMargin
	}{
		{
			name: "test_with_balance",
			args: args{
				marginBalance: &binance.MarginAccount{
					UserAssets: []binance.UserAsset{
						binance.UserAsset{
							Asset:    "",
							Borrowed: "",
							Free:     "",
							Interest: "",
							Locked:   "",
							NetAsset: "1",
						},
					},
				},
			},
			want: []service.BalanceMargin{
				service.BalanceMargin{
					NetAsset: decimal.NewFromInt(1),
				},
			},
		},
		{
			name: "test_without_balance",
			args: args{
				marginBalance: &binance.MarginAccount{
					UserAssets: []binance.UserAsset{
						binance.UserAsset{
							Asset:    "",
							Borrowed: "",
							Free:     "",
							Interest: "",
							Locked:   "",
							NetAsset: "0",
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, service.GetMarginBalance(tt.args.marginBalance), "getMarginBalance(%v)", tt.args.marginBalance)
		})
	}
}

func Test_getPrice(t *testing.T) {
	type args struct {
		priceRaw decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "price > 10",
			args: args{
				priceRaw: decimal.NewFromFloat(11.11),
			},
			want: "11.1",
		},
		{
			name: "price > 1",
			args: args{
				priceRaw: decimal.NewFromFloat(1.11),
			},
			want: "1.11",
		},
		{
			name: "price > 0",
			args: args{
				priceRaw: decimal.NewFromFloat(0.11),
			},
			want: "0.11",
		},
		{
			name: "price == 0",
			args: args{
				priceRaw: decimal.NewFromFloat(0),
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, service.GetPrice(tt.args.priceRaw), "getPrice(%v)", tt.args.priceRaw)
		})
	}
}

func Test_getQuantity(t *testing.T) {
	type args struct {
		quantityRaw decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "quantity > 10",
			args: args{
				quantityRaw: decimal.NewFromFloat(11.11),
			},
			want: "11.1",
		},
		{
			name: "quantity > 1",
			args: args{
				quantityRaw: decimal.NewFromFloat(1.11),
			},
			want: "1.11",
		},
		{
			name: "quantity > 0",
			args: args{
				quantityRaw: decimal.NewFromFloat(0.11),
			},
			want: "0.11",
		},
		{
			name: "quantity == 0",
			args: args{
				quantityRaw: decimal.NewFromInt(0),
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, service.GetQuantity(tt.args.quantityRaw), "getQuantity(%v)", tt.args.quantityRaw)
		})
	}
}

func TestExchange_composeOrder(t *testing.T) {
	type fields struct {
		binance    *binance.Client
		exchangeID int
	}
	type args struct {
		orderRaw  interface{}
		orderType int
		side      int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.OrderOut
	}{
		{
			name:   "bad_type",
			fields: fields{},
			args: args{
				orderRaw:  nil,
				orderType: 0,
				side:      0,
			},
			want: service.OrderOut{ErrorCode: errors.InternalError},
		},
		{
			name:   "bad_status_sell_limit",
			fields: fields{},
			args: args{
				orderRaw: &binance.CreateOrderResponse{
					Status: "",
				},
				orderType: service.TypeSellLimit,
				side:      0,
			},
			want: service.OrderOut{ErrorCode: errors.InternalError},
		},
		{
			name:   "bad_status_cansel",
			fields: fields{},
			args: args{
				orderRaw: &binance.CancelOrderResponse{
					Status: "",
				},
				orderType: service.TypeCancelAll,
				side:      0,
			},
			want: service.OrderOut{ErrorCode: errors.InternalError},
		},

		{
			name:   "type_sell_limit",
			fields: fields{},
			args: args{
				orderRaw: &binance.CreateOrderResponse{
					Fills: []*binance.Fill{
						&binance.Fill{
							Price: "1.23",
						},
					},
					Status: binance.OrderStatusTypeNew,
				},
				orderType: service.TypeSellLimit,
				side:      0,
			},
			want: service.OrderOut{
				OrderID: "0",
				Price:   decimal.NewFromFloatWithExponent(1.23, -8),
				Status:  service.OrderStatusNew,
				Type:    service.TypeSellLimit,
			},
		},

		{
			name:   "type_get_order",
			fields: fields{},
			args: args{
				orderRaw: &binance.Order{
					Status: binance.OrderStatusTypeNew,
				},
				orderType: service.TypeGetOrder,
				side:      0,
			},
			want: service.OrderOut{
				OrderID: "0",
				Status:  service.OrderStatusNew,
				Type:    service.TypeGetOrder,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &service.Exchange{
				Binance:    tt.fields.binance,
				ExchangeID: tt.fields.exchangeID,
			}
			out := e.ComposeOrder(tt.args.orderRaw, tt.args.orderType, tt.args.side)
			require.True(t, decimal.NewFromFloat(0).Equal(out.Amount))

			out.Amount = tt.want.Amount

			assert.Equalf(t, tt.want, out, "composeOrder(%v, %v, %v)", tt.args.orderRaw, tt.args.orderType, tt.args.side)
		})
	}
}
