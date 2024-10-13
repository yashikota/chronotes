package provider

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func WakatimeHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := utils.GetQueryParam(r, "apiKey", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	startDate, err := utils.GetQueryParam(r, "startDate", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	
	endDate, err := utils.GetQueryParam(r, "endDate", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	if apiKey == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("apiKey is not set"))
		return
	}

	if startDate == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("startDate is not set"))
		return
	}

	if endDate == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("endDate is not set"))
		return
	}

	data, err := provider.WakatimeProvider(apiKey, startDate, endDate)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		utils.SuccessJSONResponse(w, data)
	}
}
