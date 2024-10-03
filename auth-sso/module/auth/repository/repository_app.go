package repository

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"
)

// AppRepository DAO repository
type AppRepository struct {
	db AppStorage
}

// NewAppRepository Конструктор
func NewAppRepository(db AppStorage) *AppRepository {
	return &AppRepository{db: db}
}

// RegisterApp Регистрация приложения в БД.
func (a *AppRepository) RegisterApp(ctx context.Context, app *models.App) error {
	return a.db.RegisterApp(ctx, app)
}

// App Получение приложения по его ID.
func (a *AppRepository) App(ctx context.Context, id int64) (models.App, error) {
	return a.db.App(ctx, id)
}
