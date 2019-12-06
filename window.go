package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/geom/geom32"
)

type TexID int

type Win struct {
	glfwWin  *glfw.Window
	slice    *glhf.VertexSlice
	shader   *glhf.Shader
	textures []*glhf.Texture
}

func (w *Win) GetFrameSize() geom.Vec2 {
	width, height := w.glfwWin.GetFramebufferSize()
	return geom.Vec2{float32(width), float32(height)}
}

func (w *Win) LoadTexture(path string) (TexID, error) {
	width, height, pixels, err := loadImage(path)
	if err != nil {
		return 0, err
	}

	/* generate mipmap */
	id := w.loadTextureFromPixels(width, height, pixels)
	tex := w.textures[id]
	tex.Begin()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.GenerateMipmap(gl.TEXTURE_2D)
	tex.End()

	return id, nil
}

func (w *Win) loadTextureFromPixels(width, height int, pixels []uint8) TexID {
	tex := glhf.NewTexture(width, height, true, pixels)
	w.textures = append(w.textures, tex)
	return TexID(len(w.textures) - 1)
}

