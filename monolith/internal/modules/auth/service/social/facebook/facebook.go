package facebook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/service/social"
)

type FacebookAuth struct {
	config *oauth2.Config
	token  *oauth2.Token
}

type UserInfo struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	FacebookID string `json:"id"`
}

func NewFacebookAuth(conf config.AppConf, redirectURL string) *FacebookAuth {
	return &FacebookAuth{
		config: &oauth2.Config{
			ClientID:     conf.Facebook.ClientID,
			ClientSecret: conf.Facebook.ClientSecret,
			Endpoint:     facebook.Endpoint,
			RedirectURL:  redirectURL,
			Scopes:       []string{"public_profile", "email"},
		},
	}
}

func (f FacebookAuth) Authenticate(ctx context.Context) (*social.SocialUser, error) {
	if f.token == nil {
		return &social.SocialUser{}, errors.New("facebook auth: token is required")
	}
	//отправка запроса к graph.facebook на получение информации о пользователе
	response, err := http.Get("https://graph.facebook.com/v17.0/me?fields=id%2Cname%2Cemail&access_token=" +
		url.QueryEscape(f.token.AccessToken))
	if err != nil {
		return &social.SocialUser{}, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &social.SocialUser{}, err
	}

	//запись информации о пользователе facebook
	var userInfo UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return &social.SocialUser{}, err
	}

	return &social.SocialUser{
		SocialID: userInfo.FacebookID,
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Provider: "Facebook",
	}, nil
}

func (f FacebookAuth) GetRedirectURL() string {
	Url, err := url.Parse(f.config.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", f.config.ClientID)
	parameters.Add("scope", strings.Join(f.config.Scopes, " "))
	parameters.Add("redirect_uri", f.config.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", "login")
	Url.RawQuery = parameters.Encode()
	redirectUrl := Url.String()
	return redirectUrl
}

func (f FacebookAuth) InitToken(code string) error {
	httpClient := http.DefaultClient

	// Создаем запрос на обмен токена авторизации на токен доступа
	tokenURL := "https://graph.facebook.com/oauth/access_token" +
		"?client_id=" + f.config.ClientID +
		"&client_secret=" + f.config.ClientSecret +
		"&redirect_uri=" + f.config.RedirectURL +
		"&code=" + code

	resp, err := httpClient.Get(tokenURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Разбираем ответ и извлекаем токен доступа
	var accessTokenResp oauth2.Token
	if err := json.Unmarshal(body, &accessTokenResp); err != nil {
		fmt.Println(err)
	}

	if accessTokenResp.AccessToken != "" {
		f.token = &accessTokenResp
	} else {
		return fmt.Errorf("Access token not found in response: %v", accessTokenResp)
	}
	return nil
}
