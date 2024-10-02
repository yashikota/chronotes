package notes

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	model "github.com/yashikota/chronotes/model/v1/db"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

func GetNoteSummaryHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.User{}
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

	slog.Info("iso8601formattedFrom: ", iso8601formattedFrom)
	slog.Info("iso8601formattedTo: ", iso8601formattedTo)

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

	slog.Info("URL Decode passed")
	slog.Info("iso8601formattedFrom:", iso8601formattedFrom)
	slog.Info("iso8601formattedTo:", iso8601formattedTo)

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

	slog.Info("from: ", from.StdTime())
	slog.Info("to: ", to.StdTime())

	// Get notes from database
	notes, err := note.GetNoteContents(user.UserID, from.StdTime(), to.StdTime())
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("notes: ", notes)

	err = godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	token := os.Getenv("GEMINI_TOKEN")
	if err != nil && !os.IsNotExist(err) {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	result, err := utils.Summary(notes, token)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Response
	res := map[string]interface{}{"result": result}
	utils.SuccessJSONResponse(w, res)
}
