package notes

import (
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get date from request
	date := r.URL.Query().Get("date")

	// URL Decode
	date, err := utils.URLDecode(date)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("URL Decode passed")
	slog.Info("date:" + date)

	// Parse ISO8601 date
	dateTime, err := utils.Iso8601ToDate(date)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Parse ISO8601 date passed")
	slog.Info("date:" + dateTime.String())

	// Get note from database
	n, err := note.GetNote(user.UserID, dateTime)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Get note from database passed")

	// Get accounts from database
	accounts, err := note.GetAccounts(user.UserID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Get accounts from database passed")

	// Check if note exists
	if n.UserID == "" {
		slog.Info("Note does not exist")
		// Generate note
		n, err = note.GenerateNote(user.UserID, date, *accounts)
		if err != nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	slog.Info("Generate note passed")

	// Response
	res := model.NoteResponse{
		Date:    dateTime,
		Title:   n.Title,
		Content: n.Content,
		Tags:    n.Tags,
	}
	utils.SuccessJSONResponseWithoutEscape(w, res)
}
