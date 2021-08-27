package main

import (
    "fmt"
    "github.com/lcarva/festoji/app"
    "log"
    "os"
    "time"
)

func main() {
    userConfigPath := os.ExpandEnv("${HOME}/.festoji.yaml")
    config, errConfig := app.GetConfig(userConfigPath)
    if errConfig != nil {
        log.Fatal("Unable to load configuration: ", errConfig)
        return
    }
    character, errChar := app.GetCharacter(time.Now(), config)
    if errChar != nil {
        log.Fatal("Unable to get character: ", errChar)
        return
    }
    fmt.Println(character)
}
