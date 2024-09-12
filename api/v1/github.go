package handler

import (
	"net/http"

	"github.com/yashikota/58hack/pkg/utils"
)

func GithubHandler(w http.ResponseWriter, r *http.Request) {
	utils.SuccessJSONResponse(w, map[string]string{"message": "pong"})
}
