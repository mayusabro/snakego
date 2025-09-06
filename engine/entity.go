package engine

const (
	SPACE  = 0
	WALL   = 1
	PLAYER = 100
	TAIL   = 101
)

var sprites = map[byte]rune{
	PLAYER: '@',
	WALL:   '+',
	SPACE:  ' ',
	TAIL:   'o',
}

type IEntity interface {
	GetPosition() *Position
	SetPosition(Position)
	Update(*Game)
	Start(*Game)
	GetId() byte
	GetRef() int
	Destroy()
}

type Entity struct {
	Position *Position
	Id       byte
	Ref      int
}

func (e *Entity) GetPosition() *Position {
	return e.Position
}

func (e *Entity) SetPosition(pos Position) {
	e.Position = &pos
}
func (e *Entity) GetId() byte {
	return e.Id
}

func (e *Entity) GetRef() int {
	return e.Ref
}

func (e *Entity) Destroy() {
	e.Ref = -1
}

func (e *Entity) Start(game *Game) {}

func (e *Entity) Update(game *Game) {}
