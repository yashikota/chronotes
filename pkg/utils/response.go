package utils

import (
	"encoding/json"
	"net/http"
)

func SuccessJsonResponse(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ErrorJsonResponse(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
}
