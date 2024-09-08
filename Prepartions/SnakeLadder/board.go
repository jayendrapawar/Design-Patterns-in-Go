package main

import "math/rand"

type Board struct {
	cells [][]Cell
}

func NewBoard(boardSize, numberOfSnakes, numberOfLadders int) Board {
	board := Board{}
	board.InitCells(boardSize)
	board.addSnakesLadders(numberOfSnakes, numberOfLadders)
	return board
}

func (b *Board) InitCells(boardSize int) {
	b.cells = make([][]Cell, boardSize)
	for i := range b.cells {
		b.cells[i] = make([]Cell, boardSize)
	}
}

func (b *Board) addSnakesLadders(numberOfSnakes, numberOfLadders int) {
	totalCells := len(b.cells) * len(b.cells)

	for numberOfSnakes > 0 {
		snakeHead := rand.Intn(totalCells-1) + 1
		snakeTail := rand.Intn(snakeHead)

		if b.getCell(snakeHead).jump == nil {
			b.getCell(snakeHead).jump = &Jump{start: snakeHead, end: snakeTail}
			numberOfSnakes--
		}
	}

	for numberOfLadders > 0 {
		ladderStart := rand.Intn(totalCells-1) + 1
		ladderEnd := rand.Intn(totalCells-ladderStart) + ladderStart

		if b.getCell(ladderStart).jump == nil {
			b.getCell(ladderStart).jump = &Jump{start: ladderStart, end: ladderEnd}
			numberOfLadders--
		}
	}
}

func (b *Board) getCell(playerCurrPos int) *Cell {
	boardRow := playerCurrPos / len(b.cells)
	boardCol := playerCurrPos % len(b.cells)
	return &b.cells[boardRow][boardCol]
}
