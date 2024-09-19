package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func DiscordHandler(w http.ResponseWriter, r *http.Request) {
	channelID, err := utils.GetQueryParam(r, "channelID", true)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	if channelID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("channelID is not set"))
		return
	}

	data, err := provider.DiscordProvider(channelID)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		utils.SuccessJSONResponse(w, data)
	}
}
