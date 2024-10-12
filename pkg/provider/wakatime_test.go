package provider_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func TestWakatimeHandler(t *testing.T) {
	w := httptest.NewRecorder()

	apiKey := os.Getenv("WAKATIME_API_KEY")
	startDate := "2024-10-13"
	endDate := "2024-10-13"

	if apiKey == "" {
		apiKey = "WAKATIME_API_KEY"
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
