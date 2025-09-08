package engine

type World struct {
	currentLevel  int
	Score         int
	ScoreListener chan int
	gameOver      bool
	levels        []*Level
}

func (w *World) GetCurrentLevel() *Level {
	return w.levels[w.currentLevel]
}

func NewWorld(levels ...*Level) *World {
	return &World{
		levels:        levels,
		currentLevel:  0,
		ScoreListener: make(chan int),
	}
}

func (w *World) update(g *Game) {
	w.GetCurrentLevel().update(g)
	w.RenderStates(g)
}

func (w *World) Spawn(e IEntity, position Position) {
	e.Get().Position = position
	w.GetCurrentLevel().spawn(e)
}

func (w *World) Despawn(entity IEntity) {
	w.GetCurrentLevel().despawn(entity)
}

func (w *World) RenderStates(g *Game) {
	g.logger("Level : %d", g.World.currentLevel)
	g.logger("Score : %d", g.World.Score)
}

func (w *World) AddScore(g *Game, s int) (int, int, int) {
	return w.GetCurrentLevel().scorer(g, s)
}
