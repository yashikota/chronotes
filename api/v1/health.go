package handler

import (
	"net/http"

	"github.com/yashikota/58hack/pkg/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
	utils.SuccessJsonResponse(w, map[string]string{"message": "pong"})
}
