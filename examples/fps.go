package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
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

func draw(w *gfx.Win, c gfx.Canvas) {
	playerUpdate()

	// 1.) Build perspective matrix which looks down Z axis for OpenGl NDC
	size := w.Size()
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

	gfx.Draw3DArrow(c, geom.Vec3{}, geom.Vec3{3, 0, 0}, gfx.Red, 2, view)
	gfx.Draw3DArrow(c, geom.Vec3{}, geom.Vec3{0, 3, 0}, gfx.Green, 2, view)
	gfx.Draw3DArrow(c, geom.Vec3{}, geom.Vec3{0, 0, 3}, gfx.Blue, 2, view)
}

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
