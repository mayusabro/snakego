package entities

import (
	"github.com/mayusabro/snakego/dict"
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
	tail  *tail
	move  func(int int, deltaTime float64) int
}

func NewPlayer() *Player {
	return &Player{
		Snake: Snake{
			Entity: engine.Entity{
				Id: dict.PLAYER,
			},
		},
		speed: 10,
		move:  movement(),
	}
}

func (p *Player) AddTail(g *engine.Game) {
	var tail *tail
	if p.tail == nil {
		tail = newTail(&p.Snake)
	} else {
		tail = newTail(&p.tail.Snake)
	}
	p.tail = tail
	g.World.Spawn(tail, tail.parent.LastPosition)

}

func (p *Player) Update(g *engine.Game) {
	p.Direction = readInput(g, p.LastDirection)
	p.LastPosition = p.Position
	p.Position = p.Position.Move(
		p.move(p.speed, g.States.DeltaTime),
		p.Direction,
	)
	if p.Position != p.LastPosition {
		p.LastDirection = p.Direction
	}
	p.CheckCollision(g)

}

func (p *Player) CheckCollision(g *engine.Game) {
	coll := p.Collision
	if coll != nil {
		switch coll.Get().Id {
		case dict.TAIL:
			g.GameOver()
		}
	}

	surfaceId := p.SurfaceId
	switch surfaceId {
	case dict.WALL:
		g.GameOver()
	}
}

func movement() func(int int, deltaTime float64) int {
	move := 0.0
	return func(value int, deltaTime float64) int {
		move += float64(value) * deltaTime
		validMove := int(move)
		move -= float64(validMove)
		return validMove
	}

}

func readInput(g *engine.Game, lastDir engine.Direction) engine.Direction {
	input := g.States.Input
	switch input {
	case 'w':
		if lastDir.Y != 0 {
			break
		}
		return engine.Direction{}.Up()
	case 's':
		if lastDir.Y != 0 {
			break
		}
		return engine.Direction{}.Down()
	case 'd':
		if lastDir.X != 0 {
			break
		}
		return engine.Direction{}.Right()
	case 'a':
		if lastDir.X != 0 {
			break
		}
		return engine.Direction{}.Left()
	}
	return lastDir
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
				Id: dict.TAIL,
			},
		},
		parent: parent,
	}
}

func (t *tail) Update(g *engine.Game) {
	lp := t.parent.LastPosition
	cp := t.parent.Position
	if lp.X == cp.X && lp.Y == cp.Y {
		return
	}

	t.LastDirection = t.Direction
	t.LastPosition = t.Position

	t.Direction = t.parent.LastDirection

	t.Position = engine.Position{
		X: cp.X - t.Direction.X,
		Y: cp.Y - t.Direction.Y,
	}

}
