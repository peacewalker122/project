package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("can't hash pass, due of: %v", err)
	}
	return string(res), nil
}

func CheckPassword(pass, hashedpass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(pass))
}
