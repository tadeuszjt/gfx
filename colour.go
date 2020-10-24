package gfx

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
