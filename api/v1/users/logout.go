package users

import (
	"log"
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

	log.Println("Validation passed")

	// Delete token
	log.Println("Logout user.ID: ", user.UserID)
	if err := utils.DeleteToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Logout successful")

	// Response
	res := map[string]interface{}{"message": "Logout successful"}
	utils.SuccessJSONResponse(w, res)
}
