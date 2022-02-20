package fyneui

import (
	"image"
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/ppu"
)

var _ fyne.Widget = &gbscreen{}
var _ fyne.CanvasObject = &gbscreen{}
var _ fyne.WidgetRenderer = &gbscreenRenderer{}

type gbscreen struct {
	widget.BaseWidget

	gb   *gameboy.DMG
	p    *ppu.PPU
	size fyne.Size

	img      image.Image
	imgMutex sync.Mutex
}

func (g *gbscreen) Size() fyne.Size {
	return g.size
}

func (g *gbscreen) MinSize() fyne.Size {
	return g.size
}

func newScreen(gb *gameboy.DMG, size fyne.Size) fyne.Widget {
	g := &gbscreen{
		gb:   gb,
		p:    gb.PPU(),
		size: size,
	}
	return g
}

type gbscreenRenderer struct {
	p *ppu.PPU

	getImage func() image.Image

	raster *canvas.Raster
	size   fyne.Size
}

func (g *gbscreen) CreateRenderer() fyne.WidgetRenderer {
	gr := &gbscreenRenderer{
		p:    g.gb.PPU(),
		size: g.size,
		getImage: func() image.Image {
			g.imgMutex.Lock()
			defer g.imgMutex.Unlock()
			return g.img
		},
	}
	gr.raster = canvas.NewRasterWithPixels(
		gr.pixelColor,
	)
	return gr
}

func (g *gbscreen) Refresh() {
	if !ppu.GetLCDControl(g.p.MMU).LCDOperation() {
		return
	}

	g.imgMutex.Lock()
	g.img = g.p.RenderScreen()
	g.imgMutex.Unlock()
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
