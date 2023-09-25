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
