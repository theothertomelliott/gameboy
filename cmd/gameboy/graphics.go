package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	sizeX, sizeY              = 160, 144
	screenWidth, screenHeight = float64(1600 / 2), float64(1440 / 2)
)

var (
	win *pixelgl.Window
)

func setupGraphics() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game Boy",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  true,
	}
	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	// Workaround for Mojave
	// https://github.com/faiface/pixel/issues/140
	win.SetPos(win.GetPos().Add(pixel.V(0, 1)))
}

func drawGraphics(graphics [][]byte, scrollX, scrollY byte) {
	win.Clear(colornames.White)
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 0)
	screenWidth := win.Bounds().W()
	width, height := screenWidth/sizeX, screenHeight/sizeY
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			value := graphics[x+int(scrollX)][y+int(scrollY)]
			if value == 0 {
				continue
			}
			imd.Push(pixel.V(width*float64(x), height*float64(y)))
			imd.Push(pixel.V(width*float64(x)+width, height*float64(y)+height))
			R := 256 / float64((value&0xE0)>>5) // rgb8 & 1110 0000  >> 5
			G := 256 / float64((value&0x1C)>>2) // rgb8 & 0001 1100  >> 2
			B := 256 / float64(value&0x03)      // rgb8 & 0000 0011

			imd.Color = pixel.RGB(R, G, B)
			imd.Rectangle(0)
		}
	}

	imd.Draw(win)
	win.Update()
}
