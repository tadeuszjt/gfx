package gfx

import (
	"github.com/faiface/glhf"
	"github.com/tadeuszjt/geom"
)

type WinDraw struct {
	window *Win
	slice  *glhf.VertexSlice
}

func (w *WinDraw) Clear(r, g, b, a float32) {
	glhf.Clear(r, g, b, a)
}

func (w *WinDraw) DrawRect(r geom.Rect, texID *TexID) {
	var tex *glhf.Texture
	if texID != nil {
		tex = w.window.textures[*texID]
	} else {
		tex = w.window.whiteTex
	}
	
	tex.Begin()
	w.slice.SetLen(6)
	w.slice.SetVertexData([]float32{
		float32(r.Min.X), float32(r.Min.Y),
		0, 0,
		float32(r.Max.X), float32(r.Min.Y),
		1, 0,
		float32(r.Min.X), float32(r.Max.Y),
		0, 1,
		float32(r.Min.X), float32(r.Max.Y),
		0, 1,
		float32(r.Max.X), float32(r.Min.Y),
		1, 0,
		float32(r.Max.X), float32(r.Max.Y),
		1, 1,
	})
	w.slice.Draw()
	tex.End()
}
