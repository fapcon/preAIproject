package social

import (
	"context"
	"fmt"
)

type User struct {
	SocialID string
	Email    string
	Name     string
	Provider string
}

type AuthSocialer interface {
	AuthenticateSocialProvider(ctx context.Context, provider string, authCode string) (*User, error)
	GetRedirectURL(provider string) string
}

type AuthProvider interface {
	Authenticate(ctx context.Context) (*User, error)
	GetRedirectURL() string
	InitToken(code string) error
}

type AuthSocial struct {
	googleAuth   AuthProvider
	facebookAuth AuthProvider
}

func NewAuthSocial(googleAuth AuthProvider, facebookAuth AuthProvider) *AuthSocial {
	return &AuthSocial{googleAuth: googleAuth,
		facebookAuth: facebookAuth}
}

func (ac *AuthSocial) AuthenticateSocialProvider(ctx context.Context, provider string, authCode string) (*User, error) {
	var userInfo *User
	var err error

	switch provider {
	case "facebook":
		err = ac.facebookAuth.InitToken(authCode)
		if err != nil {
			return &User{}, err
		}
		userInfo, err = ac.facebookAuth.Authenticate(ctx)
		if err != nil {
			return &User{}, err
		}
	case "google":
		err = ac.googleAuth.InitToken(authCode)
		if err != nil {
			return &User{}, err
		}
		userInfo, err = ac.googleAuth.Authenticate(ctx)
		if err != nil {
			return &User{}, err
		}
	default:
		return &User{}, fmt.Errorf("not allowed social authentication provider")
	}
	return userInfo, err
}

func (ac *AuthSocial) GetRedirectURL(provider string) string {
	var url string
	switch provider {
	case "google":
		url = ac.googleAuth.GetRedirectURL()
	case "facebook":
		url = ac.facebookAuth.GetRedirectURL()
	}
	return url
}
