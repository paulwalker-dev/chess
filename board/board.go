package board

import (
	"github.com/paulwalker-dev/chess/piece"
)

type Board [][]piece.Piece

func New() Board {
	b := make(Board, 8)
	for r := range b {
		b[r] = make([]piece.Piece, 8)
	}

	return b
}
