package gfx

import (
	"fmt"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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

	var (
		win Win
		err error
	)

	mainthread.Call(func() {
		glfw.Init()

		err = win.createGlfwWindow(winConfig)
		if err != nil {
			return
		}

		/* OpenGL context setup */
		glhf.Init()
		win.makeContextCurrent()
		gl.Enable(gl.BLEND)
		glhf.BlendFunc(glhf.SrcAlpha, glhf.OneMinusSrcAlpha)
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Gfx Error:", err)
		return
	}

	mainthread.Call(func() {
		win.textInit()
		
		err = win.setup(winConfig)
		if err != nil {
			return
		}

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
			if win.glfwWin.ShouldClose() {
				shouldQuit = true
			}

			winDraw := WinDraw{window: &win}
			winDraw.begin()
			winConfig.DrawFunc(&winDraw)
			winDraw.end()

			win.glfwWin.SwapBuffers()
			glfw.PollEvents()
		})
	}

	winConfig.CloseFunc()
}
