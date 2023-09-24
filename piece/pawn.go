package piece

import (
	"github.com/paulwalker-dev/chess/move"
)

type Pawn struct {
	hasMoved bool
}

func (*Pawn) Name() rune {
	return 'P'
}

func pawnMoveValid(m move.Move, hasMoved bool) bool {
	p, err := m.At(move.Pos{0, 0})
	if err != nil {
		return false
	}

	if m.Dest.X != 1 {
		if m.Dest.X != 2 {
			return false
		}
		if hasMoved {
			return false
		}
	}

	o, err := m.At(m.Dest)
	if err != nil {
		return false
	}

	if p.White && m.North {
		return false
	}

	if !p.White && !m.North {
		return false
	}

	switch m.Dest.Y {
	case 0:
		return !o.PiecePresent
	case 1:
		return o.PiecePresent
	default:
		return false
	}
}

func (p *Pawn) MoveValid(m move.Move) bool {
	valid := pawnMoveValid(m, p.hasMoved)
	if valid {
		p.hasMoved = true
	}

	return valid
}
