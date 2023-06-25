package utils

import (
	"encoding/base64"
)

func Encode64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Decode64(text string) string {
	data, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		return ""
	}
	return string(data)
}
