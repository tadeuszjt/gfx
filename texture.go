package gfx

import (
	"os"
	"fmt"
	"bufio"
	"image"
	_ "image/png"
)

func loadImage(path string) (int, int, []uint8, error) {
	file, err := os.Open(path)
	defer file.Close()
	
	if err != nil {
		return 0, 0, nil, err
	}

	img, _, err := image.Decode(bufio.NewReader(file))
	if err != nil {
		return 0, 0, nil, err
	}
	
	size := img.Bounds().Size()

	switch trueim := img.(type) {
	case *image.RGBA:
		return size.X, size.Y, trueim.Pix, nil
	case *image.NRGBA:
		return size.X, size.Y, trueim.Pix, nil
	}
	
	return 0, 0, nil, fmt.Errorf("unhandled image format")
}

