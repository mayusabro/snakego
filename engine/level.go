package engine

import (
	. "github.com/mayusabro/snakego/dict"
)

type Level struct {
	Size            Size
	garbageEntities garbageEntities
	Bytes           [][]any
	entities        []IEntity
	scorer          func(*Game, int) (int, int, int)
}

func NewLevel(size Size) *Level {
	return &Level{
		Size:            size,
		garbageEntities: make(garbageEntities, 0),
		Bytes:           make([][]any, size.Height+1),
		entities:        make([]IEntity, 0),
		scorer:          scoreRule(20),
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
	for _, e := range lvl.entities {
		if e.Get().isDestroyed {
			continue
		}
		e.Update(g)
		pos := e.Get().Position

		func() {
			if _t, ok := lvl.Bytes[pos.Y][pos.X].(*IEntity); ok {
				t := *_t
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
				return
			} else {
				e.Get().SurfaceId = -1
			}

		}()
		if e, ok := lvl.Bytes[pos.Y][pos.X].(*IEntity); ok {
			if (*e).Get().Id == PLAYER {
				return
			}
		}
		lvl.Bytes[pos.Y][pos.X] = &e

	}

}

func (lvl *Level) spawn(e IEntity) {
	var p IEntity
	lvl.garbageEntities, p = lvl.garbageEntities.Pop()
	if p != nil {
		p.Set(e.Get())
		return
	}
	lvl.entities = append(lvl.entities, e)
}

func (lvl *Level) despawn(e IEntity) {
	lvl.garbageEntities = lvl.garbageEntities.Push(e)
	e.Destroy()
}

func scoreRule(target int) func(g *Game, s int) (int, int, int) {
	scorer := 0
	stage := 1
	return func(g *Game, s int) (int, int, int) {
		g.World.Score += s
		scorer += s
		inc := 0
		for scorer >= target {
			scorer -= target
			stage++
			inc++
		}
		return stage, scorer, inc
	}
}

type garbageEntities []IEntity

func (s garbageEntities) Push(v IEntity) garbageEntities {
	return append(s, v)
}

func (s garbageEntities) Pop() (garbageEntities, IEntity) {
	if len(s) == 0 {
		return s, nil
	}
	l := len(s)
	return s[:l-1], s[l-1]
}
