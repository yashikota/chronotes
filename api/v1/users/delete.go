package users

import (
	"log/slog"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/db"
	users "github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.User{}
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	// Check if token exists
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Validation passed")

	// Delete the user
	err := users.DeleteUser(&user)
	if err != nil {
		slog.Error("Login failed")
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	slog.Info("Delete user successful")

	// Delete token
	slog.Info("Logout user.UserID: " + user.UserID)
	if err := utils.DeleteToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Delete toke successful")

	// Response
	res := map[string]interface{}{"message": "delete user successful"}
	utils.SuccessJSONResponse(w, res)
}
