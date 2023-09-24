package move

import "fmt"

type Pos struct {
	X, Y int
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Pos) DiffTo(o Pos) Pos {
	return Pos{p.X - o.X, p.Y - o.Y}
}

func (p Pos) Abs() Pos {
	if p.X < 0 {
		p.X = -p.X
	}
	if p.Y < 0 {
		p.Y = -p.Y
	}

	return p
}
