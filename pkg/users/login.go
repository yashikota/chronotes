package users

import (
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func LoginUser(u *model.Login, identity model.Identity) (r *model.User, err error) {
	if db.DB == nil {
		return nil, errors.New("database connection is not initialized")
	}

	var result *gorm.DB
	if identity == model.Email {
		result = db.DB.Where("email = ?", u.Email).First(&r)
	} else {
		result = db.DB.Where("user_id = ?", u.UserID).First(&r)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	slog.Info("User found")

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(u.Password)); err != nil {
		return nil, errors.New("password does not match")
	}

	slog.Info("Password matched")

	return r, nil
}
