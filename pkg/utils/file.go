package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/oklog/ulid/v2"
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
	filename := ulid.Make().String() + strings.Split(extension, "/")[1]

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

func ReadFormFile(r *http.Request) ([]byte, error) {
	file, _, err := r.FormFile("image")
	if err != nil {
		return nil, errors.New("failed to read the form file")
	}
	defer file.Close()

	// Read the form file
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, errors.New("failed to read the form file")
	}

	return buf.Bytes(), nil
}
