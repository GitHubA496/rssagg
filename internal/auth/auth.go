package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header found")
	}
	vals := strings.Split(val, " ")
	// fmt.Print(vals)
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("UnAuthorized user")
	}
	return vals[1], nil
}
