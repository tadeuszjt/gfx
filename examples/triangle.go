package main

import (
	"github.com/tadeuszjt/gfx"
)

func draw(w *gfx.WinDraw) {
	data := []float32{
		-0, 0,
		0, 0,
		0, 0, 0, 1,
		100, 0,
		0, 0,
		0, 0, 1, 1,
		0, 100,
		0, 0,
		0, 1, 0, 1,
	}

	w.DrawVertexData(data, nil, nil)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		DrawFunc: draw,
		Title:    "Triangle",
	})
}
