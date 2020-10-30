package gfx

import (
	"fmt"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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
		win *Win
		err error
	)

	mainthread.Call(func() {
		glfw.Init()

        win = &Win{}

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
		err = win.setup(winConfig)
		if err != nil {
			return
		}

		frame := win.GetFrameRect()
		winConfig.ResizeFunc(int(frame.Width()), int(frame.Height()))
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Gfx Error:", err)
		return
	}

	shouldQuit := false
	winDraw := WinCanvas{window: win}

	for !shouldQuit {
		mainthread.Call(func() {
			if win.glfwWin.ShouldClose() {
				shouldQuit = true
			}

			winDraw.Clear(White)
			winConfig.DrawFunc(&winDraw)

			win.glfwWin.SwapBuffers()
			glfw.PollEvents()
		})
	}

	winConfig.CloseFunc()
}
