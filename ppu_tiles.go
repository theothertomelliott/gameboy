package gameboy

import (
	"image"
)

func NewTile(tile []byte) Tile {
	return Tile(tile)
}

type Tile []byte

func (t Tile) At(x, y int) byte {
	high := t[y*2+1]
	low := t[y*2]

	h := bitValue(7-byte(x), high)
	l := bitValue(7-byte(x), low)
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
