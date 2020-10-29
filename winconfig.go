package gfx

import (
	"fmt"
)

type WinConfig struct {
	Title         string
	Width, Height int
	Resizable     bool
	SetupFunc     func(*Win) error
	DrawFunc      func(*WinCanvas)
	CloseFunc     func()
	MouseFunc     func(*Win, MouseEvent)
	KeyFunc       func(*Win, KeyEvent)
	ResizeFunc    func(width, height int)
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
		c.DrawFunc = func(w *WinCanvas) { w.Clear(White) }
	}

	if c.SetupFunc == nil {
		c.SetupFunc = func(*Win) error { return nil }
	}

	if c.CloseFunc == nil {
		c.CloseFunc = func() { fmt.Println("gfx goodbye") }
	}

	if c.MouseFunc == nil {
		c.MouseFunc = func(*Win, MouseEvent) {}
	}

	if c.KeyFunc == nil {
		c.KeyFunc = func(*Win, KeyEvent) {}
	}

	if c.ResizeFunc == nil {
		c.ResizeFunc = func(int, int) {}
	}
}
