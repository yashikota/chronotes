package morph

import (
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GetMorphHandler(w http.ResponseWriter, r *http.Request) {
	// Validate token
	user := model.NewUser()
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).UserID

	// Check if token exists
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Validation passed")

	sentence, err := utils.GetQueryParam(r, "sentence", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Parse request body passed")

	result, err := utils.GetMorph(sentence)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("morph passed")

	// Response
	utils.SuccessJSONResponse(w, result)
}
