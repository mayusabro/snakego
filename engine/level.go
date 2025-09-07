package engine

type Level struct {
	size            Size
	garbageEntities garbageEntities
	bytes           [][]any
	entities        []IEntity
}

func NewLevel(size Size) *Level {
	return &Level{
		size:            size,
		garbageEntities: make(garbageEntities, 0),
		bytes:           make([][]any, size.Height+1),
		entities:        make([]IEntity, 0),
	}
}

func (lvl *Level) Init() *Level {
	var height = len(lvl.bytes)
	var width = lvl.size.Width + 1
	for y := range height {
		lvl.bytes[y] = make([]any, width)
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
	var height = lvl.size.Height + 1
	var width = lvl.size.Width + 1
	for y := range height {
		for x := range width {
			if y == 0 || y == height-1 {
				lvl.bytes[y][x] = WALL
				continue
			}

			if x == 0 || x == width-1 {
				lvl.bytes[y][x] = WALL
				continue
			}

			lvl.bytes[y][x] = SPACE
		}
	}

}

func (lvl *Level) updateEntities(g *Game) {
	for _, e := range lvl.entities {
		if e.Get().Ref == nil {
			continue
		}

		e.Update(g)
		pos := e.Get().Position

		func() {
			if _t, ok := lvl.bytes[pos.Y][pos.X].(*IEntity); ok {
				t := *_t
				if t.Get().Ref == nil || t.Get().Ref == e.Get().Ref {
					return
				}
				t.Get().Collision = e
				e.Get().Collision = t
				return
			} else {
				e.Get().Collision = nil
			}

			if int, ok := lvl.bytes[pos.Y][pos.X].(int); ok {
				e.Get().SurfaceId = int
				return
			} else {
				e.Get().SurfaceId = -1
			}

		}()
		if e, ok := lvl.bytes[pos.Y][pos.X].(*IEntity); ok {
			if (*e).Get().Id == PLAYER {
				return
			}
		}
		lvl.bytes[pos.Y][pos.X] = &e

	}

}

func (lvl *Level) spawn(e IEntity) {
	_, p := lvl.garbageEntities.Pop()
	if p != nil {
		*p = e
		e.Get().Ref = p
		return
	}
	lvl.entities = append(lvl.entities, e)
	e.Get().Ref = &e
}

func (lvl *Level) despawn(e *IEntity) {
	lvl.garbageEntities = lvl.garbageEntities.Push(e)
	(*e).Destroy()
}

type garbageEntities []*IEntity

func (s garbageEntities) Push(v *IEntity) garbageEntities {
	return append(s, v)
}

func (s garbageEntities) Pop() (garbageEntities, *IEntity) {
	if len(s) == 0 {
		return s, nil
	}
	l := len(s)
	return s[:l-1], s[l-1]
}
