package utils

import (
	"errors"
	"net/http"
)

func ValidateMaxFileSize(r *http.Request) error {
	const maxFileSize = 10 * 1024 * 1024 // 10MB

	r.Body = http.MaxBytesReader(nil, r.Body, maxFileSize)
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		return errors.New("file size exceeds the limit")
	}

	return nil
}

func ValidateFileType(r *http.Request) (string, error) {
	file, _, err := r.FormFile("image")
	if err != nil {
		return "", errors.New("failed to read the form file")
	}
	defer file.Close()

	// Validate the file type
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		return "", errors.New("failed to read the form file")
	}

	// Check the file type. jpeg or png
	fileType := http.DetectContentType(buff)
	if fileType != "image/jpeg" && fileType != "image/png" && fileType != "image/webp" {
		return "", errors.New("invalid file type")
	}

	return fileType, nil
}
