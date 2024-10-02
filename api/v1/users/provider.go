package users

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func UpdateAccountsHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.User{}
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	// Check if token exists
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Validation passed")

	// Parse request
	var accounts model.Account
	accounts.UserID = user.UserID
	err := json.NewDecoder(r.Body).Decode(&accounts)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Parsed request: ", accounts)

	// Update accounts
	err = users.UpdateAccounts(&accounts)
	if err != nil {
		log.Println("Update accounts failed")
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Update accounts successful")

	// Response
	res := map[string]interface{}{"message": "update accounts successful"}
	utils.SuccessJSONResponse(w, res)
}
