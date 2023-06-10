package utils

import (
	"encoding/base64"
)

func encode64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func decode64(text string) string {
	data, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		return ""
	}
	return string(data)
}
