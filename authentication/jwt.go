package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/todo_application/model"
)

var secretKey = []byte("secret")

func GenerateTokenAndRefreshToken(emp model.Employee) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": emp.UserName,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken.Claims = jwt.MapClaims{
		"username": emp.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}
	return "", fmt.Errorf("invalid token")
}
