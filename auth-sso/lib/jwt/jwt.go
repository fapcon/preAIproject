package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"
)

// NewToken Создание нового JWT токена, используя данные пользователя и приложения
func NewToken(user models.User, signingMethod, secretKey string, duration time.Duration) (string, error) {
	if signingMethod == "" {
		signingMethod = "HS256"
	}
	jwtSigningMethod := jwt.GetSigningMethod(signingMethod)
	token := jwt.New(jwtSigningMethod)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetUserIDFromToken Получение ID пользователя из JWT токена, используя ключ шифрования и токен
func GetUserIDFromToken(tokenString string, secretKey string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("incorrect token format")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := fmt.Sprintf("%v", claims["uid"])
		return userID, nil
	} else {
		return "", fmt.Errorf("token is invalid")
	}
}
