package utils

import (
	"errors"
	"fmt"
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

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return 0, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	fmt.Println("err:", err)

	if err != nil {
		return 0, errors.New("invalid token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	// email := claims["email"].(string)
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid token: userId is not a number")
	}
	userId := int64(userIdFloat) // Convert float64 to int64

	return userId, nil

}
