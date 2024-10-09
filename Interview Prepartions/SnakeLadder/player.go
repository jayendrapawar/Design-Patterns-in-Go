package main

type Player struct {
	name    string
	currPos int
}

func NewPlayer(name string, currPos int) Player{
	return Player{
		name: name,
		currPos: currPos,
	}
}

func (p *Player) GetName() string{
	return p.name
}

func (p *Player) GetCurrentPos() int {
	return p.currPos
}

// getters and setters
