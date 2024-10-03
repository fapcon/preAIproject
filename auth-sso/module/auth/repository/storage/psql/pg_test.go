package psql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func setup(_ *testing.T) (
	UserStorage,
	context.Context,
	sqlmock.Sqlmock,
	models.User,
) {
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	ctx := context.Background()

	repo := NewStorage(sqlxDB)

	user := models.User{
		Email:        "X9wH6@example.com",
		Password:     []byte("password"),
		DeleteStatus: false,
	}

	return repo, ctx, mock, user
}

func TestStorage_GetUser(t *testing.T) {
	t.Parallel()

	r, ctx, dbMock, dummyUser := setup(t)

	dummyUser.ID = 100

	rows := sqlmock.NewRows([]string{"id", "email", "password", "delete_status"}).AddRow(
		dummyUser.ID,
		dummyUser.Email,
		dummyUser.Password,
		dummyUser.DeleteStatus,
	)

	dbMock.ExpectQuery("SELECT \\* FROM users WHERE delete_status = \\$1 AND email = \\$2").
		WithArgs(
			dummyUser.DeleteStatus,
			dummyUser.Email,
		).WillReturnRows(rows)

	user, err := r.GetUserByEmail(ctx, dummyUser.Email)

	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.Equal(t, dummyUser.ID, user.ID)
}

func TestStorage_CreateUserSuccess(t *testing.T) {
	t.Parallel()

	r, ctx, dbMock, dummyUser := setup(t)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(dummyUser.ID + 1)

	dbMock.ExpectQuery("INSERT INTO users \\(email,password\\) VALUES \\(\\$1,\\$2\\) RETURNING id").
		WithArgs(
			dummyUser.Email,
			dummyUser.Password,
		).WillReturnRows(rows)

	id, err := r.SaveUser(ctx, &dummyUser)

	assert.NotNil(t, dummyUser)
	assert.Equal(t, id, dummyUser.ID) // int64(1)
	assert.NoError(t, err)
}

func TestStorage_CreateUserFail(t *testing.T) {
	t.Parallel()

	r, ctx, dbMock, dummyUser := setup(t)

	dbMock.ExpectQuery("INSERT INTO users \\(email,password\\) VALUES \\(\\$1,\\$2\\) RETURNING id").
		WithArgs(
			dummyUser.Email,
			dummyUser.Password,
		)

	_, err := r.SaveUser(ctx, &dummyUser)

	assert.Error(t, err)
}

func TestStorage_IsAdmin(t *testing.T) {
	t.Parallel()

	r, ctx, dbMock, dummyUser := setup(t)

	rows := sqlmock.NewRows([]string{"is_admin"}).AddRow(true)
	dummyUser.ID = 100
	dummyUser.IsAdmin = true

	dbMock.ExpectQuery("SELECT is_admin FROM users WHERE id = \\$1").
		WithArgs(
			dummyUser.ID,
		).WillReturnRows(rows)

	isAdmin, err := r.IsAdmin(ctx, dummyUser.ID)

	assert.NoError(t, err)
	assert.True(t, isAdmin)
}

func TestStorage_IsAdminFail(t *testing.T) {
	t.Parallel()

	r, ctx, dbMock, dummyUser := setup(t)

	dbMock.ExpectQuery("SELECT is_admin FROM users WHERE id = \\$1").
		WithArgs(
			dummyUser.ID,
		)

	isAdmin, err := r.IsAdmin(ctx, dummyUser.ID)

	assert.Error(t, err)
	assert.False(t, isAdmin)
}

func TestStorage_IsAdminNoRows(t *testing.T) {
	t.Parallel()

	r, ctx, dbMock, dummyUser := setup(t)

	dbMock.ExpectQuery("SELECT is_admin FROM users WHERE id = \\$1").
		WithArgs(
			dummyUser.ID,
		).WillReturnError(sql.ErrNoRows)

	isAdmin, err := r.IsAdmin(ctx, dummyUser.ID)

	assert.Error(t, err)
	assert.False(t, isAdmin)
}

func TestStorage_App(t *testing.T) {
	t.Parallel()

	r, ctx, dbMock, _ := setup(t)

	dbMock.ExpectQuery("SELECT \\* FROM apps WHERE id = \\$1").
		WithArgs(
			1,
		).WillReturnError(sql.ErrNoRows)

	_, err := r.App(ctx, 1)

	assert.Error(t, err)
}
