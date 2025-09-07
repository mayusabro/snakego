package engine

type World struct {
	currentLevel int
	score        int
	gameOver     bool
	levels       []*Level
}

func (w *World) getCurrentLevel() *Level {
	return w.levels[w.currentLevel]
}

func NewWorld(levels ...*Level) *World {
	return &World{
		levels:       levels,
		currentLevel: 0,
	}
}

func (w *World) update(g *Game) {
	w.getCurrentLevel().update(g)
}

func (w *World) Spawn(e IEntity, position Position) {
	e.Get().Position = position
	w.getCurrentLevel().spawn(e)
}

func (w *World) Despawn(entity *IEntity) {
	w.getCurrentLevel().despawn(entity)
}
