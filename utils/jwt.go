package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "secret"

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Use []byte(secret) instead of a string
	return token.SignedString([]byte(secret))
}
