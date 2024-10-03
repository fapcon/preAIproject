package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"reflect"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service"
	"testing"
)

func TestBot_Create(t *testing.T) {
	UID := 1
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.BotCreateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.BotOut
	}{
		{
			name:   "create test",
			fields: fields{Mock: mock.Mock{}},
			args: args{
				ctx: context.Background(),
				in:  service.BotCreateIn{UserID: UID},
			},
			want: service.BotOut{
				ErrorCode: 0,
				Bot:       models.Bot{UserID: UID},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boter{
				Mock: tt.fields.Mock,
			}
			b.Mock.On("Create", tt.args.ctx, tt.args.in).
				Return(tt.want)

			if got := b.Create(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_Delete(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.BotDeleteIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.BOut
	}{
		{
			name:   "delete test",
			fields: fields{Mock: mock.Mock{}},
			args: args{
				ctx: context.Background(),
				in: service.BotDeleteIn{
					UserID: 0,
				},
			},
			want: service.BOut{
				ErrorCode: 0,
				Success:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boter{
				Mock: tt.fields.Mock,
			}
			b.Mock.On("Delete", tt.args.ctx, tt.args.in).
				Return(tt.want)
			if got := b.Delete(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_Get(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.BotGetIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.BotOut
	}{
		{
			name:   "get test",
			fields: fields{Mock: mock.Mock{}},
			args: args{
				ctx: context.Background(),
				in: service.BotGetIn{
					ID: 0,
				},
			},
			want: service.BotOut{
				ErrorCode: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boter{
				Mock: tt.fields.Mock,
			}
			b.Mock.On("Get", tt.args.ctx, tt.args.in).
				Return(tt.want)
			if got := b.Get(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_List(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.BotListIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.BotListOut
	}{
		{
			name:   "list test",
			fields: fields{Mock: mock.Mock{}},
			args: args{
				ctx: context.Background(),
				in: service.BotListIn{
					UserID: 0,
				},
			},
			want: service.BotListOut{
				Success:   true,
				ErrorCode: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boter{
				Mock: tt.fields.Mock,
			}
			b.Mock.On("List", tt.args.ctx, tt.args.in).
				Return(tt.want)
			if got := b.List(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_Toggle(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.BotToggleIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.BOut
	}{
		{
			name:   "toggle test",
			fields: fields{Mock: mock.Mock{}},
			args: args{
				ctx: context.Background(),
				in: service.BotToggleIn{
					UserID: 1,
					Active: true,
				},
			},
			want: service.BOut{
				Success:   false,
				ErrorCode: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boter{
				Mock: tt.fields.Mock,
			}
			b.Mock.On("Toggle", tt.args.ctx, tt.args.in).
				Return(tt.want)
			got := b.Toggle(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Toggle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_Update(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.BotUpdateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.BOut
	}{
		{
			name:   "Update test",
			fields: fields{Mock: mock.Mock{}},
			args: args{
				ctx: context.Background(),
				in: service.BotUpdateIn{
					Bot: models.Bot{
						ID:     2,
						UserID: 1,
					},
				},
			},
			want: service.BOut{
				Success:   true,
				ErrorCode: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boter{
				Mock: tt.fields.Mock,
			}
			b.Mock.On("Update", tt.args.ctx, tt.args.in).
				Return(tt.want)
			if got := b.Update(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_WebhookSignal(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
		in  service.WebhookSignalIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   service.WebhookSignalOut
	}{
		{
			name:   "WebhookSignal test",
			fields: fields{Mock: mock.Mock{}},
			args: args{
				ctx: context.Background(),
				in: service.WebhookSignalIn{
					PairID: 2,
				},
			},
			want: service.WebhookSignalOut{
				ErrorCode: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boter{
				Mock: tt.fields.Mock,
			}
			b.Mock.On("WebhookSignal", tt.args.ctx, tt.args.in).
				Return(tt.want)
			if got := b.WebhookSignal(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebhookSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}
