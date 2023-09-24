package logic

import (
	"bufio"
	"fmt"
	"github.com/paulwalker-dev/chess/board"
	"github.com/paulwalker-dev/chess/move"
	"github.com/paulwalker-dev/chess/piece"
	"os"
	"strconv"
	"strings"
)

type Logic struct {
	b      board.Board
	winner move.ViewPiece
}

func New() Logic {
	l := Logic{b: board.New()}

	setupHeavies := func(r []piece.Piece, isWhite bool) {
		r[0] = piece.New(piece.Rook{}, isWhite)
		r[1] = piece.New(piece.Knight{}, isWhite)
		r[2] = piece.New(piece.Bishop{}, isWhite)
		r[3] = piece.New(piece.Queen{}, isWhite)
		r[4] = piece.New(piece.King{&l.winner}, isWhite)
		r[5] = piece.New(piece.Bishop{}, isWhite)
		r[6] = piece.New(piece.Knight{}, isWhite)
		r[7] = piece.New(piece.Rook{}, isWhite)
	}

	setupPawns := func(r []piece.Piece, isWhite bool) {
		for x := range r {
			r[x] = piece.New(&piece.Pawn{}, isWhite)
		}
	}

	setupHeavies(l.b[0], false)
	setupPawns(l.b[1], false)

	setupHeavies(l.b[7], true)
	setupPawns(l.b[6], true)

	return l
}

func (l *Logic) Loop() {
	scanner := bufio.NewScanner(os.Stdin)
	l.b.Show()
	fmt.Println("Please enter your move")
	for scanner.Scan() {
		input := scanner.Text()

		splitInput := strings.Split(input, ",")
		if len(splitInput) != 4 {
			fmt.Println("Invalid move")
			continue
		}

		coords := make([]int, 4)

		for i, num := range splitInput {
			cord, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Invalid move")
				continue
			}
			coords[i] = cord
		}

		l.MakeMove(move.Pos{coords[0], coords[1]}, move.Pos{coords[2], coords[3]})

		l.b.Show()
		fmt.Println("Please enter your move")
	}
}
