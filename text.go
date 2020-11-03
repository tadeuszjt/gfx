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
	textDPI         = 72
    textDefaultSize = 12
)

var (
	trueTypeFont, _ = truetype.Parse(gomono.TTF)
)

type Text struct {
	str   string
	size  int
	face  font.Face
	img   *image.RGBA

    isCurrent bool
    texID     *TexID
}

func MakeText() Text {
    face := truetype.NewFace(trueTypeFont, &truetype.Options{
        Size:    float64(textDefaultSize),
        DPI:     textDPI,
        Hinting: font.HintingFull,
    })

    return Text{
        size: textDefaultSize,
        face: face,
    }
}

func (t *Text) redrawImg() {
    width := font.MeasureString(t.face, t.str).Ceil()
	t.img = image.NewRGBA(image.Rect(0, 0, width, t.Height()))

	d := &font.Drawer{
		Dst:  t.img,
		Src:  image.Black,
		Face: t.face,
        Dot:  fixed.Point26_6{X: 0, Y: t.face.Metrics().Ascent},
	}

	d.DrawString(t.str)
}

func (t *Text) SetString(str string) {
	t.str = str
    t.redrawImg()
    t.isCurrent = false
}

func (t *Text) SetSize(size int) {
	t.size = size
	t.face = truetype.NewFace(trueTypeFont, &truetype.Options{
		Size:    float64(size),
		DPI:     textDPI,
		Hinting: font.HintingFull,
	})

    t.redrawImg()
    t.isCurrent = false
}

func (t *Text) Height() int {
    metrics := t.face.Metrics()
    return (metrics.Ascent + metrics.Descent).Ceil()
}

func (t *Text) Size() int {
	return t.size
}

func (t *Text) GetString() string {
    return t.str
}

func DrawText(c Canvas, text *Text, pos geom.Vec2) {
	if text.img == nil || text.str == "" {
		return
	}

	win := c.getWindow()

    bounds := text.img.Bounds()
    if !text.isCurrent {
        if text.texID != nil {
            win.FreeTexture(*text.texID)
            text.texID = nil
        }

        id := win.LoadTextureFromPixels(bounds.Max.X, bounds.Max.Y, text.img.Pix)
        text.texID = &id
        text.isCurrent = true
    }


	W, H := float32(bounds.Max.X), float32(bounds.Max.Y)
	strRect := geom.MakeRect(W, H, pos)
    DrawRect(c, text.texID, strRect, geom.RectOrigin(1, 1))
}
