package dict

const (
	SPACE  = 0
	WALL   = 1
	PLAYER = 100
	TAIL   = 101
)

var Sprites = map[int]rune{
	PLAYER: '●',
	WALL:   '#',
	SPACE:  ' ',
	TAIL:   '○',
}
