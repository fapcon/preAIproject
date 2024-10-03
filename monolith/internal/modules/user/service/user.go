package service

import (
	"context"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	iservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/storage"
	"time"
)

type UserService struct {
	storage storage.Userer
	logger  *zap.Logger
	conf    config.AppConf
	notify  iservice.Notifier
	hash    cryptography.Hasher
}

func NewUserService(storage storage.Userer, logger *zap.Logger, components *component.Components) *UserService {
	return &UserService{storage: storage, logger: logger, notify: components.Notify, hash: components.Hash, conf: components.Conf}
}

func (u *UserService) Create(ctx context.Context, in UserCreateIn) UserCreateOut {
	var dto models.UserDTO
	dto.SetName(in.Name).
		SetPhone(in.Phone).
		SetEmail(in.Email).
		SetPassword(in.Password).
		SetRole(in.Role)

	userID, err := u.storage.Create(ctx, dto)
	if err != nil {
		if v, ok := err.(*pq.Error); ok && v.Code == "23505" {
			return UserCreateOut{
				ErrorCode: errors.UserServiceUserAlreadyExists,
			}
		}
		return UserCreateOut{
			ErrorCode: errors.UserServiceCreateUserErr,
		}
	}

	return UserCreateOut{
		UserID: userID,
	}
}

func (u *UserService) Update(ctx context.Context, in UserUpdateIn) UserUpdateOut {
	panic("implement me")
}

func (u *UserService) ChangePassword(ctx context.Context, in ChangePasswordIn) ChangePasswordOut {

	userDTO, err := u.storage.GetByEmail(ctx, in.Email)
	if err != nil {
		u.logger.Error("user: GetByEmail err", zap.Error(err))
		return ChangePasswordOut{
			ErrorCode: errors.UserServiceRetrieveUserErr,
		}
	}

	if in.NewPassword != in.ConfirmNewPassword {
		return ChangePasswordOut{
			ErrorCode: errors.UserServicePasswordMismatch,
		}
	}

	if in.OldPassword == in.NewPassword {
		u.logger.Error("user: old password matches new password")
		return ChangePasswordOut{
			ErrorCode: errors.UserServicePasswordMismatch,
		}
	}

	if !cryptography.CheckPassword(userDTO.GetPassword(), in.OldPassword) {
		u.logger.Error("user: incorrect old password")
		return ChangePasswordOut{
			ErrorCode: errors.UserServiceIncorrectOldPassword,
		}
	}

	hashedNewPassword, err := cryptography.HashPassword(in.NewPassword)
	if err != nil {
		u.logger.Error("user: hash new password err", zap.Error(err))
		return ChangePasswordOut{
			ErrorCode: errors.HashPasswordError,
		}
	}

	userDTO.SetPassword(hashedNewPassword)

	if err := u.storage.Update(ctx, userDTO); err != nil {
		u.logger.Error("user: update user err", zap.Error(err))
		return ChangePasswordOut{
			ErrorCode: errors.UserServiceUpdateErr,
		}
	}

	return ChangePasswordOut{
		Success: true,
	}
}

func (u *UserService) VerifyEmail(ctx context.Context, in UserVerifyEmailIn) UserUpdateOut {
	dto, err := u.storage.GetByID(ctx, in.UserID)
	if err != nil {
		u.logger.Error("user: GetByEmail err", zap.Error(err))
		return UserUpdateOut{
			ErrorCode: errors.UserServiceRetrieveUserErr,
		}
	}
	dto.SetEmailVerified(true)
	err = u.storage.Update(ctx, dto)
	if err != nil {
		u.logger.Error("user: update err", zap.Error(err))
		return UserUpdateOut{
			ErrorCode: errors.UserServiceUpdateErr,
		}
	}

	return UserUpdateOut{
		Success: true,
	}
}

func (u *UserService) GetByEmail(ctx context.Context, in GetByEmailIn) UserOut {
	userDTO, err := u.storage.GetByEmail(ctx, in.Email)
	if err != nil {
		u.logger.Error("user: GetByEmail err", zap.Error(err))
		return UserOut{
			ErrorCode: errors.UserServiceRetrieveUserErr,
		}
	}

	return UserOut{
		User: &models.User{
			ID:            userDTO.GetID(),
			Name:          userDTO.GetName(),
			Phone:         userDTO.GetPhone(),
			Email:         userDTO.GetEmail(),
			Password:      userDTO.GetPassword(),
			Role:          userDTO.GetRole(),
			Verified:      userDTO.Verified,
			EmailVerified: userDTO.EmailVerified,
			PhoneVerified: userDTO.PhoneVerified,
		},
	}
}

