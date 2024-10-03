package service

import (
	"context"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/grpc/auth"
)

type AuthGRPC struct {
	user   service.Userer
	client auth.AuthServiceGRPCClient
	logger *zap.Logger
}

func NewAuthGRPC(user service.Userer, client auth.AuthServiceGRPCClient, components *component.Components) *AuthGRPC {
	return &AuthGRPC{
		user:   user,
		client: client,
		logger: components.Logger,
	}
}

func (a AuthGRPC) Register(ctx context.Context, in RegisterIn, field int) RegisterOut {
	regOut, err := a.client.Register(ctx, &auth.RegisterIn{
		Email:          in.Email,
		Phone:          in.Phone,
		Password:       in.Password,
		IdempotencyKey: in.IdempotencyKey,
		Field:          int32(field),
	})
	if err != nil {
		a.logger.Error("Error register", zap.Error(err))
		return RegisterOut{
			ErrorCode: int(regOut.GetErrorCode()),
		}
	}
	return RegisterOut{
		Status:    int(regOut.GetStatus()),
		ErrorCode: int(regOut.GetErrorCode()),
	}
}

func (a AuthGRPC) AuthorizeEmail(ctx context.Context, in AuthorizeEmailIn) AuthorizeOut {
	authOut, err := a.client.AuthorizeEmail(ctx, &auth.AuthorizeEmailIn{
		Email:          in.Email,
		Password:       in.Password,
		RetypePassword: in.RetypePassword,
	})

	if err != nil {
		a.logger.Error("Error authorize email", zap.Error(err))
		return AuthorizeOut{
			ErrorCode: int(authOut.GetErrorCode()),
		}
	}

	return AuthorizeOut{
		UserID:       int(authOut.UserId),
		AccessToken:  authOut.GetAccessToken(),
		RefreshToken: authOut.GetRefreshToken(),
	}
}

func (a AuthGRPC) AuthorizeRefresh(ctx context.Context, in AuthorizeRefreshIn) AuthorizeOut {
	authOut, err := a.client.AuthorizeRefresh(ctx, &auth.AuthorizeRefreshIn{UserId: int32(in.UserID)})

	if err != nil {
		a.logger.Error("Error authorize refresh", zap.Error(err))
		return AuthorizeOut{
			ErrorCode: int(authOut.GetErrorCode()),
		}
	}

	return AuthorizeOut{
		UserID:       int(authOut.UserId),
		AccessToken:  authOut.GetAccessToken(),
		RefreshToken: authOut.GetRefreshToken(),
	}
}

func (a AuthGRPC) AuthorizePhone(ctx context.Context, in AuthorizePhoneIn) AuthorizeOut {
	authOut, err := a.client.AuthorizePhone(ctx, &auth.AuthorizePhoneIn{
		Phone: in.Phone,
		Code:  int32(in.Code),
	})

	if err != nil {
		a.logger.Error("Error authorize phone", zap.Error(err))
		return AuthorizeOut{
			ErrorCode: int(authOut.GetErrorCode()),
		}
	}

	return AuthorizeOut{
		UserID:       int(authOut.UserId),
		AccessToken:  authOut.GetAccessToken(),
		RefreshToken: authOut.GetRefreshToken(),
	}
}

func (a AuthGRPC) SendPhoneCode(ctx context.Context, in SendPhoneCodeIn) SendPhoneCodeOut {

	authOut, err := a.client.SendPhoneCode(ctx, &auth.SendPhoneCodeIn{Phone: in.Phone})
	if err != nil {
		a.logger.Error("Error send phone code", zap.Error(err))
		return SendPhoneCodeOut{
			ErrorCode: int(authOut.GetErrorCode()),
		}
	}

	return SendPhoneCodeOut{
		Phone: authOut.GetPhone(),
		Code:  int(authOut.GetCode()),
	}
}

func (a AuthGRPC) VerifyEmail(ctx context.Context, in VerifyEmailIn) VerifyEmailOut {
	authOut, err := a.client.VerifyEmail(ctx, &auth.VerifyEmailIn{
		Hash:  in.Hash,
		Email: in.Email,
	})
	if err != nil {
		a.logger.Error("Error verify email", zap.Error(err))
		return VerifyEmailOut{
			ErrorCode: int(authOut.GetErrorCode()),
		}
	}

	return VerifyEmailOut{
		Success: authOut.GetSuccess(),
	}
}

func (a AuthGRPC) SocialCallback(ctx context.Context, in SocialCallbackIn) AuthorizeOut {
	authOut, err := a.client.SocialCallback(ctx, &auth.SocialCallbackIn{
		Code:     in.Code,
		Provider: in.Provider,
	})

	if err != nil {
		a.logger.Error("Error social callback", zap.Error(err))
		return AuthorizeOut{
			ErrorCode: int(authOut.GetErrorCode()),
		}
	}

	return AuthorizeOut{
		UserID:       int(authOut.GetUserId()),
		AccessToken:  authOut.GetAccessToken(),
		RefreshToken: authOut.GetRefreshToken(),
	}
}

func (a AuthGRPC) SocialGetRedirectURL(ctx context.Context, in SocialGetRedirectUrlIn) SocialGetRedirectUrlOut {

	authOut, err := a.client.SocialGetRedirectURL(ctx, &auth.SocialGetRedirectUrlIn{
		Provider: in.Provider,
	})
	if err != nil {
		a.logger.Error("Error social get redirect URL", zap.Error(err))
		return SocialGetRedirectUrlOut{
			ErrorCode: int(authOut.GetErrorCode()),
		}
	}

	return SocialGetRedirectUrlOut{
		Url: authOut.GetUrl(),
	}
}
