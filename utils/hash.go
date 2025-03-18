package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	// if err != nil {
	// 	panic(err)
	// }
	return string(bytes), err
}
