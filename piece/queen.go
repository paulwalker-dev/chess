package piece

import "github.com/paulwalker-dev/chess/move"

type Queen struct{}

func (Queen) Name() rune {
	return 'Q'
}

func (Queen) MoveValid(m move.Move) bool {
	r := Rook{}
	b := Bishop{}

	return r.MoveValid(m) || b.MoveValid(m)
}
