package app

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Default string
	Extend  bool
	Rules   []struct {
		// Required
		Name  string
		Emoji string
		Span  int

		Month int // TODO: Month string
		// Either Day or Week+Weekday is required
		Day     int
		Week    int
		Weekday int // TODO: Month string
	}
}

//go:embed config.yaml
var defaultYamlData string

func NewConfig(userConfigPath string) (Config, error) {
	defaultData := []byte(defaultYamlData)
	defaultConfig := Config{}
	if err := yaml.UnmarshalStrict(defaultData, &defaultConfig); err != nil {
		return Config{}, err
	}

	userData, err := os.ReadFile(userConfigPath)
	if err != nil {
		return defaultConfig, nil
	}
	userConfig := Config{}
	if err := yaml.UnmarshalStrict(userData, &userConfig); err != nil {
		return userConfig, err
	}

	if userConfig.Extend {
		if userConfig.Default == "" {
			userConfig.Default = defaultConfig.Default
		}
		userConfig.Rules = append(defaultConfig.Rules, userConfig.Rules...)
	}
	return userConfig, nil
}
