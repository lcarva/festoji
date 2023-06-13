package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lcarva/festoji/app"
)

func main() {
	userConfigPath := os.ExpandEnv("${HOME}/.festoji.yaml")
	config, errConfig := app.NewConfig(userConfigPath)
	if errConfig != nil {
		log.Fatal("Unable to load configuration: ", errConfig)
		return
	}
	character, errChar := app.Character(time.Now(), config)
	if errChar != nil {
		log.Fatal("Unable to get character: ", errChar)
		return
	}
	fmt.Println(character)
}
