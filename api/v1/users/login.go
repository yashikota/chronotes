package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
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

	// Validate password
	// Rule: Required, Min 8, Max 32
	if err := validation.Validate(user.Password, validation.Required, validation.Length(8, 32)); err != nil {
		log.Printf("password error: %+v", err)
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Validation passed")

	// Login user
	token, err := users.LoginUser(&user)
	if err != nil {
		log.Println("Login failed")
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	log.Println("token" + token)

	// Response
	res := map[string]interface{}{"token": token}
	utils.SuccessJSONResponse(w, res)
}
