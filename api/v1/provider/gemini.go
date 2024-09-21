package provider

import (
	"encoding/json"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/provider"
)

// GeminiHandler handles the Gemini API requests.
func GeminiHandler(w http.ResponseWriter, r *http.Request) {
	var input model.Gemini

	// リクエストボディからデータをデコード
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Gemini関数を呼び出し
	result, err := provider.Gemini(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 成功時のレスポンス
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
