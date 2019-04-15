package main

import (
	"image"

	"github.com/faiface/pixel"
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

func drawGraphics(graphics image.Image) {
	win.Clear(colornames.White)
	pg := pixel.PictureDataFromImage(graphics)
	sprite := pixel.NewSprite(pg, pg.Bounds())
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	win.Update()
}
