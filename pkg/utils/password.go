package utils

import (
	"log/slog"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/yashikota/chronotes/model/v1"

	"github.com/joho/godotenv"
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

func GetAdminPassword() string {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}
	return os.Getenv("ADMIN_PASS")
}
