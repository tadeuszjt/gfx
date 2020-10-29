package gfx

import (
	"github.com/tadeuszjt/geom/32"
)

type Canvas interface {
	Clear(col Colour)
	//Size() geom.Vec2
	//DrawOn(c Canvas, mat *geom.Mat3)
	Draw2DVertexData(data []float32, texID *TexID, mat *geom.Mat3)
	Draw3DVertexData(data []float32, texID *TexID, mat *geom.Mat4)
}

func Draw3DArrow(c Canvas, start, end geom.Vec3, colour Colour, scale float32, view geom.Mat4) {
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
	headMat := view.Product(headModel)
	c.Draw3DVertexData(headData, nil, &headMat)

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
	tailMat := view.Product(tailModel)
	c.Draw3DVertexData(tailData, nil, &tailMat)
}

func DrawSprite(c Canvas, ori geom.Ori2, rec geom.Rect, col Colour, mat geom.Mat3, tex TexID) {
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

	c.Draw2DVertexData(data, &tex, &mat)
}
