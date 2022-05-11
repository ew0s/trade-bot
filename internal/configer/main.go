package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const applicationConfName = "application.yaml"

func parseEnvironments(envs string) ([]EnvName, error) {
	names := strings.Split(envs, ",")

	envNames := make([]EnvName, 0, len(names))

	for _, name := range names {
		envName, err := parseEnv(name)
		if err != nil {
			return nil, fmt.Errorf("parsing env name: %w", err)
		}

		envNames = append(envNames, envName)
	}

	return envNames, nil
}

func parseEnv(name string) (EnvName, error) {
	for env, envName := range EnvironmentNames {
		if name == envName {
			return env, nil
		}
	}

	return 0, fmt.Errorf("can't find environment for name='%s'", name)
}

type configCtx struct {
	envName EnvName
	appName string
}

var (
	appName      = flag.String("app-name", "", "Name of application")
	configFolder = flag.String("o", "./deploy/", "Configs folder path")
	configEnv    = flag.String("env", "local", "List of environments to be generated")
)

func main() {
	flag.Parse()

	envNames, err := parseEnvironments(*configEnv)
	if err != nil {
		log.Fatalf("parsing environments: %s", err)
	}

	if *appName == "" {
		log.Fatalf("application name can't be empty")
	}

	log.Printf("strart generating config...\n")
	log.Printf("appName='%s', configFolder='%s', env='%s'", *appName, *configFolder, *configEnv)

	for _, envName := range envNames {
		ctx := configCtx{
			envName: envName,
			appName: *appName,
		}

		config, ok := appConfigs[*appName]
		if !ok {
			log.Fatalf("can't find config for application")
		}

		out, err := yaml.Marshal(config(ctx))
		if err != nil {
			log.Fatalf("can't marshal config: %s", err)
		}

		filePath := filepath.Join(*configFolder, envName.String(), "config")
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			log.Fatalf("can't make folder directory: %s", err)
		}

		filePath = filepath.Join(filePath, applicationConfName)
		if err = ioutil.WriteFile(filePath, out, os.ModePerm); err != nil {
			log.Fatalf("can't write file: %s", err)
		}
	}
}
