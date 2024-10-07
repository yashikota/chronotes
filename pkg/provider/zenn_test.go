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

func TestZennHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}
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
