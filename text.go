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
	trueTypeFont, _ = truetype.Parse(goregular.TTF)
)

func (w *Win) setupText() {
	w.textTexID = w.loadTextureFromPixels(
		textTexWidth,
		textTexHeight,
		false,
		image.NewRGBA(image.Rect(0, 0, textTexWidth, textTexHeight)).Pix)
}

type Text struct {
	str  string
	size float64
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
	t.w, t.h = bounds.Max.X.Ceil(), (-bounds.Min.Y).Ceil()+int(t.size*0.4)
	t.img = image.NewRGBA(image.Rect(0, 0, t.w, t.h))

	d := &font.Drawer{
		Dst:  t.img,
		Src:  image.Black,
		Face: t.face,
		Dot:  fixed.Point26_6{X: 0, Y: -bounds.Min.Y},
	}
	d.DrawString(t.str)
}

func (t *Text) SetSize(size float64) {
	t.size = size
	t.face = truetype.NewFace(trueTypeFont, &truetype.Options{
		Size:    size,
		DPI:     textDPI,
        Hinting: font.HintingFull,
	})

	t.SetString(t.str)
}

func (w *WinDraw) DrawText(text *Text, pos geom.Vec2) {
	if text.img == nil {
		return
	}

	tex := w.window.textures[w.window.textTexID]
    tex.Begin()
	tex.SetPixels(0, 0, text.w, text.h, text.img.Pix)
    tex.End()

	W, H := float32(text.w), float32(text.h)
	strRect := geom.MakeRect(W, H, pos)
	texRect := geom.RectOrigin(W/textTexWidth, H/textTexHeight)

    texCoords := texRect.Verts()
    verts := strRect.Verts()
    col := White
    mat := geom.Mat3Identity()
	data := make([]float32, 0, 6*8)

	for _, j := range [6]int{0, 1, 2, 0, 2, 3} {
		data = append(
			data,
			verts[j].X, verts[j].Y,
			texCoords[j].X, texCoords[j].Y,
			col.R, col.G, col.B, col.A,
		)
	}

	w.DrawVertexData(data, &w.window.textTexID, &mat)
}
