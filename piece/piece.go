package piece

import "github.com/paulwalker-dev/chess/move"

type Piece struct {
	White bool
	Logic PieceLogic
}

type PieceLogic interface {
	Name() rune
	MoveValid(move move.Move) bool
}

func New(logic PieceLogic, isWhite bool) Piece {
	return Piece{
		White: isWhite,
		Logic: logic,
	}
}
