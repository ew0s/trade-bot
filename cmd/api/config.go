package main

import (
	"flag"
	"fmt"

	"github.com/ew0s/trade-bot/internal/configer/appcofig"
	"github.com/ew0s/trade-bot/pkg/config"
	"github.com/joho/godotenv"
)

func mustParseAppConfig() (appcofig.API, error) {
	configPath := flag.String("config-path", "/opt/app/config/application.yaml", "application config path")
	envFilePath := flag.String("env-file-path", "", "application env file path")

	flag.Parse()

	if *envFilePath != "" {
		if err := godotenv.Load(*envFilePath); err != nil {
			return appcofig.API{}, fmt.Errorf("loading .env file variables: %w", err)
		}
	}

	configurator, err := config.NewConfigurator(config.YAML)
	if err != nil {
		return appcofig.API{}, fmt.Errorf("creating configurator: %w", err)
	}

	var cfg appcofig.API
	if err = configurator.UnmarshalAndSetup(*configPath, &cfg); err != nil {
		return appcofig.API{}, fmt.Errorf("unmarshaling and setting up config: %w", err)
	}

	return cfg, nil
}
