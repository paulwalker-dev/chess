package logic

import (
	"bufio"
	"fmt"
	"github.com/paulwalker-dev/chess/move"
	"os"
	"strconv"
	"strings"
)

func (l *Logic) Loop() {
	scanner := bufio.NewScanner(os.Stdin)
	l.board.Show()
	fmt.Println("Please enter your move")

loop:
	for scanner.Scan() {
		input := scanner.Text()

		splitInput := strings.Split(input, ",")
		if len(splitInput) != 4 {
			fmt.Println("Invalid command")
			continue
		}

		coords := make([]int, 4)

		for i, num := range splitInput {
			cord, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Number #%d not valid\n", i+1)
				continue
			}
			coords[i] = cord
		}

		fromPos := move.Pos{coords[1], coords[0]}
		toPos := move.Pos{coords[3], coords[2]}
		if !l.MakeMove(fromPos, toPos) {
			fmt.Println("Invalid move")
		}

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
