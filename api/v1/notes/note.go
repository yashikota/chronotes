package notes

import (
	"log"
	"net/http"

	modelDB "github.com/yashikota/chronotes/model/v1/db"
	modelProvider "github.com/yashikota/chronotes/model/v1/provider"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := modelDB.User{}
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	// Check if token exists
	key := "jwt:" + user.UserID
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

	log.Println("URL Decode passed")
	log.Println("date:", date)

	// Parse ISO8601 date
	dateTime, err := utils.Iso8601ToDate(date)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Parse ISO8601 date passed")
	log.Println("date:", dateTime)

	// Get note from database
	n, err := note.GetNote(user.UserID, dateTime)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Get note from database passed")

	// Get accounts from database
	// accounts, err := note.GetAccounts(user.UserID)
	// if err != nil {
	// 	utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
	// 	return
	// }

	// DEBUG
	accounts := modelProvider.Gemini{
		GitHubUserID: "yashikota",
	}

	log.Println("Get accounts from database passed")

	// Check if note exists
	if n.UserID == "" {
		log.Println("Note does not exist")
		// Generate note
		n, err = note.GenerateNote(user.UserID, date, accounts)
		if err != nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	log.Println("Generate note passed")

	// Response
	res := modelDB.NoteResponse{
		Date:    dateTime.String(),
		Title:   n.Title,
		Content: n.Content,
		Tags:    n.Tags,
	}
	utils.SuccessJSONResponseWithoutEscape(w, res)
}
