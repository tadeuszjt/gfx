package main

import (
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
)

const (
	width  = 640
	height = 480
)

func draw(w *gfx.Win, c gfx.Canvas) {
	start := geom.Vec2[float32]{50, 100}
	end := geom.Vec2[float32]{200, 300}
	colour := gfx.Red
	scale := float32(10)
	view := geom.Mat3Identity[float32]()

	gfx.Draw2DArrow(c, start, end, colour, scale, view)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		Width:    width,
		Height:   height,
		DrawFunc: draw,
		Title:    "Arrow",
	})
}
