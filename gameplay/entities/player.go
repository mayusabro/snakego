package entities

import (
	"github.com/mayusabro/snakego/engine"
)

type Snake struct {
	engine.Entity
	LastPosition  engine.Position
	LastDirection engine.Direction
	Direction     engine.Direction
}

type Player struct {
	Snake
	speed int
	tail  []*tail
}

func NewPlayer() *Player {
	return &Player{
		Snake: Snake{
			Direction: engine.Direction{}.Zero(),
			Entity: engine.Entity{
				Id: engine.PLAYER,
			},
		},
		speed: 4, //fps
		tail:  make([]*tail, 0),
	}
}

func (p *Player) AddTail(g *engine.Game) {
	var tail *tail
	if len(p.tail) == 0 {
		tail = newTail(&p.Snake)
	} else {
		tail = newTail(&p.tail[len(p.tail)-1].Snake)
	}
	p.tail = append(p.tail, tail)
	g.World.Spawn(tail, tail.parent.LastPosition)

}

func (p *Player) Update(g *engine.Game) {
	p.LastDirection = p.Direction
	p.LastPosition = *p.Position
	p.Direction = p.readInput(g)
	p.Position = p.Position.Move(0.2, p.Direction)
	g.Logf("LP : %v", p.LastPosition)
	logx, logy := p.Position.Floor()
	g.Logf("P : %v ", logx, logy)
}

func (p *Player) readInput(g *engine.Game) engine.Direction {
	input := g.States.Input
	switch input {
	case 'w':
		return engine.Direction{}.Up()
	case 's':
		return engine.Direction{}.Down()
	case 'd':
		return engine.Direction{}.Right()
	case 'a':
		return engine.Direction{}.Left()
	}
	return p.Direction
}

//===== TAIL

type tail struct {
	Snake
	parent *Snake
}

func newTail(parent *Snake) *tail {
	return &tail{
		Snake: Snake{
			Entity: engine.Entity{
				Id: engine.TAIL,
			},
		},
		parent: parent,
	}
}

func (t *tail) Update(g *engine.Game) {
	nlx, nly := t.parent.LastPosition.Floor()
	olx, oly := t.parent.Position.Floor()
	if nlx == olx && nly == oly {
		return
	}

	t.LastDirection = t.Direction
	t.LastPosition = *t.Position

	t.Direction = t.parent.LastDirection

	nx, ny := t.parent.Position.Floor()
	next := engine.Position{
		X: float64(nx - t.Direction.X),
		Y: float64(ny - t.Direction.Y),
	}

	nextX, nextY := next.Floor()
	t.Position = &engine.Position{
		X: float64(nextX),
		Y: float64(nextY),
	}
}
