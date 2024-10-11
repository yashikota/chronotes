package notes

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
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
	iso8601formattedFrom, err := utils.GetQueryParam(r, "from", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	iso8601formattedTo, err := utils.GetQueryParam(r, "to", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	if iso8601formattedFrom == "" || iso8601formattedTo == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("from and to are required"))
		return
	}

	// Get fields from query parameters
	formattedFields, err := utils.GetQueryParam(r, "fields", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// URL Decode
	iso8601formattedFrom, err = utils.URLDecode(iso8601formattedFrom)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	iso8601formattedTo, err = utils.URLDecode(iso8601formattedTo)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	fields, err := utils.URLDecode(formattedFields)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("URL Decode passed")
	slog.Info("iso8601formattedFrom:" + iso8601formattedFrom)
	slog.Info("iso8601formattedTo:" + iso8601formattedTo)
	slog.Info("fields:" + fields)

	// Parse ISO8601 date
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
	if to.After(synchro.Now[tz.AsiaTokyo]()) {
		to = synchro.Now[tz.AsiaTokyo]()
	}

	// Parse fields
	fieldArray, err := utils.ParseFields(fields)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Parse fields passed")
	slog.Info("from:" + from.String())
	slog.Info("to:" + to.String())
	slog.Info("fields:" + fields)

	// Get notes from database
	notes, err := note.GetNotes(user.UserID, from.StdTime(), to.StdTime(), fieldArray)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("notes: ", slog.Any("%v", notes))

	// Response
	res := map[string]interface{}{"notes": notes}
	utils.SuccessJSONResponse(w, res)
}
