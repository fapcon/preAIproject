package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/config"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/lib/jwt"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/social"
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth/repository/storage"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppID       = errors.New("invalid app id")
	ErrUserExists         = errors.New("user already exists")
	ErrAppExists          = errors.New("app already exists")
)

// UserStorage Интерфейс хранилища для пользователя
type UserStorage interface {
	SaveUser(ctx context.Context, user *models.User) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id int64) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

// AppStorage Интерфейс хранилища для приложения
type AppStorage interface {
	RegisterApp(ctx context.Context, app *models.App) error
	App(ctx context.Context, id int64) (models.App, error)
}

// CodeAuthCache Интерфейс кэша для кода авторизации.
type CodeAuthCache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, value interface{}) error
}

type Auth struct {
	log        *zerolog.Logger
	usrStorage UserStorage
	appStorage AppStorage
	social     social.AuthSocialer
	config     *config.Config
}

// NewAuth Конструктор
func NewAuth(
	log *zerolog.Logger,
	usrStorage UserStorage,
	appStorage AppStorage,
	social social.AuthSocialer,
	config *config.Config,
) *Auth {
	return &Auth{
		log:        log,
		usrStorage: usrStorage,
		appStorage: appStorage,
		social:     social,
		config:     config}
}

func (a *Auth) Login(ctx context.Context, email string, password string) (token string, err error) {
	const op = "service.Auth.Login"

	log := a.log.With().Str("op", op).Str("email", email).Logger()

	log.Info().Msgf("login attempt")

	user, err := a.usrStorage.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		log.Error().Err(err).Msgf("failed to get user")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		log.Error().Err(err).Msgf("failed to compare password")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info().Msgf("login successful")

	token, err = jwt.NewToken(user, a.config.SigningMethod, a.config.SecretKey, a.config.TokenTTL)
	if err != nil {
		log.Error().Err(err).Msgf("failed to generate token")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info().Msgf("token generated")

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error) {
	const op = "service.Auth.Register"

	log := a.log.With().Str("op", op).Str("email", email).Logger()

	log.Info().Msgf("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msgf("failed to hash password")

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	user := &models.User{
		Email:    email,
		Password: passHash,
	}

	userID, err = a.usrStorage.SaveUser(ctx, user)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		if errors.Is(err, ErrInvalidAppID) {
			return 0, fmt.Errorf("%s: %w", op, ErrInvalidAppID)
		}
		log.Error().Err(err).Msgf("failed to save user")

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info().Msgf("user registered")

	return userID, nil
}

func (a *Auth) Profile(ctx context.Context, token string) (user models.User, err error) {
	const op = "service.Auth.Profile"

	log := a.log.With().Str("op", op).Logger()

	log.Info().Msgf("getting user profile")

	userID, err := jwt.GetUserIDFromToken(token, a.config.SecretKey)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get userID")

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user, err = a.usrStorage.GetUserByID(ctx, int64(id))
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		log.Error().Err(err).Msgf("failed to get user")

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "service.Auth.IsAdmin"

	log := a.log.With().Str("op", op).Int64("userID", userID).Logger()

	log.Info().Msgf("checking if user is admin")

	isAdmin, err := a.usrStorage.IsAdmin(ctx, userID)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			return false, fmt.Errorf("%s: %w", op, ErrInvalidAppID)
		}
		log.Error().Err(err).Msgf("failed to check admin status")

		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info().Msgf("admin status checked - %t", isAdmin)

	return isAdmin, nil
}

func (a *Auth) RegisterNewApp(ctx context.Context, appID int64, name, redirectUrl string) (err error) {
	const op = "service.Auth.RegisterNewApp"

	log := a.log.With().Str("op", op).Int64("appID", appID).Logger()

	log.Info().Msgf("registering new app")

	app := &models.App{
		ID:          appID,
		Name:        name,
		RedirectURL: redirectUrl,
	}

	err = a.appStorage.RegisterApp(ctx, app)
	if err != nil {
		if errors.Is(err, storage.ErrAppExists) {
			return fmt.Errorf("%s: %w", op, ErrAppExists)
		}
		log.Error().Err(err).Msgf("failed to register new app")

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *Auth) AppByID(ctx context.Context, appID int64) (app models.App, err error) {
	const op = "service.Auth.RegisterNewApp"

	log := a.log.With().Str("op", op).Int64("appID", appID).Logger()

	log.Info().Msgf("checking if app exist")

	app, err = a.appStorage.App(ctx, appID)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			return models.App{}, fmt.Errorf("%s: %w", op, ErrInvalidAppID)
		}

		log.Error().Err(err).Msgf("failed to get app")

		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}

func (a *Auth) SocialCallback(ctx context.Context, provider, code string) (token string, err error) {
	const op = "service.Auth.SocialCallback"

	userInfo, err := a.social.AuthenticateSocialProvider(ctx, provider, code)
	if err != nil {
		return token, err
	}
	//поиск пользователя в базе
	user, err := a.usrStorage.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			_, err = a.RegisterNewUser(ctx, userInfo.Email, "")
			if err != nil {
				return "", fmt.Errorf("%s: %w", op, err)
			}
		} else {
			log.Error().Err(err).Msgf("failed to get user")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
	}

	user, err = a.usrStorage.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Генерируем токен
	token, err = jwt.NewToken(user, a.config.SigningMethod, a.config.SecretKey, a.config.TokenTTL)
	if err != nil {
		log.Error().Err(err).Msgf("failed to generate token")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info().Msgf("token generated")

	return token, nil
}

func (a *Auth) SocialRedirectURL(ctx context.Context, provider string) string {
	return a.social.GetRedirectURL(provider)
}
