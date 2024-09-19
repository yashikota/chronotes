package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yashikota/chronotes/pkg/provider"
)

func SlackProvider(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("channelID")
	if channelID == "" {
		http.Error(w, "channelID is required", http.StatusBadRequest)
		return
	}
	data, err := provider.GitHubProvider(channelID)
	if err != nil {
		http.Error(w, "Failed to fetch data from GitHub", http.StatusInternalServerError)
		log.Println("Error fetching data:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("error encoding response", err)
	}
}
