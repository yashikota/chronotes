package utils

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

func URLDecode(s string) (string, error) {
	decodedStr, err := url.QueryUnescape(s)
	if err != nil {
		slog.Error("Error decoding URL:", err)
		return "", err
	}
	// TODO: Fix this hack
	decodedStr = strings.Replace(decodedStr, " ", "+", -1)
	return decodedStr, nil
}

func Iso8601ToDate(t string) (time.Time, error) {
	d, err := synchro.ParseISO[tz.AsiaTokyo](t)
	if err != nil {
		slog.Error("Error parsing ISO8601 date:", err)
		return time.Time{}, err
	}
	return d.StdTime(), nil
}

func CustomJSONEncoder(v interface{}) (string, error) {
	var resultBytes bytes.Buffer
	enc := json.NewEncoder(&resultBytes)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	if err != nil {
		slog.Error("Error encoding JSON:", err)
		return "", err
	}
	result := strings.TrimSpace(resultBytes.String())
	return result, nil
}
