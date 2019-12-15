package main

import (
	"github.com/tadeuszjt/geom/32"
	"github.com/tadeuszjt/gfx"
)

var (
	corners = []geom.Vec3{
		{-1, -1, -1},
		{1, -1, -1},
		{1, 1, -1},
		{-1, 1, -1},
		{-1, -1, 1},
		{1, -1, 1},
		{1, 1, 1},
		{-1, 1, 1},
	}

	colours = []gfx.Colour{
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{0, 1, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
	}

	angleX, angleY geom.Angle
)

func draw(w *gfx.WinDraw) {
	data := make([]float32, 0, 9*2*6)
	elem := []int{
		0, 1, 2,
		0, 1, 5,
		0, 2, 3,
		0, 3, 7,
		0, 4, 5,
		0, 4, 7,
		2, 1, 5,
		2, 3, 7,
		2, 5, 6,
		2, 6, 7,
		5, 4, 7,
		5, 6, 7,
	}

	for _, i := range elem {
		data = append(data,
			corners[i].X, corners[i].Y, corners[i].Z,
			0, 0,
			colours[i].R, colours[i].G, colours[i].B, colours[i].A)
	}

	size := w.GetFrameSize()
	p := geom.Mat4Perspective(geom.RectCentred(size.X, size.Y), 0, 0, 75, 0.1, 100)

	t := geom.Mat4Translation(geom.Vec3{0, 0, 3})
	rx := geom.Mat4RotationX(angleX)
	ry := geom.Mat4RotationY(angleY)
	m := t.Product(rx).Product(ry)

	w.Draw3DVertexData(data, nil, &m, &p)

	angleX = angleX.Plus(geom.Angle(0.02))
	angleY = angleY.Plus(geom.Angle(0.03))
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		DrawFunc: draw,
		Title:    "Cube",
	})
}
