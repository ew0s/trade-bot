package entities

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type TokenDetails struct {
	UserUID    string
	AccessUUID string
	AtExpires  int64
}

func NewTokenDetails(userUID string, d time.Duration) (TokenDetails, error) {
	var td TokenDetails

	td.UserUID = userUID

	td.AtExpires = time.Now().Add(d).Unix()

	accessUUID, err := uuid.NewV4()
	if err != nil {
		return TokenDetails{}, fmt.Errorf("generating uuid v4: %w", err)
	}

	td.AccessUUID = accessUUID.String()

	return td, nil
}
