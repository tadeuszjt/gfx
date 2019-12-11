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

func init() {
	var err error
	trueTypeFont, err = truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
}

func (w *Win) textInit() {
	textTexID = w.loadTextureFromPixels(
		textTexWidth,
		textTexHeight,
		false,
		image.NewRGBA(image.Rect(0, 0, textTexWidth, textTexHeight)).Pix)
}

type Text struct {
	str   string
	size  float64
	w, h  int
	face  font.Face
	img   *image.RGBA
}

func (t *Text) SetString(str string) {
	t.str = str
	
	if t.face == nil {
		t.size = 12
		t.face = truetype.NewFace(trueTypeFont, &truetype.Options{
			Size:    12,
			DPI:     textDPI,
			Hinting: font.HintingNone,
		})
	}
	
	bounds, _ := font.BoundString(t.face, t.str)
	t.w, t.h = bounds.Max.X.Ceil(), (-bounds.Min.Y).Ceil() + int(t.size * 0.4)
	t.img = image.NewRGBA(image.Rect(0, 0, t.w, t.h))
	
	d := &font.Drawer{
		Dst:  t.img,
		Src:  image.Black,
		Face: t.face,
		Dot:  fixed.Point26_6{X: -bounds.Min.X, Y: -bounds.Min.Y},
	}
	d.DrawString(t.str)
}

func (t *Text) SetSize(size float64) {
	t.size = size
	t.face = truetype.NewFace(trueTypeFont, &truetype.Options{
		Size:    size,
		DPI:     textDPI,
		Hinting: font.HintingNone,
	})
	
	t.SetString(t.str)
}

func (w *WinDraw) DrawText(text *Text, pos geom.Vec2) {
	if text.img == nil {
		return
	}
	
	w.setActiveTexture(textTexID)
	w.activeTexture.SetPixels(0, 0, text.w, text.h, text.img.Pix)

	W, H := float32(text.w), float32(text.h)
	strRect := geom.MakeRect(W, H, pos)
	texRect := geom.RectOrigin(W/textTexWidth, H/textTexHeight)

	w.DrawRect(strRect, &textTexID, nil, nil, &texRect)
}
