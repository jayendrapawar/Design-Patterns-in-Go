package logic

import (
	"fmt"
	entity "ticktactoe/pkg"
)

type IGame interface {
	StartGame()
}

type game struct {
	board     entity.Board
	Players   []entity.Player
	status    entity.GameStatus
	moveCount int
	winner    entity.Player
}

func InitGame(row int, col int) IGame {
    players := make([]entity.Player, 2)
    players[0] = entity.NewPlayer("P1", entity.PieceX, 1)
    players[1] = entity.NewPlayer("P2", entity.PieceO, 2)

    board := entity.NewBoard(row, col)
    return &game{
        board:   board,
        Players: players,
        status:  entity.InProgress,
    }
}

func (g *game) StartGame(){

	for g.status == entity.InProgress {
		g.board.DisplayBoard()
		g.makeMove()
		g.checkGameStatus()
		g.moveCount++
	}
}

func (g *game) makeMove(){
	player := g.Players[g.moveCount % len(g.Players)]
	println("Player is", player.GetName())
	println("Enter the row and col")
	var row, col int
	_, err := fmt.Scanln(&row,&col)
	if err != nil {
		println("Invalid Move")
		g.makeMove()
		return
	}

	placingErr := g.board.PlacePiece(row, col, player.GetPlayingPiece())
	if placingErr != nil{
		println("Err :",placingErr)
		println("Try Again")
		g.makeMove()
	}

}

func (g *game) checkGameStatus(){
	// winning check status
	if g.board.CheckWin(g.Players[g.moveCount%len(g.Players)].GetPlayingPiece()){
		g.status = entity.Win
		g.winner = g.Players[g.moveCount%len(g.Players)]
		println("Winner is :", g.winner.GetName())
		return
	}	

	if !g.board.CheckFreeSpace(){
		g.status = entity.Draw
		println("Game Draw :(")
	}
}