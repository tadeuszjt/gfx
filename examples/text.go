package main

import (
	"github.com/tadeuszjt/geom/32"
	"github.com/tadeuszjt/gfx"
)

var (
    text gfx.Text
    tex gfx.TexID
)

func setup(w *gfx.Win) error {
	text.SetSize(31)
	text.SetString("benis")
    tex = w.LoadTextureBlank(100, 100)
	return nil
}

func draw(w *gfx.WinCanvas) {
    tc := w.GetTextureCanvas(tex)
    tc.Clear(gfx.Red)
    gfx.DrawSprite(tc, geom.Ori2{}, geom.RectOrigin(10, 10), gfx.Green, nil, nil)
    gfx.DrawText(tc, &text, geom.Vec2{})
    gfx.DrawSprite(w, geom.Ori2{}, geom.RectOrigin(100, 100), gfx.White, nil, &tex)
    gfx.DrawSprite(w, geom.Ori2{100, 100, 0}, geom.RectOrigin(100, 100), gfx.White, nil, &tex)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		SetupFunc: setup,
		DrawFunc:  draw,
		Title:     "Triangle",
	})
}
