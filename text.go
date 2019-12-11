package gfx

import (
	"image"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"	
	"github.com/tadeuszjt/geom/32"
	"golang.org/x/image/font/gofont/goregular"
)
	
var (
	textTexID     TexID
	textTexWidth  = 2000
	textTexHeight = 200
	trueTypeFace  font.Face
)

func (w *Win) textInit() {
	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	
	trueTypeFace = truetype.NewFace(ttf, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	
	textTexID = w.loadTextureFromPixels(
		textTexWidth,
		textTexHeight,
		false,
		image.NewRGBA(image.Rect(0, 0, textTexWidth, textTexHeight)).Pix)
}

func (w *WinDraw) DrawText(str string, pos geom.Vec2, size float32) {
	w.setActiveTexture(textTexID)
	
	strBounds, _ := font.BoundString(trueTypeFace, str)
	strW, strH := strBounds.Max.X.Ceil(), -strBounds.Min.Y.Floor()
	
	img := image.NewRGBA(image.Rect(0, 0, strW, strH))
	d := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: trueTypeFace,
		Dot: fixed.Point26_6{X: fixed.I(0), Y: fixed.I(strH)},
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
