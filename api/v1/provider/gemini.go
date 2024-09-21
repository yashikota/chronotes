package provider

import (
	"encoding/json"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

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

	utils.SuccessJSONResponse(w, response)
}
