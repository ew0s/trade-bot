package request

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
)

const (
	authorizationHeader = "Authorization"
)

const (
	uidQueryParam = "uid"
)

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

func GetUIDFromQuery(r *http.Request, required bool) (string, error) {
	if !r.URL.Query().Has(uidQueryParam) && required {
		return "", fmt.Errorf("does not contain uid key")
	}

	uid := r.URL.Query().Get(uidQueryParam)

	if err := ValidateUID(uid); err != nil {
		return "", fmt.Errorf("validating uid: %w", err)
	}

	return uid, nil
}

func ValidateUID(uid string) error {
	_, err := uuid.FromString(uid)

	return err
}
