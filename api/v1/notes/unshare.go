package notes

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"
)

func UnShareNoteHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get date from request
	note := model.NewNote()
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	slog.Info("note_id: " + note.NoteID)

	// UnShare Note
	err := notes.UnShareNote(note.NoteID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	res := map[string]string{"message": "success"}
	utils.SuccessJSONResponse(w, res)
}
