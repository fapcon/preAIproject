package service

import (
	"context"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user_sub/storage"
)

const (
	UserSubNotFound = "user_sub storage: GetByID not found"
)

type UserSubService struct {
	storage storage.UserSuber
	logger  *zap.Logger
}

func NewUserSubService(storage storage.UserSuber, logger *zap.Logger) *UserSubService {
	return &UserSubService{storage: storage, logger: logger}
}

func (u *UserSubService) Add(ctx context.Context, in UserSubAddIn) UserSubOut {
	// подписаться на себя нельзя
	if in.UserID == in.SubUserID {
		return UserSubOut{
			Success:   false,
			ErrorCode: errors.UserSubError,
		}
	}

	var dto models.UserSubDTO
	var err error

	dto.SetUserID(in.UserID).
		SetSubscriberID(in.SubUserID)

	// проверяем если уже подписка
	ckeck, err := u.storage.GetByID(ctx, dto)
	if err != nil {
		if err.Error() != UserSubNotFound {
			return UserSubOut{
				Success:   false,
				ErrorCode: errors.UserSubError,
			}
		}
	}
	//подписка найдена
	if ckeck.UserID == dto.UserID {
		return UserSubOut{
			Success:   false,
			ErrorCode: errors.UserSubAlreadyExists,
		}
	}

	err = u.storage.Create(ctx, dto)
	if err != nil {
		return UserSubOut{
			Success:   false,
			ErrorCode: errors.UserSubAddError,
		}
	}

	return UserSubOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}

func (u *UserSubService) List(ctx context.Context, in UserSubListIn) UserSubListOut {
	var err error

	res, err := u.storage.GetList(ctx)
	if err != nil {
		return UserSubListOut{
			Success:   false,
			ErrorCode: errors.UserSubListError,
		}
	}

	subUsersIDs := make([]int, 0, len(res))
	for _, v := range res {
		if in.UserID == v.UserID {
			subUsersIDs = append(subUsersIDs, v.SubUserID)
		}
	}

	return UserSubListOut{
		SubUserIDs: subUsersIDs,
		Success:    true,
		ErrorCode:  errors.NoError,
	}
}

func (u *UserSubService) Delete(ctx context.Context, in UserSubDelIn) UserSubOut {
	var dto models.UserSubDTO
	var err error

	dto.SetUserID(in.UserID).
		SetSubscriberID(in.SubUserID)

	err = u.storage.Delete(ctx, dto)
	if err != nil {
		if err.Error() == UserSubNotFound {
			return UserSubOut{
				Success:   false,
				ErrorCode: errors.UserSubNotFound,
			}
		} else {
			return UserSubOut{
				Success:   false,
				ErrorCode: errors.UserSubDeleteError,
			}
		}
	}

	return UserSubOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}
