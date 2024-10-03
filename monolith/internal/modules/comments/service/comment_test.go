package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/comments/storage/mocks"
	"testing"
)

type fields struct {
	Storage *mocks.Commenter
}

func Test_commentService_CreateComment(t *testing.T) {

	type args struct {
		ctx context.Context
		in  CommentCreateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CommentCreateOut
	}{{
		name:   "Create test",
		fields: fields{Storage: mocks.NewCommenter(t)},
		args: args{
			ctx: context.Background(),
			in: CommentCreateIn{
				Comment:  "",
				AuthorID: 1,
				Source:   1,
				SourceID: 1,
			},
		},
		want: CommentCreateOut{
			Status:    http.StatusOK,
			ErrorCode: errors.NoError,
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &commentService{
				storage: tt.fields.Storage,
			}
			tt.fields.Storage.On("Create", context.Background(), mock.Anything).Return(nil)
			if got := p.CreateComment(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_UpdateComment(t *testing.T) {

	type args struct {
		ctx context.Context
		in  CommentUpdateIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CommentUpdateOut
	}{{
		name:   "Update test",
		fields: fields{Storage: mocks.NewCommenter(t)},
		args: args{
			ctx: context.Background(),
			in: CommentUpdateIn{
				Id:       2,
				Comment:  "2",
				AuthorID: 2,
			},
		},
		want: CommentUpdateOut{
			Success:   true,
			ErrorCode: errors.NoError,
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &commentService{
				storage: tt.fields.Storage,
			}
			tt.fields.Storage.On("GetByID", mock.Anything, tt.args.in.Id).
				Return(models.CommentDTO{
					ID:       tt.args.in.Id,
					AuthorID: tt.args.in.AuthorID,
					Comment:  tt.args.in.Comment,
				}, nil)

			tt.fields.Storage.On("Update", mock.Anything, mock.Anything).
				Return(nil)

			if got := p.UpdateComment(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_GetCommentByID(t *testing.T) {

	type args struct {
		ctx context.Context
		in  CommentGetByIdIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CommentGetByIdOut
	}{{
		name:   "Valid test",
		fields: fields{Storage: mocks.NewCommenter(t)},
		args: args{
			ctx: context.Background(),
			in: CommentGetByIdIn{
				Id:       3,
				AuthorID: 3,
			},
		},
		want: CommentGetByIdOut{
			Success:   true,
			ErrorCode: errors.NoError,
		},
	},
		{
			name:   "Invalid test",
			fields: fields{Storage: mocks.NewCommenter(t)},
			args: args{
				ctx: context.Background(),
				in: CommentGetByIdIn{
					Id:       4,
					AuthorID: 4,
				},
			},
			want: CommentGetByIdOut{
				Success:   false,
				ErrorCode: errors.CommentServiceCreatedOtherUser,
			},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &commentService{
				storage: tt.fields.Storage,
			}
			if tt.name == "Valid test" {
				tt.fields.Storage.On("GetByID", mock.Anything, tt.args.in.Id).
					Return(models.CommentDTO{
						ID:       tt.args.in.Id,
						AuthorID: tt.args.in.AuthorID,
					}, nil)
			} else if tt.name == "Invalid test" {
				tt.fields.Storage.On("GetByID", mock.Anything, tt.args.in.Id).
					Return(models.CommentDTO{}, nil)
			}
			got := p.GetCommentByID(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_DeleteComment(t *testing.T) {

	type args struct {
		ctx context.Context
		in  CommentDeleteIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CommentDeleteOut
	}{{
		name:   "Delete test",
		fields: fields{Storage: mocks.NewCommenter(t)},
		args: args{
			ctx: context.Background(),
			in: CommentDeleteIn{
				Id:       5,
				AuthorID: 5,
			},
		},
		want: CommentDeleteOut{
			Success:   true,
			ErrorCode: errors.NoError,
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &commentService{
				storage: tt.fields.Storage,
			}
			tt.fields.Storage.On("GetByID", mock.Anything, tt.args.in.Id).
				Return(models.CommentDTO{
					ID:       tt.args.in.Id,
					AuthorID: tt.args.in.AuthorID,
				}, nil)

			tt.fields.Storage.On("Update", mock.Anything, mock.Anything).
				Return(nil)
			if got := p.DeleteComment(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_GetCommentList(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CommentGetTapeOut
	}{
		{
			name:   "GetList test",
			fields: fields{Storage: mocks.NewCommenter(t)},
			args: args{
				ctx: context.Background(),
			},
			want: CommentGetTapeOut{
				Success:   true,
				ErrorCode: errors.NoError,
				Body:      []models.CommentDTO{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &commentService{
				storage: tt.fields.Storage,
			}
			tt.fields.Storage.On("GetList", mock.Anything).
				Return([]models.CommentDTO{}, nil)

			if got := p.GetCommentList(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommentList() = %v, want %v", got, tt.want)
			}
		})
	}
}
