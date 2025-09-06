package engine

import (
	"bytes"
	"fmt"
)

type Renderer struct {
	buf *bytes.Buffer
}

func NewRenderer() *Renderer {
	return &Renderer{
		buf: bytes.NewBufferString(""),
	}
}

func (r *Renderer) render(startRender func(*Renderer)) {
	r.buf.Reset()
	fmt.Println("\033[2J\033[1;1H")
	startRender(r)
	fmt.Println(r.buf.String())

}
