package main

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

var (
	win fyne.Window
	img image.Image
	a   fyne.App
)

func setupGraphics() {
	a = app.New()
	win = a.NewWindow("Game Boy")

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			if img == nil {
				return color.Black
			}
			xPos := int((float64(x) / float64(w)) * float64(img.Bounds().Dx()))
			yPos := int((float64(y) / float64(h)) * float64(img.Bounds().Dy()))
			return img.At(xPos, yPos)
		})
	raster.SetMinSize(fyne.Size{
		Width:  600,
		Height: 600,
	})

	win.SetContent(raster)
}

func drawGraphics(graphics image.Image) {
	img = graphics
	win.Content().Refresh()
}
