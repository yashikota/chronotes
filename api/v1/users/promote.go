package users

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func PromoteHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.NewUser()
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	// Check if token exists
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("JWT Validation passed")

	// Check enter password
	enter := model.NewPassword()
	if err := json.NewDecoder(r.Body).Decode(&enter); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Password entered")

	// Validate password
	adminPassword := utils.GetAdminPassword()
	if err := utils.ComparePassword(adminPassword, enter.Password); err != nil {
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	slog.Info("Password matched")

	// Promote user
	err := users.PromoteUser(user)
	if err != nil {
		slog.Error("Promote failed")
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Promote successful")

	// Delete token
	if err := utils.DeleteToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// New token
	token, err := utils.GenerateToken(user.UserID, true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Save the token in Redis
	if err := utils.SaveToken(key, token); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Generated Token: " + token)

	// Response
	res := map[string]interface{}{"new token": token}
	utils.SuccessJSONResponse(w, res)
}
