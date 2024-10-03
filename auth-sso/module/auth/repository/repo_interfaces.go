package repository

import (
	"context"
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"
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
