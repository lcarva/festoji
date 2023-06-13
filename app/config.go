package app

import (
	_ "embed"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type FestojiConfig struct {
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

func GetConfig(userConfigPath string) (FestojiConfig, error) {
	defaultConfig := FestojiConfig{}
	defaultData := []byte(defaultYamlData)
	if err := yaml.UnmarshalStrict(defaultData, &defaultConfig); err != nil {
		return defaultConfig, err
	}

	userConfig := FestojiConfig{}
	userData, readUserConfigErr := ioutil.ReadFile(userConfigPath)
	if readUserConfigErr != nil {
		return defaultConfig, nil
	}

	if err := yaml.UnmarshalStrict(userData, &userConfig); err != nil {
		return userConfig, err
	}

	if userConfig.Extend {
		if userConfig.Default == "" {
			userConfig.Default = defaultConfig.Default
		}
		newRules := defaultConfig.Rules
		for _, rule := range userConfig.Rules {
			newRules = append(newRules, rule)
		}
		userConfig.Rules = newRules
	}
	return userConfig, nil
}
