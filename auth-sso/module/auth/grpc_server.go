package auth

import (
	"context"
	"database/sql"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"

	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth/repository/storage"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth/service"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/pb/sso"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auther интерфейс gRPC-сервера авторизации
type Auther interface {
	Login(ctx context.Context, email string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error)
	Profile(ctx context.Context, token string) (user models.User, err error)
	AppByID(ctx context.Context, appID int64) (app models.App, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
	RegisterNewApp(ctx context.Context, appID int64, name, redirectUrl string) (err error)
	SocialCallback(ctx context.Context, provider, code string) (token string, err error)
	SocialRedirectURL(ctx context.Context, provider string) string
}

// serverAPI структура gRPC-сервера
type serverAPI struct {
	sso.UnimplementedAuthServer
	auth Auther
}

// Register регистрация gRPC-сервера
func Register(gRPCServer *grpc.Server, auth Auther) {
	sso.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

// Login вход пользователя в систему
func (s *serverAPI) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty email")
	}

	if req.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty password")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid email or password")
		}

		return nil, status.Errorf(codes.Internal, "failed to login")
	}

	return &sso.LoginResponse{Token: token}, nil
}

// Register регистрация нового пользователя
func (s *serverAPI) Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty email")
	}

	if req.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty password")
	}

	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, service.ErrUserExists) {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}

		if errors.Is(err, service.ErrInvalidAppID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid app id")
		}
		return nil, status.Errorf(codes.Internal, "failed to register new user")
	}

	return &sso.RegisterResponse{UserId: userID}, nil
}

// Profile профиль пользователя
func (s *serverAPI) Profile(ctx context.Context, req *sso.ProfileRequest) (*sso.ProfileResponse, error) {
	if req.GetToken() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty token")
	}

	user, err := s.auth.Profile(ctx, req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	resp := &sso.ProfileResponse{
		Id:           user.ID,
		Email:        user.Email,
		Password:     user.Password,
		IsAdmin:      user.IsAdmin,
		DeleteStatus: user.DeleteStatus,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    ConvertNullTimeToProtoTimestamp(user.DeletedAt),
	}

	return resp, nil
}

// IsAdmin проверка пользователя на администратора
func (s *serverAPI) IsAdmin(ctx context.Context, req *sso.IsAdminRequest) (*sso.IsAdminResponse, error) {
	if req.GetUserId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}

	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to check if user is admin")
	}

	return &sso.IsAdminResponse{IsAdmin: isAdmin}, nil
}

func (s *serverAPI) RegisterApp(ctx context.Context, req *sso.RegisterAppRequest) (*emptypb.Empty, error) {
	if req.GetAppId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty app id")
	}

	if req.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	if req.GetRedirectUrl() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty redirect url")
	}

	err := s.auth.RegisterNewApp(ctx, req.GetAppId(), req.GetName(), req.GetRedirectUrl())
	if err != nil {
		if errors.Is(err, service.ErrAppExists) {
			return nil, status.Errorf(codes.AlreadyExists, "app already exists")
		}

		return nil, status.Errorf(codes.Internal, "failed to register new app")
	}

	return &emptypb.Empty{}, nil
}

func (s *serverAPI) SocialCallback(ctx context.Context, req *sso.SocialCallbackRequest) (*sso.LoginResponse, error) {
	if req.GetProvider() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty provider")
	}

	if req.GetCode() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty code")
	}

	token, err := s.auth.SocialCallback(ctx, req.GetProvider(), req.GetCode())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid credentials")
		}

		return nil, status.Errorf(codes.Internal, "failed to login")
	}

	return &sso.LoginResponse{Token: token}, nil
}

func (s *serverAPI) SocialGetRedirectURL(ctx context.Context, req *sso.SocialGetRedirectUrlRequest) (*sso.SocialGetRedirectUrlResponse, error) {
	if req.GetProvider() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty provider")
	}

	redirectURL := s.auth.SocialRedirectURL(ctx, req.GetProvider())

	return &sso.SocialGetRedirectUrlResponse{Url: redirectURL}, nil
}

func ConvertNullTimeToProtoTimestamp(nt sql.NullTime) *timestamppb.Timestamp {
	if nt.Valid {
		return timestamppb.New(nt.Time)
	}
	return nil
}
