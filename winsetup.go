package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/glfw/v3.2/glfw"
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
	w.setupText()
	w.setupInput(&c)

	/* load default white texture */
	w.whiteTexID = w.loadTextureFromPixels(1, 1, false, []uint8{255, 255, 255, 255})

	/* load shaders */
	var err error
	w.w2D.shader, err = newShader(&shader2D)
	if err != nil {
		return err
	}

	w.w3D.shader, err = newShader(&shader3D)
	if err != nil {
		return err
	}

	/* create slice */
	w.w2D.slice = glhf.MakeVertexSlice(w.w2D.shader, 0, 0)
	w.w3D.slice = glhf.MakeVertexSlice(w.w3D.shader, 0, 0)

	/* call user setup function */
	return c.SetupFunc(w)
}
