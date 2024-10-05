package auth

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/admin"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	user := model.NewLogin()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	var identity model.Identity

	if user.UserID == "" {
		identity = model.Email
		// Validate email
		// Rule: Required, Email, Unique
		if err := validation.Validate(user.Email, is.Email); err != nil {
			slog.Error("email error: " + err.Error())
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
		// Check if email is already taken
		if taken, err := users.IsEmailTaken(user.Email); err != nil {
			utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
			return
		} else if !taken {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("email is not registered"))
			return
		}
		slog.Info("Email validated, email: " + user.Email)
	} else {
		identity = model.UserID
		// Validate user_id
		// Rule: Required, Min 4, Max 32
		if err := validation.Validate(user.UserID, validation.Length(4, 32)); err != nil {
			slog.Error("user_id error: " + err.Error())
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
		slog.Info("UserID validated, user_id: " + user.UserID)
	}

	// Validate password
	// Rule: Required, Min 8, Max 32
	if err := validation.Validate(user.Password, validation.Required, validation.Length(8, 32)); err != nil {
		slog.Error("password error: " + err.Error())
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Validation passed")

	// Login user
	loginUser, err := users.LoginUser(user, identity)
	if err != nil {
		slog.Error("Login failed")
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	slog.Info("Login user.UserID: " + loginUser.UserID)

	// Generate a new token
	isAdmin, err := admin.IsAdmin(loginUser.UserID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	var token string
	if isAdmin {
		slog.Info("Admin user.UserID: " + loginUser.UserID)
		token, err = utils.GenerateToken(loginUser.UserID, true)
	} else {
		slog.Info("Normal user.UserID: " + loginUser.UserID)
		token, err = utils.GenerateToken(loginUser.UserID, false)
	}
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Save the token in Redis
	key := "jwt:" + loginUser.UserID
	if err := utils.SaveToken(key, token); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Generated Token: " + token)

	// Response
	res := map[string]interface{}{"token": token}
	utils.SuccessJSONResponse(w, res)
}
