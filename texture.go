package gfx

import (
	"bufio"
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tadeuszjt/geom/32"
	"image"
	_ "image/png"
	"os"
)

type texCanvas struct {
	win   *Win
	texID TexID
}

func (t texCanvas) Clear(col Colour) {
	tex := t.win.getTexture(&t.texID)

	tex.frame.Begin()
	gl.ClearColor(col.R, col.G, col.B, col.A)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	tex.frame.End()
}

func (t texCanvas) Draw2DVertexData(data []float32, texID *TexID, mat *geom.Mat3) {
	win := t.getWindow()
	tex := win.getTexture(texID)
	frame := win.getTexture(&t.texID).frame
	width := frame.Texture().Width()
	height := frame.Texture().Height()

	win.w2D.shader.Begin()
	win.w2D.slice.Begin()
	tex.frame.Texture().Begin()
	frame.Begin()
	defer frame.End()
	defer tex.frame.Texture().End()
	defer win.w2D.slice.End()
	defer win.w2D.shader.End()

	toNDC := geom.Mat3Camera2D(
		geom.RectOrigin(float32(width), float32(height)),
		geom.RectCentred(2, 2),
	)

	var m geom.Mat3
	if mat != nil {
		m = toNDC.Product(*mat)
	} else {
		m = toNDC
	}
	win.w2D.shader.SetUniformAttr(0, mgl32.Mat3{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8],
	})

	win.w2D.slice.SetLen(len(data) / 8)
	win.w2D.slice.SetVertexData(data)

	/* draw with texture viewport, replace with previous */
	var i [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &i[0])
	gl.Viewport(0, 0, int32(width), int32(height))
	win.w2D.slice.Draw()
	gl.Viewport(i[0], i[1], i[2], i[3])
}

func (t texCanvas) Draw3DVertexData(data []float32, texID *TexID, mat *geom.Mat4) {
	panic("")
}

func (t texCanvas) getWindow() *Win {
	return t.win
}

func loadImage(path string) (int, int, []uint8, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return 0, 0, nil, err
	}

	img, _, err := image.Decode(bufio.NewReader(file))
	if err != nil {
		return 0, 0, nil, err
	}

	size := img.Bounds().Size()

	switch trueim := img.(type) {
	case *image.RGBA:
		return size.X, size.Y, trueim.Pix, nil
	case *image.NRGBA:
		return size.X, size.Y, trueim.Pix, nil
	}

	return 0, 0, nil, fmt.Errorf("unhandled image format")
}
