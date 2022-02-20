package fyneui

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (u *UI) gbScreen() fyne.CanvasObject {
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			u.imgMutex.Lock()
			defer u.imgMutex.Unlock()
			if u.img == nil {
				return color.Black
			}
			xPos := int((float64(x) / float64(w)) * float64(u.img.Bounds().Dx()))
			yPos := int((float64(y) / float64(h)) * float64(u.img.Bounds().Dy()))
			return u.img.At(xPos, yPos)
		})
	raster.SetMinSize(fyne.Size{
		Width:  600,
		Height: 600,
	})
	return raster
}

func (u *UI) drawGraphics(graphics image.Image) {
	u.imgMutex.Lock()
	defer u.imgMutex.Unlock()

	u.img = graphics
	u.win.Content().Refresh()
}
