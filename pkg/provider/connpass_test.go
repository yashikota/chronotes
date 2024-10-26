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

func TestConnpassHandler(t *testing.T) {
	w := httptest.NewRecorder()

	userID := os.Getenv("CONNPASS_USER_ID")

	if userID == "" {
		userID = "taueikumi"
	}

	if userID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("CONNPASS_USER_ID is not set"))
		return
	}

	summaries, err := provider.ConnpassProvider(userID)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if summaries == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}

	fmt.Println(summaries)

	utils.SuccessJSONResponse(w, summaries)
}