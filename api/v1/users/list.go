package users

import (
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GetUsersListHandler(w http.ResponseWriter, r *http.Request) {
	// Get the users
	users, err := users.GetUsersList()
	if err != nil {
		slog.Warn("Fetch users failed")
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}
	slog.Info("Get users successful")

	// Response
	utils.SuccessJSONResponse(w, users)
}
