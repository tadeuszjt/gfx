# gfx
Yet another attempt at creating a graphics API

## Example
```go
package main

import (
	"github.com/tadeuszjt/gfx"
	"github.com/tadeuszjt/geom"
)

var (
	texID gfx.TexID
)

func setup(w *gfx.Win) error {
	var err error
	texID, err = w.LoadTexture("dog.png")
	return err
}

func draw(w *gfx.WinDraw) {
	colour := gfx.Colour{0, 1, 0, 1}
	w.DrawRect(geom.RectCentered(1, 1, geom.Vec2{}), &texID, &colour)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		SetupFunc: setup,
		DrawFunc: draw,
	})
}
```
