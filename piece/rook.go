package piece

import "github.com/paulwalker-dev/chess/move"

type Rook struct{}

func (Rook) Name() rune {
	return 'R'
}

func (Rook) MoveValid(m move.Move) bool {
	var f func(int) move.Pos = nil
	var max int

	if m.Dest.X == 0 {
		max = m.Dest.Y
		f = func(i int) move.Pos {
			return move.Pos{0, i}
		}
	}

	if m.Dest.Y == 0 {
		max = m.Dest.X
		f = func(i int) move.Pos {
			return move.Pos{i, 0}
		}
	}

	if f == nil {
		return false
	}

	for i := 1; i < max; i++ {
		p, err := m.At(f(i))
		if err != nil {
			return false
		}

		if p.PiecePresent {
			return false
		}
	}

	return true
}
