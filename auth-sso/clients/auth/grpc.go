package auth

import (
	"context"
	"database/sql"
	"fmt"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/config"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/pb/sso"
)

type ClientService interface {
	Login(ctx context.Context, email string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error)
	Profile(ctx context.Context, token string) (user models.User, err error)
	SocialCallback(ctx context.Context, provider string, code string) (token string, err error)
	SocialGetRedirectURL(ctx context.Context, provider string) (url string, err error)
}

type Client struct {
	api sso.AuthClient
}

func NewClient(ctx context.Context, log *zerolog.Logger, addr string, cfg config.Client) (*Client, error) {
	const op = "clients.NewClient"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(cfg.RetriesCount)),
		grpcretry.WithPerRetryTimeout(cfg.Timeout),
	}

	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}

	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{api: sso.NewAuthClient(conn)}, nil
}

func InterceptorLogger(l *zerolog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.WithLevel(zerolog.Level(lvl)).Fields(fields).Msg(msg)
	})
}

func (c *Client) Login(ctx context.Context, email string, password string) (token string, err error) {
	const op = "clients.Login"

	response, err := c.api.Login(ctx, &sso.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token = response.GetToken()

	return token, nil
}

func (c *Client) RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error) {
	const op = "clients.RegisterNewUser"

	response, err := c.api.Register(ctx, &sso.RegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	userID = response.GetUserId()

	return userID, nil
}

func (c *Client) Profile(ctx context.Context, token string) (user models.User, err error) {
	const op = "clients.Profile"

	response, err := c.api.Profile(ctx, &sso.ProfileRequest{Token: token})
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user = models.User{
		ID:           response.Id,
		Email:        response.Email,
		Password:     response.Password,
		IsAdmin:      response.IsAdmin,
		DeleteStatus: response.DeleteStatus,
		CreatedAt:    response.CreatedAt,
		UpdatedAt:    response.UpdatedAt,
		DeletedAt:    ConvertProtoTimestampToNullTime(response.DeletedAt),
	}

	return user, nil
}

func (c *Client) SocialCallback(ctx context.Context, provider string, code string) (token string, err error) {
	const op = "clients.SocialCallback"

	response, err := c.api.SocialCallback(ctx, &sso.SocialCallbackRequest{
		Code:     code,
		Provider: provider,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return response.GetToken(), nil
}

func (c *Client) SocialGetRedirectURL(ctx context.Context, provider string) (url string, err error) {
	const op = "clients.SocialGetRedirectURL"

	response, err := c.api.SocialGetRedirectURL(ctx, &sso.SocialGetRedirectUrlRequest{Provider: provider})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return response.GetUrl(), err
}

func ConvertProtoTimestampToNullTime(ts *timestamppb.Timestamp) sql.NullTime {
	if ts != nil {
		return sql.NullTime{Time: ts.AsTime(), Valid: true}
	}
	return sql.NullTime{Valid: false}
}
