package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(util.TokenKey).(model.Token).ID

	roomID, err := util.GetUserValue(userID)
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}
	if roomID == "" {
		util.ErrorJsonResponse(w, http.StatusBadRequest, errors.New("you are not join the room"))
		return
	}

	err = r.ParseMultipartForm(1024)
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusBadRequest, err)
		return
	}

	// Create directory
	originalPhotoUploadDir := filepath.Join("uploads", roomID, "original")
	err = util.MakeDir(originalPhotoUploadDir)
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	convertedPhotoUploadDir := filepath.Join("uploads", roomID, "converted")
	err = util.MakeDir(convertedPhotoUploadDir)
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Validate Max File Size
	err = util.ValidateMaxFileSize(r)
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate File Type
	_, err = util.ValidateFileType(r)
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusBadRequest, err)
		return
	}

	// multipart.File to bytes.Buffer
	img, err := util.ReadFormFile(r)
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusBadRequest, err)
		return
	}

	// Save the original photo
	originalFileName, err := util.SaveFile(img, originalPhotoUploadDir, ".jpg")
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusBadRequest, err)
		return
	}
	originalFilePath := fmt.Sprintf("%s/%s", originalPhotoUploadDir, originalFileName)

	// Resize the photo
	resizedImg := util.Resize(img, 720)
	convertedFileName, err := util.SaveFile(resizedImg, convertedPhotoUploadDir, ".jpg")
	if err != nil {
		util.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}
	convertedFilePath := fmt.Sprintf("%s/%s", convertedPhotoUploadDir, convertedFileName)

	uploadLog := map[string]string{
		"original":  originalFilePath,
		"converted": convertedFilePath,
	}
	log.Println(uploadLog)I

	res := map[string]string{"message": "photo uploaded"}
	util.SuccessJsonResponse(w, res)
}