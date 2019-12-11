package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/tadeuszjt/geom/32"
)

/* Creates the internal glfw window.
 * Must be called after glfw.Init()
 */
func (w *Win) createGlfwWindow(c WinConfig) error {
	resizable := glfw.False
	if c.Resizable {
		resizable = glfw.True
	}
	
	glfw.WindowHint(glfw.Resizable, resizable)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	w.glfwWin, err = glfw.CreateWindow(c.Width, c.Height, c.Title, nil, nil)
	return err
}

/* Called after createGlfwWindow to activate context
 */
func (w *Win) makeContextCurrent() {
	w.glfwWin.MakeContextCurrent()
}

/* Window setup after OpenGL initialised and makeContextCurrent()
 */
func (w *Win) setup(c WinConfig) error {
	/* setup callbacks */
	w.glfwWin.SetFramebufferSizeCallback(
		func(_ *glfw.Window, width, height int) {
			gl.Viewport(0, 0, int32(width), int32(height))
			c.ResizeFunc(width, height)
		})
	
	w.glfwWin.SetCursorPosCallback(
		func(_ *glfw.Window, xpos, ypos float64) {
			c.MouseFunc(w, MouseMove{geom.Vec2{float32(xpos), float32(ypos)}})
		})
	
	w.glfwWin.SetScrollCallback(
		func(_ *glfw.Window, dx, dy float64) {
			c.MouseFunc(w, MouseScroll{float32(dx), float32(dy)})
		})
	
	w.glfwWin.SetMouseButtonCallback(
		func(_ *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
			c.MouseFunc(w, MouseButton{
				button,
				action,
				mods,
			})
		})
		
	/* load default white texture into slot 0 */
	w.textures = append(w.textures, glhf.NewTexture(1, 1, false, []uint8{255, 255, 255, 255}))

	/* load shader */
	var err error
	w.shader, err = newShader(&shader2D)
	if err != nil {
		return err
	}
	
	/* create slice */
	w.slice = glhf.MakeVertexSlice(w.shader, 0, 0)
	
	/* call user setup function */
	return c.SetupFunc(w)
}
