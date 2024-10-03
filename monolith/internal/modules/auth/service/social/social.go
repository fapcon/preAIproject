package social

import (
	"context"
	"fmt"
)

type SocialUser struct {
	SocialID string
	Email    string
	Name     string
	Provider string
}

type AuthSocialer interface {
	AuthenticateSocialProvider(ctx context.Context, provider string, authCode string) (*SocialUser, error)
	GetRedirectURL(provider string) string
}

type AuthProvider interface {
	Authenticate(ctx context.Context) (*SocialUser, error)
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

func (ac *AuthSocial) AuthenticateSocialProvider(ctx context.Context, provider string, authCode string) (*SocialUser, error) {
	var userInfo *SocialUser
	var err error

	switch provider {
	case "facebook":
		err = ac.facebookAuth.InitToken(authCode)
		if err != nil {
			return &SocialUser{}, err
		}
		userInfo, err = ac.facebookAuth.Authenticate(ctx)
		if err != nil {
			return &SocialUser{}, err
		}
	case "google":
		err = ac.googleAuth.InitToken(authCode)
		if err != nil {
			return &SocialUser{}, err
		}
		userInfo, err = ac.googleAuth.Authenticate(ctx)
		if err != nil {
			return &SocialUser{}, err
		}
	default:
		return &SocialUser{}, fmt.Errorf("not allowed social authentication provider")
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
