package main

import (
	"github.com/tadeuszjt/gfx"
)

func draw(w *gfx.WinCanvas) {
	data := []float32{
		-0, 0,
		0, 0,
		1, 0, 0, 1,
		100, 0,
		0, 0,
		0, 1, 0, 1,
		0, 100,
		0, 0,
		0, 0, 1, 1,
	}

	w.Draw2DVertexData(data, nil, nil)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		DrawFunc: draw,
		Title:    "Triangle",
	})
}
