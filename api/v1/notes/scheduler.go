package notes

import (
	"net/http"
	"time"

	// "github.com/yashikota/chronotes/api/v1/provider"
	"github.com/go-co-op/gocron/v2"
	"github.com/yashikota/chronotes/model/v1"
	note "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/utils"
)

func ScheduleNote(w http.ResponseWriter, r *http.Request) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	ns, err := gocron.NewScheduler(gocron.WithLocation(jst))
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	user := model.NewUser()
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).ID

	date := r.URL.Query().Get("date")
	date, err = utils.URLDecode(date)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	accounts, err := note.GetAccounts(user.UserID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	ns.Start()

	_, err = ns.NewJob(
		gocron.CronJob("50 23 * * *", false),
		gocron.NewTask(note.GenerateNote(user.UserID, date, *accounts)),
	)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

}
