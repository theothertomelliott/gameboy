package ppu

import (
	"image"

	"github.com/theothertomelliott/gameboy/binary"
)

func NewTile(tile []byte) Tile {
	return Tile(tile)
}

type Tile []byte

func (t Tile) At(x, y int) byte {
	if y*2+1 >= len(t) {
		return 0
	}
	high := t[y*2+1]
	low := t[y*2]

	h := binary.Bit(7-byte(x), high)
	l := binary.Bit(7-byte(x), low)
	return l + (h << 1)
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
