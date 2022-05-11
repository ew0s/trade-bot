package main

import (
	"time"

	"github.com/ew0s/trade-bot/internal/configer/appcofig"
)

const (
	jwtSigningKeyEnvName = "${JWT_SIGNING_KEY}"
)

var (
	jwtConfiguration = func(ctx configCtx) appcofig.JWTConfiguration {
		switch ctx.envName {
		case Local:
			return appcofig.JWTConfiguration{
				SigningKey:         jwtSigningKeyEnvName,
				ExpirationDuration: jwtTokenExpirationDuration(ctx),
			}
		default:
			return appcofig.JWTConfiguration{}
		}
	}

	jwtTokenExpirationDuration = func(ctx configCtx) time.Duration {
		switch ctx.envName {
		case Local:
			return 5 * time.Minute
		default:
			return 0
		}
	}
)
