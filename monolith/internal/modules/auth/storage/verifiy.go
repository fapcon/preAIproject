package storage

import (
	"context"
	"fmt"

	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
)

//go:generate easytags $GOFILE json

type SQLAdapter interface {
	Create(ctx context.Context, entity utils.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, tableName string, condition utils.Condition, opts ...interface{}) error
	Update(ctx context.Context, entity utils.Tabler, condition utils.Condition, operation string, opts ...interface{}) error
}

type EmailVerify struct {
	adapter SQLAdapter
}

func NewEmailVerify(sqlAdapter SQLAdapter) Verifier {
	return &EmailVerify{adapter: sqlAdapter}
}

func (e *EmailVerify) GetByEmail(ctx context.Context, email, hash string) (models.EmailVerifyDTO, error) {
	var emailVerifies []models.EmailVerifyDTO
	err := e.adapter.List(ctx, &emailVerifies, "email_verify", utils.Condition{
		Equal: map[string]interface{}{
			"email":    email,
			"hash":     hash,
			"verified": false,
		},
	})
	if err != nil {
		return models.EmailVerifyDTO{}, err
	}
	if len(emailVerifies) < 1 {
		return models.EmailVerifyDTO{}, fmt.Errorf("email verify %s not found", email)
	}
	return emailVerifies[0], nil
}

func (e *EmailVerify) GetByUserID(ctx context.Context, userID int) (models.EmailVerifyDTO, error) {
	var dto []models.EmailVerifyDTO
	err := e.adapter.List(ctx, &dto, "email_verify", utils.Condition{
		Equal: map[string]interface{}{
			"user_id": userID,
		},
	})
	if err != nil {
		return models.EmailVerifyDTO{}, err
	}
	if len(dto) < 1 {
		return models.EmailVerifyDTO{}, fmt.Errorf("email verify with user_id %d not found", userID)
	}
	return dto[0], nil
}

func (e *EmailVerify) Verify(ctx context.Context, userID int) error {
	dto, err := e.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	dto.SetVerified(true)
	err = e.adapter.Update(ctx, &dto, utils.Condition{
		Equal: map[string]interface{}{
			"id": dto.GetID(),
		},
	}, utils.Update)

	return err
}

func (e *EmailVerify) VerifyEmail(ctx context.Context, email, hash string) error {
	dto, err := e.GetByEmail(ctx, email, hash)
	if err != nil {
		return err
	}
	dto.SetVerified(true)
	err = e.adapter.Update(ctx, &dto, utils.Condition{
		Equal: map[string]interface{}{
			"email":    email,
			"hash":     hash,
			"verified": false,
		},
	}, utils.Update)

	return err
}

func (e *EmailVerify) Create(ctx context.Context, email, hash string, userID int) error {
	emailVerify := &models.EmailVerifyDTO{}
	emailVerify.Email = email
	emailVerify.Hash = hash
	emailVerify.UserID = userID
	err := e.adapter.Create(ctx, emailVerify)
	return err
}
