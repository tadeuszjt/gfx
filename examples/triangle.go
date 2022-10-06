package main

import (
	"github.com/tadeuszjt/gfx"
)

const (
	width  = 640
	height = 480
)

func draw(w *gfx.Win, c gfx.Canvas) {
	triangle := []float32{
		width * 0.5, height * 0.4,
		0, 0,
		1, 0, 0, 1,
		width * 0.4, height * 0.6,
		0, 0,
		0, 1, 0, 1,
		width * 0.6, height * 0.6,
		0, 0,
		0, 0, 1, 1,
	}

	c.Draw2DVertexData(triangle, nil, nil)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		Width:    width,
		Height:   height,
		DrawFunc: draw,
		Title:    "Triangle",
	})
}
