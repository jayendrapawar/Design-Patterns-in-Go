Entities:
- Playing Piece
- Board
- Player
- GameStatus
- Game 

PlayingPiece:
- ENUM(X,O)

Player:
- Name
- ID
- PlayingPiece

Board:
- row 
- col
- cells

GameStatus:
- ENUM(InProgress, Win, Draw)

Game:
- board
- status
- players
- winner

