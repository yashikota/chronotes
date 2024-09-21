package notes

import (
	"log"
	"net/http"
	"time"

	model "github.com/yashikota/chronotes/model/v1/db"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

func GetNoteListHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.User{}
	user.ID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	// Check if token exists
	key := "jwt:" + user.ID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Validation passed")

	// Get date from request
	iso8601formattedFrom := r.PathValue("from")
	iso8601formattedTo := r.PathValue("to")
	from, err := synchro.ParseISO[tz.AsiaTokyo](iso8601formattedFrom)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	to, err := synchro.ParseISO[tz.AsiaTokyo](iso8601formattedTo)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Get notes from database
	var notes []model.Note
	for c := from; c.Before(to) || c.Equal(to); c = c.Add(24 * time.Hour) {
		date := c.Format("2006-01-02")
		note, err := note.GetNoteIgnoreContent(user.ID, date)
		if err != nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
		notes = append(notes, note)
	}

	// Response
	utils.SuccessJSONResponse(w, notes)
}
