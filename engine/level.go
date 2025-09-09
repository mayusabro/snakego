package engine

import (
	"math/rand"

	"github.com/mayusabro/snakego/dict"
	. "github.com/mayusabro/snakego/dict"
)

type Level struct {
	Size         Size
	Bytes        [][]any
	entities     []IEntity
	tempEntities []IEntity
	entityQueue  Queue[IEntity]
	stage        int
	scorer       func(*Game, int) (int, int, int)
}

func NewLevel(size Size) *Level {
	return &Level{
		Size:        size,
		Bytes:       make([][]any, size.Height+1),
		entities:    []IEntity{},
		entityQueue: Queue[IEntity]{},
		scorer:      scoreRule(10),
	}
}

func (lvl *Level) Init() *Level {
	var height = len(lvl.Bytes)
	var width = lvl.Size.Width + 1
	for y := range height {
		lvl.Bytes[y] = make([]any, width)
	}
	return lvl
}

func (lvl *Level) update(g *Game) {
	g.logger("Entities : %v", len(lvl.entities))
	if !g.World.gameOver {
		lvl.updateLevel()
		lvl.updateEntities(g)
	}
}

func (lvl *Level) updateLevel() {
	var height = lvl.Size.Height + 1
	var width = lvl.Size.Width + 1
	for y := range height {
		for x := range width {
			if y == 0 || y == height-1 {
				lvl.Bytes[y][x] = WALL
				continue
			}

			if x == 0 || x == width-1 {
				lvl.Bytes[y][x] = WALL
				continue
			}

			lvl.Bytes[y][x] = SPACE
		}
	}

}

func (lvl *Level) updateEntities(g *Game) {
	lvl.entities = lvl.tempEntities
	lvl.spawnQueue()
	lvl.tempEntities = lvl.entities[0:]
	for i, e := range lvl.entities {
		if e.Get().isDestroyed {
			lvl.tempEntities = removeElement(lvl.entities, i)
			continue
		}
		e.Update(g)
		pos := e.Get().Position

		func() {
			if t, ok := lvl.Bytes[pos.Y][pos.X].(IEntity); ok {
				if t.Get().isDestroyed {
					return
				}
				t.Get().Collision = e
				e.Get().Collision = t
				return
			} else {
				e.Get().Collision = nil
			}

			if int, ok := lvl.Bytes[pos.Y][pos.X].(int); ok {
				e.Get().SurfaceId = int
			} else {
				e.Get().SurfaceId = -1
			}
		}()

		if e, ok := lvl.Bytes[pos.Y][pos.X].(IEntity); ok {
			if e.Get().Id == PLAYER {
				return
			}
		}
		lvl.Bytes[pos.Y][pos.X] = e
	}

}

func (lvl *Level) spawnQueue() {
	e, err := lvl.entityQueue.Dequeue()
	if err != nil {
		return
	}

	if (*e).Get().Position.Equals(Position{}.Undefined()) {
		pos := lvl.getRandomPosition()
		lb := lvl.Bytes
		for lb[pos.X][pos.Y] != dict.SPACE {
			pos = lvl.getRandomPosition()
		}
		(*e).Get().Position = pos
	}
	lvl.entities = append(lvl.entities, *e)

}

func (l *Level) getRandomPosition() Position {
	return Position{X: rand.Intn(2 + l.Size.Width - 2), Y: rand.Intn(2 + l.Size.Height - 2)}
}

func (lvl *Level) spawn(e IEntity) {
	lvl.entityQueue.Enqueue(&e)
}

func (lvl *Level) despawn(e IEntity) {
	e.Destroy()
}

func scoreRule(target int) func(g *Game, s int) (int, int, int) {
	scorer := 0
	return func(g *Game, s int) (int, int, int) {
		g.World.Score += s
		scorer += s
		inc := 0
		for scorer >= target {
			scorer -= target
			g.World.GetCurrentLevel().stage++
			inc++
		}
		return g.World.GetCurrentLevel().stage, scorer, inc
	}
}
