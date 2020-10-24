package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/geom/32"
	"github.com/tadeuszjt/gfx"
)

func playerUpdate() {
	forward := geom.Vec3NormPitchYaw(player.pitch, player.bearing)
	right := geom.Vec3NormPitchYaw(0, player.bearing.Plus(geom.Angle90Deg))

	if keys.w {
		player.position.PlusEquals(forward.ScaledBy(playerSpeed))
	}
	if keys.a {
		player.position.MinusEquals(right.ScaledBy(playerSpeed))
	}
	if keys.s {
		player.position.MinusEquals(forward.ScaledBy(playerSpeed))
	}
	if keys.d {
		player.position.PlusEquals(right.ScaledBy(playerSpeed))
	}
}

const (
	playerSpeed           = 0.4
	playerLookSensitivity = 0.015
)

var (
	mouseWin geom.Vec2
	player   = struct {
		position geom.Vec3
		bearing  geom.Angle
		pitch    geom.Angle
	}{
		position: geom.Vec3{0, 0, -10},
		pitch:    0,
		bearing:  0,
	}

	keys struct{ w, a, s, d bool }
)

func setup(w *gfx.Win) error {
	w.GetGlfwWindow().SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	w.GetGlfwWindow().SetCursorPos(0, 0)
	fmt.Println("running")
	return nil
}

func draw(w *gfx.WinDraw) {
	playerUpdate()

	// 1.) Build perspective matrix which looks down Z axis for OpenGl NDC
	size := w.GetFrameSize()
	ar := size.X / size.Y
	near := float32(0.1)
	perspective := geom.Mat4Perspective(-ar*near, ar*near, -near, near, near, 1000).Product(geom.Mat4Scalar(-1, 1, -1))

	// 2.) Translate 'world by opposite of player position
	translation := geom.Mat4Translation(player.position.ScaledBy(-1))

	// 2.) Rotate world by opposite of pitch and yaw
	rx := geom.Mat4RotationX(-player.pitch)
	ry := geom.Mat4RotationY(-player.bearing)

	// 3.) Sequence transformations
	view := perspective.Product(rx).Product(ry).Product(translation)

	verts := [5]geom.Vec3{
		{3, 0, 0},
		{-1, 1, 1},
		{-1, 1, -1},
		{-1, -1, -1},
		{-1, -1, 1},
	}

	elem := [6 * 3]int{
		0, 1, 2,
		0, 2, 3,
		0, 3, 4,
		0, 1, 4,
		1, 2, 3,
		1, 3, 4,
	}

	arrows := []struct {
		model    geom.Mat4
		colour   gfx.Colour
		position geom.Vec3
	}{
		{geom.Mat4Identity(), gfx.Red, geom.Vec3{3, 0, 0}},
		{geom.Mat4RotationZ(geom.Angle90Deg), gfx.Green, geom.Vec3{0, 3, 0}},
		{geom.Mat4RotationY(-geom.Angle90Deg), gfx.Blue, geom.Vec3{0, 0, 3}},
	}

	for i := range arrows {
		data := make([]float32, 0, len(elem)*(3+2+4))

		for _, j := range elem {
			colour := arrows[i].colour

			data = append(
				data,
				verts[j].X, verts[j].Y, verts[j].Z,
				0, 0,
				colour.R, colour.G, colour.B, colour.A,
			)
		}
		translation := geom.Mat4Translation(arrows[i].position)
		m := translation.Product(arrows[i].model)
		w.Draw3DVertexData(data, nil, &m, &view)
	}
}

var first = true

func mouse(w *gfx.Win, ev gfx.MouseEvent) {
	switch e := ev.(type) {
	case gfx.MouseScroll:
	case gfx.MouseMove:
		{
			player.bearing.PlusEquals(geom.MakeAngle(e.Position.X * playerLookSensitivity))
			player.pitch.MinusEquals(geom.MakeAngle(e.Position.Y * playerLookSensitivity))
			player.pitch.Clamp(-geom.Angle90Deg, geom.Angle90Deg)
			w.GetGlfwWindow().SetCursorPos(0, 0)
		}
	case gfx.MouseButton:
	}
}

func keyboard(w *gfx.Win, ev gfx.KeyEvent) {
	switch ev.Key {
	case glfw.KeyW:
		{
			if ev.Action == glfw.Press {
				keys.w = true
			} else if ev.Action == glfw.Release {
				keys.w = false
			}
		}

	case glfw.KeyA:
		{
			if ev.Action == glfw.Press {
				keys.a = true
			} else if ev.Action == glfw.Release {
				keys.a = false
			}
		}

	case glfw.KeyS:
		{
			if ev.Action == glfw.Press {
				keys.s = true
			} else if ev.Action == glfw.Release {
				keys.s = false
			}
		}

	case glfw.KeyD:
		{
			if ev.Action == glfw.Press {
				keys.d = true
			} else if ev.Action == glfw.Release {
				keys.d = false
			}
		}
	}
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		DrawFunc:  draw,
		MouseFunc: mouse,
		SetupFunc: setup,
		KeyFunc:   keyboard,
		Title:     "Boids",
		Width:     1024,
		Height:    768,
	})
}
