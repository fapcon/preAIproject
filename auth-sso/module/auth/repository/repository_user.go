package repository

import (
	"context"

	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"
)

// UserRepository DAO repository
type UserRepository struct {
	db UserStorage
}

// NewUserRepository Конструктор
func NewUserRepository(db UserStorage) *UserRepository {
	return &UserRepository{db: db}
}

// SaveUser Сохранение пользователя
func (r *UserRepository) SaveUser(ctx context.Context, user *models.User) (int64, error) {
	return r.db.SaveUser(ctx, user)
}

// GetUserByEmail Поиск пользователя по email
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return r.db.GetUserByEmail(ctx, email)
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID int64) (models.User, error) {
	return r.db.GetUserByID(ctx, userID)
}

// IsAdmin Проверка на администратора
func (r *UserRepository) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	return r.db.IsAdmin(ctx, userID)
}
