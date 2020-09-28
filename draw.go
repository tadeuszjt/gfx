package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tadeuszjt/geom/32"
)

type Colour struct {
	R, G, B, A float32
}

type WinDraw struct {
	window *Win
}

func (w *WinDraw) Clear(r, g, b, a float32) {
	w.window.w2D.shader.Begin()
	defer w.window.w2D.shader.End()

	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *WinDraw) DrawVertexData(data []float32, texID *TexID, mat *geom.Mat3) {
	id := w.window.whiteTexID
	if texID != nil {
		id = *texID
	}
	tex := w.window.textures[id]

	w.window.w2D.shader.Begin()
	w.window.w2D.slice.Begin()
	tex.Begin()
	defer tex.End()
	defer w.window.w2D.slice.End()
	defer w.window.w2D.shader.End()

	if mat != nil {
		w.setMatrix2D(*mat)
	} else {
		w.setMatrix2D(geom.Mat3Identity())
	}

	w.window.w2D.slice.SetLen(len(data) / 8)
	w.window.w2D.slice.SetVertexData(data)
	w.window.w2D.slice.Draw()
}

func (w *WinDraw) Draw3DVertexData(data []float32, texID *TexID, model, view *geom.Mat4) {
	id := w.window.whiteTexID
	if texID != nil {
		id = *texID
	}
	tex := w.window.textures[id]

	w.window.w3D.shader.Begin()
	w.window.w3D.slice.Begin()
	tex.Begin()
	defer tex.End()
	defer w.window.w3D.slice.End()
	defer w.window.w3D.shader.End()

	if model == nil {
		w.setModelMatrix3D(geom.Mat4Identity())
	} else {
		w.setModelMatrix3D(*model)
	}

	if view == nil {
		w.setViewMatrix3D(geom.Mat4Identity())
	} else {
		w.setViewMatrix3D(*view)
	}

	w.window.w3D.slice.SetLen(len(data) / 9)
	w.window.w3D.slice.SetVertexData(data)
	w.window.w3D.slice.Draw()
}

func (w *WinDraw) GetFrameSize() geom.Vec2 {
	return w.window.GetFrameSize()
}

func (w *WinDraw) setMatrix2D(m geom.Mat3) {
	frameSize := w.GetFrameSize()
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

func geomToMgl32Mat4(m geom.Mat4) mgl32.Mat4 {
	return mgl32.Mat4{
		m[0], m[1], m[2], m[3],
		m[4], m[5], m[6], m[7],
		m[8], m[9], m[10], m[11],
		m[12], m[13], m[14], m[15],
	}
}

func (w *WinDraw) setViewMatrix3D(m geom.Mat4) {
	w.window.w3D.shader.SetUniformAttr(0, geomToMgl32Mat4(m))
}

func (w *WinDraw) setModelMatrix3D(m geom.Mat4) {
	w.window.w3D.shader.SetUniformAttr(1, geomToMgl32Mat4(m))
}
