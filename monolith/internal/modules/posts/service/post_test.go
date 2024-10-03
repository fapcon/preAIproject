package service

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"reflect"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/posts/storage/mocks"
	"testing"
)

func newTestLogger(conf config.AppConf) *zap.Logger {
	ws := bytes.NewBuffer(make([]byte, 0, 1000))
	logger := logs.NewLogger(conf, zapcore.AddSync(ws))
	return logger
}

type fields struct {
	posterStorage *mocks.Poster
	logger        *zap.Logger
}

func TestPoster_Create(t *testing.T) {

	type args struct {
		ctx context.Context
		in  PostCreateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PostCreateOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				posterStorage: mocks.NewPoster(t),
				logger:        newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: PostCreateIn{
					Title:            "Test Title",
					ShortDescription: "Test Short Description",
					FullDescription:  "Test Full Description",
					Author:           1,
				},
			},
			want: PostCreateOut{
				Status:    http.StatusOK,
				ErrorCode: errors.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostService{
				storage: tt.fields.posterStorage,
				logger:  tt.fields.logger,
			}
			tt.fields.posterStorage.On("Create", mock.Anything, mock.Anything).Return(nil)

			if got := p.Create(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoster_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		in  PostDeleteIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PostDeleteOut
	}{
		{
			name: "Valid Deletion",
			fields: fields{
				posterStorage: mocks.NewPoster(t),
				logger:        newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: PostDeleteIn{
					Id:     1,
					Author: 1,
				},
			},
			want: PostDeleteOut{
				Success:   true,
				ErrorCode: errors.NoError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostService{
				storage: tt.fields.posterStorage,
				logger:  tt.fields.logger,
			}

			tt.fields.posterStorage.On("GetById", mock.Anything, tt.args.in.Id).
				Return(models.PostDTO{ID: tt.args.in.Id, Author: tt.args.in.Author}, nil)

			tt.fields.posterStorage.On("Update", mock.Anything, mock.Anything).
				Return(nil)

			if got := p.Delete(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoster_GetById(t *testing.T) {

	type args struct {
		ctx context.Context
		in  PostGetByIdIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PostGetByIdOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				posterStorage: mocks.NewPoster(t),
				logger:        newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: PostGetByIdIn{
					Id:     1,
					Author: 1,
				},
			},
			want: PostGetByIdOut{
				Success:   true,
				ErrorCode: errors.NoError,
				Body: Data{
					Title:            "Test Title",
					ShortDescription: "Short Description",
					FullDescription:  "Full Description",
				},
			},
		},
		{
			name: "Record Not Found Case",
			fields: fields{
				posterStorage: mocks.NewPoster(t),
				logger:        newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: PostGetByIdIn{
					Id:     2,
					Author: 1,
				},
			},
			want: PostGetByIdOut{
				Success:   false,
				ErrorCode: errors.PostServiceCreatedOtherUser,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostService{
				storage: tt.fields.posterStorage,
				logger:  tt.fields.logger,
			}

			if tt.name == "Valid Case" {
				tt.fields.posterStorage.On("GetById", mock.Anything, tt.args.in.Id).
					Return(models.PostDTO{
						ID:               tt.args.in.Id,
						Author:           tt.args.in.Author,
						Title:            "Test Title",
						ShortDescription: "Short Description",
						FullDescription:  "Full Description",
					}, nil)
			} else if tt.name == "Record Not Found Case" {
				tt.fields.posterStorage.On("GetById", mock.Anything, tt.args.in.Id).
					Return(models.PostDTO{}, nil)
			}
			got := p.GetById(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoster_GetListTape(t *testing.T) {
	type args struct {
		ctx context.Context
		in  PostGetTapeIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PostGetTapeOut
	}{
		{
			name: "Valid Case",
			fields: fields{
				posterStorage: mocks.NewPoster(t),
				logger:        newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: PostGetTapeIn{
					Limit:  10,
					Offset: 1,
				},
			},
			want: PostGetTapeOut{

				Body: []models.PostDTO{
					{
						ID:               1,
						Title:            "Test Post 1",
						ShortDescription: "Short Description 1",
						FullDescription:  "Full Description 1",
						Author:           1,
					},
				},
				Success:   true,
				ErrorCode: errors.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostService{
				storage: tt.fields.posterStorage,
				logger:  tt.fields.logger,
			}

			tt.fields.posterStorage.On("GetList", mock.Anything).
				Return([]models.PostDTO{
					{
						ID:               1,
						Title:            "Test Post 1",
						ShortDescription: "Short Description 1",
						FullDescription:  "Full Description 1",
						Author:           1,
					},
				}, nil)

			got := p.GetListTape(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListTape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoster_UpdatePost(t *testing.T) {
	type args struct {
		ctx context.Context
		in  PostUpdateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PostUpdateOut
	}{
		{
			name: "Successful Update",
			fields: fields{
				posterStorage: mocks.NewPoster(t),
				logger:        newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: PostUpdateIn{
					Id:               1,
					Author:           1,
					Title:            "Updated Title",
					ShortDescription: "Updated Short Description",
					FullDescription:  "Updated Full Description",
				},
			},
			want: PostUpdateOut{
				Success:   true,
				ErrorCode: errors.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostService{
				storage: tt.fields.posterStorage,
				logger:  tt.fields.logger,
			}

			tt.fields.posterStorage.On("GetById", mock.Anything, tt.args.in.Id).
				Return(models.PostDTO{
					ID:               tt.args.in.Id,
					Title:            tt.args.in.FullDescription,
					ShortDescription: tt.args.in.Title,
					FullDescription:  tt.args.in.ShortDescription,
					Author:           tt.args.in.Author,
				}, nil)
			tt.fields.posterStorage.On("Update", mock.Anything, mock.AnythingOfType("models.PostDTO")).
				Return(nil)

			got := p.UpdatePost(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
