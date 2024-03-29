package gfx

import (
	"fmt"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"os"
)

var winConfig WinConfig

func RunWindow(config WinConfig) {
	winConfig = config
	winConfig.loadDefaults()
	mainthread.Run(run)
}

func run() {
	defer func() {
		glfw.Terminate()
	}()

	var (
		win *Win = &Win{}
		err error
	)

	mainthread.Call(func() {
		glfw.Init()

		err = win.createGlfwWindow(winConfig)
		if err != nil {
			return
		}

		/* OpenGL context setup */
		win.makeContextCurrent()
		glfw.SwapInterval(1) // vsync enabled
		fmt.Println("vsync set enabled")
		glhf.Init()
		gl.Enable(gl.BLEND)
		glhf.BlendFunc(glhf.SrcAlpha, glhf.OneMinusSrcAlpha)
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Gfx Error:", err)
		return
	}

	mainthread.Call(func() {
		err = win.setup(winConfig)
		if err != nil {
			return
		}
		winConfig.ResizeFunc(win)
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Gfx Error:", err)
		return
	}

	shouldQuit := false
	for !shouldQuit {
		mainthread.Call(func() {
			if win.glfwWin.ShouldClose() {
				shouldQuit = true
			}

			c := &WinCanvas{window: win}
			c.Clear(White)
			winConfig.DrawFunc(win, c)

			win.glfwWin.SwapBuffers()
			glfw.PollEvents()
		})
	}

	winConfig.CloseFunc()
}
