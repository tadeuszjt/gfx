package main

import (
	//"math"
	"github.com/tadeuszjt/data"
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
)

const (
	boidsCount       = 200
	boidsSpeed       = 0.8
	boidsSightRadius = 10
	boidsAvoid       = 0.1
	boidsAlign       = 0.01
	boidsCohere      = 0.01
)

var (
	arena = geom.CuboidCentred[float32](160, 160, 160)

	boids struct {
		data.Table
		positions  data.RowT[geom.Vec3[float32]]
		directions data.RowT[geom.Vec3[float32]]
		colours    data.RowT[gfx.Colour]
	}
)

func init() {
	boids.Table = data.Table{&boids.positions, &boids.directions, &boids.colours}
}

func updateBoids() {

	for i, iPos := range boids.positions {

		jPositions := make([]geom.Vec3[float32], 0, 4)

		for j, jPos := range boids.positions {
			if i == j {
				continue
			}
			if jPos.Minus(iPos).Len2() > boidsSightRadius*boidsSightRadius {
				continue
			}
			jPositions = append(jPositions, jPos)
		}

		var avoid, align, deltaAvg geom.Vec3[float32]

		for j, jPos := range jPositions {
			delta := jPos.Minus(iPos)
			delta01 := delta.ScaledBy(1. / boidsSightRadius)

			avoid.PlusEquals(delta01.Normal().ScaledBy(boidsAvoid * -(1./delta01.Len() - 1)))
			align.PlusEquals(boids.directions[j].ScaledBy(boidsAlign))
			deltaAvg.PlusEquals(delta)
		}

		if len(jPositions) > 0 {
			deltaAvg = deltaAvg.ScaledBy(1. / float32(len(jPositions)))
			cohere := deltaAvg.ScaledBy(boidsCohere)

			boids.directions[i].PlusEquals(avoid)
			boids.directions[i].PlusEquals(align)
			boids.directions[i].PlusEquals(cohere)
			boids.directions[i] = boids.directions[i].Normal()
		}
	}

	for i, pos := range boids.positions {
		if pos.X > arena.Max.X {
			boids.positions[i].X = arena.Max.X
			boids.directions[i].X *= -1
		} else if pos.X < arena.Min.X {
			boids.positions[i].X = arena.Min.X
			boids.directions[i].X *= -1
		}
		if pos.Y > arena.Max.Y {
			boids.positions[i].Y = arena.Max.Y
			boids.directions[i].Y *= -1
		} else if pos.Y < arena.Min.Y {
			boids.positions[i].Y = arena.Min.Y
			boids.directions[i].Y *= -1
		}
		if pos.Z > arena.Max.Z {
			boids.positions[i].Z = arena.Max.Z
			boids.directions[i].Z *= -1
		} else if pos.Z < arena.Min.Z {
			boids.positions[i].Z = arena.Min.Z
			boids.directions[i].Z *= -1
		}

		boids.positions[i].PlusEquals(boids.directions[i].ScaledBy(boidsSpeed))
	}

}

func spawnBoids() {
	boids.Append(geom.Vec3[float32]{}, geom.Vec3[float32]{1, 0, 0}, gfx.Red)

	for i := 0; i < boidsCount; i++ {
		pos := geom.Vec3Rand(arena)
		dir := geom.Vec3NormRand[float32]()
		boids.Append(pos, dir, gfx.ColourRand())
	}
}

func drawBoids(w gfx.Canvas, view geom.Mat4[float32]) {
	verts := [5]geom.Vec3[float32]{
		{0, 0, 3},
		{1, 1, -1},
		{-1, 1, -1},
		{-1, -1, -1},
		{1, -1, -1},
	}

	elem := [6 * 3]int{
		0, 1, 2,
		0, 2, 3,
		0, 3, 4,
		0, 1, 4,
		1, 2, 3,
		1, 3, 4,
	}

	for i := range boids.positions {
		colours := []gfx.Colour{
			gfx.Black,
			boids.colours[i],
			boids.colours[i],
			boids.colours[i],
			boids.colours[i],
		}

		data := make([]float32, 0, len(elem)*(3+2+4))
		for _, j := range elem {
			data = append(
				data,
				verts[j].X, verts[j].Y, verts[j].Z,
				0, 0,
				colours[j].R, colours[j].G, colours[j].B, colours[j].A,
			)
		}

		t := geom.Mat4Translation(boids.positions[i])
		r := geom.Mat4RollPitchYaw(0, boids.directions[i].Pitch(), boids.directions[i].Yaw())

		model := t.Product(r)
		mat := view.Product(model)
		w.Draw3DVertexData(data, nil, &mat)
	}
}
