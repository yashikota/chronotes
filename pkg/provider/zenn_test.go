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

func TestZennHandler(t *testing.T) {
	w := httptest.NewRecorder()

	ZennUsername := os.Getenv("ZENN_USERNAME")
	summaries, err := provider.ZennProvider(ZennUsername)
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
