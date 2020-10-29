package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/geom/32"
)

type TexID int

type Win struct {
	glfwWin  *glfw.Window
	w2D, w3D struct {
		slice  *glhf.VertexSlice
		shader *glhf.Shader
	}
	textures   []*glhf.Frame
	whiteTexID TexID
	textTexID  TexID
}

func (w *Win) GetFrameRect() geom.Rect {
	width, height := w.glfwWin.GetFramebufferSize()
	return geom.RectOrigin(float32(width), float32(height))
}

func (w *Win) GetGlfwWindow() *glfw.Window {
	return w.glfwWin
}

func (w *Win) LoadTexture(path string) (TexID, error) {
	width, height, pixels, err := loadImage(path)
	if err != nil {
		return 0, err
	}

	/* generate mipmap */
	id := w.loadTextureFromPixels(width, height, true, pixels)
	tex := w.textures[id].Texture()
	tex.Begin()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.GenerateMipmap(gl.TEXTURE_2D)
	tex.End()

	return id, nil
}

func (w *Win) loadTextureFromPixels(width, height int, smooth bool, pixels []uint8) TexID {
	frame := glhf.NewFrame(width, height, smooth)
	tex := frame.Texture()
	tex.Begin()
	tex.SetPixels(0, 0, width, height, pixels)
	tex.End()
	w.textures = append(w.textures, frame)
	return TexID(len(w.textures) - 1)
}
