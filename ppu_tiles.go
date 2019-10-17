package gameboy

import (
	"image"
)

func NewTile(tile []byte) Tile {
	var out = make(Tile, 8*8)
	if len(tile) < 16 {
		return out
	}
	for line := 0; line < 8; line++ {
		high := tile[line*2+1]
		low := tile[line*2]
		for bit := byte(0); bit < 8; bit++ {
			h := bitValue(7-bit, high)
			l := bitValue(7-bit, low)
			colorValue := l + (h << 1)
			out.set(int(bit), line, colorValue)
		}
	}
	return out
}

type Tile []byte

func (t Tile) At(x, y int) byte {
	if i := x*8 + y; i >= 0 && i < len(t) {
		return t[i]
	}
	return 0
}

func (t Tile) set(x, y int, value byte) {
	t[x*8+y] = value
}

// ToImage converts this tile into an image.Image using the
// default pallete
func (t Tile) ToImage() image.Image {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 8, Y: 8},
	})
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			img.Set(i, j, colorForValue(t.At(i, j)))
		}
	}
	return img
}
