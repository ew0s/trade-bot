package main

import (
	"github.com/ew0s/trade-bot/internal/configer/appcofig"
)

const (
	appAPI = "api"
)

var (
	appConfigs = map[string]func(ctx configCtx) interface{}{
		appAPI: func(ctx configCtx) interface{} {
			return appcofig.API{
				ListenAddr: ":5000",
				BasePath:   "/trade-bot/api/v1",
				DocsPath:   "/trade-bot/api/v1/docs",
				Log:        logConfiguration(ctx),
				Postgres:   postgresConfiguration(ctx),
				Redis:      redisConfiguration(ctx),
				JWT:        jwtConfiguration(ctx),
			}
		},
	}
)
