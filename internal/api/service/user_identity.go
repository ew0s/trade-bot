package service

import (
	"context"
	"fmt"

	"github.com/ew0s/trade-bot/internal/domain/entities"
)

type IdentityService interface {
	ExtractTokenMetadata(token string) (entities.TokenDetails, error)
}

type IdentityRepo interface {
	AccessUIDExists(ctx context.Context, accessUID string) (bool, error)
}

type userIdentity struct {
	service IdentityService
	repo    IdentityRepo
}

func NewUserIdentity(service IdentityService, repo IdentityRepo) *userIdentity {
	return &userIdentity{
		service: service,
		repo:    repo,
	}
}

func (s userIdentity) ValidAccessToken(ctx context.Context, accessToken string) (bool, error) {
	td, err := s.service.ExtractTokenMetadata(accessToken)
	if err != nil {
		return false, fmt.Errorf("extracting token metadata: %w", err)
	}

	if td.Expired() {
		return false, nil
	}

	found, err := s.repo.AccessUIDExists(ctx, td.AccessUUID)
	if err != nil {
		return false, fmt.Errorf("checking access uid exists: %w", err)
	}

	return found, nil
}
