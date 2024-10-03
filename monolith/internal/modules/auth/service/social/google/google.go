package google

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/service/social"
)

type GoogleAuth struct {
	config *oauth2.Config
	token  *oauth2.Token
}

type UserInfo struct {
	GoogleID string `json:"sub"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func NewGoogleAuth(conf config.AppConf, redirectURL string) *GoogleAuth {
	return &GoogleAuth{
		config: &oauth2.Config{
			ClientID:     conf.Google.ClientID,
			ClientSecret: conf.Google.ClientSecret,
			Endpoint:     google.Endpoint,
			RedirectURL:  "http://localhost:8080/api/1/auth/google/callback",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
	}
}

func (ga *GoogleAuth) Authenticate(ctx context.Context) (*social.SocialUser, error) {
	if ga.token == nil {
		return &social.SocialUser{}, fmt.Errorf("google auth token is required")
	}

	client := ga.config.Client(ctx, ga.token)

	// Получение информации о пользователе и запись ответа в струтуру UserInfo
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return &social.SocialUser{}, fmt.Errorf("get google user error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var userInfo UserInfo

	if err := json.Unmarshal(body, &userInfo); err != nil {
		fmt.Println(err)
	}

	return &social.SocialUser{
		SocialID: userInfo.GoogleID,
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Provider: "Google",
	}, nil
}

func (ga *GoogleAuth) GetRedirectURL() string {
	url := ga.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url
}

func (ga *GoogleAuth) InitToken(code string) error {
	var err error

	ctx := context.Background()
	token, err := ga.config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	ga.token = token
	return nil
}
