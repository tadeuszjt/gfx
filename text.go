package gfx

import (
	"github.com/golang/freetype/truetype"
	"github.com/tadeuszjt/geom/32"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
	"image"
)

const (
	textDPI       = 72
	textTexWidth  = 2048
	textTexHeight = 256
)

var (
	textTexID    TexID
	trueTypeFont *truetype.Font
)

func (w *Win) textInit() {
	var err error
	trueTypeFont, err = truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	textTexID = w.loadTextureFromPixels(
		textTexWidth,
		textTexHeight,
		false,
		image.NewRGBA(image.Rect(0, 0, textTexWidth, textTexHeight)).Pix)
}

func (w *WinDraw) DrawText(str string, pos geom.Vec2, size float64) {
	w.setActiveTexture(textTexID)

	trueTypeFace := truetype.NewFace(trueTypeFont, &truetype.Options{
		Size:    size,
		DPI:     textDPI,
		Hinting: font.HintingNone,
	})

	bounds, _ := font.BoundString(trueTypeFace, str)
	strW, strH := bounds.Max.X.Ceil(), (-bounds.Min.Y).Ceil() + int(size * 0.4)

	img := image.NewRGBA(image.Rect(0, 0, strW, strH))
	d := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: trueTypeFace,
		Dot:  fixed.Point26_6{X: -bounds.Min.X, Y: -bounds.Min.Y},
	}
	d.DrawString(str)

	w.activeTexture.SetPixels(0, 0, strW, strH, img.Pix)

	W, H := float32(strW), float32(strH)
	TW, TH := float32(textTexWidth), float32(textTexHeight)

	strRect := geom.MakeRect(W, H, pos)
	texRect := geom.RectOrigin(W/TW, H/TH)

	verts := strRect.Verts()
	texCoords := texRect.Verts()

	data := make([]float32, 0, 8*6)

	for _, i := range []int{0, 1, 2, 0, 2, 3} {
		data = append(data,
			verts[i].X, verts[i].Y,
			texCoords[i].X, texCoords[i].Y,
			1, 1, 1, 1)
	}

	w.DrawVertexData(data, &textTexID, nil)
}
