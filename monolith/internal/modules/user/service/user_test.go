package service

import (
	"context"
	"errors"
	"os"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	interr "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	mockuser "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/storage/mocks"
)

const (
	name     = "John"
	phone    = "+0-000-000-00-00"
	email    = "test@test.test"
	password = "test123"
	role     = 0
)

type UserServiceSuite struct {
	suite.Suite
	*require.Assertions
	ctrl            *gomock.Controller
	mockUserStorage *mockuser.MockUserer
	service         *UserService
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

func (u *UserServiceSuite) SetupTest() {
	u.Assertions = require.New(u.T())

	u.ctrl = gomock.NewController(u.T())
	u.mockUserStorage = mockuser.NewMockUserer(u.ctrl)

	logger := logs.NewLogger(config.NewAppConf(), os.Stdout)
	components := &component.Components{
		Notify: nil,
		Hash:   nil,
		Conf:   config.NewAppConf(),
	}

	u.service = NewUserService(u.mockUserStorage, logger, components)
}

func (u *UserServiceSuite) TearDownTest() {
	u.ctrl.Finish()
}

func (u *UserServiceSuite) TestUserService_Create() {
	var dto models.UserDTO
	ctx := context.Background()
	dto.SetName(name).
		SetPhone(phone).
		SetEmail(email).
		SetPassword(password).
		SetRole(role)

	u.mockUserStorage.EXPECT().Create(ctx, dto).Return(1, nil).Times(1)

	expectedOut := UserCreateOut{
		UserID: 1,
	}

	result := u.service.Create(ctx, UserCreateIn{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
		Role:     role,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_CreateNegative() {
	ctx := context.Background()
	err := errors.New("create error")
	u.mockUserStorage.EXPECT().Create(ctx, gomock.Any()).Return(0, err).Times(1)

	expectedOut := UserCreateOut{
		ErrorCode: interr.UserServiceCreateUserErr,
	}

	result := u.service.Create(ctx, UserCreateIn{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
		Role:     role,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_CreateNegativePq() {
	ctx := context.Background()
	err := &pq.Error{
		Code: "23505",
	}
	u.mockUserStorage.EXPECT().Create(ctx, gomock.Any()).Return(0, err).Times(1)

	expectedOut := UserCreateOut{
		ErrorCode: interr.UserServiceUserAlreadyExists,
	}

	result := u.service.Create(ctx, UserCreateIn{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
		Role:     role,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_VerifyEmail() {
	var dto models.UserDTO
	ctx := context.Background()
	dto.SetName(name).
		SetPhone(phone).
		SetEmail(email).
		SetPassword(password).
		SetRole(role).
		SetID(1)

	u.mockUserStorage.EXPECT().GetByID(ctx, 1).Return(dto, nil).Times(1)
	dto.SetEmailVerified(true)
	u.mockUserStorage.EXPECT().Update(ctx, dto).Return(nil).Times(1)

	expectedOut := UserUpdateOut{
		Success:   true,
		ErrorCode: 0,
	}

	result := u.service.VerifyEmail(ctx, UserVerifyEmailIn{
		UserID: 1,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_VerifyEmailNegativeGetByID() {
	ctx := context.Background()

	err := errors.New("verify email err")
	u.mockUserStorage.EXPECT().GetByID(ctx, 1).Return(models.UserDTO{}, err).Times(1)

	expectedOut := UserUpdateOut{
		Success:   false,
		ErrorCode: interr.UserServiceRetrieveUserErr,
	}

	result := u.service.VerifyEmail(ctx, UserVerifyEmailIn{
		UserID: 1,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_VerifyEmailNegativeUpdate() {
	var dto models.UserDTO
	ctx := context.Background()
	dto.SetName(name).
		SetPhone(phone).
		SetEmail(email).
		SetPassword(password).
		SetRole(role).
		SetID(1)

	u.mockUserStorage.EXPECT().GetByID(ctx, 1).Return(dto, nil).Times(1)
	err := errors.New("update err")
	dto.SetEmailVerified(true)
	u.mockUserStorage.EXPECT().Update(ctx, dto).Return(err).Times(1)

	expectedOut := UserUpdateOut{
		Success:   false,
		ErrorCode: interr.UserServiceUpdateErr,
	}

	result := u.service.VerifyEmail(ctx, UserVerifyEmailIn{
		UserID: 1,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_GetByEmail() {
	var dto models.UserDTO
	ctx := context.Background()
	dto.SetName(name).
		SetPhone(phone).
		SetEmail(email).
		SetPassword(password).
		SetRole(role).
		SetID(1)

	u.mockUserStorage.EXPECT().GetByEmail(ctx, email).Return(dto, nil).Times(1)

	expectedOut := UserOut{
		User: &models.User{
			ID:       1,
			Name:     name,
			Phone:    phone,
			Email:    email,
			Password: password,
			Role:     role,
		},
	}

	result := u.service.GetByEmail(ctx, GetByEmailIn{
		Email: email,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_GetByEmailNegative() {
	ctx := context.Background()

	err := errors.New("get by email err")
	u.mockUserStorage.EXPECT().GetByEmail(ctx, email).Return(models.UserDTO{}, err).Times(1)

	expectedOut := UserOut{
		ErrorCode: interr.UserServiceRetrieveUserErr,
	}

	result := u.service.GetByEmail(ctx, GetByEmailIn{
		Email: email,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_GetByID() {
	var dto models.UserDTO
	ctx := context.Background()
	dto.SetName(name).
		SetPhone(phone).
		SetEmail(email).
		SetPassword(password).
		SetRole(role).
		SetID(1)

	u.mockUserStorage.EXPECT().GetByID(ctx, 1).Return(dto, nil).Times(1)

	expectedOut := UserOut{
		User: &models.User{
			ID:       1,
			Name:     name,
			Phone:    phone,
			Email:    email,
			Password: password,
			Role:     role,
		},
	}

	result := u.service.GetByID(ctx, GetByIDIn{
		UserID: 1,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_GetByIDNegative() {
	ctx := context.Background()

	err := errors.New("get by ID err")
	u.mockUserStorage.EXPECT().GetByID(ctx, 1).Return(models.UserDTO{}, err).Times(1)

	expectedOut := UserOut{
		ErrorCode: interr.UserServiceRetrieveUserErr,
	}

	result := u.service.GetByID(ctx, GetByIDIn{
		UserID: 1,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_ChangePassword_ValidInput() {
	var userDTO models.UserDTO
	userDTO.SetID(1).
		SetEmail(email).
		SetName(name).
		SetPhone(phone).
		SetPassword(password)

	hashedPassword, err := cryptography.HashPassword(password)
	if err != nil {
	}
	userDTO.SetPassword(hashedPassword)

	ctx := context.Background()

	newPassword := "new_password"
	confirmNewPassword := "new_password"

	u.mockUserStorage.EXPECT().GetByEmail(ctx, email).Return(userDTO, nil).Times(1)

	u.mockUserStorage.EXPECT().Update(ctx, gomock.Any()).Return(nil).Times(1)

	expectedOut := ChangePasswordOut{
		Success: true,
	}

	result := u.service.ChangePassword(ctx, ChangePasswordIn{
		Email:              email,
		OldPassword:        "test123", // Используйте фактический пароль
		NewPassword:        newPassword,
		ConfirmNewPassword: confirmNewPassword,
	})

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_ChangePassword_Incorrect_Old_Password() {
	var userDTO models.UserDTO
	userDTO.SetID(1).
		SetEmail(email).
		SetName(name).
		SetPhone(phone).
		SetPassword(password)

	hashedPassword, err := cryptography.HashPassword(password)
	if err != nil {
	}
	userDTO.SetPassword(hashedPassword)

	ctx := context.Background()

	newPassword := "new_password"
	confirmNewPassword := "new_password"

	u.mockUserStorage.EXPECT().GetByEmail(ctx, email).Return(userDTO, nil).Times(1)

	result := u.service.ChangePassword(ctx, ChangePasswordIn{
		Email:              email,
		OldPassword:        "wrong_password",
		NewPassword:        newPassword,
		ConfirmNewPassword: confirmNewPassword,
	})

	expectedOut := ChangePasswordOut{
		Success:   false,
		ErrorCode: interr.UserServiceIncorrectOldPassword,
	}

	u.Equal(expectedOut, result)
}

func (u *UserServiceSuite) TestUserService_ChangePassword_Password_Mismatch() {
	var userDTO models.UserDTO
	userDTO.SetID(1).
		SetName(email).
		SetName(name).
		SetPhone(phone).
		SetPassword(password)

	hashedPassword, err := cryptography.HashPassword(password)
	if err != nil {
	}
	userDTO.SetPassword(hashedPassword)

	ctx := context.Background()

	u.mockUserStorage.EXPECT().GetByEmail(ctx, email).Return(userDTO, nil).Times(1)

	result := u.service.ChangePassword(ctx, ChangePasswordIn{
		Email:              email,
		OldPassword:        password,
		NewPassword:        "new_password",
		ConfirmNewPassword: "different_password",
	})

	expectedOut := ChangePasswordOut{
		Success:   false,
		ErrorCode: interr.UserServicePasswordMismatch,
	}

	u.Equal(expectedOut, result)
}
