package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(userPassword, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(currentPassword))

	return err == nil
}

func CheckPasswordConfirmation(password, password_confirmation string) error {
	var err error

	if password != password_confirmation {
		err = errors.New("password and password confirmation doesn't match")
	}

	return err
}
