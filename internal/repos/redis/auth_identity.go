package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/ew0s/trade-bot/internal/domain/entities"
)

type Identity struct {
	client *redis.Client
}

func NewIdentity(client *redis.Client) *Identity {
	return &Identity{client: client}
}

func (r *Identity) SetAccessToken(ctx context.Context, userUID string, tokenDetails entities.TokenDetails) error {
	at := time.Unix(tokenDetails.AtExpires, 0)
	now := time.Now()

	expiration := at.Sub(now)

	if err := r.client.Set(ctx, tokenDetails.AccessUUID, userUID, expiration).Err(); err != nil {
		return fmt.Errorf("setting to redis: %w", err)
	}

	return nil
}

func (r *Identity) GetJWTUserUID(td entities.TokenDetails) (string, error) {
	userUID, err := r.client.Get(context.Background(), td.AccessUUID).Result()
	if err != nil {
		return "", fmt.Errorf("getting from redis: %w", err)
	}

	return userUID, nil
}

func (r *Identity) RemoveAccessToken(ctx context.Context, accessUID string) error {
	_, err := r.client.Del(ctx, accessUID).Result()
	if err != nil {
		return fmt.Errorf("deleting from redis: %w", err)
	}

	return nil
}
