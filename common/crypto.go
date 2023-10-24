package common

import (
	"golang.org/x/crypto/bcrypt"
)

func Crypto(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), bytes)
	return string(bytes), nil
}
