package move

import (
	"fmt"
)

type ViewPiece struct {
	PiecePresent bool
	White        bool
}

type Move struct {
	Dest  Pos
	View  [][]ViewPiece
	North bool
}

type OutsideViewErr struct {
	dest      Pos
	attempted Pos
}

func (err OutsideViewErr) Error() string {
	return fmt.Sprintf("%v is not within bounds: %v-%v", err.attempted, Pos{0, 0}, err.dest)
}

func (m *Move) At(pos Pos) (*ViewPiece, error) {
	if m.Dest.X < pos.X || m.Dest.Y < pos.Y {
		return nil, OutsideViewErr{m.Dest, pos}
	}

	return &m.View[pos.X][pos.Y], nil
}
