package provider_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func TestWakatimeHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}

	apiKey := os.Getenv("WAKATIME_API_KEY")
	startDate := "2024-10-13"
	endDate := "2024-10-13"

	if apiKey == "" {
		apiKey = "WaKATIME_API_KEY"
	}

	if apiKey == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("WAKATIME_API_KEY is not set"))
		return
	}

	stats, err := provider.WakatimeProvider(apiKey, startDate, endDate)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(stats)

	utils.SuccessJSONResponse(w, stats)

}
