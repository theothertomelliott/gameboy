package gameboy

import (
	"image"
	"image/color"
)

var _ image.Image = LCD{}

type LCD [][]byte

// NewLCD initializes a 160x144 screen
func NewLCD() LCD {
	lcd := make(LCD, 144)
	for row := range lcd {
		lcd[row] = make([]byte, 160)
	}
	return lcd
}

func (l LCD) RenderLine(line byte, values []byte) {
	l[line] = values
}

func (l LCD) At(x int, y int) color.Color {
	return colorForValue(l[y][x])
}

func (l LCD) Bounds() image.Rectangle {
	return image.Rect(0, 0, 160, 144)
}

func (l LCD) ColorModel() color.Model {
	return color.RGBA64Model
}
