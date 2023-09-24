package logic

import (
	"fmt"
	"github.com/paulwalker-dev/chess/board"
	"github.com/paulwalker-dev/chess/move"
	"github.com/paulwalker-dev/chess/piece"
	"log"
)

func (l *Logic) MakeMove(from, to move.Pos) bool {
	fromPiece := l.b[from.X][from.Y]
	toPiece := l.b[to.X][to.Y]

	if fromPiece.Logic == nil {
		return false
	}

	if toPiece.Logic != nil && toPiece.White == fromPiece.White {
		return false
	}

	m := l.b.ToMove(from, to)
	moveBoard := board.FromMove(m)
	moveBoard[0][0] = fromPiece
	moveBoard.Show()
	if fromPiece.Logic.MoveValid(m) {
		l.b[to.X][to.Y] = fromPiece
		l.b[from.X][from.Y] = piece.Piece{}
		if _, ok := toPiece.Logic.(piece.King); ok {
			winner := "White"
			if toPiece.White {
				winner = "Black"
			}
			log.Fatalf("The winner is %v", winner)
		}
		fmt.Println("Valid Move")
		return true
	}
	fmt.Println("Invalid Move")

	return false
}
