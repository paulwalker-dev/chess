package board

import (
	"github.com/paulwalker-dev/chess/move"
	"github.com/paulwalker-dev/chess/piece"
	"log"
)

func FromMove(m move.Move) Board {
	b := make(Board, m.Dest.X+1)
	for x := range b {
		b[x] = make([]piece.Piece, m.Dest.Y+1)
		for y := range b[x] {
			var p piece.Piece

			movePiece, err := m.At(move.Pos{x, y})
			if err != nil {
				log.Fatal(err)
			}

			var logic piece.PieceLogic
			if movePiece.PiecePresent {
				logic = piece.Other{}
			}
			p = piece.New(logic, movePiece.White)

			b[x][y] = p
		}
	}

	return b
}

func (b *Board) ToMove(from, to move.Pos) move.Move {
	diff := from.DiffTo(to)
	size := diff.Abs()

	transform := func(x, y int) (int, int) {
		if diff.X > 0 {
			x = -x
		}
		if diff.Y > 0 {
			y = -y
		}

		x += from.X
		y += from.Y

		return x, y
	}

	v := make([][]move.ViewPiece, size.X+1)
	for x := range v {
		v[x] = make([]move.ViewPiece, size.Y+1)
		for y := range v[x] {
			bX, bY := transform(x, y)

			p := (*b)[bX][bY]
			if p.Logic != nil {
				v[x][y].PiecePresent = true
				v[x][y].White = p.White
			}
		}
	}

	m := move.Move{
		Dest:  size,
		View:  v,
		North: diff.X < 0,
	}

	return m
}
