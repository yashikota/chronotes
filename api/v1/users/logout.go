package users

import (
	"log/slog"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
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

	// Delete token
	slog.Info("Logout user.UserID: " + user.UserID)
	if err := utils.DeleteToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Logout successful")

	// Response
	res := map[string]interface{}{"message": "Logout successful"}
	utils.SuccessJSONResponse(w, res)
}
