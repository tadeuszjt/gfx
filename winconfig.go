package gfx

import (
	"github.com/tadeuszjt/geom"
)

type WinConfig struct {
	Title         string
	Width, Height int
	DrawFunc      func(*WinDraw)
}

func defaultDraw(w *WinDraw) {
	w.Clear(1, 1, 1, 1)
	w.DrawRect(geom.Rect{})
}

func (c *WinConfig) loadDefaults() {
	if c.Title == "" {
		c.Title = "Gfx"
	}

	if c.Width == 0 {
		c.Width = 640
	}

	if c.Height == 0 {
		c.Height = 480
	}

	if c.DrawFunc == nil {
		c.DrawFunc = defaultDraw
	}
}
