package notes

import (
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GetSharedNoteHandler(w http.ResponseWriter, r *http.Request) {
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
	shareURL, err := utils.GetQueryParam(r, "share_url", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Get sharedNote from database
	sharedNote, err := note.GetNoteByNoteShareURL(shareURL)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("note: ", slog.Any("%v", sharedNote))

	// Response
	res := map[string]interface{}{"shared_note": sharedNote}
	utils.SuccessJSONResponse(w, res)
}
