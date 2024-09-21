package notes

import (
	"log"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/db"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
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
	date := r.URL.Query().Get("date")

	// URL Decode
	date, err := utils.URLDecode(date)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Parse ISO8601 date
	date, err = utils.Iso8601ToDateString(date)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Get note from database
	n, err := note.GetNote(user.ID, date)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Get accounts from database
	accounts, err := note.GetAccounts(user.ID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Check if note exists
	if n.ID == "" {
		// Generate note
		n, err = note.GenerateNote(user.ID, date, accounts)
		if err != nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	// Response
	utils.SuccessJSONResponse(w, n)
}
