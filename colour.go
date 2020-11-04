package gfx

import (
	"math/rand"
    "image/color"
)

var (
	Red   = Colour{1, 0, 0, 1}
	Green = Colour{0, 1, 0, 1}
	Blue  = Colour{0, 0, 1, 1}
	White = Colour{1, 1, 1, 1}
	Black = Colour{0, 0, 0, 1}
)

type Colour struct {
	R, G, B, A float32
}

func (c Colour) RGBA() (r, g, b, a uint32) {
    return color.RGBA64{
        uint16(c.R*65535.),
        uint16(c.G*65535.),
        uint16(c.B*65535.),
        uint16(c.A*65535.),
    }.RGBA()
}

func ColourRGBA(r, g, b, a uint8) Colour {
    return Colour{
        float32(r)/255.,
        float32(g)/255.,
        float32(b)/255.,
        float32(a)/255.,
    }
}

func ColourRand() Colour {
	return Colour{
		float32(rand.Float64()),
		float32(rand.Float64()),
		float32(rand.Float64()),
		1.0,
	}
}

type SliceColour []Colour

func (s *SliceColour) Len() int {
	return len(*s)
}

func (s *SliceColour) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SliceColour) Delete(i int) {
	end := s.Len() - 1
	if i < end {
		s.Swap(i, end)
	}

	*s = (*s)[:end]
}

func (s *SliceColour) Append(item interface{}) {
	i, _ := item.(Colour)
	*s = append(*s, i)
}
