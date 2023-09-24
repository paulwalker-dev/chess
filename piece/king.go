package piece

import "github.com/paulwalker-dev/chess/move"

type King struct {
	WinChannel chan<- bool
}

func (King) Name() rune {
	return 'K'
}

func (King) MoveValid(m move.Move) bool {
	return m.Dest.X < 2 && m.Dest.Y < 2
}
