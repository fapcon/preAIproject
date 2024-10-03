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
	nservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/service"
	nmock "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/service/mocks"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
	cryptomock "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography/mocks"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	vmock "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/storage/mocks"
	uservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/service"
	umock "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/service/mocks"
	"testing"
)

type fields struct {
	conf         config.AppConf
	user         *umock.Userer
	verify       *vmock.Verifier
	notify       *nmock.Notifier
	tokenManager *cryptomock.TokenManager
	hash         *cryptomock.Hasher
	logger       *zap.Logger
}

func newTestLogger(conf config.AppConf) *zap.Logger {
	ws := bytes.NewBuffer(make([]byte, 0, 1000))
	logger := logs.NewLogger(conf, zapcore.AddSync(ws))
	return logger
}

func TestAuth_AuthorizeEmail(t *testing.T) {
	type args struct {
		ctx          context.Context
		in           AuthorizeEmailIn
		userVerified bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   AuthorizeOut
	}{
		{
			name: "test_authorize_email_positive",
			fields: fields{
				conf:         config.AppConf{},
				user:         umock.NewUserer(t),
				verify:       vmock.NewVerifier(t),
				notify:       nmock.NewNotifier(t),
				tokenManager: cryptomock.NewTokenManager(t),
				hash:         cryptomock.NewHasher(t),
				logger:       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: AuthorizeEmailIn{
					Email:          "test@example.com",
					Password:       "AnyTestPassword12345678",
					RetypePassword: "AnyTestPassword12345678",
				},
				userVerified: true,
			},
			want: AuthorizeOut{
				UserID:       1,
				AccessToken:  "access",
				RefreshToken: "refresh",
				ErrorCode:    errors.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, _ := cryptography.HashPassword(tt.args.in.Password)
			// Установка ожидаемого результата для мока сервиса User
			tt.fields.user.On("GetByEmail", mock.Anything, mock.Anything).Return(uservice.UserOut{
				User: &models.User{
					ID:            1,
					Verified:      tt.args.userVerified,
					EmailVerified: tt.args.userVerified,
					Password:      hashedPassword,
				},
				ErrorCode: errors.NoError,
			})
			tt.fields.tokenManager.On("CreateToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return("access", nil).Once()
			tt.fields.tokenManager.On("CreateToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return("refresh", nil).Once()
			a := &Auth{
				conf:         tt.fields.conf,
				user:         tt.fields.user,
				verify:       tt.fields.verify,
				notify:       tt.fields.notify,
				tokenManager: tt.fields.tokenManager,
				hash:         tt.fields.hash,
				logger:       tt.fields.logger,
			}

			if got := a.AuthorizeEmail(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_AuthorizeRefresh(t *testing.T) {
	type args struct {
		ctx context.Context
		in  AuthorizeRefreshIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   AuthorizeOut
	}{
		{
			name: "test_authorize_refresh",
			fields: fields{
				conf:         config.AppConf{},
				user:         umock.NewUserer(t),
				verify:       vmock.NewVerifier(t),
				notify:       nmock.NewNotifier(t),
				tokenManager: cryptomock.NewTokenManager(t),
				hash:         cryptomock.NewHasher(t),
				logger:       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: AuthorizeRefreshIn{
					1,
				},
			},
			want: AuthorizeOut{
				UserID:       1,
				AccessToken:  "access",
				RefreshToken: "refresh",
				ErrorCode:    errors.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Установка ожидаемого результата для мока сервиса User
			tt.fields.user.On("GetByID", mock.Anything, mock.Anything).Return(uservice.UserOut{
				User: &models.User{
					ID: 1,
				},
				ErrorCode: errors.NoError,
			})
			tt.fields.tokenManager.On("CreateToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return("access", nil).Once()
			tt.fields.tokenManager.On("CreateToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return("refresh", nil).Once()
			a := &Auth{
				conf:         tt.fields.conf,
				user:         tt.fields.user,
				verify:       tt.fields.verify,
				notify:       tt.fields.notify,
				tokenManager: tt.fields.tokenManager,
				hash:         tt.fields.hash,
				logger:       tt.fields.logger,
			}
			if got := a.AuthorizeRefresh(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeRefresh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_Register(t *testing.T) {
	type args struct {
		ctx   context.Context
		in    RegisterIn
		field int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   RegisterOut
	}{
		{
			name: "test_auth_register",
			fields: fields{
				conf:         config.AppConf{},
				user:         umock.NewUserer(t),
				verify:       vmock.NewVerifier(t),
				notify:       nmock.NewNotifier(t),
				tokenManager: cryptomock.NewTokenManager(t),
				hash:         cryptomock.NewHasher(t),
				logger:       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: RegisterIn{
					Email:          "test@example.com",
					Phone:          "79171937114",
					Password:       "anypassword",
					IdempotencyKey: "key",
				},
				field: 0,
			},
			want: RegisterOut{
				Status:    http.StatusOK,
				ErrorCode: errors.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.user.On("Create", mock.Anything, mock.Anything).Return(
				uservice.UserCreateOut{
					ErrorCode: errors.NoError,
				})
			tt.fields.user.On("GetByEmail", mock.Anything, mock.Anything).Return(uservice.UserOut{
				User: &models.User{
					ID: 1,
				},
				ErrorCode: errors.NoError,
			})
			tt.fields.hash.On("GenHashString", []byte(nil), 1).Return("hashedString")
			tt.fields.verify.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
			//Функция Maybe в конце вызова метода Push ввиду вызова функции внутри горутины, то есть функция вызывается после завершения тестируемой функции Register
			tt.fields.notify.On("Push", mock.Anything).Return(nservice.PushOut{ErrorCode: errors.NoError}).Maybe()
			a := &Auth{
				conf:         tt.fields.conf,
				user:         tt.fields.user,
				verify:       tt.fields.verify,
				notify:       tt.fields.notify,
				tokenManager: tt.fields.tokenManager,
				hash:         tt.fields.hash,
				logger:       tt.fields.logger,
			}
			if got := a.Register(tt.args.ctx, tt.args.in, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
			tt.fields.verify.AssertExpectations(t)
		})
	}
}

func TestAuth_VerifyEmail(t *testing.T) {
	type args struct {
		ctx context.Context
		in  VerifyEmailIn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   VerifyEmailOut
	}{
		{
			name: "test_auth_verify",
			fields: fields{
				conf:         config.AppConf{},
				user:         umock.NewUserer(t),
				verify:       vmock.NewVerifier(t),
				notify:       nmock.NewNotifier(t),
				tokenManager: cryptomock.NewTokenManager(t),
				hash:         cryptomock.NewHasher(t),
				logger:       newTestLogger(config.AppConf{}),
			},
			args: args{
				ctx: context.Background(),
				in: VerifyEmailIn{
					Hash:  "anyhash",
					Email: "test@test.com",
				},
			},
			want: VerifyEmailOut{
				Success:   true,
				ErrorCode: errors.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.verify.On("GetByEmail", mock.Anything, mock.Anything, mock.Anything).Return(models.EmailVerifyDTO{
				Email: tt.args.in.Email,
				Hash:  tt.args.in.Hash,
			}, nil)
			tt.fields.verify.On("VerifyEmail", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			tt.fields.user.On("VerifyEmail", mock.Anything, mock.Anything).Return(uservice.UserUpdateOut{
				Success:   true,
				ErrorCode: errors.NoError,
			})
			a := &Auth{
				conf:         tt.fields.conf,
				user:         tt.fields.user,
				verify:       tt.fields.verify,
				notify:       tt.fields.notify,
				tokenManager: tt.fields.tokenManager,
				hash:         tt.fields.hash,
				logger:       tt.fields.logger,
			}
			if got := a.VerifyEmail(tt.args.ctx, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
