package notes

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	n "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.NewUser()
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).UserID

	// Check if token exists
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Validation passed")

	notes := []model.Note{}
	if err := json.NewDecoder(r.Body).Decode(&notes); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Parse request body passed")

	// Get user
	userInfo, err := users.GetUser(user)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	for _, note := range notes {
		err := n.DeleteNote(note, userInfo)
		if err != nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	slog.Info("Delete notes passed")

	// Response
	res := map[string]string{"message": "delete note successful"}
	utils.SuccessJSONResponse(w, res)
}