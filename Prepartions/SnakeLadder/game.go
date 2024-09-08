package main

import "fmt"

type IGame interface {
	StartGame()
}

type Game struct {
	board   Board
	dice    Dice
	players []Player
	winner  *Player
	turn    int
}

func InitGame(boardSize, numberOfSnakes, numberOfLadders, diceCnt int, playerNames []string) *Game {
	board := NewBoard(boardSize, numberOfSnakes, numberOfLadders)
	dice := NewDice(diceCnt)
	players := make([]Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = NewPlayer(name, 0)
	}
	return &Game{
		board:   board,
		dice:    dice,
		players: players,
		winner:  nil,
		turn:    0,
	}
}

func (g *Game) StartGame() {
	totalCells := len(g.board.cells) * len(g.board.cells)

	for g.winner == nil {
		playerTurn := &g.players[g.turn]
		fmt.Printf("Player turn: %s, current position: %d\n", playerTurn.name, playerTurn.currPos)

		diceNumbers := g.dice.RollDice()
		playerNewPosition := playerTurn.currPos + diceNumbers

		if playerNewPosition >= totalCells-1 {
			playerNewPosition = totalCells - 1
		}

		playerNewPosition = g.jumpCheck(playerNewPosition)
		playerTurn.currPos = playerNewPosition

		fmt.Printf("Player %s new position: %d\n", playerTurn.name, playerNewPosition)

		if playerNewPosition == totalCells-1 {
			g.winner = playerTurn
			break
		}

		g.turn = (g.turn + 1) % len(g.players)
	}

	fmt.Printf("WINNER IS: %s\n", g.winner.name)
}

func (g *Game) jumpCheck(playerNewPosition int) int {
	if playerNewPosition >= len(g.board.cells)*len(g.board.cells) {
		return playerNewPosition
	}

	cell := g.board.getCell(playerNewPosition)
	if cell.jump != nil {
		jumpType := "ladder"
		if cell.jump.start > cell.jump.end {
			jumpType = "snake"
		}
		fmt.Printf("Hit a %s! Moving from %d to %d\n", jumpType, cell.jump.start, cell.jump.end)
		return cell.jump.end
	}

	return playerNewPosition
}
