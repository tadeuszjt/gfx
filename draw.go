package gfx

import (
	"github.com/tadeuszjt/geom"
	"github.com/faiface/glhf"
)

type Colour struct {
	R, G, B, A float32
}

type WinDraw struct {
	window *Win
	slice  *glhf.VertexSlice
}

func (w *WinDraw) Clear(r, g, b, a float32) {
	glhf.Clear(r, g, b, a)
}

func (w *WinDraw) DrawRect(r geom.Rect, texID *TexID, colour *Colour) {
	tex := w.window.whiteTex
	if texID != nil {
		tex = w.window.textures[*texID]
	}
	
	col := Colour{1, 1, 1, 1}
	if colour != nil {
		col = *colour
	}
	
	coords := [4][2]float32{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}
	
	verts := [4][2]float32{
		{float32(r.Min.X), float32(r.Min.Y)},
		{float32(r.Max.X), float32(r.Min.Y)},
		{float32(r.Max.X), float32(r.Max.Y)},
		{float32(r.Min.X), float32(r.Max.Y)},
	}
	
	index := [6]int{0, 1, 2, 0, 2, 3}
	
	data := make([]float32, 0, 6*8)
	
	for i := range index {
		idx := index[i]
		data = append(data,
			verts[idx][0], verts[idx][1],
			coords[idx][0], coords[idx][1],
			col.R, col.G, col.B, col.A,
		)
	}
	
	tex.Begin()
	w.slice.SetLen(6)
	w.slice.SetVertexData(data)
	w.slice.Draw()
	tex.End()
}
