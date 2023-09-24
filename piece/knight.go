package piece

import "github.com/paulwalker-dev/chess/move"

type Knight struct{}

func (Knight) Name() rune {
	return 'K'
}

func (Knight) MoveValid(m move.Move) bool {
	if m.Dest.X == 2 && m.Dest.Y == 1 {
		return true
	}
	if m.Dest.X == 1 && m.Dest.Y == 2 {
		return true
	}

	return false
}
