package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tadeuszjt/geom/generic"
)

type MouseEvent interface {
}

type MouseScroll struct {
	Dx, Dy float32
}

type MouseMove struct {
	Position geom.Vec2[float32]
}

type MouseButton struct {
	Button glfw.MouseButton
	Action glfw.Action
	Mods   glfw.ModifierKey
}

type KeyEvent struct {
	Key    glfw.Key
	Action glfw.Action
	Mods   glfw.ModifierKey
}

func (w *Win) setupInput(c *WinConfig) {
	w.glfwWin.SetFramebufferSizeCallback(
		func(_ *glfw.Window, width, height int) {
			gl.Viewport(0, 0, int32(width), int32(height))
			c.ResizeFunc(w)
		})

	w.glfwWin.SetCursorPosCallback(
		func(_ *glfw.Window, xpos, ypos float64) {
			c.MouseFunc(w, MouseMove{geom.Vec2[float32]{float32(xpos), float32(ypos)}})
		})

	w.glfwWin.SetScrollCallback(
		func(_ *glfw.Window, dx, dy float64) {
			c.MouseFunc(w, MouseScroll{float32(dx), float32(dy)})
		})

	w.glfwWin.SetMouseButtonCallback(
		func(_ *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
			c.MouseFunc(w, MouseButton{button, action, mods})
		})

	w.glfwWin.SetKeyCallback(
		func(_ *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			c.KeyFunc(w, KeyEvent{key, action, mods})

		})

}
