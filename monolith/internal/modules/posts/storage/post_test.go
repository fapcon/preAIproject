package storage

import (
	"context"
	"errors"
	"testing"
	"time"

	"gitlab.com/golight/orm/utils"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSQLAdapter struct {
	mock.Mock
}

func (m *MockSQLAdapter) Create(ctx context.Context, entity utils.Tabler, opts ...interface{}) error {
	args := m.Called(ctx, entity)
	return args.Error(0)
}

func (m *MockSQLAdapter) List(ctx context.Context, dest interface{}, tableName string, condition utils.Condition, opts ...interface{}) error {
	args := m.Called(ctx, dest, tableName, condition)
	return args.Error(0)
}

func (m *MockSQLAdapter) Update(ctx context.Context, entity utils.Tabler, condition utils.Condition, operation string, opts ...interface{}) error {
	args := m.Called(ctx, entity, condition, operation)
	return args.Error(0)
}

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Get(ctx context.Context, key string, ptrValue interface{}) error {
	args := m.Called(ctx, key, ptrValue)
	return args.Error(0)
}

func (m *MockCache) Set(ctx context.Context, key string, value interface{}, expires time.Duration) {
	m.Called(ctx, key, value, expires)
}

func (m *MockCache) KeyExists(ctx context.Context, key string) (bool, error) {
	args := m.Called(ctx, key)
	return args.Bool(0), args.Error(1)
}

func (m *MockCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	args := m.Called(ctx, key, expiration)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	mockSQLAdapter := new(MockSQLAdapter)
	mockCache := new(MockCache)

	t.Run("Success", func(t *testing.T) {
		mockPost := models.PostDTO{Title: "Title", Author: 1}

		mockSQLAdapter.On("Create", context.Background(), &mockPost).Return(nil)

		storage := NewPostStorage(mockSQLAdapter, mockCache)

		err := storage.Create(context.Background(), mockPost)

		assert.NoError(t, err)
		mockSQLAdapter.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {

		mockSQLAdapter.On("Create", mock.Anything, mock.Anything).Return(errors.New("error"))

		storage := NewPostStorage(mockSQLAdapter, mockCache)

		err := storage.Create(context.Background(), models.PostDTO{})
		assert.Equal(t, err, errors.New("error"))
		mockSQLAdapter.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockSQLAdapter := new(MockSQLAdapter)
	mockCache := new(MockCache)

	t.Run("Success", func(t *testing.T) {
		mockPost := models.PostDTO{
			Title:            "Title",
			ShortDescription: "ShortDescription",
			FullDescription:  "FullDescription",
			ID:               1,
			Author:           1}

		mockSQLAdapter.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		storage := NewPostStorage(mockSQLAdapter, mockCache)

		err := storage.Update(context.Background(), mockPost)

		assert.NoError(t, err)
		mockSQLAdapter.AssertExpectations(t)
	})
	t.Run("Error", func(t *testing.T) {
		mockSQLAdapter.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))

		storage := NewPostStorage(mockSQLAdapter, mockCache)

		err := storage.Update(context.Background(), models.PostDTO{})

		assert.NoError(t, err, errors.New("error"))
		mockSQLAdapter.AssertExpectations(t)
	})
}
