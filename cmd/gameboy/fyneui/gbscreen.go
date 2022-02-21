package fyneui

import (
	"image"
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/ppu"
)

var _ fyne.Widget = &gbscreen{}
var _ fyne.WidgetRenderer = &gbscreenRenderer{}

type gbscreen struct {
	widget.BaseWidget

	size fyne.Size

	img ScreenBinding
}

func (g *gbscreen) Size() fyne.Size {
	return g.size
}

func (g *gbscreen) MinSize() fyne.Size {
	return g.size
}

func newScreen(img ScreenBinding, size fyne.Size) fyne.Widget {
	g := &gbscreen{
		img:  img,
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
		getImage: g.img.Get,
	}
	gr.raster = canvas.NewRasterWithPixels(
		gr.pixelColor,
	)
	g.img.AddListener(gr)
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

type ScreenBinding interface {
	binding.DataItem
	Get() image.Image
	Set(gb *gameboy.DMG)
}

var _ ScreenBinding = &screenData{}

func NewScreenBinding() ScreenBinding {
	return &screenData{}
}

type screenData struct {
	img image.Image

	listenersMtx sync.Mutex
	listeners    []binding.DataListener
}

func (s *screenData) Get() image.Image {
	return s.img
}

func (s *screenData) Set(gb *gameboy.DMG) {
	if !ppu.GetLCDControl(gb.MMU()).LCDOperation() {
		return
	}
	s.img = gb.PPU().RenderScreen()

	s.listenersMtx.Lock()
	defer s.listenersMtx.Unlock()
	for _, ln := range s.listeners {
		ln.DataChanged()
	}
}

func (s *screenData) AddListener(l binding.DataListener) {
	s.listenersMtx.Lock()
	defer s.listenersMtx.Unlock()

	s.listeners = append(s.listeners, l)
	l.DataChanged()
}

func (s *screenData) RemoveListener(l binding.DataListener) {
	s.listenersMtx.Lock()
	defer s.listenersMtx.Unlock()

	var newListeners []binding.DataListener
	for _, ln := range s.listeners {
		if ln != l {
			newListeners = append(newListeners, ln)
		}
	}

	s.listeners = newListeners
}
