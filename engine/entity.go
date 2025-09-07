package engine

const (
	SPACE  = 0
	WALL   = 1
	PLAYER = 100
	TAIL   = 101
)

var sprites = map[int]rune{
	PLAYER: '@',
	WALL:   '+',
	SPACE:  ' ',
	TAIL:   'o',
}

type IEntity interface {
	Get() *Entity

	Update(*Game)
	Start(*Game)

	Destroy()
}

type Entity struct {
	Position  Position
	Collision IEntity
	SurfaceId int
	Id        int
	Ref       *IEntity
}

func (e *Entity) Get() *Entity {
	return e
}

func (e *Entity) Destroy() {
	e.Ref = nil
}

func (e *Entity) Start(game *Game) {}

func (e *Entity) Update(game *Game) {}