func (u *UserService) GetByPhone(ctx context.Context, in GetByPhoneIn) UserOut {
	panic("implement me")
}

func (u *UserService) GetByID(ctx context.Context, in GetByIDIn) UserOut {
	userDTO, err := u.storage.GetByID(ctx, in.UserID)
	if err != nil {
		u.logger.Error("user: GetByEmail err", zap.Error(err))
		return UserOut{
			ErrorCode: errors.UserServiceRetrieveUserErr,
		}
	}

	return UserOut{
		User: &models.User{
			ID:            userDTO.GetID(),
			Name:          userDTO.GetName(),
			Phone:         userDTO.GetPhone(),
			Email:         userDTO.GetEmail(),
			Password:      userDTO.GetPassword(),
			Role:          userDTO.Role,
			Verified:      userDTO.Verified,
			EmailVerified: userDTO.EmailVerified,
			PhoneVerified: userDTO.PhoneVerified,
		},
	}
}

func (u *UserService) GetByIDs(ctx context.Context, in GetByIDsIn) UsersOut {
	panic("implement me")
}

func (u *UserService) ResetPassword(ctx context.Context, in ResetPasswordIn) ResetPasswordOut {
	userDTO, err := u.storage.GetByEmail(ctx, in.Email)
	if err != nil {
		u.logger.Error("user: GetByEmail err", zap.Error(err))
		return ResetPasswordOut{
			Success:   false,
			ErrorCode: errors.UserServiceRetrieveUserErr,
		}
	}

	if in.NewPassword != in.ConfirmNewPassword {
		return ResetPasswordOut{
			Success:   false,
			ErrorCode: errors.UserServicePasswordMismatch,
		}
	}

	if in.ResetCode != userDTO.ResetCode {
		return ResetPasswordOut{
			Success:   false,
			ErrorCode: errors.UserServiceInvalidResetCode,
		}
	}

	hashPass, err := cryptography.HashPassword(in.NewPassword)
	if err != nil {
		return ResetPasswordOut{
			Success:   false,
			ErrorCode: errors.HashPasswordError,
		}
	}

	userDTO.SetPassword(hashPass)

	userDTO.ResetCode = 0

	if err := u.storage.Update(ctx, userDTO); err != nil {
		u.logger.Error("user: update user err", zap.Error(err))
		return ResetPasswordOut{
			Success:   false,
			ErrorCode: errors.UserServiceUpdateErr,
		}
	}

	return ResetPasswordOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}

func (u *UserService) SendResetCodeEmail(ctx context.Context, in SendResetCodeEmailIn) SendResetCodeEmailOut {

	userDTO, err := u.storage.GetByEmail(ctx, in.Email)
	if err != nil {
		u.logger.Error("user: GetByEmail err", zap.Error(err))
		return SendResetCodeEmailOut{
			Success:   false,
			ErrorCode: errors.UserServiceRetrieveUserErr,
		}
	}

	rand.Seed(time.Now().UnixNano())
	resetCode := rand.Intn(90000) + 10000
	userDTO.ResetCode = resetCode

	if err := u.storage.Update(ctx, userDTO); err != nil {
		u.logger.Error("user: update user err", zap.Error(err))
		return SendResetCodeEmailOut{
			Success:   false,
			ErrorCode: errors.UserServiceUpdateErr,
		}
	}

	resetCodeStr := strconv.Itoa(resetCode)

	u.notifyEmail(iservice.PushIn{
		Identifier: in.Email,
		Type:       iservice.PushEmail,
		Title:      "Diamond Trade Reset Password Link",
		Data:       []byte(resetCodeStr),
		Options:    nil,
	})

	return SendResetCodeEmailOut{
		Success:   true,
		ErrorCode: errors.NoError,
	}
}

func (u *UserService) notifyEmail(p iservice.PushIn) {
	res := u.notify.Push(p)
	if res.ErrorCode != errors.NoError {
		time.Sleep(1 * time.Minute)
		go u.notifyEmail(p)
	}
}
