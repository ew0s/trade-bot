package main

import "github.com/ew0s/trade-bot/internal/configer/appcofig"

var (
	postgresConfiguration = func(ctx configCtx) appcofig.PostgresConfiguration {
		switch ctx.envName {
		default:
			return appcofig.PostgresConfiguration{
				Host: "localhost",
				Port: "5432",
				Credentials: appcofig.PostgresCredentials{
					Username: "trade_bot",
					Password: "trade_bot",
				},
				DBName:  "trade_bot",
				SSLMode: "disable",
			}
		}
	}
)
