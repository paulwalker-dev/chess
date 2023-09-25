package logic

import (
	"github.com/paulwalker-dev/chess/board"
	"github.com/paulwalker-dev/chess/move"
	"github.com/paulwalker-dev/chess/piece"
)

func (l *Logic) MakeMove(from, to move.Pos) bool {
	fromPiece := l.board[from.X][from.Y]
	toPiece := l.board[to.X][to.Y]

	if fromPiece.Logic == nil {
		return false
	}

	m := l.board.ToMove(from, to)
	moveBoard := board.FromMove(m)
	moveBoard[0][0] = fromPiece
	defer moveBoard.Show()

	if toPiece.Logic != nil && toPiece.White == fromPiece.White {
		return false
	}

	if !fromPiece.Logic.MoveValid(m) {
		return false
	}

	l.board[to.X][to.Y] = fromPiece
	l.board[from.X][from.Y] = piece.Piece{}
	if king, ok := toPiece.Logic.(piece.King); ok {
		king.WinChannel <- toPiece.White
	}
	return true
}
