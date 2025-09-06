package engine

import "math"

type Size struct {
	Width, Height int
}

type Position struct {
	X float64
	Y float64
}

func (p *Position) Move(value float64, d Direction) *Position {
	return &Position{
		X: p.X + value*float64(d.X),
		Y: p.Y + value*float64(d.Y),
	}
}

func (p *Position) Floor() (x int, y int) {
	return int(math.Floor(p.X)), int(math.Floor(p.Y))
}

func (p *Position) Add(p2 *Position) *Position {
	p.X = p.X + p2.X
	p.Y = p.Y + p2.Y
	return p
}

func (p *Position) Add2(x, y float64) *Position {
	p.X = p.X + x
	p.Y = p.Y + y
	return p
}

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
