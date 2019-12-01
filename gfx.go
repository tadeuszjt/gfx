package gfx

import (
	"fmt"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tadeuszjt/geom"
	"os"
)

func Mat3GeomToMgl32(m geom.Mat3) mgl32.Mat3 {
	return mgl32.Mat3{
		float32(m[0]), float32(m[3]), float32(m[6]),
		float32(m[1]), float32(m[4]), float32(m[7]),
		float32(m[2]), float32(m[5]), float32(m[8]),
	}
}

func RunWindow(config WinConfig) {
	mainthread.Run(func() {
		defer func() {
			mainthread.Call(func() {
				glfw.Terminate()
			})
		}()

		config.loadDefaults()

		glfw.Init()

		glfw.WindowHint(glfw.Resizable, glfw.False)
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

		win, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		win.SetFramebufferSizeCallback(func(_ *glfw.Window, width, height int) {
			gl.Viewport(0, 0, int32(width), int32(height))
		})

		win.MakeContextCurrent()
		
		
		glhf.Init()
		
		vertexFmt := glhf.AttrFormat{
			{Name: "position", Type: glhf.Vec2},
		}
		uniformFmt := glhf.AttrFormat{{Name: "matrix", Type: glhf.Mat3}}
		shader, err := glhf.NewShader(vertexFmt, uniformFmt, shader2D.vertex, shader2D.fragment)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		
		slice := glhf.MakeVertexSlice(shader, 0, 0)

		for !win.ShouldClose() {
			glfw.PollEvents()
			
			shader.Begin()
			shader.SetUniformAttr(0, Mat3GeomToMgl32(geom.Mat3{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			}))
			slice.Begin()
			config.DrawFunc(&WinDraw{win, slice})
			slice.Draw()
			slice.End()
			shader.End()
			
			win.SwapBuffers()
		}
	})
}
