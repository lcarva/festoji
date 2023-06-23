package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lcarva/festoji/app"
)

var (
	n    = flag.Int("n", 1, "number of characters/days to display")
	day  = flag.Bool("day", false, "display day for the character")
	rule = flag.Bool("rule", false, "display matching rule for the character")
)

func main() {
	flag.Parse()

	userConfigPath := os.ExpandEnv("${HOME}/.festoji.yaml")
	config, errConfig := app.NewConfig(userConfigPath)
	if errConfig != nil {
		log.Fatal("Unable to load configuration: ", errConfig)
		return
	}

	t := time.Now()
	for i := 0; i < *n; i++ {
		character, ruleName, errChar := app.Character(t, config)
		if errChar != nil {
			log.Fatal("Unable to get character: ", errChar)
			return
		}

		var parts []string
		if *day {
			parts = append(parts, t.Format("02/Jan/2006"))
		}
		parts = append(parts, character)
		if *rule {
			parts = append(parts, ruleName)
		}
		fmt.Println(strings.Join(parts, " "))

		t = t.Add(time.Hour * 24)
	}

}
