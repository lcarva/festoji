package app

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type FestojiConfig struct {
    Default string
    Rules []struct {
        // Required
        Name string
        Emoji string
        Span int

        Month int  // TODO: Month string
        // Either Day or Week+Weekday is required
        Day int
        Week int
        Weekday int  // TODO: Month string
    }
}

var defaultConfig string = `
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
`

func GetConfig(userConfigPath string) (FestojiConfig, error) {
    config := FestojiConfig{}

    data, err := ioutil.ReadFile(userConfigPath)
    if err != nil {
        data = []byte(defaultConfig)
    }

    if err:= yaml.UnmarshalStrict(data, &config); err != nil {
        return config, err
    }
    return config, nil
}
