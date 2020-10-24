package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tadeuszjt/geom/32"
)

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

func (w *WinDraw) DrawSprite(ori geom.Ori2, rec geom.Rect, col Colour, mat geom.Mat3, tex TexID) {
	texCoords := [4]geom.Vec2{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	data := make([]float32, 0, 6*8)

	m := ori.Mat3Transform()
	verts := rec.Verts()
	for i := range verts {
		verts[i] = m.TimesVec2(verts[i], 1).Vec2()
	}

	for _, j := range [6]int{0, 1, 2, 0, 2, 3} {
		data = append(
			data,
			verts[j].X, verts[j].Y,
			texCoords[j].X, texCoords[j].Y,
			col.R, col.G, col.B, col.A,
		)
	}

	w.DrawVertexData(data, &tex, &mat)
}

func (w *WinDraw) Draw3DVertexData(data []float32, texID *TexID, model, view *geom.Mat4) {
	gl.Enable(gl.DEPTH_TEST)
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

	gl.Disable(gl.DEPTH_TEST)
}

func (w *WinDraw) Draw3DArrow(start, end geom.Vec3, colour Colour, scale float32, view geom.Mat4) {
	headLength := 1. * scale
	headWidth := 0.3 * scale
	tailWidth := 0.1 * scale

	headVerts := []geom.Vec3{
		{0, 0, 0},
		{1, 1, -1},
		{1, -1, -1},
		{-1, -1, -1},
		{-1, 1, -1},
	}
	headElem := []int{0, 1, 2, 0, 2, 3, 0, 3, 4, 0, 4, 1, 1, 2, 3, 0, 1, 3}
	headData := make([]float32, 0, len(headElem)*(3+2+4))

	for _, j := range headElem {
		headData = append(
			headData,
			headVerts[j].X, headVerts[j].Y, headVerts[j].Z,
			0, 0,
			colour.R, colour.G, colour.B, colour.A,
		)
	}

	delta := end.Minus(start)
	rot := geom.Mat4RollPitchYaw(0, delta.Pitch(), delta.Yaw())

	headScale := geom.Mat4Scalar(headWidth/2, headWidth/2, headLength)
	headModel := geom.Mat4Translation(end).Product(rot).Product(headScale)
	w.Draw3DVertexData(headData, nil, &headModel, &view)

	tailVerts := []geom.Vec3{
		{1, 1, 0},
		{1, -1, 0},
		{-1, -1, 0},
		{-1, 1, 0},
		{1, 1, 1},
		{1, -1, 1},
		{-1, -1, 1},
		{-1, 1, 1},
	}
	tailElem := []int{
		0, 1, 2,
		0, 1, 4,
		0, 2, 3,
		0, 3, 4,
		1, 2, 5,
		1, 4, 5,
		2, 3, 6,
		2, 5, 6,
		3, 4, 7,
		3, 6, 7,
		4, 5, 6,
		4, 6, 7,
	}
	tailData := make([]float32, 0, len(tailElem)*(3+2+4))

	for _, j := range tailElem {
		tailData = append(
			tailData,
			tailVerts[j].X, tailVerts[j].Y, tailVerts[j].Z,
			0, 0,
			colour.R, colour.G, colour.B, colour.A,
		)
	}

	tailScale := geom.Mat4Scalar(tailWidth/2, tailWidth/2, delta.Len()-headLength)
	tailModel := geom.Mat4Translation(start).Product(rot).Product(tailScale)
	w.Draw3DVertexData(tailData, nil, &tailModel, &view)
}

func (w *WinDraw) GetFrameSize() geom.Vec2 {
	return geom.Vec2{
		w.window.GetFrameRect().Width(),
		w.window.GetFrameRect().Height(),
	}
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

func (w *WinDraw) setViewMatrix3D(m geom.Mat4) {
	w.window.w3D.shader.SetUniformAttr(0, mgl32.Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	})
}

func (w *WinDraw) setModelMatrix3D(m geom.Mat4) {
	w.window.w3D.shader.SetUniformAttr(1, mgl32.Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	})
}
