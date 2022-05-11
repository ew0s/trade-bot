package main

import "github.com/ew0s/trade-bot/internal/configer/appcofig"

const (
	apiProject = "api"
)

type EnvName int

const (
	Local EnvName = iota
)

var EnvironmentNames = map[EnvName]string{
	Local: "local",
}

func (n EnvName) String() string {
	return EnvironmentNames[n]
}

var (
	logConfiguration = func(ctx configCtx) appcofig.Logger {
		switch ctx.envName {
		case Local:
			return appcofig.Logger{
				Project:           apiProject,
				Format:            appcofig.Text,
				Level:             appcofig.Info,
				Env:               Local.String(),
				DisableStackTrace: false,
			}

		default:
			return appcofig.Logger{}
		}
	}
)
