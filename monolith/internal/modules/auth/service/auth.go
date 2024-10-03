package service

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	iservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/service/social"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/service/social/facebook"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/service/social/google"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/storage"
	uservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/service"
	"time"
)

type Auth struct {
	conf         config.AppConf
	user         uservice.Userer
	verify       storage.Verifier
	notify       iservice.Notifier
	tokenManager cryptography.TokenManager
	hash         cryptography.Hasher
	logger       *zap.Logger
	social       social.AuthSocialer
}

func NewAuth(user uservice.Userer, verify storage.Verifier, components *component.Components) *Auth {
	googleAuth := google.NewGoogleAuth(components.Conf, "http://127.0.0.1:8080/api/1/auth/google/callback")
	facebookAuth := facebook.NewFacebookAuth(components.Conf, "http://localhost:8080/api/1/auth/facebook/callback")
	socialService := social.NewAuthSocial(googleAuth, facebookAuth)
	return &Auth{conf: components.Conf,
		user:         user,
		verify:       verify,
		notify:       components.Notify,
		tokenManager: components.TokenManager,
		hash:         components.Hash,
		logger:       components.Logger,
		social:       socialService,
	}
}

func (a *Auth) Register(ctx context.Context, in RegisterIn, field int) RegisterOut {
	hashPass, err := cryptography.HashPassword(in.Password)
	if err != nil {
		return RegisterOut{
			Status:    http.StatusInternalServerError,
			ErrorCode: errors.HashPasswordError,
		}
	}

	userCreate := uservice.UserCreateIn{
		Email:          in.Email,
		Password:       hashPass,
		IdempotencyKey: in.IdempotencyKey,
	}

	userOut := a.user.Create(ctx, userCreate)
	if userOut.ErrorCode != errors.NoError {
		if userOut.ErrorCode == errors.UserServiceUserAlreadyExists {
			return RegisterOut{
				Status:    http.StatusConflict,
				ErrorCode: userOut.ErrorCode,
			}
		}
		return RegisterOut{
			Status:    http.StatusInternalServerError,
			ErrorCode: userOut.ErrorCode,
		}
	}
	user := a.user.GetByEmail(ctx, uservice.GetByEmailIn{Email: in.Email})
	if user.ErrorCode != errors.NoError {
		return RegisterOut{
			Status:    http.StatusInternalServerError,
			ErrorCode: userOut.ErrorCode,
		}
	}

	hash := a.hash.GenHashString(nil, cryptography.UUID)
	err = a.verify.Create(ctx, in.Email, hash, user.User.ID)
	if err != nil {
		return RegisterOut{
			Status:    http.StatusInternalServerError,
			ErrorCode: http.StatusInternalServerError,
		}
	}
	go a.sendEmailVerifyLink(in.Email, hash)

	return RegisterOut{
		Status:    http.StatusOK,
		ErrorCode: errors.NoError,
	}
}

func (a *Auth) sendEmailVerifyLink(email, hash string) int {
	userOut := a.user.GetByEmail(context.Background(), uservice.GetByEmailIn{Email: email})
	if userOut.ErrorCode != errors.NoError {
		return userOut.ErrorCode
	}

	u, err := url.Parse("http://bing.com/verify?email=sample&hash=sample")
	if err != nil {
		a.logger.Fatal("auth: url parse err", zap.Error(err))

		return errors.AuthUrlParseErr
	}
	u.Scheme = "https"
	if a.conf.Environment != "production" {
		u.Scheme = "http"
	}
	u.Host = a.conf.Domain
	q := u.Query()
	q.Set("email", email)
	q.Set("hash", hash)
	u.RawQuery = q.Encode()
	a.notifyEmail(iservice.PushIn{
		Identifier: email,
		Type:       iservice.PushEmail,
		Title:      "Diamond Trade Activation Link",
		Data:       []byte(u.String()),
		Options:    nil,
	})

	return errors.NoError
}

// TODO: Refactor
func (a *Auth) notifyEmail(p iservice.PushIn) {
	res := a.notify.Push(p)
	if res.ErrorCode != errors.NoError {
		time.Sleep(1 * time.Minute)
		go a.notifyEmail(p)
	}
}

