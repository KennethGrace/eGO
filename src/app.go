package main

import (
	"eGO/src/engine"
	"eGO/src/sample"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	game := sample.NewGame()
	e := engine.NewEngine("GAME", 640, 360, false, game)
	e.Run()
}
