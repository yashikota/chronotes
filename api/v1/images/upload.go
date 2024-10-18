package images

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/yashikota/chronotes/pkg/utils"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(utils.TokenKey).(utils.Token).UserID

	err := r.ParseMultipartForm(1024)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Create directory
	originalPhotoUploadDir := filepath.Join("img", userID)
	err = utils.MakeDir(originalPhotoUploadDir)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	convertedPhotoUploadDir := filepath.Join("img", userID)
	err = utils.MakeDir(convertedPhotoUploadDir)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Validate Max File Size
	err = utils.ValidateMaxFileSize(r)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate File Type
	mimeType, err := utils.ValidateFileType(r)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// multipart.File to bytes.Buffer
	img, err := utils.ReadFormFile(r)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Save the original photo
	originalFileName, err := utils.SaveFile(img, originalPhotoUploadDir, mimeType)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}
	originalFilePath := fmt.Sprintf("%s/%s", originalPhotoUploadDir, originalFileName)

	// Resize the photo
	resizedImg := utils.Resize(img, 720)
	convertedFileName, err := utils.SaveFile(resizedImg, convertedPhotoUploadDir, "png")
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	convertedFilePath := fmt.Sprintf("%s/%s", convertedPhotoUploadDir, convertedFileName)

	uploadLog := map[string]string{
		"original":  originalFilePath,
		"converted": convertedFilePath,
	}

	utils.SuccessJSONResponse(w, uploadLog)
}
