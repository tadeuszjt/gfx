package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
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

	mouseWin geom.Vec2
)

func setup(w *gfx.Win) error {
	w.GetGlfwWindow().SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	return nil
}

func draw(w *gfx.WinDraw) {
    playerUpdate()

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

	t := geom.Mat4Translation(player.position)
	a := geom.Mat4RotationY(player.bearing)
	b := geom.Mat4RotationX(player.pitch)
	p := geom.Mat4Perspective(size.X / size.Y, 1, -1, 90, 0.1, 100)
	view := p.Product(b).Product(a).Product(t)

	rx := geom.Mat4RotationX(angleX)
	ry := geom.Mat4RotationY(angleY)
	model := rx.Product(ry)

	w.Draw3DVertexData(data, nil, &model, &view)

	angleX.PlusEquals(geom.Angle(0.02))
	angleY.PlusEquals(geom.Angle(0.03))
}

func mouse(w *gfx.Win, ev gfx.MouseEvent) {
	switch e := ev.(type) {
	case gfx.MouseScroll:
	case gfx.MouseMove:
		{
            playerLook(e.Position.X, -e.Position.Y)
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
		Title:     "Cube",
        Width:     1024,
        Height:    768,
	})
}
