package appcofig

import "time"

type JWTConfiguration struct {
	SigningKey         string        `yaml:"signing_key"`
	ExpirationDuration time.Duration `yaml:"expiration_duration"`
}
