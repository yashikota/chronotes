package utils

import (
	"errors"
	"net/http"
)

func GetQueryParam(r *http.Request, key string, required bool) (string, error) {
	value := r.URL.Query().Get(key)
	if required && value == "" {
		return "", errors.New(key + " is required")
	}
	return value, nil
}
