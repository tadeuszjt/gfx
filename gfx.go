package gfx

import (
	"fmt"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/geom/geom32"
	"os"
)

func RunWindow(config WinConfig) {
	winConfig = config
	winConfig.loadDefaults()
	mainthread.Run(run)
}

var (
	winConfig WinConfig
)

func run() {
	defer func() {
		glfw.Terminate()
	}()

	var glfwWin *glfw.Window

	mainthread.Call(func() {
		glfw.Init()

		if winConfig.Resizable {
			glfw.WindowHint(glfw.Resizable, glfw.True)
		} else {
			glfw.WindowHint(glfw.Resizable, glfw.False)
		}
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

		var err error
		glfwWin, err = glfw.CreateWindow(
			winConfig.Width, winConfig.Height, winConfig.Title, nil, nil,
		)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		glhf.Init()

		glfwWin.MakeContextCurrent()
		
		glfwWin.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
			gl.Viewport(0, 0, int32(width), int32(height))
			winConfig.ResizeFunc(width, height)
		})

		gl.Enable(gl.BLEND)
		glhf.BlendFunc(glhf.SrcAlpha, glhf.OneMinusSrcAlpha)
	})

	var (
		win    = Win{glfwWin: glfwWin}
		slice  *glhf.VertexSlice
		shader *glhf.Shader
		err    error
	)

	mainthread.Call(func() {
		shader, err = newShader(&shader2D)
		if err != nil {
			return
		}

		win.loadWhiteTex()
		err = winConfig.SetupFunc(&win)
		if err != nil {
			return
		}
		
		glfwWin.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
			winConfig.MouseFunc(&win, MouseMove{geom.Vec2{float32(xpos), float32(ypos)}})
		})
		
		glfwWin.SetScrollCallback(func(w *glfw.Window, dx, dy float64) {
			winConfig.MouseFunc(&win, MouseScroll{float32(dx), float32(dy)})
		})
		
		glfwWin.SetMouseButtonCallback(func(
			w *glfw.Window, 
			button glfw.MouseButton,
			action glfw.Action,
			mods glfw.ModifierKey,
		) {
			winConfig.MouseFunc(&win, MouseButton{})
		})

		slice = glhf.MakeVertexSlice(shader, 0, 0)
		
		size := win.GetFrameSize()
		winConfig.ResizeFunc(int(size.X), int(size.Y))
	})
	
	if err != nil {
		fmt.Fprintln(os.Stderr, "Gfx Error:", err)
		return
	}

	shouldQuit := false
	for !shouldQuit {
		mainthread.Call(func() {
			if glfwWin.ShouldClose() {
				shouldQuit = true
			}

			winDraw := makeWinDraw(slice, shader, &win)
			winDraw.begin()
			winConfig.DrawFunc(&winDraw)
			winDraw.end()

			glfwWin.SwapBuffers()
			glfw.PollEvents()
		})
	}

	winConfig.CloseFunc()
}
