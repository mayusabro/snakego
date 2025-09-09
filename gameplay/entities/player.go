package entities

import (
	"github.com/mayusabro/snakego/dict"
	"github.com/mayusabro/snakego/engine"
)

type IPlayer interface {
	GetPlayer() *Player
}

type Snake struct {
	engine.Entity
	LastPosition  engine.Position
	LastDirection engine.Direction
	Direction     engine.Direction
}

type Player struct {
	Snake
	Speed int
	tail  *tail
	move  func(int int, deltaTime float64) int
}

func NewPlayer(position engine.Position) *Player {
	return &Player{
		Snake: Snake{
			Entity: engine.Entity{
				Id:       dict.PLAYER,
				Position: position,
			},
			LastPosition:  position,
			Direction:     engine.Direction{}.Right(),
			LastDirection: engine.Direction{}.Right(),
		},
		Speed: 5,
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
	tail.Direction = tail.parent.Direction
	tail.Position = engine.Position{
		X: tail.parent.Position.X - tail.Direction.X,
		Y: tail.parent.Position.Y - tail.Direction.Y,
	}
	g.World.Spawn(tail, tail.Position)

}

func (p *Player) Update(g *engine.Game) {
	p.Move(g)
	p.CheckCollision(g)
	p.CheckSurface(g)
	g.Logf("Speed: %v", p.Speed)

}

func (p *Player) Move(g *engine.Game) {
	levelSize := g.World.GetCurrentLevel().Size
	p.Direction = readInput(g, p.LastDirection)
	p.LastPosition = p.Position
	newPosition := p.Position.Move(
		p.move(p.Speed, g.States.DeltaTime),
		p.Direction,
	)
	p.Position = engine.Position{
		X: min(max(0, newPosition.X), levelSize.Width),
		Y: min(max(0, newPosition.Y), levelSize.Height),
	}

	if !p.Position.Equals(p.LastPosition) {
		p.LastDirection = p.Direction
	}
}

func (p *Player) CheckCollision(g *engine.Game) {
	coll := p.Collision
	if coll != nil {
		if entity, ok := coll.(engine.IEntity); ok {
			switch (entity).Get().Id {

			case dict.TAIL:
				g.GameOver()
			case dict.WALL:
				g.GameOver()
			}
			if item, ok := coll.(IItem); ok {
				p.checkItemCollision(g, item)
			}

		}

	}
}

func (p *Player) checkItemCollision(g *engine.Game, item IItem) {
	switch item.Get().Id {
	case dict.SMALL_FOOD:
		item.StartEffect(p)

	case dict.BIG_FOOD:
		item.StartEffect(p)

	case dict.SPEED_FOOD:
		item.StartEffect(p)
	}
	g.World.Despawn(item.Get())
}

func (p *Player) CheckSurface(g *engine.Game) {
	surfaceId := p.SurfaceId
	switch surfaceId {
	case dict.WALL:
		g.GameOver()
	}
}

func (p *Player) GetPlayer() *Player {
	return p
}

func movement() func(int int, deltaTime float64) int {
	move := 0.0
	return func(value int, deltaTime float64) int {
		move += min(float64(value)*deltaTime, 1.0)
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
				Id:       dict.TAIL,
				Position: parent.Position,
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
