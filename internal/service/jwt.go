package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/ew0s/trade-bot/internal/domain/entities"
)

const (
	accessUUIDTokenClaim = "access_UUID"
	authorizedTokenClaim = "authorized"
	userUIDTokenClaim    = "user_uid"
	expiresTokenClaim    = "exp"
)

type jwtService struct {
	signingKey         string
	expirationDuration time.Duration
}

func NewJWTService(signingKey string, expirationDuration time.Duration) *jwtService {
	return &jwtService{
		signingKey:         signingKey,
		expirationDuration: expirationDuration,
	}
}

func (s *jwtService) GenerateToken(userUID string) (string, entities.TokenDetails, error) {
	td, err := entities.NewTokenDetails(userUID, s.expirationDuration)
	if err != nil {
		return "", entities.TokenDetails{}, fmt.Errorf("creating token details: %w", err)
	}

	token, err := newSignedToken(userUID, s.signingKey, td)
	if err != nil {
		return "", entities.TokenDetails{}, fmt.Errorf("creating new signed token: %w", err)
	}

	return token, td, nil
}

func (s *jwtService) ExtractTokenMetadata(token string) (entities.TokenDetails, error) {
	parsed, err := newParsedToken(token, s.tokenKeyFunc)
	if err != nil {
		return entities.TokenDetails{}, fmt.Errorf("creating parsed token: %w", err)
	}

	details, err := parsed.getDetails()
	if err != nil {
		return entities.TokenDetails{}, fmt.Errorf("getting details: %w", err)
	}

	return details, nil
}

func (s *jwtService) tokenKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("getting token signing method: unexpected method: %s", token.Method.Alg())
	}

	return []byte(s.signingKey), nil
}

type jwtToken jwt.Token

func newParsedToken(token string, keyFunc jwt.Keyfunc) (*jwtToken, error) {
	verified, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}

	if ok := verified.Valid; !ok {
		return nil, fmt.Errorf("invalid token passed")
	}

	return (*jwtToken)(verified), nil
}

func (t *jwtToken) getDetails() (entities.TokenDetails, error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return entities.TokenDetails{}, fmt.Errorf("cant't get token claims")
	}

	userUID, ok := claims[userUIDTokenClaim].(string)
	if !ok {
		return entities.TokenDetails{}, fmt.Errorf("can't get userUID token claim")
	}

	accessUUID, ok := claims[accessUUIDTokenClaim].(string)
	if !ok {
		return entities.TokenDetails{}, fmt.Errorf("can't get accessUUID token claim")
	}

	atExpires, ok := claims[expiresTokenClaim].(string)
	if !ok {
		return entities.TokenDetails{}, fmt.Errorf("can't get expires token claim")
	}

	expTime, err := time.Parse(time.RFC3339, atExpires)
	if err != nil {
		return entities.TokenDetails{}, fmt.Errorf("parsing expires token claim: %w", err)
	}

	return entities.TokenDetails{
		UserUID:    userUID,
		AccessUUID: accessUUID,
		ExpiresAt:  expTime,
	}, nil
}

func newSignedToken(userUID string, signingKey string, td entities.TokenDetails) (string, error) {
	claims := jwt.MapClaims{}
	claims[authorizedTokenClaim] = true
	claims[userUIDTokenClaim] = userUID
	claims[accessUUIDTokenClaim] = td.AccessUUID
	claims[expiresTokenClaim] = td.ExpiresAt

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("creating jwt token signed with string: %w", err)
	}

	return token, nil
}
