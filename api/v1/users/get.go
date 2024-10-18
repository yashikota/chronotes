package users

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

type UserResponse struct {
	UserID    string          `json:"user_id"`
	UserName  string          `json:"user_name"`
	Email     string          `json:"email"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Role      int             `json:"role"`
	Accounts  *model.Accounts `json:"accounts"`
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get the user
	user, err := users.GetUser(req)
	if err != nil {
		slog.Warn("Login failed")
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	slog.Info("Get user successful")

	// Get user accounts
	accounts, err := users.GetAccounts(req.UserID)
	if err != nil {
		slog.Warn("Get accounts failed")
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Get accounts successful, accounts: ", slog.Any("%v", accounts))

	userResponse := UserResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role:      int(user.Role),
		Accounts:  accounts,
	}

	// Response
	utils.SuccessJSONResponse(w, userResponse)
}
