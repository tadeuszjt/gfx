package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/geom"
)

type WinDraw struct {
	window *glfw.Window
	slice  *glhf.VertexSlice
}

func (w *WinDraw) Clear(r, g, b, a float32) {
	glhf.Clear(r, g, b, a)
}

func (w *WinDraw) WindowRect() geom.Rect {
	width, height := w.window.GetFramebufferSize()
	return geom.RectOrigin(float64(width), float64(height))
}

func (w *WinDraw) DrawRect(r geom.Rect) {
	w.slice.SetLen(6)
	w.slice.SetVertexData([]float32{
		-0.5, -0.5,
		0.5, -0.5,
		-0.5, 0.5,
		0.5, -0.5,
		-0.5, 0.5,
		0.5, 0.5,
	})
}
