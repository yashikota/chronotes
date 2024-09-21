package provider

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func ZennHandler(w http.ResponseWriter, r *http.Request) {
	userName, err := utils.GetQueryParam(r, "userName", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	if userName == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("userName is not set"))
		return
	}

	data, err := provider.ZennProvider(userName)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		utils.SuccessJSONResponse(w, data)
	}
}
