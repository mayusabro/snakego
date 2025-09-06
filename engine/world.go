package engine

type World struct {
	currentLevel int
	levels       []*Level
}

func (w *World) getCurrentLevel() *Level {
	return w.levels[w.currentLevel]
}

func (w *World) render(g *Game) {
	data := w.getCurrentLevel().bytes
	for y := range data {
		for x := range data[y] {
			g.renderer.buf.WriteRune(sprites[data[y][x]])
		}
		g.renderer.buf.WriteString("\r\n")
	}

	for _, e := range w.getCurrentLevel().entities {
		e.Update(g)
	}
}

func NewWorld(levels ...*Level) *World {
	return &World{
		levels:       levels,
		currentLevel: 0,
	}
}

func (w *World) update() {
	w.getCurrentLevel().update()
}

func (w *World) Spawn(e IEntity, position Position) {
	e.SetPosition(position)
	w.getCurrentLevel().spawn(e)
}

func (w *World) Despawn(entity *IEntity) {
	w.getCurrentLevel().despawn(entity)
}
