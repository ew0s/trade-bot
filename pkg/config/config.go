package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type FileType string

var (
	YAML FileType = "yaml"
)

var availableFileTypes = map[FileType]struct{}{
	YAML: {},
}

func (t FileType) valid() bool {
	_, ok := availableFileTypes[t]
	return ok
}

type Configurator struct {
	fileType FileType
}

func NewConfigurator(fileType FileType) (*Configurator, error) {
	if !fileType.valid() {
		return nil, fmt.Errorf("invalid file type passed")
	}

	return &Configurator{fileType: fileType}, nil
}

func (c Configurator) UnmarshalAndSetup(configPath string, cfg interface{}) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	if err = c.setup(data, cfg); err != nil {
		return fmt.Errorf("setting up config: %w", err)
	}

	return nil
}

func (c Configurator) setup(data []byte, cfg interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook:  c.decodeHook,
		ErrorUnused: true,
		Result:      cfg,
		TagName:     string(c.fileType),
	})
	if err != nil {
		return fmt.Errorf("creating new decoder: %w", err)
	}

	var configMap map[string]interface{}
	if err = yaml.Unmarshal(data, &configMap); err != nil {
		return fmt.Errorf("unmarshaling config: %w", err)
	}

	if err = decoder.Decode(configMap); err != nil {
		return fmt.Errorf("decoding configuration map to config struct: %w", err)
	}

	return nil
}

func (c Configurator) decodeHook(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() == reflect.String {
		strData := data.(string)

		switch to {
		case reflect.TypeOf(time.Duration(0)):
			return time.ParseDuration(strData)

		case reflect.TypeOf(""):
			return os.ExpandEnv(strData), nil
		}
	}

	return data, nil
}
