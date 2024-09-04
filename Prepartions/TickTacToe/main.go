package main

import "ticktactoe/logic"

func main() {
	game := logic.InitGame(3,3)
	game.StartGame()
}