package gfx

import (
	"github.com/golang/freetype/truetype"
	"github.com/tadeuszjt/geom/32"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
    "io/ioutil"
)

const (
	textDPI         = 72
    textDefaultSize = 12
)

var (
	trueTypeFont *truetype.Font
)

func init() {
    data, err := ioutil.ReadFile("/usr/share/fonts/TTF/LiberationMono-Regular.ttf")
    if err != nil {
        panic(err)
    }
    trueTypeFont, _ = truetype.Parse(data)
}

type Text struct {
	str    string
	size   int
	face   font.Face
    colour Colour
	img    *image.RGBA

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
        size:   textDefaultSize,
        face:   face,
        colour: Black,
    }
}

func (t *Text) SetColour(col Colour) {
    t.colour = col
    t.redrawImg()
}

func (t *Text) redrawImg() {
    width := font.MeasureString(t.face, t.str).Ceil()
	t.img = image.NewRGBA(image.Rect(0, 0, width, t.Height()))

	d := &font.Drawer{
		Dst:  t.img,
		Src:  image.NewUniform(t.colour),
		Face: t.face,
        Dot:  fixed.Point26_6{X: 0, Y: t.face.Metrics().Ascent},
	}

	d.DrawString(t.str)
    t.isCurrent = false
}

func (t *Text) SetString(str string) {
	t.str = str
    t.redrawImg()
}

func (t *Text) SetSize(size int) {
	t.size = size
	t.face = truetype.NewFace(trueTypeFont, &truetype.Options{
		Size:    float64(size),
		DPI:     textDPI,
		Hinting: font.HintingFull,
	})

    t.redrawImg()
}

func (t *Text) Height() int {
    metrics := t.face.Metrics()
    return (metrics.Ascent + metrics.Descent).Ceil()
}

func (t *Text) CharWidth() fixed.Int26_6 {
    f, _ := t.face.GlyphAdvance(' ')
    return f
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
