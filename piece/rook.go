package piece

import "github.com/paulwalker-dev/chess/move"

type Rook struct{}

func (Rook) Name() rune {
	return 'R'
}

func (Rook) MoveValid(m move.Move) bool {
	var f func(int) move.Pos = nil

	if m.Dest.X == 0 {
		f = func(i int) move.Pos {
			return move.Pos{0, i}
		}
	}

	if m.Dest.Y == 0 {
		f = func(i int) move.Pos {
			return move.Pos{i, 0}
		}
	}

	if f == nil {
		return false
	}

	for i := 1; i < m.Dest.Y; i++ {
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
