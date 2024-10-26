package provider

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func ConnpassHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetQueryParam(r, "channelID", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	if userID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("connpass userID is not set"))
		return
	}
	
	data, err := provider.ConnpassProvider(userID)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		utils.SuccessJSONResponse(w, data)
	}
}
