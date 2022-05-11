package service

import (
	"fmt"

	"github.com/ew0s/trade-bot/internal/domain/entities"
)

type JWTIdentityService interface {
	ExtractTokenMetadata(token string) (entities.TokenDetails, error)
}

type UserIdentity struct {
	service JWTIdentityService
}

func NewUserIdentity(service JWTIdentityService) *UserIdentity {
	return &UserIdentity{
		service: service,
	}
}

func (s UserIdentity) GetUserUID(bearerToken string) (string, error) {
	td, err := s.service.ExtractTokenMetadata(bearerToken)
	if err != nil {
		return "", fmt.Errorf("extracting token metadata: %w", err)
	}

	return td.UserUID, nil
}
