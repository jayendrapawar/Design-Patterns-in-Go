package entity

import "fmt"

type Board struct {
	cells     [][]PlayingPiece
	row       int
	col       int
	freeSpace int
}

func NewBoard(row int, col int) Board {
	cells := make([][]PlayingPiece, row)
	for i := range cells {
		cells[i] = make([]PlayingPiece, col)
		for j := range cells[i] {
			cells[i][j] = PieceEmpty
		}
	}

	return Board{
		cells:     cells,
		row:       row,
		col:       col,
		freeSpace: row * col,
	}
}

func (b *Board) DisplayBoard() {
	for i := 0; i < b.row; i++ {
		for j := 0; j < b.row; j++ {
			print(b.cells[i][j])
			if j != b.col-1 {
				print(" | ")
			}
		}
		println()
	}
}

func (b *Board) PlacePiece(r int, c int, piece PlayingPiece) error {
	if !b.isValidMove(r, c) {
		return fmt.Errorf("Not Valid Move")
	}
	b.cells[r-1][c-1] = piece
	b.freeSpace--;
	return nil
}

func (b *Board) isValidMove(r int, c int) bool{
	return r>=1 && r<= b.row && c>=1 && c <= b.col && b.cells[r-1][c-1] == PieceEmpty
}

func (b *Board) CheckFreeSpace() bool{
	return b.freeSpace > 0
}

func (b *Board) CheckWin(piece PlayingPiece) bool{
	// check row 
	for i:=0; i<b.row; i++{
		for j:=0; j<b.col; j++{
			if b.cells[i][j] != piece{
				break
			}
			if j == b.col-1 {
				return true
			}
		}
	}

	// check col
	for j:=0; j<b.col; j++{
		for i:=0; i<b.row; i++{
			if b.cells[i][j] != piece{
				break
			}
			if i == b.row-1 {
				return true
			}
		}
	}

	// check diagonal
	for i:=0; i<b.row; i++{
		if b.cells[i][b.row-i-1] != piece{
			break
		}
		if i == b.row-1 {
			return true
		} 
	}

	return false
}