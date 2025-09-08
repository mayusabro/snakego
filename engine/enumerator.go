package engine

type Size struct {
	Width, Height int
}

type Position struct {
	X int
	Y int
}

func (p *Position) Move(value int, d Direction) Position {
	return Position{
		X: p.X + value*d.X,
		Y: p.Y + value*d.Y,
	}
}

func (p *Position) Equals(p2 Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}

//-=-==========

type Direction struct {
	X, Y int
}

func NewDirection(x, y int) Direction {
	return Direction{X: x, Y: y}
}

func (d Direction) Zero() Direction {
	return Direction{X: 0, Y: 0}
}

func (d Direction) Up() Direction {
	return Direction{X: 0, Y: -1}
}

func (d Direction) Left() Direction {
	return Direction{X: -1, Y: 0}
}

func (d Direction) Right() Direction {
	return Direction{X: 1, Y: 0}
}

func (d Direction) Down() Direction {
	return Direction{X: 0, Y: 1}
}

func (d Direction) Multiply(m int) Direction {
	return Direction{X: d.X * m, Y: d.Y * m}
}
