package entity

type GameStatus string

const (
	InProgress GameStatus = "InProgress"
	Win        GameStatus = "Win"
	Draw       GameStatus = "Draw"
)
