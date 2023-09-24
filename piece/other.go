package piece

import "github.com/paulwalker-dev/chess/move"

type Other struct{}

func (Other) Name() rune {
	return 'O'
}

func (Other) MoveValid(_ move.Move) bool {
	return false
}