func (a *Auth) SocialCallback(ctx context.Context, in SocialCallbackIn) AuthorizeOut {
	//отправка запроса к соответствующему сервису на получение данных о пользователе
	userInfo, err := a.social.AuthenticateSocialProvider(ctx, in.Provider, in.Code)
	if err != nil {
		return AuthorizeOut{
			ErrorCode: errors.AuthServiceSocialAuthErr,
		}
	}
	//поиск пользователя в базе
	var userOut uservice.UserOut
	userOut = a.user.GetByEmail(ctx, uservice.GetByEmailIn{Email: userInfo.Email})

	// в случае если пользователя не удалось получить его необходимо внести в систему используя данные сервиса ( !=  полностью зарегистрировать, так как не пароля )
	if userOut.ErrorCode != errors.NoError {
		registerOut := a.Register(ctx, RegisterIn{
			Email: userInfo.Email}, RegisterEmail)

		if registerOut.ErrorCode != errors.NoError {
			return AuthorizeOut{
				ErrorCode: registerOut.ErrorCode,
			}
		}

		userOut = a.user.GetByEmail(ctx, uservice.GetByEmailIn{Email: userInfo.Email})
		if userOut.ErrorCode != errors.NoError {
			return AuthorizeOut{
				ErrorCode: userOut.ErrorCode,
			}
		}
	}

	user := userOut.User
	accessToken, refreshToken, errorCode := a.generateTokens(user)
	if errorCode != errors.NoError {
		return AuthorizeOut{
			ErrorCode: errorCode,
		}
	}
	return AuthorizeOut{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (a *Auth) SocialGetRedirectURL(ctx context.Context, in SocialGetRedirectUrlIn) SocialGetRedirectUrlOut {
	return SocialGetRedirectUrlOut{
		Url: a.social.GetRedirectURL(in.Provider),
	}
}

func (a *Auth) AuthorizeEmail(ctx context.Context, in AuthorizeEmailIn) AuthorizeOut {
	userOut := a.user.GetByEmail(ctx, uservice.GetByEmailIn{Email: in.Email})
	if userOut.ErrorCode != errors.NoError {
		return AuthorizeOut{
			ErrorCode: userOut.ErrorCode,
		}
	}
	user := userOut.User
	if !cryptography.CheckPassword(user.Password, in.Password) {
		return AuthorizeOut{
			ErrorCode: errors.AuthServiceWrongPasswordErr,
		}
	}
	if !user.EmailVerified {
		return AuthorizeOut{
			ErrorCode: errors.AuthServiceUserNotVerified,
		}
	}

	accessToken, refreshToken, errorCode := a.generateTokens(user)
	if errorCode != errors.NoError {
		return AuthorizeOut{
			ErrorCode: errorCode,
		}
	}

	return AuthorizeOut{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (a *Auth) AuthorizeRefresh(ctx context.Context, in AuthorizeRefreshIn) AuthorizeOut {
	userOut := a.user.GetByID(ctx, uservice.GetByIDIn{UserID: in.UserID})
	if userOut.ErrorCode != errors.NoError {
		return AuthorizeOut{
			ErrorCode: userOut.ErrorCode,
		}
	}
	user := userOut.User

	accessToken, refreshToken, errorCode := a.generateTokens(user)
	if errorCode != errors.NoError {
		return AuthorizeOut{
			ErrorCode: errorCode,
		}
	}

	return AuthorizeOut{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (a *Auth) generateTokens(user *models.User) (string, string, int) {
	accessToken, err := a.tokenManager.CreateToken(
		strconv.Itoa(user.ID),
		strconv.Itoa(user.Role),
		"",
		a.conf.Token.AccessTTL,
		cryptography.AccessToken,
	)
	if err != nil {
		a.logger.Error("auth: create access token err", zap.Error(err))
		return "", "", errors.AuthServiceAccessTokenGenerationErr
	}
	refreshToken, err := a.tokenManager.CreateToken(
		strconv.Itoa(user.ID),
		strconv.Itoa(user.Role),
		"",
		a.conf.Token.RefreshTTL,
		cryptography.RefreshToken,
	)
	if err != nil {
		a.logger.Error("auth: create access token err", zap.Error(err))
		return "", "", errors.AuthServiceRefreshTokenGenerationErr
	}

	return accessToken, refreshToken, errors.NoError
}

func (a *Auth) AuthorizePhone(ctx context.Context, in AuthorizePhoneIn) AuthorizeOut {
	return AuthorizeOut{}
}

func (a *Auth) SendPhoneCode(ctx context.Context, in SendPhoneCodeIn) SendPhoneCodeOut {
	panic("asfasf")
}

func (a *Auth) VerifyEmail(ctx context.Context, in VerifyEmailIn) VerifyEmailOut {
	dto, err := a.verify.GetByEmail(ctx, in.Email, in.Hash)
	if err != nil {
		return VerifyEmailOut{
			ErrorCode: errors.AuthServiceVerifyErr,
		}
	}
	err = a.verify.VerifyEmail(ctx, dto.GetEmail(), dto.GetHash())
	if err != nil {
		return VerifyEmailOut{
			ErrorCode: errors.AuthServiceVerifyErr,
		}
	}
	out := a.user.VerifyEmail(ctx, uservice.UserVerifyEmailIn{
		UserID: dto.GetUserID(),
	})

	return VerifyEmailOut{
		Success:   out.Success,
		ErrorCode: out.ErrorCode,
	}
}
