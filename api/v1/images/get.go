package images

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/minio"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
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

	objectName, err := utils.GetQueryParam(r, "object_name", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate Object Name
	objectOwner := strings.Split(objectName, "/")
	if objectOwner[0] != user.UserID {
		utils.ErrorJSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	// Get Object URL
	url, err := minio.GetObjectURL(objectName)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Replace URL
	url = strings.Replace(url, "http://minio:9000", "https://chronotes.yashikota.com/file", 1)

	res := map[string]string{
		"url": url,
	}
	utils.SuccessJSONResponse(w, res)
}
