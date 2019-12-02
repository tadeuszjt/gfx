package gfx

import (
	"os"
	"fmt"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tadeuszjt/geom"
)

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
		
		if winConfig.Resisable {
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
		
		glfwWin.MakeContextCurrent()
		glhf.Init()
		
		gl.Enable(gl.BLEND)
		glhf.BlendFunc(glhf.SrcAlpha, glhf.OneMinusSrcAlpha)
	})
	
	var (
		win    Win
		slice  *glhf.VertexSlice
		shader *glhf.Shader
	)
	
	mainthread.Call(func() {
		var err error
		shader, err = newShader(&shader2D)
		if err != nil {
			panic(err)
		}
		
		win.loadWhiteTex()
		winConfig.SetupFunc(&win)
		
		slice = glhf.MakeVertexSlice(shader, 0, 0)
	})
	
	fmt.Println(win, slice, shader)
	
	shouldQuit := false
	for !shouldQuit {
		mainthread.Call(func() {
			if glfwWin.ShouldClose() {
				shouldQuit = true
			}

			shader.Begin()
			shader.SetUniformAttr(0, Mat3GeomToMgl32(geom.Mat3{
				1, 0, 0,
				0, -1, 0,
				0, 0, 1,
			}))
			slice.Begin()
			glhf.Clear(1, 1, 1, 1)
			winConfig.DrawFunc(&WinDraw{window: &win, slice: slice})
			slice.End()
			shader.End()

			glfwWin.SwapBuffers()
			glfw.PollEvents()
		})
	}
	
	winConfig.CloseFunc()
}

func RunWindow(config WinConfig) {
	winConfig = config
	winConfig.loadDefaults()
	mainthread.Run(run)
}


func Mat3GeomToMgl32(m geom.Mat3) mgl32.Mat3 {
	return mgl32.Mat3{
		float32(m[0]), float32(m[3]), float32(m[6]),
		float32(m[1]), float32(m[4]), float32(m[7]),
		float32(m[2]), float32(m[5]), float32(m[8]),
	}
}
