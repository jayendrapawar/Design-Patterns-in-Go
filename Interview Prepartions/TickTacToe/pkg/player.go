package entity

type Player struct {
	name         string
	playingPiece PlayingPiece
	id           int
}

func NewPlayer(name string, piece PlayingPiece, id int) Player {
	return Player{
		name:         name,
		playingPiece: piece,
		id:           id,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPlayingPiece() PlayingPiece {
	return p.playingPiece
}

func (p *Player) GetID() int {
	return p.id
}
