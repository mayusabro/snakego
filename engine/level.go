package engine

type Level struct {
	size            Size
	garbageEntities garbageEntities
	bytes           [][]byte
	entities        []IEntity
}

func NewLevel(size Size) *Level {
	return &Level{
		size:            size,
		garbageEntities: make(garbageEntities, 0),
		bytes:           make([][]byte, size.Height+1),
		entities:        make([]IEntity, 0),
	}
}

func (lvl *Level) Init() *Level {
	var height = len(lvl.bytes)
	var width = lvl.size.Width + 1
	for y := range height {
		lvl.bytes[y] = make([]byte, width)
	}
	return lvl
}

func (lvl *Level) update() {
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

	for _, e := range lvl.entities {
		if e.GetRef() == -1 {
			continue
		}
		pos := e.GetPosition()

		lvl.bytes[pos.Y][pos.X] = e.GetId()
	}
}

func (lvl *Level) spawn(e IEntity) {
	_, p := lvl.garbageEntities.Pop()
	if p != nil {
		*p = e
		return
	}
	lvl.entities = append(lvl.entities, e)
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
