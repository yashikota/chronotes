package notes

import (
	"errors"
	"log"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/db"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

func GetNoteListHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.User{}
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	// Check if token exists
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Validation passed")

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

	log.Println("iso8601formattedFrom: ", iso8601formattedFrom)
	log.Println("iso8601formattedTo: ", iso8601formattedTo)

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

	log.Println("URL Decode passed")
	log.Println("iso8601formattedFrom:", iso8601formattedFrom)
	log.Println("iso8601formattedTo:", iso8601formattedTo)

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

	log.Println("from: ", from.StdTime())
	log.Println("to: ", to.StdTime())

	// Get notes from database
	notes, err := note.GetNoteList(user.UserID, from.StdTime(), to.StdTime())
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("notes: ", notes)

	// Response
	res := map[string]interface{}{"notes": notes}
	utils.SuccessJSONResponse(w, res)
}
