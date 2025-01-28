package token17

import (
	"time"

	models "cmd/main.go/internal/models/user"

	"github.com/golang-jwt/jwt"
)

func NewAccessToken(user *models.GetUserProfileResponse) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user
	claims["email"] = user.Email
	claims["username"]=user.Username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte("unknown17"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateTokens(user *models.GetUserProfileResponse) (string, error) {
	accessToken, err := NewAccessToken(user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
