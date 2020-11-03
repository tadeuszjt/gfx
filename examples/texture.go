package main

import (
	"github.com/tadeuszjt/gfx"
)

var (
    texID gfx.TexID
)

func setup(w *gfx.Win) error {
    var err error
    texID, err = w.LoadTextureFromFile("picture.png")

    return err
}

func draw(w *gfx.Win, c gfx.Canvas) {
	data := []float32{
		0, 0,
		0, 0,
		1, 1, 1, 1,

		100, 0,
		1, 0,
		1, 1, 1, 1,

		0, 100,
		0, 1,
		1, 1, 1, 1,
	}

	c.Draw2DVertexData(data, &texID, nil)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
        SetupFunc: setup,
		DrawFunc: draw,
		Title:    "Triangle",
	})
}
