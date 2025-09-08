package engine

import (
	"bytes"
	"fmt"

	"github.com/mayusabro/snakego/dict"
)

type Renderer struct {
	buf      *bytes.Buffer
	messages []string
}

func NewRenderer() *Renderer {
	return &Renderer{
		buf:      bytes.NewBufferString(""),
		messages: make([]string, 0),
	}
}

func (r *Renderer) render(startRender func(*Renderer)) {
	r.buf.Reset()
	startRender(r)
	fmt.Printf("%s%s", "\033[2J\033[1;1H", r.buf.String())

}

func (r *Renderer) renderGame(g *Game) {
	w := g.World
	data := w.GetCurrentLevel().Bytes
	for y := range data {
		for _x := 0; _x < (len(data[y])-1)*2; _x++ {
			if _x%2 == 1 {
				g.renderer.buf.WriteRune(
					dict.Sprites[dict.SPACE],
				)
				_x++
			}
			x := _x / 2
			if r.writeMessage(g, x, y) {
				continue
			}
			var sprite rune
			if e, ok := data[y][x].(*IEntity); ok {
				if (*e).Get().Ref == nil {
					sprite = dict.Sprites[dict.SPACE]
				}
				sprite = dict.Sprites[(*e).Get().Id]
			} else {
				sprite = dict.Sprites[data[y][x].(int)]
			}
			g.renderer.buf.WriteRune(sprite)
		}
		g.renderer.buf.WriteString("\r\n")
	}

	g.States.render(r)

}

func (r *Renderer) writeMessage(g *Game, x, y int) bool {
	messagesLen := len(r.messages)
	if messagesLen == 0 {
		return false
	}
	var centerV = g.World.GetCurrentLevel().Size.Height / 2
	var centerH = g.World.GetCurrentLevel().Size.Width / 2

	var startIndexV = centerV - messagesLen/2
	var currentIndexV = y - startIndexV
	if currentIndexV < 0 || currentIndexV >= messagesLen {
		return false
	}

	message := r.messages[currentIndexV]

	if len(message) == 0 {
		return false
	}

	var startIndexH = centerH - len(message)/2
	var currentIndexH = x - startIndexH
	if currentIndexH < 0 || currentIndexH >= len(message) {
		return false
	}
	g.renderer.buf.WriteRune(rune(r.messages[currentIndexV][currentIndexH]))
	return true

}

func (r *Renderer) addMessageLine(s string, args ...any) {
	r.messages = append(r.messages, fmt.Sprintf(s, args...))
}
