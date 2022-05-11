package main

import "github.com/ew0s/trade-bot/internal/configer/appcofig"

var (
	redisConfiguration = func(ctx configCtx) appcofig.RedisConfiguration {
		switch ctx.envName {
		default:
			return appcofig.RedisConfiguration{
				Host:     "localhost",
				Port:     "6379",
				Password: "trade-bot",
			}
		}
	}
)
