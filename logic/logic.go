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

func (l *Logic) Loop() {
	scanner := bufio.NewScanner(os.Stdin)
	l.board.Show()
	fmt.Println("Please enter your move")

loop:
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

		l.board.Show()

		select {
		case winner := <-l.winChannel:
			name := "White"
			if winner {
				name = "Black"
			}
			fmt.Printf("The winner is %s\n", name)
			break loop
		default:
		}

		fmt.Println("Please enter your move")
	}
}
