package service

import (
	"context"
	"fmt"

	"github.com/ew0s/trade-bot/internal/api/mapper"
	"github.com/ew0s/trade-bot/internal/api/request"
	"github.com/ew0s/trade-bot/internal/api/response"
	"github.com/ew0s/trade-bot/internal/domain/entities"
	"github.com/ew0s/trade-bot/pkg/constant"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, user entities.User) (string, error)
	GetUserByUID(ctx context.Context, username string) (entities.User, bool, error)
}

type AuthIdentityRepo interface {
	SetAccessToken(ctx context.Context, userUID string, tokenDetails entities.TokenDetails) error
	RemoveAccessToken(ctx context.Context, token string) error
}

type JWTAuthService interface {
	GenerateToken(userUID string) (string, entities.TokenDetails, error)
	ExtractTokenMetadata(token string) (entities.TokenDetails, error)
}

type auth struct {
	repo         AuthRepo
	identityRepo AuthIdentityRepo
	jwtService   JWTAuthService
	mapper       mapper.Auth
}

func NewAuth(repo AuthRepo, identityRepo AuthIdentityRepo, jwtService JWTAuthService) *auth {
	return &auth{
		repo:         repo,
		identityRepo: identityRepo,
		jwtService:   jwtService,
		mapper:       mapper.Auth{},
	}
}

func (s *auth) SignUp(ctx context.Context, req request.SignUp) (response.SignUp, error) {
	user := s.mapper.MakeUser(req)

	uid, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return response.SignUp{}, fmt.Errorf("creating user: %w", err)
	}

	return s.mapper.MakeSignUpResponse(uid), nil
}

func (s *auth) SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error) {
	user, found, err := s.repo.GetUserByUID(ctx, req.UID)
	if err != nil {
		return response.SignIn{}, fmt.Errorf("getting user by uid: %w", err)
	}

	if !found {
		return response.SignIn{}, fmt.Errorf("getting user by uid: %w", constant.ErrNotFound)
	}

	if !user.ValidPassword(req.Password) {
		return response.SignIn{}, fmt.Errorf("validating password: invalid password: %w", constant.ErrBadRequest)
	}

	token, tokenDetails, err := s.jwtService.GenerateToken(req.UID)
	if err != nil {
		return response.SignIn{}, fmt.Errorf("generating jwt token: %w", err)
	}

	if err = s.identityRepo.SetAccessToken(ctx, user.UID, tokenDetails); err != nil {
		return response.SignIn{}, fmt.Errorf("setting access token")
	}

	return s.mapper.MakeSignInResponse(token), nil
}

func (s *auth) Logout(ctx context.Context, req request.Logout) error {
	tokenDetails, err := s.jwtService.ExtractTokenMetadata(req.AccessToken)
	if err != nil {
		return fmt.Errorf("extracting token metadata: %w", err)
	}

	if err = s.identityRepo.RemoveAccessToken(ctx, tokenDetails.AccessUUID); err != nil {
		return fmt.Errorf("removing access token: %w", err)
	}

	return nil
}
