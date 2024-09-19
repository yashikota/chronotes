package users

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/oklog/ulid/v2"

	model "github.com/yashikota/chronotes/model/v1/db"
	users "github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate username
	// Rule: Required, Min 3, Max 32, Alphanumeric
	if err := validation.Validate(user.Name, validation.Required, validation.Length(3, 32), is.Alphanumeric); err != nil {
		log.Printf("name error: %+v", err)
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
	} else if taken {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("email is already taken"))
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

	// Generate a new UserID
	user.ID = ulid.MustNew(ulid.Now(), nil).String()

	log.Println("Generated UserID: " + user.ID)

	// Create a new user
	if err := users.CreateUser(&user); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Created User: " + user.ID)

	// Generate a new token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Generated Token: " + token)

	// Response
	res := map[string]interface{}{"token": token}
	utils.SuccessJSONResponse(w, res)
}
