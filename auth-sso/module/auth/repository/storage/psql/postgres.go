package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/database/postgres"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth/repository/storage"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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

// Storage Структура хранилища PostgreSQL
type Storage struct {
	db *sqlx.DB
}

// NewStorage Конструктор
func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// SaveUser Сохранение пользователя
func (s *Storage) SaveUser(ctx context.Context, user *models.User) (int64, error) {
	const op = "storage.psql.SaveUser"

	query := postgres.StatementBuilder.
		Insert("users").
		Columns("email", "password").
		Values(user.Email, user.Password).
		Suffix("RETURNING id").
		RunWith(s.db)

	if err := query.QueryRowContext(ctx).Scan(&user.ID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return 0, errors.Wrap(storage.ErrUserExists, op)
			}
		}

		return 0, errors.Wrap(err, op)
	}

	return user.ID, nil
}

// GetUserByEmail Поиск пользователя по email
func (s *Storage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "storage.psql.GetUserByEmail"

	query := postgres.StatementBuilder.
		Select("*"). // id, email, password, delete_status
		From("users").
		Where(sq.Eq{"email": email, "delete_status": false})

	queryText, queryArgs, err := query.ToSql()
	if err != nil {
		return models.User{}, errors.Wrap(err, op)
	}

	rows, err := s.db.QueryxContext(ctx, queryText, queryArgs...)
	if err != nil {
		return models.User{}, errors.Wrap(err, op)
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return models.User{}, errors.Wrap(err, op)
		}
	}

	if user.ID == 0 {
		return models.User{}, errors.Wrap(storage.ErrUserNotFound, op)
	}

	return user, nil
}

// GetUserByID Поиск пользователя по email
func (s *Storage) GetUserByID(ctx context.Context, userID int64) (models.User, error) {
	const op = "storage.psql.GetUserByEmail"

	query := postgres.StatementBuilder.
		Select("*").
		From("users").
		Where(sq.Eq{"id": userID, "delete_status": false})

	queryText, queryArgs, err := query.ToSql()
	if err != nil {
		return models.User{}, errors.Wrap(err, op)
	}

	rows, err := s.db.QueryxContext(ctx, queryText, queryArgs...)
	if err != nil {
		return models.User{}, errors.Wrap(err, op)
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return models.User{}, errors.Wrap(err, op)
		}
	}

	if user.ID == 0 {
		return models.User{}, errors.Wrap(storage.ErrUserNotFound, op)
	}

	return user, nil
}

// IsAdmin проверка пользователя на администратора
func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "storage.psql.IsAdmin"

	query := postgres.StatementBuilder.
		Select("is_admin").
		From("users").
		Where(sq.Eq{"id": userID}).
		RunWith(s.db)

	var isAdmin bool
	row := query.QueryRowContext(ctx)
	err := row.Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return false, errors.Wrap(err, op)
	}

	return isAdmin, nil
}

// RegisterApp Регистрация приложения в хранилище.
func (s *Storage) RegisterApp(ctx context.Context, app *models.App) error {
	const op = "storage.psql.RegisterApp"

	query := postgres.StatementBuilder.
		Insert("apps").
		Columns("id", "name", "redirect_url").
		Values(app.ID, app.Name, app.RedirectURL).
		Suffix("RETURNING id").
		RunWith(s.db)

	if err := query.QueryRowContext(ctx).Scan(&app.ID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return errors.Wrap(storage.ErrAppExists, op)
			}
		}

		return errors.Wrap(err, op)
	}

	return nil
}

// App Поиск приложения по ID
func (s *Storage) App(ctx context.Context, appID int64) (models.App, error) {
	const op = "storage.psql.App"

	query := postgres.StatementBuilder.
		Select("*").
		From("apps").
		Where(sq.Eq{"id": appID})

	queryText, queryArgs, err := query.ToSql()
	if err != nil {
		return models.App{}, errors.Wrap(err, op)
	}

	rows, err := s.db.QueryxContext(ctx, queryText, queryArgs...)
	if err != nil {
		return models.App{}, errors.Wrap(err, op)
	}
	defer rows.Close()

	var app models.App
	for rows.Next() {
		if err := rows.StructScan(&app); err != nil {
			return models.App{}, errors.Wrap(err, op)
		}
	}

	if app.ID == 0 {
		return models.App{}, errors.Wrap(storage.ErrAppNotFound, op)
	}

	return app, nil
}
