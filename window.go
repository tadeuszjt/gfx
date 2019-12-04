package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Win struct {
	textures []*glhf.Texture
	glfwWin  *glfw.Window
}

type TexID int

func (w *Win) loadWhiteTex() {
	w.textures = append(w.textures, glhf.NewTexture(1, 1, false, []uint8{255, 255, 255, 255}))
}

func (w *Win) loadTextureFromPixels(width, height int, pixels []uint8) TexID {
	tex := glhf.NewTexture(width, height, true, pixels)
	w.textures = append(w.textures, tex)
	return TexID(len(w.textures) - 1)
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
