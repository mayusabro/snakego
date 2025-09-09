package engine

type IEntity interface {
	Get() *Entity
	Update(*Game)
	Start(*Game)

	Destroy()
}

type Entity struct {
	Position    Position
	Collision   any
	SurfaceId   int
	isDestroyed bool
	Id          int
}

func (e *Entity) Get() *Entity {
	return e
}

func (e *Entity) Destroy() {
	e.isDestroyed = true
}

func (e *Entity) Start(game *Game) {}

func (e *Entity) Update(game *Game) {}
