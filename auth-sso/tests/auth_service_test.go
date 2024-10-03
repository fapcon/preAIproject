package tests

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/pb/sso"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/tests/suite"
)

const (
	emptyAppID = 0
	appID      = 1
	appSecret  = "test-secret"
	wrongAppID = 2

	passDefaultLen = 10
)

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}

// TestRegisterLogin_Login_HappyPath тестирование регистрации и входа
func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	email := gofakeit.Email()
	pass := randomFakePassword()

	respReg, err := st.AuthClient.Register(
		ctx, &sso.RegisterRequest{
			Email:    email,
			Password: pass,
		},
	)
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respLogin, err := st.AuthClient.Login(
		ctx, &sso.LoginRequest{
			Email:    email,
			Password: pass,
			AppId:    appID,
		},
	)
	require.NoError(t, err)
	token := respLogin.GetToken()
	assert.NotEmpty(t, token)

	loginTime := time.Now()

	tokenParsed, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return []byte(appSecret), nil
		},
	)
	require.NoError(t, err)

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	require.True(t, ok)

	assert.Equal(t, respReg.GetUserId(), int64(claims["uid"].(float64)))
	assert.Equal(t, email, claims["email"].(string))
	assert.Equal(t, appID, int(claims["app_id"].(float64)))

	const deltaSeconds = 1

	// check if exp of token is in correct range, ttl gets from cfg
	assert.InDelta(t, loginTime.Add(st.Cfg.TokenTTL).Unix(), claims["exp"].(float64), deltaSeconds)
}

// TestRegisterLogin_DuplicatedRegistration проверяет регистрацию пользователя с уже существующей почтой
func TestRegisterLogin_DuplicatedRegistration(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	email := gofakeit.Email()
	pass := randomFakePassword()

	respReg, err := st.AuthClient.Register(
		ctx, &sso.RegisterRequest{
			Email:    email,
			Password: pass,
		},
	)
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respReg, err = st.AuthClient.Register(
		ctx, &sso.RegisterRequest{
			Email:    email,
			Password: pass,
		},
	)
	require.Error(t, err)
	assert.Emptyf(t, respReg.GetUserId(), "expected empty user id on duplicated registration")
	assert.ErrorContains(t, err, "user already exists")
}

// TestRegister_FailCases проверяет некорректные входные данные
func TestRegister_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	testCases := []struct {
		name        string
		email       string
		password    string
		expectedErr string
	}{
		{
			name:        "empty email",
			email:       "",
			password:    randomFakePassword(),
			expectedErr: "empty email",
		},
		{
			name:        "empty password",
			email:       gofakeit.Email(),
			password:    "",
			expectedErr: "empty password",
		},
		{
			name:        "empty pass and email",
			email:       "",
			password:    "",
			expectedErr: "empty email",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(
			tc.name, func(t *testing.T) {

				_, err := st.AuthClient.Register(
					ctx, &sso.RegisterRequest{
						Email:    tc.email,
						Password: tc.password,
					},
				)
				require.Errorf(t, err, tc.expectedErr)
				assert.ErrorContains(t, err, tc.expectedErr)
			},
		)
	}
}

// TestLogin_FailCases проверяет некорректные входные данные при входе
func TestLogin_FailCases(t *testing.T) {
	ctx, st := suite.NewSuite(t)

	testCases := []struct {
		name        string
		email       string
		password    string
		appID       int32
		expectedErr string
	}{
		{
			name:        "Login with empty email",
			email:       "",
			password:    randomFakePassword(),
			appID:       appID,
			expectedErr: "empty email",
		},
		{
			name:        "Login with empty password",
			email:       gofakeit.Email(),
			password:    "",
			appID:       appID,
			expectedErr: "empty password",
		},
		{
			name:        "Login with empty email and password",
			email:       "",
			password:    "",
			appID:       appID,
			expectedErr: "empty email",
		},
		{
			name:        "Login with empty app id",
			email:       gofakeit.Email(),
			password:    randomFakePassword(),
			appID:       emptyAppID,
			expectedErr: "empty app id",
		},
		{
			name:        "Login with invalid app id",
			email:       gofakeit.Email(),
			password:    randomFakePassword(),
			appID:       wrongAppID,
			expectedErr: "invalid email or password",
		},
		{
			name:        "Login with wrong password",
			email:       gofakeit.Email(),
			password:    "wrong",
			appID:       appID,
			expectedErr: "invalid email or password",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(
			tc.name, func(t *testing.T) {

				_, err := st.AuthClient.Login(
					ctx, &sso.LoginRequest{
						Email:    tc.email,
						Password: tc.password,
						AppId:    tc.appID,
					},
				)
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
			},
		)
	}
}
