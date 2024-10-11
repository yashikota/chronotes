package notes

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	n "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.NewUser()
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

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

	for _, note := range notes {
		note.NoteID = utils.GenerateULID()
		if note.UserID == "" {
			note.UserID = user.UserID
		}
		note.Length = utils.GetCharacterLength(note.Content)

		err := n.CreateNote(note)
		if err != nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	slog.Info("Create notes passed")

	// Response
	res := map[string]string{"message": "create note successful"}
	utils.SuccessJSONResponse(w, res)
}
