package gfx

import (
	"github.com/golang/freetype/truetype"
	"github.com/tadeuszjt/geom/32"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
	"image"
)

const (
	textDPI       = 72
	textTexWidth  = 1024
	textTexHeight = 256
)

var (
	trueTypeFont, _ = truetype.Parse(gomono.TTF)
)

type Text struct {
	str  string
	size int
	w, h int
	face font.Face
	img  *image.RGBA
}

func (t *Text) SetString(str string) {
	t.str = str

	if t.face == nil {
		t.size = 12
		t.face = truetype.NewFace(trueTypeFont, &truetype.Options{
			Size:    12,
			DPI:     textDPI,
			Hinting: font.HintingFull,
		})
	}

	bounds, _ := font.BoundString(t.face, t.str)
	t.w, t.h = bounds.Max.X.Ceil(), (-bounds.Min.Y).Ceil()+(t.size*2)/5
	t.img = image.NewRGBA(image.Rect(0, 0, t.w, t.h))

	d := &font.Drawer{
		Dst:  t.img,
		Src:  image.Black,
		Face: t.face,
		Dot:  fixed.Point26_6{X: 0, Y: -bounds.Min.Y},
	}
	d.DrawString(t.str)
}

func (t *Text) SetSize(size int) {
	t.size = size
	t.face = truetype.NewFace(trueTypeFont, &truetype.Options{
		Size:    float64(size),
		DPI:     textDPI,
		Hinting: font.HintingFull,
	})

	t.SetString(t.str)
}

func (t *Text) Size() int {
	return t.size
}

func (t *Text) CharWidth() float64 {
    return float64(font.MeasureString(t.face, "          ").Ceil()) / 10.
}

func DrawText(c Canvas, text *Text, pos geom.Vec2) {
	if text.img == nil || text.str == "" {
		return
	}

	win := c.getWindow()
	tex := win.getTexture(&win.textTexID)
	tex.frame.Texture().Begin()
	tex.frame.Texture().SetPixels(0, 0, text.w, text.h, text.img.Pix)
	tex.frame.Texture().End()

	W, H := float32(text.w), float32(text.h)
	strRect := geom.MakeRect(W, H, pos)
	texRect := geom.RectOrigin(W/float32(textTexWidth), H/float32(textTexHeight))

	texCoords := texRect.Verts()
	verts := strRect.Verts()
	col := White
	data := make([]float32, 0, 6*8)

	for _, j := range [6]int{0, 1, 2, 0, 2, 3} {
		data = append(
			data,
			verts[j].X, verts[j].Y,
			texCoords[j].X, texCoords[j].Y,
			col.R, col.G, col.B, col.A,
		)
	}

	c.Draw2DVertexData(data, &win.textTexID, nil)
}
