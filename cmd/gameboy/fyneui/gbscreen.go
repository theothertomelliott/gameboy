package fyneui

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &gbscreen{}
var _ fyne.WidgetRenderer = &gbscreenRenderer{}

type gbscreen struct {
	widget.BaseWidget

	size fyne.Size

	data DataTransport
}

func (g *gbscreen) Size() fyne.Size {
	return g.size
}

func (g *gbscreen) MinSize() fyne.Size {
	return g.size
}

func newScreen(data DataTransport, size fyne.Size) fyne.Widget {
	g := &gbscreen{
		data: data,
		size: size,
	}
	return g
}

type gbscreenRenderer struct {
	getImage func() image.Image

	raster *canvas.Raster
	size   fyne.Size
}

func (g *gbscreenRenderer) DataChanged() {
	g.raster.Refresh()
}

func (g *gbscreen) CreateRenderer() fyne.WidgetRenderer {
	gr := &gbscreenRenderer{
		size:     g.size,
		getImage: g.data.Screen,
	}
	gr.raster = canvas.NewRasterWithPixels(
		gr.pixelColor,
	)
	g.data.AddListener(gr)
	return gr
}

// Destroy is for internal use.
func (g *gbscreenRenderer) Destroy() {}

// Layout is a hook that is called if the widget needs to be laid out.
// This should never call Refresh.
func (g *gbscreenRenderer) Layout(s fyne.Size) {
	g.raster.Move(fyne.Position{X: 0, Y: 0})
	g.raster.Resize(g.size)
}

// MinSize returns the minimum size of the widget that is rendered by this renderer.
func (g *gbscreenRenderer) MinSize() fyne.Size {
	return g.size
}

// Objects returns all objects that should be drawn.
func (g *gbscreenRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		g.raster,
	}
}

func (g *gbscreenRenderer) Refresh() {}

func (g *gbscreenRenderer) pixelColor(x, y, w, h int) color.Color {
	img := g.getImage()
	if img == nil {
		return color.Black
	}
	xPos := int((float64(x) / float64(w)) * float64(img.Bounds().Dx()))
	yPos := int((float64(y) / float64(h)) * float64(img.Bounds().Dy()))
	return img.At(xPos, yPos)
}
