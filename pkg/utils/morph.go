package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/yashikota/chronotes/model/v1"
)

func GetMorph(sentence string) (model.MorphResponse, error) {
	appID := os.Getenv("GOO_LAB_TOKEN")
	if appID == "" {
		return model.MorphResponse{}, errors.New("GOO_LAB_TOKEN is required")
	}

	req := model.MorphRequest{
		AppID:    appID,
		Sentence: sentence,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		slog.Error("Error json marshal")
		return model.MorphResponse{}, err
	}

	url := "https://labs.goo.ne.jp/api/morph"
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return model.MorphResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.MorphResponse{}, err
	}

	var morph model.MorphResponse
	if err := json.Unmarshal(body, &morph); err != nil {
		slog.Error("json unmarshal is error")
	}

	fmt.Println("morph", morph)

	return morph, nil
}
