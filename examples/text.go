package main

import (
	"github.com/tadeuszjt/geom/32"
	"github.com/tadeuszjt/gfx"
)

var text gfx.Text

func setup(w *gfx.Win) error {
	text.SetSize(31)
	text.SetString("benis")
	return nil
}

func draw(w *gfx.WinDraw) {
	w.DrawText(&text, geom.Vec2{})
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		SetupFunc: setup,
		DrawFunc:  draw,
		Title:     "Triangle",
	})
}
