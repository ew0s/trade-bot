package redis

import (
	"context"
	"fmt"

	"github.com/ew0s/trade-bot/internal/configer/appcofig"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(ctx context.Context, cfg appcofig.RedisConfiguration) (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     cfg.Host + ":" + cfg.Port,
			Password: cfg.Password,
			DB:       0,
		})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("pinging redis client: %w", err)
	}

	return client, nil
}
