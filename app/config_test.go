package app

import (
    "io/ioutil"
    "os"
    "testing"
)

func TestGetConfigDefault(t *testing.T) {
    config, err := GetConfig("/this/path/does/not/exit")
    if err != nil {
        t.Error(err)
    }
    if config.Default == "" {
        t.Error("Expected non-empty .default value in config")
    }
    if len(config.Rules) == 0 {
        t.Error("Expected non empty rules")
    }
}

var testConfigNoRules string = `
---
default: 😳
`

func TestGetConfigFromPathWithNoRules(t *testing.T) {
    tmpFile, errFile := ioutil.TempFile(os.TempDir(), "festoji-test-*.yaml")
    if errFile != nil {
        t.Error(errFile)
    }
    tmpFile.Write([]byte(testConfigWithRules))
    config, errConfig := GetConfig(tmpFile.Name())
    if errConfig != nil {
        t.Error(errConfig)
    }
    if config.Default == "" {
        t.Error("Expected non-empty .default value in config")
    }
    if len(config.Rules) == 0 {
        t.Error("Expected non empty rules")
    }
}

var testConfigWithRules string = `
---
default: 😏
rules:
- name: one
  emoji: 🐲
  month: 1
  day: 2
- name: two
  emoji: 🦈
  month: 2
  week: 1
  weekday: 0
`

func TestGetConfigFromPathWithRules(t *testing.T) {
    tmpFile, errFile := ioutil.TempFile(os.TempDir(), "festoji-test-*.yaml")
    if errFile != nil {
        t.Error(errFile)
    }
    tmpFile.Write([]byte(testConfigWithRules))
    config, errConfig := GetConfig(tmpFile.Name())
    if errConfig != nil {
        t.Error(errConfig)
    }
    if config.Default == "" {
        t.Error("Expected non-empty .default value in config")
    }
    if len(config.Rules) != 2 {
        t.Error("Expected exactly two rules")
    }

    ruleOne := config.Rules[0]
    if ruleOne.Name != "one" {
        t.Error("Unexpected rule one name", ruleOne.Name)
    }
    if ruleOne.Month != 1 {
        t.Error("Unexpected rule one month", ruleOne.Month)
    }
    if ruleOne.Day != 2 {
        t.Error("Unexpected rule one day", ruleOne.Day)
    }
    if ruleOne.Emoji != "🐲" {
        t.Error("Unexpected rule one emoji", ruleOne.Emoji)
    }
    if ruleOne.Week != 0 {
        t.Error("Unexpected rule one week", ruleOne.Week)
    }
    if ruleOne.Weekday != 0 {
        t.Error("Unexpected rule one weekday", ruleOne.Weekday)
    }

    ruleTwo := config.Rules[1]
    if ruleTwo.Name != "two" {
        t.Error("Unexpected rule two name", ruleTwo.Name)
    }
    if ruleTwo.Month != 2 {
        t.Error("Unexpected rule two month", ruleTwo.Month)
    }
    if ruleTwo.Day != 0 {
        t.Error("Unexpected rule two day", ruleTwo.Day)
    }
    if ruleTwo.Emoji != "🦈" {
        t.Error("Unexpected rule two emoji", ruleTwo.Emoji)
    }
    if ruleTwo.Week != 1 {
        t.Error("Unexpected rule two week", ruleTwo.Week)
    }
    if ruleTwo.Weekday != 0 {
        t.Error("Unexpected rule two weekday", ruleTwo.Weekday)
    }
}

var testConfigWithExtend string = `
---
default: 😏
extend: true
rules:
- name: five
  emoji: 🐲
  month: 1
  day: 2
`

func TestGetConfigFromPathWithExtend(t *testing.T) {
    tmpFile, errFile := ioutil.TempFile(os.TempDir(), "festoji-test-*.yaml")
    if errFile != nil {
        t.Error(errFile)
    }
    tmpFile.Write([]byte(testConfigWithExtend))
    config, errConfig := GetConfig(tmpFile.Name())
    if errConfig != nil {
        t.Error(errConfig)
    }
    if config.Default == "" {
        t.Error("Expected non-empty .default value in config")
    }
    if len(config.Rules) != 5 {
        t.Error("Expected exactly five rules")
    }

    ruleOne := config.Rules[4]
    if ruleOne.Name != "five" {
        t.Error("Unexpected rule five name", ruleOne.Name)
    }
    if ruleOne.Month != 1 {
        t.Error("Unexpected rule one month", ruleOne.Month)
    }
    if ruleOne.Day != 2 {
        t.Error("Unexpected rule one day", ruleOne.Day)
    }
}
