package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tadeuszjt/geom/32"
)

type Colour struct {
	R, G, B, A float32
}

type WinDraw struct {
	window        *Win
	activeTexture *glhf.Texture
}

func (w *WinDraw) Clear(r, g, b, a float32) {
	glhf.Clear(r, g, b, a)
}

func (w *WinDraw) DrawVertexData(data []float32, texID *TexID, mat *geom.Mat3) {
	tex := TexID(0)
	if texID != nil {
		tex = *texID
	}
	
	if mat != nil {
		w.setMatrix(*mat)
	} else {
		w.setMatrix(geom.Mat3Identity())
	}

	w.setActiveTexture(tex)
	w.window.slice.SetLen(len(data) / 8)
	w.window.slice.SetVertexData(data)
	w.window.slice.Draw()
}

func (w *WinDraw) DrawRect(r geom.Rect, texID *TexID, colour *Colour, mat *geom.Mat3) {
	col := Colour{1, 1, 1, 1}
	if colour != nil {
		col = *colour
	}

	texCoords := [4]geom.Vec2{{0, 0}, {1, 0}, {1, 1}, {0, 1}}

	verts := [4]geom.Vec2{
		r.Min,
		{r.Max.X, r.Min.Y},
		r.Max,
		{r.Min.X, r.Max.Y},
	}

	data := make([]float32, 0, 6*8)

	for _, i := range [6]int{0, 1, 2, 0, 2, 3} {
		data = append(data,
			verts[i].X, verts[i].Y,
			texCoords[i].X, texCoords[i].Y,
			col.R, col.G, col.B, col.A,
		)
	}

	w.DrawVertexData(data, texID, mat)
}

func (w *WinDraw) GetFrameSize() geom.Vec2 {
	return w.window.GetFrameSize()
}

func (w *WinDraw) setMatrix(m geom.Mat3) {
	frameSize := w.GetFrameSize()

	worldToGL := geom.Mat3Camera2D(
		geom.RectOrigin(frameSize.X, frameSize.Y),
		geom.RectCentred(2, -2),
	)

	m = worldToGL.Product(m)

	w.window.shader.SetUniformAttr(0, mgl32.Mat3{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8],
	})
}

func (w *WinDraw) setActiveTexture(tex TexID) {
	if w.window.textures[tex] != w.activeTexture {
		if w.activeTexture != nil {
			w.activeTexture.End()
		}

		w.activeTexture = w.window.textures[tex]
		w.activeTexture.Begin()
	}
}

func (w *WinDraw) begin() {
	w.window.slice.Begin()
	w.window.shader.Begin()
	w.setMatrix(geom.Mat3Identity())
	glhf.Clear(1, 1, 1, 1)
}

func (w *WinDraw) end() {
	w.window.slice.End()
	w.window.shader.End()

	if w.activeTexture != nil {
		w.activeTexture.End()
	}
}
