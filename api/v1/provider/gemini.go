package provider

import (
	"encoding/json"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

// Response is the structure for the API response.
type Response struct {
	Result string `json:"result"`
	Day    string `json:"day"`
}

// GeminiHandler handles the Gemini API requests.
func GeminiHandler(w http.ResponseWriter, r *http.Request) {
	var input model.Gemini
	// リクエストボディからデータをデコード
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Gemini関数を呼び出し
	response, err := provider.Gemini(input)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// 成功時のレスポンス
	apiResponse := Response{
		Result: response.Result,
		Day:    response.Day,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
}
