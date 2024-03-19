package services

import (
	"fmt"
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService interface {
	CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error)
}

var secretKey = []byte("secret-key")

func CreateToken(userId uint, expiry int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  userId,
			"exp": time.Now().Add(time.Minute * time.Duration(expiry)).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Couldn't parse claims")
		return nil, nil
	}
	id := uint(claims["id"].(float64))
	return &id, nil
}
