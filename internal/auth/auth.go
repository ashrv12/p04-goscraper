package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts an api key from the headers of an http request

// example Authorization: ApiKey {inser apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed key of Auth value")
	}

	return vals[1], nil
}
