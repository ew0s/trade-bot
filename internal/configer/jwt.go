package main

import (
	"time"

	"github.com/ew0s/trade-bot/internal/configer/appcofig"
)

const (
	jwtHidedSigningKey = "${JWT_SIGNING_KEY}"
)

var (
	jwtConfiguration = func(ctx configCtx) appcofig.JWTConfiguration {
		switch ctx.envName {
		case Local:
			return appcofig.JWTConfiguration{
				SigningKey:         jwtSigningKey(ctx),
				ExpirationDuration: jwtTokenExpirationDuration(ctx),
			}
		default:
			return appcofig.JWTConfiguration{}
		}
	}

	jwtSigningKey = func(ctx configCtx) string {
		return jwtHidedSigningKey
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
