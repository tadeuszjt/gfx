package gfx

import (
	"fmt"
)

type WinConfig struct {
	Title         string
	Width, Height int
	Resizable     bool
	SetupFunc     func(*Win) error
	DrawFunc      func(*WinDraw)
	CloseFunc     func()
}

func defaultDraw(w *WinDraw) {
	w.Clear(1, 1, 1, 1)
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

	if c.SetupFunc == nil {
		c.SetupFunc = func(*Win) error { return nil }
	}

	if c.CloseFunc == nil {
		c.CloseFunc = func() {
			fmt.Println("gfx goodbye")
		}
	}
}
