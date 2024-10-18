package notes

import (
	"log/slog"
	"time"
    "github.com/go-co-op/gocron/v2"
	"github.com/yashikota/chronotes/pkg/users"
)

func Cron() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Error("Error loading location")
		return
	}

	ns, err := gocron.NewScheduler(gocron.WithLocation(jst))
	if err != nil {
		slog.Error("Error creating scheduler")
		return
	}

	ns.Start()

	_, err = ns.NewJob(
		gocron.CronJob("50 23 * * *", false),
		gocron.NewTask(GenerateNoteCron),
	)
	if err != nil {
		slog.Error("Error creating job")
		return
	}
}

func GenerateNoteCron(){
	var usersList []string
	today := time.Now().Format("2006-01-02")
	usersList, err := users.GetUsersList()
	if err != nil {
		slog.Error("Error getting user list")
		return
	}

	for _, user := range usersList {
		accounts, err := users.GetAccounts(user)
		if err != nil {
			slog.Error("Error getting accounts")
			continue
		}

		_, err = GenerateNote(user, today, *accounts)
		if err != nil {
			slog.Error("Error generating note")
			continue
		}
	}

}
