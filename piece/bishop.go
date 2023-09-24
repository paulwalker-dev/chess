package piece

import "github.com/paulwalker-dev/chess/move"

type Bishop struct{}

func (Bishop) Name() rune {
	return 'B'
}

func (Bishop) MoveValid(m move.Move) bool {
	if m.Dest.X != m.Dest.Y {
		return false
	}

	for i := 1; i < m.Dest.X; i++ {
		p, err := m.At(move.Pos{i, i})
		if err != nil {
			return false
		}
		if p.PiecePresent {
			return false
		}
	}

	return true
}
