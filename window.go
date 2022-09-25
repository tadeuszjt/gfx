package gfx

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/glfw/v3.3/glfw"
	geom "github.com/tadeuszjt/geom/32"
)

type TexID struct {
	texIdx int
	texNum int
}

type texture struct {
	texNum      int // unique id
	isMipmapped bool
	frame       *glhf.Frame // nil if dead
}

type Win struct {
	glfwWin *glfw.Window

	w2D, w3D struct {
		slice  *glhf.VertexSlice
		shader *glhf.Shader
	}

	textureCount int
	textures     []texture
	whiteTexID   TexID
}

func (w *Win) Size() geom.Vec2 {
	width, height := w.glfwWin.GetFramebufferSize()
	return geom.Vec2{float32(width), float32(height)}
}

func (w *Win) GetGlfwWindow() *glfw.Window {
	return w.glfwWin
}

func (w *Win) getTexture(id *TexID) *texture {
	if id == nil {
		return w.getTexture(&w.whiteTexID)
	}
	tex := &w.textures[(*id).texIdx]
	if (*id).texNum != tex.texNum {
		panic("invalid TexID")
	}
	return tex
}

func (w *Win) addTexture(isMipmapped bool, frame *glhf.Frame) TexID {
	num := w.textureCount
	w.textureCount++

	tex := texture{
		texNum:      num,
		isMipmapped: isMipmapped,
		frame:       frame,
	}

	for i := range w.textures {
		if w.textures[i].frame == nil {
			w.textures[i] = tex
			return TexID{texNum: num, texIdx: i}
		}
	}

	id := TexID{texIdx: len(w.textures), texNum: num}
	w.textures = append(w.textures, tex)
	return id
}

func (w *Win) GetTextureCanvas(texID TexID) texCanvas {
	w.getTexture(&texID)
	return texCanvas{
		win:   w,
		texID: texID,
	}
}

func (win *Win) SetTexturePixels(id TexID, x, y, w, h int, pixels []uint8) {
	texture := win.getTexture(&id)
	texture.frame.Texture().Begin()
	texture.frame.Texture().SetPixels(x, y, w, h, pixels)
	texture.frame.Texture().End()
}

func (w *Win) FreeTexture(texID TexID) {
	for i := range w.textures {
		if w.textures[i].texNum == texID.texNum {
			w.textures[i].frame = nil
			w.textures[i].texNum = -1
			return
		}
	}
	panic("invalid TexID")
}

func (w *Win) LoadTextureFromFile(path string) (TexID, error) {
	width, height, pixels, err := loadImage(path)
	if err != nil {
		return TexID{}, err
	}

	frame := glhf.NewFrame(width, height, false)
	frame.Texture().Begin()
	frame.Texture().SetPixels(0, 0, width, height, pixels)
	frame.Texture().End()

	return w.addTexture(false, frame), nil

	//	tex := w.textures[id].Texture()
	//	tex.Begin()
	//	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	//	gl.GenerateMipmap(gl.TEXTURE_2D)
	//	tex.End()
}

func (w *Win) LoadTextureFromPixels(width, height int, pixels []uint8) TexID {
	frame := glhf.NewFrame(width, height, false)
	frame.Texture().Begin()
	frame.Texture().SetPixels(0, 0, width, height, pixels)
	frame.Texture().End()

	return w.addTexture(false, frame)
}

func (w *Win) LoadTextureBlank(width, height int) TexID {
	frame := glhf.NewFrame(width, height, false)
	return w.addTexture(false, frame)
}
