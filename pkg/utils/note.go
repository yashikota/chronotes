package utils

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

func URLDecode(s string) (string, error) {
	decodedStr, err := url.QueryUnescape(s)
	if err != nil {
		log.Println("Error decoding URL:", err)
		return "", err
	}
	// TODO: Fix this hack
	decodedStr = strings.Replace(decodedStr, " ", "+", -1)
	return decodedStr, nil
}

func Iso8601ToDate(t string) (time.Time, error) {
	d, err := synchro.ParseISO[tz.AsiaTokyo](t)
	if err != nil {
		log.Println("Error parsing ISO8601 date:", err)
		return time.Time{}, err
	}
	return d.StdTime(), nil
}

func CustomJSONEncoder(v interface{}) (string, error) {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return "", err
	}

	return string(jsonBytes), nil
}
