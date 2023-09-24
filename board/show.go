package board

import (
	"fmt"
	"github.com/fatih/color"
)

func (b *Board) Show() {
	lightBg := true
	fmt.Print("  ")
	for i := range (*b)[0] {
		fmt.Printf(" %1d", i)
	}
	fmt.Println()
	for i, r := range *b {
		fmt.Printf("%1d ", i)
		for _, p := range r {
			var c *color.Color

			if lightBg {
				c = color.New(color.BgWhite)
			} else {
				c = color.New(color.BgHiBlack)
			}

			var symbol string

			symbol = c.Sprint(" ")

			if p.Logic != nil {
				name := p.Logic.Name()

				if p.White {
					c = color.New(color.BgHiWhite, color.FgBlack)
				} else {
					c = color.New(color.BgBlack, color.FgHiWhite)
				}

				symbol += c.Sprintf("%c", name)
			} else {
				symbol += symbol
			}

			fmt.Print(symbol)
			lightBg = !lightBg
		}
		fmt.Println()
		if len(r)%2 == 0 {
			lightBg = !lightBg
		}
	}
}
