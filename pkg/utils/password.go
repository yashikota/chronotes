package utils

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/yashikota/chronotes/model/v1"
)

func GeneratePassword(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func ComparePassword(r, u string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(r), []byte(u)); err != nil {
		return err
	}
	return nil
}
