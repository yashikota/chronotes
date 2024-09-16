package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yashikota/chronotes/pkg/provider"
)

func GithubHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	data, err := provider.GitHubProvider(userID)
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
