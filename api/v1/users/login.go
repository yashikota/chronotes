package users

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	model "github.com/yashikota/chronotes/model/v1/db"
	users "github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate email
	// Rule: Required, Email, Unique
	if err := validation.Validate(user.Email, validation.Required, is.Email); err != nil {
		log.Printf("email error: %+v", err)
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

	// Validate password
	// Rule: Required, Min 8, Max 32
	if err := validation.Validate(user.Password, validation.Required, validation.Length(8, 32)); err != nil {
		log.Printf("password error: %+v", err)
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Validation passed")

	// Login user
	err := users.LoginUser(&user)
	if err != nil {
		log.Println("Login failed")
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	// Check if user is already logged in
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("already logged in"))
		return
	}

	// Generate a new token
	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Save the token in Redis
	log.Println("Login user.ID: ", user.UserID)
	if err := utils.SaveToken(key, token); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Generated Token: " + token)

	// Response
	res := map[string]interface{}{"token": token}
	utils.SuccessJSONResponse(w, res)
}
