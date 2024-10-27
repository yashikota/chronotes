package images

import (
	"net/http"
	"path/filepath"

	"github.com/yashikota/chronotes/pkg/minio"
	"github.com/yashikota/chronotes/pkg/utils"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(utils.TokenKey).(utils.Token).UserID

	uid := utils.GenerateULID()

	err := r.ParseMultipartForm(1024)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
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
	file, filename, fileSize, err := utils.ReadFormFile(r)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	filePath := filepath.Join(userID, uid+"_"+filename)

	// Save the original photo
	uploadInfo, err := minio.SaveObject(file, filePath, fileSize, mimeType)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// Get SignedURL
	signedURL, err := minio.GetObjectURL(uploadInfo.Key)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Resize the photo
	// file, fileSize = utils.Resize(file.Bytes(), 720)
	// _, err = minio.SaveFile(file, "c_" + filePath, fileSize, mimeType)
	// if err != nil {
	// 	utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
	// 	return
	// }

	res := map[string]string{
		// "object_name": uploadInfo.Key,
		"url": signedURL,
	}
	utils.SuccessJSONResponse(w, res)
}
