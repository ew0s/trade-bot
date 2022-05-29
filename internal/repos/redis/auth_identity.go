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
	expiration := time.Until(tokenDetails.ExpiresAt)

	if err := r.client.Set(ctx, tokenDetails.AccessUUID, userUID, expiration).Err(); err != nil {
		return fmt.Errorf("setting to redis: %w", err)
	}

	return nil
}

func (r *Identity) RemoveAccessToken(ctx context.Context, accessUID string) error {
	_, err := r.client.Del(ctx, accessUID).Result()
	if err != nil {
		return fmt.Errorf("deleting from redis: %w", err)
	}

	return nil
}

func (r *Identity) AccessUIDExists(ctx context.Context, accessUID string) (bool, error) {
	count, err := r.client.Exists(ctx, accessUID).Result()
	if err != nil {
		return false, fmt.Errorf("checking key exists in redis: %w", err)
	}

	return count != 0, nil
}
