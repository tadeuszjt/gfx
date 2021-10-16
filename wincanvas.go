package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	geom "github.com/tadeuszjt/geom/32"
)

type WinCanvas struct {
	window *Win
}

func (w *WinCanvas) Clear(col Colour) {
	w.window.w2D.shader.Begin()
	defer w.window.w2D.shader.End()

	gl.ClearColor(col.R, col.G, col.B, col.A)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *WinCanvas) Size() geom.Vec2 {
	return w.window.Size()
}

func (w *WinCanvas) Draw2DVertexData(data []float32, texID *TexID, mat *geom.Mat3) {
	tex := w.window.getTexture(texID)

	w.window.w2D.shader.Begin()
	w.window.w2D.slice.Begin()
	tex.frame.Texture().Begin()
	defer tex.frame.Texture().End()
	defer w.window.w2D.slice.End()
	defer w.window.w2D.shader.End()

	if mat != nil {
		w.setMatrix2D(*mat)
	} else {
		w.setMatrix2D(geom.Mat3Identity())
	}

	w.window.w2D.slice.SetLen(len(data) / 8)
	w.window.w2D.slice.SetVertexData(data)
	//gl.Viewport(0, 0, 640, 480)
	w.window.w2D.slice.Draw()
}

func (w *WinCanvas) Draw3DVertexData(data []float32, texID *TexID, mat *geom.Mat4) {
	gl.Enable(gl.DEPTH_TEST)
	tex := w.window.getTexture(texID)

	w.window.w3D.shader.Begin()
	w.window.w3D.slice.Begin()
	tex.frame.Texture().Begin()
	defer tex.frame.Texture().End()
	defer w.window.w3D.slice.End()
	defer w.window.w3D.shader.End()

	if mat == nil {
		w.setMatrix3D(geom.Mat4Identity())
	} else {
		w.setMatrix3D(*mat)
	}

	w.window.w3D.slice.SetLen(len(data) / 9)
	w.window.w3D.slice.SetVertexData(data)
	w.window.w3D.slice.Draw()

	gl.Disable(gl.DEPTH_TEST)
}

func (w *WinCanvas) getWindow() *Win {
	return w.window
}

func (w *WinCanvas) setMatrix2D(m geom.Mat3) {
	frameSize := w.Size()
	worldToGL := geom.Mat3Camera2D(
		geom.RectOrigin(frameSize.X, frameSize.Y),
		geom.RectCentred(2, -2),
	)
	m = worldToGL.Product(m)

	w.window.w2D.shader.SetUniformAttr(0, mgl32.Mat3{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8],
	})
}

func (w *WinCanvas) setMatrix3D(m geom.Mat4) {
	w.window.w3D.shader.SetUniformAttr(0, mgl32.Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	})
}
