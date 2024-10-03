package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"reflect"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/strategy/service"
	"testing"
)

func TestStrateger_Create(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.StrategyCreateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.StrategyCreateOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyCreateIn{
					ID:          1,
					Name:        "ValidStrategy",
					UUID:        "12345678",
					Description: "A valid strategy",
					ExchangeID:  2,
					Bots: []models.Bot{
						{
							ID:          1,
							Kind:        2,
							UserID:      3,
							Name:        "ValidBot",
							Description: "A valid bot",
							PairID:      4,
						},
					},
				},
			},
			want: service.StrategyCreateOut{
				StrategyID: 123,
				ErrorCode:  0,
			},
		},
		{
			name: "Empty Name",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyCreateIn{
					ID:          2,
					Name:        "",
					UUID:        "12345678",
					Description: "An invalid strategy",
					ExchangeID:  3,
					Bots: []models.Bot{
						{
							ID:          2,
							Kind:        1,
							UserID:      4,
							Name:        "InvalidBot",
							Description: "An invalid bot",
							PairID:      5,
						},
					},
				},
			},
			want: service.StrategyCreateOut{
				ErrorCode: errors.StrategyServiceCreateErr,
			},
		},
		{
			name: "Strategy Already Exists Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyCreateIn{
					ID:          3,
					Name:        "DuplicateStrategyName",
					UUID:        "98765432",
					Description: "A duplicate strategy",
					ExchangeID:  4,
					Bots: []models.Bot{
						{
							ID:          3,
							Kind:        1,
							UserID:      5,
							Name:        "DuplicateBotName",
							Description: "A duplicate bot",
							PairID:      6,
						},
					},
				},
			},
			want: service.StrategyCreateOut{
				ErrorCode: errors.StrategyServiceStrategyAlreadyExists,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &Strateger{
				Mock: tt.fields.Mock,
			}

			_m.Mock.On("Create", mock.Anything, mock.Anything).
				Return(tt.want)

			got := _m.Create(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrateger_Delete(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.StrategyDeleteIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.StrategyDeleteOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in:  service.StrategyDeleteIn{ID: 1},
			},
			want: service.StrategyDeleteOut{Success: true},
		},
		{
			name: "Strategy Doesnt Exist Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in:  service.StrategyDeleteIn{ID: 2},
			},
			want: service.StrategyDeleteOut{ErrorCode: errors.StrategyServiceStrategyDoesntExist},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Mock.On("Delete", mock.Anything, tt.args.in).Return(tt.want, nil)

			_m := &Strateger{
				Mock: tt.fields.Mock,
			}
			if got := _m.Delete(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrateger_GetByID(t *testing.T) {

	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.StrategyGetByIDIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.StrategyOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyGetByIDIn{
					ID: 123,
				},
			},
			want: service.StrategyOut{
				Strategy: &models.Strategy{
					ID:          123,
					Name:        "TestStrategy",
					UUID:        "UUID123",
					Description: "Test Description",
					ExchangeID:  456,
					Bots:        []models.Bot{},
				},
			},
		},
		{
			name: "Strategy Doesnt Exist Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyGetByIDIn{
					ID: 999,
				},
			},
			want: service.StrategyOut{
				ErrorCode: errors.StrategyServiceRetrieveErr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &Strateger{
				Mock: tt.fields.Mock,
			}
			_m.Mock.On("GetByID", tt.args.ctx, tt.args.in).Return(tt.want)

			if got := _m.GetByID(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrateger_GetByName(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.StrategyGetByNameIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.StrategyOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyGetByNameIn{
					Name: "TestStrategy",
				},
			},
			want: service.StrategyOut{
				Strategy: &models.Strategy{
					ID:          123,
					Name:        "TestStrategy",
					UUID:        "UUID123",
					Description: "Test Description",
					ExchangeID:  456,
					Bots:        []models.Bot{},
				},
			},
		},
		{
			name: "Strategy Doesnt Exist Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyGetByNameIn{
					Name: "Strategy Doesnt Exist",
				},
			},
			want: service.StrategyOut{
				ErrorCode: errors.StrategyServiceRetrieveErr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &Strateger{
				Mock: tt.fields.Mock,
			}
			_m.Mock.On("GetByName", tt.args.ctx, tt.args.in).Return(tt.want)

			if got := _m.GetByName(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrateger_GetList(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.StrategiesOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
			},
			want: service.StrategiesOut{
				Strategy: []models.Strategy{
					{
						ID:          123,
						Name:        "Strategy1",
						UUID:        "UUID123",
						Description: "Description1",
						ExchangeID:  456,
						Bots:        []models.Bot{},
					},
					{
						ID:          456,
						Name:        "Strategy2",
						UUID:        "UUID456",
						Description: "Description2",
						ExchangeID:  789,
						Bots:        []models.Bot{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &Strateger{
				Mock: tt.fields.Mock,
			}
			_m.Mock.On("GetList", tt.args.ctx).Return(tt.want)

			if got := _m.GetList(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrateger_Update(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.StrategyUpdateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.StrategyUpdateOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyUpdateIn{
					Strategy: models.Strategy{
						ID:          123,
						Name:        "UpdatedStrategy",
						UUID:        "UpdatedUUID",
						Description: "Updated Description",
						ExchangeID:  789,
						Bots:        []models.Bot{},
					},
					Fields: []int{1, 2, 3},
				},
			},
			want: service.StrategyUpdateOut{
				Success: true,
			},
		},
		{
			name: "Failed To Update Case",
			fields: fields{
				Mock: mock.Mock{},
			},
			args: args{
				ctx: context.TODO(),
				in: service.StrategyUpdateIn{
					Strategy: models.Strategy{
						ID: 123,
					},
					Fields: []int{1, 2, 3},
				},
			},
			want: service.StrategyUpdateOut{
				Success:   false,
				ErrorCode: errors.StrategyServiceUpdateErr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &Strateger{
				Mock: tt.fields.Mock,
			}
			_m.Mock.On("Update", tt.args.ctx, tt.args.in).Return(tt.want)

			if got := _m.Update(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
