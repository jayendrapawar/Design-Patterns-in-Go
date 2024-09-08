package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	playerNames := []string{"P1", "P2"}
	game := InitGame(10, 5, 5, 1, playerNames)
	game.StartGame()
}
