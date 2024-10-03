package service

import (
	"context"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/grpc/user"
)

type UserGRPC struct {
	client user.UserServiceGRPCClient
	logger *zap.Logger
}

func NewUserGRPC(client user.UserServiceGRPCClient) *UserGRPC {
	return &UserGRPC{
		client: client,
	}
}

func (u UserGRPC) Create(ctx context.Context, in UserCreateIn) UserCreateOut {
	userOut, err := u.client.Create(ctx, &user.UserCreateIn{
		Name:           in.Name,
		Phone:          in.Phone,
		Email:          in.Email,
		Password:       in.Password,
		Role:           int32(in.Role),
		IdempotencyKey: in.IdempotencyKey,
	})

	if err != nil {
		u.logger.Error("Error create user", zap.Error(err))
		return UserCreateOut{
			ErrorCode: errors.UserServiceCreateUserErr,
		}
	}

	return UserCreateOut{
		UserID:    int(userOut.UserId),
		ErrorCode: int(userOut.ErrorCode),
	}
}
func (u UserGRPC) Update(ctx context.Context, in UserUpdateIn) UserUpdateOut {
	var fields []int32
	for _, v := range in.Fields {
		fields = append(fields, int32(v))
	}

	userOut, err := u.client.Update(ctx, &user.UserUpdateIn{
		User: &user.User{
			Id:            int32(in.User.ID),
			Name:          in.User.Name,
			Phone:         in.User.Phone,
			Email:         in.User.Email,
			Password:      in.User.Password,
			Role:          int32(in.User.Role),
			Verified:      in.User.Verified,
			EmailVerified: in.User.EmailVerified,
			PhoneVerified: in.User.PhoneVerified,
		},
		Fields: fields,
	})
	if err != nil {
		u.logger.Error("Error update user", zap.Error(err))

		return UserUpdateOut{
			Success:   false,
			ErrorCode: errors.UserServiceUpdateErr,
		}
	}

	return UserUpdateOut{
		Success: userOut.Success,
	}
}

func (u UserGRPC) VerifyEmail(ctx context.Context, in UserVerifyEmailIn) UserUpdateOut {
	userOut, err := u.client.VerifyEmail(ctx, &user.UserVerifyEmailIn{UserId: int32(in.UserID)})
	if err != nil {
		u.logger.Error("Error verify email", zap.Error(err))
		return UserUpdateOut{
			Success:   false,
			ErrorCode: errors.UserServiceVerifyEmailErr,
		}
	}
	return UserUpdateOut{
		Success: userOut.Success,
	}
}

func (u UserGRPC) ChangePassword(ctx context.Context, in ChangePasswordIn) ChangePasswordOut {
	userOut, err := u.client.ChangePassword(ctx, &user.ChangePasswordIn{
		Email:              in.Email,
		OldPassword:        in.OldPassword,
		NewPassword:        in.NewPassword,
		ConfirmNewPassword: in.ConfirmNewPassword,
	})
	if err != nil {
		u.logger.Error("Error change password", zap.Error(err))
		return ChangePasswordOut{
			Success:   false,
			ErrorCode: errors.UserServiceChangePasswordErr,
		}
	}

	return ChangePasswordOut{
		Success: userOut.Success,
	}
}

func (u UserGRPC) GetByEmail(ctx context.Context, in GetByEmailIn) UserOut {
	userOut, err := u.client.GetByEmail(ctx, &user.GetByEmailIn{Email: in.Email})

	if err != nil {
		u.logger.Error("Error get by email", zap.Error(err))
		return UserOut{
			ErrorCode: errors.UserServiceGetByEmailErr,
		}
	}

	return UserOut{
		User: &models.User{
			ID:            int(userOut.User.Id),
			Name:          userOut.User.Name,
			Phone:         userOut.User.Phone,
			Email:         userOut.User.Email,
			Password:      userOut.User.Password,
			Role:          int(userOut.User.Role),
			Status:        int(userOut.User.Status),
			Verified:      userOut.User.Verified,
			EmailVerified: userOut.User.EmailVerified,
			PhoneVerified: userOut.User.PhoneVerified,
		},
	}
}

