package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/url"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

func URLDecode(s string) (string, error) {
	decodedStr, err := url.QueryUnescape(s)
	if err != nil {
		log.Println("Error decoding URL:", err)
		return "", err
	}
	return decodedStr, nil
}

func Iso8601ToDateString(t string) (string, error) {
	d, err := synchro.ParseISO[tz.AsiaTokyo](t)
	if err != nil {
		log.Println("Error parsing ISO8601 date:", err)
		return "", err
	}
	return d.Format("2006-01-02"), nil
}

func CustomJSONEncoder(v interface{}) (string, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(v); err != nil {
		log.Println("Error encoding JSON:", err)
		return "", err
	}

	return buffer.String(), nil
}
