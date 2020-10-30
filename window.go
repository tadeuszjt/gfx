package gfx

import (
	"github.com/faiface/glhf"
	//"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/geom/32"
)

type TexID int

type texture struct {
	isMipmapped bool
	isSmooth    bool
	frame       *glhf.Frame
}

type Win struct {
	glfwWin *glfw.Window

	w2D, w3D struct {
		slice  *glhf.VertexSlice
		shader *glhf.Shader
	}

	textures   []texture
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

func (w *Win) getTexture(id *TexID) *texture {
	if id == nil {
		return &w.textures[w.whiteTexID]
	}

	if *id <= 0 || int(*id) >= len(w.textures) {
		panic("invalid texture ID")
	}

	return &w.textures[*id]
}

func (w *Win) LoadTextureFromFile(path string) (TexID, error) {
	width, height, pixels, err := loadImage(path)
	if err != nil {
		return 0, err
	}

	tex := texture{
		isMipmapped: false,
		frame:       glhf.NewFrame(width, height, false),
	}

	tex.frame.Texture().Begin()
	tex.frame.Texture().SetPixels(0, 0, width, height, pixels)
	tex.frame.Texture().End()

	w.textures = append(w.textures, tex)
	return TexID(len(w.textures) - 1), nil

	//	tex := w.textures[id].Texture()
	//	tex.Begin()
	//	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	//	gl.GenerateMipmap(gl.TEXTURE_2D)
	//	tex.End()
}

func (w *Win) LoadTextureFromPixels(width, height int, pixels []uint8) TexID {
	tex := texture{
		isMipmapped: false,
		frame:       glhf.NewFrame(width, height, false),
	}

	tex.frame.Texture().Begin()
	tex.frame.Texture().SetPixels(0, 0, width, height, pixels)
	tex.frame.Texture().End()

	w.textures = append(w.textures, tex)
	return TexID(len(w.textures) - 1)
}

func (w *Win) LoadTextureBlank(width, height int) TexID {
	tex := texture{
		isMipmapped: false,
		isSmooth:    false,
		frame:       glhf.NewFrame(width, height, false),
	}

	w.textures = append(w.textures, tex)
	return TexID(len(w.textures) - 1)
}
