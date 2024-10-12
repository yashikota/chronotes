package users

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	req := model.NewUser()
	req.UserID = r.Context().Value(utils.TokenKey).(utils.Token).UserID

	// Check if token exists
	key := "jwt:" + req.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Validation passed")

	// Parse request
	user := model.NewUser()
	user.UserID = req.UserID
	user.Accounts.UserID = req.UserID

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Parsed request: ", slog.Any("%v", user))

	// Update user
	err = users.UpdateUser(user)
	if err != nil {
		slog.Error("Update user failed")
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Update accounts
	err = users.UpdateAccounts(&user.Accounts)
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
