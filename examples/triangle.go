package main

import (
	"github.com/tadeuszjt/gfx"
)

var (
	verts = []float32{
		300, 80, // position
		0, 0, // texCoord
		1, 0, 0, 1, // colour

		500, 380,
		0, 0,
		0, 1, 0, 1,

		100, 380,
		0, 0,
		0, 0, 1, 1,
	}
)

func draw(w *gfx.WinDraw) {
	w.DrawVertexData(verts, nil, nil)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		DrawFunc: draw,
	})
}
