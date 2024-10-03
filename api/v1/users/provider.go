package users

import (
	"encoding/json"
	"log/slog"
	"net/http"

	db "github.com/yashikota/chronotes/model/v1/db"
	provider "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func UpdateAccountsHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := db.NewUser()
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	// Check if token exists
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Validation passed")

	// Parse request
	var accounts provider.Gemini
	accounts.UserID = user.UserID
	err := json.NewDecoder(r.Body).Decode(&accounts)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Parsed request: ", slog.Any("%v", accounts))

	// Update accounts
	err = users.UpdateAccounts(accounts)
	if err != nil {
		slog.Error("Update accounts failed")
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Update accounts successful")

	// Response
	res := map[string]interface{}{"message": "update accounts successful"}
	utils.SuccessJSONResponse(w, res)
}
