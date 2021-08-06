package main

import (
    "fmt"
    "time"
    "github.com/lcarva/festoji/app"
)

func main() {
    fmt.Println(app.GetCharacter(time.Now()))
}
