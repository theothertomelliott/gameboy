package fyneui

import (
	"fmt"
	"image"
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/input"
	"github.com/theothertomelliott/gameboy/ppu"
)

func NewUI() *UI {
	return &UI{}
}

type UI struct {
	win      fyne.Window
	img      image.Image
	a        fyne.App
	imgMutex sync.Mutex
}

func (u *UI) Run(
	gb *gameboy.DMG,
) {
	p := gb.PPU()

	u.setupApp()

	if deskCanvas, ok := u.win.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			k, exists := keyByName[ke.Name]
			if !exists {
				fmt.Println("key not found: ", ke.Name)
			}
			gb.Input().Press(k)
		})

		deskCanvas.SetOnKeyUp(func(ke *fyne.KeyEvent) {
			k, exists := keyByName[ke.Name]
			if !exists {
				fmt.Println("key not found: ", ke.Name)
			}
			gb.Input().Release(k)
		})
	}

	go func() {
		for {
			if ppu.GetLCDControl(p.MMU).LCDOperation() {
				u.drawGraphics(p.RenderScreen())
			}

			time.Sleep(time.Second / 60)
		}
	}()

	u.win.ShowAndRun()

}

func (u *UI) setupApp() {
	u.a = app.New()
	u.win = u.a.NewWindow("Game Boy")

	raster := u.gbScreen()

	u.win.SetContent(raster)
}

var (
	keyByName = map[fyne.KeyName]input.Key{
		fyne.KeyZ:      input.KeyA,
		fyne.KeyX:      input.KeyB,
		fyne.KeyReturn: input.KeyStart,
		fyne.KeySpace:  input.KeySelect,
		fyne.KeyUp:     input.KeyUp,
		fyne.KeyDown:   input.KeyDown,
		fyne.KeyLeft:   input.KeyLeft,
		fyne.KeyRight:  input.KeyRight,
	}
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
