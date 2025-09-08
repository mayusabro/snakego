package dict

const (
	SPACE  = 0
	WALL   = 1
	PLAYER = 100
	TAIL   = 101

	//FOOD
	SMALL_FOOD = 200
	BIG_FOOD   = 201
	SPEED_FOOD = 202
)

var Foods = [...]int{SMALL_FOOD, BIG_FOOD, SPEED_FOOD}

var Sprites = map[int]rune{
	WALL:  '#',
	SPACE: ' ',

	//PLAYER
	PLAYER: '●',
	TAIL:   '○',

	//FOOD
	SMALL_FOOD: '◈',
	BIG_FOOD:   '◆',
	SPEED_FOOD: '✦',
}
