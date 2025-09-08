package engine

type IEntity interface {
	Get() *Entity
	Set(e *Entity)
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
	Ref         *IEntity
}

func (e *Entity) Get() *Entity {
	return e
}

func (e *Entity) Set(e2 *Entity) {
	*e = *e2
}

func (e *Entity) Destroy() {
	e.isDestroyed = true
}

func (e *Entity) Start(game *Game) {}

func (e *Entity) Update(game *Game) {}