func (u UserGRPC) GetByPhone(ctx context.Context, in GetByPhoneIn) UserOut {
	userOut, err := u.client.GetByPhone(ctx, &user.GetByPhoneIn{Phone: in.Phone})
	if err != nil {
		u.logger.Error("Error get by phone", zap.Error(err))
		return UserOut{
			ErrorCode: errors.UserServiceGetByPhoneErr,
		}
	}

	return UserOut{User: &models.User{
		ID:            int(userOut.User.Id),
		Name:          userOut.User.Name,
		Phone:         userOut.User.Phone,
		Email:         userOut.User.Email,
		Password:      userOut.User.Password,
		Role:          int(userOut.User.Role),
		Status:        int(userOut.User.Status),
		Verified:      userOut.User.Verified,
		EmailVerified: userOut.User.EmailVerified,
		PhoneVerified: userOut.User.PhoneVerified,
	}}
}

func (u UserGRPC) GetByID(ctx context.Context, in GetByIDIn) UserOut {
	userOut, err := u.client.GetByID(ctx, &user.GetByIDIn{UserId: int32(in.UserID)})
	if err != nil {
		u.logger.Error("Error get by ID", zap.Error(err))
		return UserOut{
			ErrorCode: errors.UserServiceGetByIDErr,
		}
	}

	return UserOut{User: &models.User{
		ID:            int(userOut.User.Id),
		Name:          userOut.User.Name,
		Phone:         userOut.User.Phone,
		Email:         userOut.User.Email,
		Password:      userOut.User.Password,
		Role:          int(userOut.User.Role),
		Status:        int(userOut.User.Status),
		Verified:      userOut.User.Verified,
		EmailVerified: userOut.User.EmailVerified,
		PhoneVerified: userOut.User.PhoneVerified,
	}}
}

func (u UserGRPC) GetByIDs(ctx context.Context, in GetByIDsIn) UsersOut {
	var ids []int32
	var users []models.User

	for _, v := range in.UserIDs {
		ids = append(ids, int32(v))
	}

	userOut, err := u.client.GetByIDs(ctx, &user.GetByIDsIn{UserIds: ids})
	if err != nil {
		u.logger.Error("Error get by IDs", zap.Error(err))
		return UsersOut{
			ErrorCode: errors.UserServiceGetByIDsErr,
		}
	}

	for _, us := range userOut.User {
		users = append(users, models.User{
			ID:            int(us.Id),
			Name:          us.Name,
			Phone:         us.Phone,
			Email:         us.Email,
			Password:      us.Password,
			Role:          int(us.Role),
			Status:        int(us.Status),
			Verified:      us.Verified,
			EmailVerified: us.EmailVerified,
			PhoneVerified: us.PhoneVerified,
		})
	}

	return UsersOut{
		User: users,
	}
}

func (u UserGRPC) BanByID(ctx context.Context, in *user.BanByIDIn) *user.BanByIDOut {
	userOut, err := u.client.BanByID(ctx, in)
	if err != nil {
		u.logger.Error("Error ban by IDs", zap.Error(err))
		return &user.BanByIDOut{
			ErrorCode: errors.UserServiceBanError,
		}
	}

	return userOut
}

func (u UserGRPC) IsBanned(ctx context.Context, in *user.IsBannedIn) *user.IsBannedOut {
	userOut, err := u.client.IsBanned(ctx, in)
	if err != nil {
		u.logger.Error("Error while checking is banned", zap.Error(err))
		return &user.IsBannedOut{
			ErrorCode: errors.UserServiceIsBannedError,
		}
	}

	return userOut
}

func (u UserGRPC) UnbanByID(ctx context.Context, in *user.UnbanByIDIn) *user.UnbanByIDOut {
	userOut, err := u.client.UnbanByID(ctx, in)
	if err != nil {
		u.logger.Error("Error unban", zap.Error(err))
		return &user.UnbanByIDOut{
			ErrorCode: errors.UserServiceUnbanErr,
		}
	}

	return userOut
}

func (u UserGRPC) ResetPassword(ctx context.Context, in ResetPasswordIn) ResetPasswordOut {
	//TODO implement me
	panic("implement me")
}

func (u UserGRPC) SendResetCodeEmail(ctx context.Context, in SendResetCodeEmailIn) SendResetCodeEmailOut {
	//TODO implement me
	panic("implement me")
}
