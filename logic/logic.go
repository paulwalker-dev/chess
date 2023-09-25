package logic

import (
	"github.com/paulwalker-dev/chess/board"
	"github.com/paulwalker-dev/chess/piece"
)

type Logic struct {
	board      board.Board
	winChannel <-chan bool
}

func New() Logic {
	winChannel := make(chan bool, 1)
	l := Logic{board: board.New(), winChannel: winChannel}

	setupHeavies := func(r []piece.Piece, isWhite bool) {
		r[0] = piece.New(piece.Rook{}, isWhite)
		r[1] = piece.New(piece.Knight{}, isWhite)
		r[2] = piece.New(piece.Bishop{}, isWhite)
		r[3] = piece.New(piece.Queen{}, isWhite)
		r[4] = piece.New(piece.King{winChannel}, isWhite)
		r[5] = piece.New(piece.Bishop{}, isWhite)
		r[6] = piece.New(piece.Knight{}, isWhite)
		r[7] = piece.New(piece.Rook{}, isWhite)
	}

	setupPawns := func(r []piece.Piece, isWhite bool) {
		for x := range r {
			r[x] = piece.New(&piece.Pawn{}, isWhite)
		}
	}

	setupHeavies(l.board[0], false)
	setupPawns(l.board[1], false)

	setupHeavies(l.board[7], true)
	setupPawns(l.board[6], true)

	return l
}
