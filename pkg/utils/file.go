package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func MakeDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return errors.New("failed to create the directory: " + err.Error())
		}
	}
	return nil
}

func DeleteDir(dirName string) error {
	err := os.RemoveAll(dirName)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(data []byte, path string, extension string) (string, error) {
	if strings.Contains(extension, "/") {
		extension = strings.Split(extension, "/")[1]
	}

	filename := GenerateULID() + "." + extension

	// Save the file
	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return "", errors.New("failed to save the file: " + err.Error())
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(data))
	if err != nil {
		return "", errors.New("failed to save the file: " + err.Error())
	}

	return filename, nil
}

func ReadFormFile(r *http.Request) (*bytes.Buffer, string, int64, error) {
	file, header, err := r.FormFile("image")
	if err != nil {
		return nil, "", 0, errors.New("failed to read the form file")
	}
	defer file.Close()

	// Read the form file
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, "", 0, errors.New("failed to read the form file")
	}

	return buf, header.Filename, header.Size, nil
}
