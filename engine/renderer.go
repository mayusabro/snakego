package engine

import (
	"bytes"
	"fmt"
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
	fmt.Println("\033[2J\033[1;1H")
	startRender(r)
	fmt.Println(r.buf.String())
}

func (r *Renderer) renderGame(g *Game) {
	w := g.World
	data := w.getCurrentLevel().bytes
	for y := range data {
		for x := range data[y] {
			if r.writeMessage(g, x, y) {
				continue
			}
			g.renderer.buf.WriteRune(sprites[data[y][x]])
		}
		g.renderer.buf.WriteString("\r\n")
	}
	for _, e := range w.getCurrentLevel().entities {
		e.Update(g)
	}

	g.States.render(r)

}

func (r *Renderer) writeMessage(g *Game, x, y int) bool {
	messageLen := len(r.messages)
	if messageLen == 0 {
		return false
	}
	var centerV = g.World.getCurrentLevel().size.Height / 2
	var centerH = g.World.getCurrentLevel().size.Width / 2

	var startIndexV = centerV - messageLen/2

	if y-startIndexV < 0 {
		return false
	}
	message := r.messages[y-startIndexV]
	var startIndexH = centerH - len(message)/2
	if x-startIndexH < 0 {
		return false
	}
	g.renderer.buf.WriteRune(rune(r.messages[y][x]))
	return true

}

func (r *Renderer) addMessageLine(s string) {
	r.messages = append(r.messages, s)
}
