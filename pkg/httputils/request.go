package httputils

import (
	"fmt"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

func GetBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get(authorizationHeader)

	if authHeader == "" {
		return "", fmt.Errorf("empty auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", fmt.Errorf("invalid elements amount in auth header")
	}

	if headerParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid auth header signature")
	}

	if len(headerParts[1]) == 0 {
		return "", fmt.Errorf("empty bearer token in auth header")
	}

	return headerParts[1], nil
}
