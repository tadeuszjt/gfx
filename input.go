package gfx

import (
	"github.com/tadeuszjt/geom/geom32"
)

type MouseEvent interface {
}

type MouseScroll struct {
	Dx, Dy float32
}

type MouseMove struct {
	Position geom.Vec2
}

type MouseButton struct {
}
