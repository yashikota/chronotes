package users

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/yashikota/chronotes/model/v1"
	users "github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	user := model.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate username
	// Rule: Required, Min 1, Max 32
	if err := validation.Validate(user.UserName, validation.Required, validation.Length(1, 32)); err != nil {
		slog.Error("name error: %+v" + err.Error())
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate email
	// Rule: Required, Email, Unique
	if err := validation.Validate(user.Email, validation.Required, is.Email); err != nil {
		slog.Error("email error: %+v" + err.Error())
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	// Check if email is already taken
	if taken, err := users.IsEmailTaken(user.Email); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	} else if taken {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("email is already taken"))
		return
	}

	// Validate password
	// Rule: Required, Min 8, Max 32
	if err := validation.Validate(user.Password, validation.Required, validation.Length(8, 32)); err != nil {
		slog.Error("password error: %+v" + err.Error())
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Validation passed")

	// Generate a new UserID
	user.UserID = utils.GenerateULID()

	slog.Info("Generated UserID: " + user.UserID)

	// Create a new user
	if err := users.CreateUser(user); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Created User: " + user.UserID)

	// Generate a new token
	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Save the token in Redis
	slog.Info("Register user.UserID: " + user.UserID)
	key := "jwt:" + user.UserID
	if err := utils.SaveToken(key, token); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Generated Token: " + token)

	// Response
	res := map[string]interface{}{"token": token}
	utils.SuccessJSONResponse(w, res)
}
