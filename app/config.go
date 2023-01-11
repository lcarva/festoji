package app

import (
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

var defaultYamlData string = `
---
default: 🐚

rules:
- name: Xmas
  emoji: 🎄
  span: 14
  month: 12
  day: 25
- name: Thanksgiving
  emoji: 🦃
  span: 7
  month: 11
  week: 4
  weekday: 4
- name: New Year's
  emoji: 🍾
  span: 5
  month: 1
  day: 1
- name: Valentine's Day
  emoji: ❤️
  span: 7
  month: 2
  day: 14
- name: Halloween
  emoji: 🎃
  span: 7
  month: 10
  day: 31
- name: Andrew's Public Birthday
  emoji: 🧔
  span: 1
  month: 1
  day: 11
`

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
