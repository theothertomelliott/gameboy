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
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/input"
	"github.com/theothertomelliott/gameboy/ppu"
)

func NewUI(gb *gameboy.DMG) *UI {
	return &UI{
		gb: gb,

		a:  stringBinding("0x00"),
		f:  stringBinding("0x00"),
		b:  stringBinding("0x00"),
		c:  stringBinding("0x00"),
		d:  stringBinding("0x00"),
		e:  stringBinding("0x00"),
		h:  stringBinding("0x00"),
		l:  stringBinding("0x00"),
		sp: stringBinding("0x00"),
		pc: stringBinding("0x00"),
	}
}

func stringBinding(starting string) binding.String {
	b := binding.NewString()
	b.Set("0x00")
	return b
}

type UI struct {
	gb *gameboy.DMG

	win      fyne.Window
	img      image.Image
	app      fyne.App
	imgMutex sync.Mutex

	// Register content
	a binding.String
	f binding.String
	b binding.String
	c binding.String
	d binding.String
	e binding.String
	h binding.String
	l binding.String

	sp binding.String
	pc binding.String
}

func newMonoLabel(s string) *widget.Label {
	l := widget.NewLabel(s)
	l.TextStyle.Monospace = true
	return l
}

func newMonoLabelWithData(b binding.String) *widget.Label {
	l := widget.NewLabelWithData(b)
	l.TextStyle.Monospace = true
	return l
}

func (u *UI) Run() {
	p := u.gb.PPU()

	u.setupApp()

	if deskCanvas, ok := u.win.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			k, exists := keyByName[ke.Name]
			if !exists {
				fmt.Println("key not found: ", ke.Name)
			}
			u.gb.Input().Press(k)
		})

		deskCanvas.SetOnKeyUp(func(ke *fyne.KeyEvent) {
			k, exists := keyByName[ke.Name]
			if !exists {
				fmt.Println("key not found: ", ke.Name)
			}
			u.gb.Input().Release(k)
		})
	}

	go func() {
		for {
			if ppu.GetLCDControl(p.MMU).LCDOperation() {
				u.drawGraphics(p.RenderScreen())
			}
			u.updateDebugInfo()

			time.Sleep(time.Second / 60)
		}
	}()

	u.win.ShowAndRun()

}

func (u *UI) setupApp() {
	u.app = app.New()
	u.win = u.app.NewWindow("Game Boy")

	raster := u.gbScreen()

	c := container.NewHBox(raster, u.registerState())
	u.win.SetContent(c)
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

func (u *UI) registerState() fyne.CanvasObject {

	return container.NewVBox(
		newMonoLabel("Registers"),
		container.NewHBox(
			newMonoLabel("A"),
			newMonoLabelWithData(u.a),
			newMonoLabel("F"),
			newMonoLabelWithData(u.f),
		),
		container.NewHBox(
			newMonoLabel("B"),
			newMonoLabelWithData(u.b),
			newMonoLabel("C"),
			newMonoLabelWithData(u.c),
		),
		container.NewHBox(
			newMonoLabel("D"),
			newMonoLabelWithData(u.d),
			newMonoLabel("E"),
			newMonoLabelWithData(u.e),
		),
		container.NewHBox(
			newMonoLabel("H"),
			newMonoLabelWithData(u.h),
			newMonoLabel("L"),
			newMonoLabelWithData(u.l),
		),
		container.NewHBox(
			newMonoLabel("SP"),
			newMonoLabelWithData(u.sp),
			newMonoLabel("PC"),
			newMonoLabelWithData(u.pc),
		),
	)
}

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

func (u *UI) updateDebugInfo() {
	u.a.Set(fmt.Sprintf("%02X", u.gb.CPU().A.Read8()))
	u.f.Set(fmt.Sprintf("%02X", u.gb.CPU().F.Read8()))

	u.b.Set(fmt.Sprintf("%02X", u.gb.CPU().B.Read8()))
	u.c.Set(fmt.Sprintf("%02X", u.gb.CPU().C.Read8()))

	u.d.Set(fmt.Sprintf("%02X", u.gb.CPU().D.Read8()))
	u.e.Set(fmt.Sprintf("%02X", u.gb.CPU().E.Read8()))

	u.h.Set(fmt.Sprintf("%02X", u.gb.CPU().H.Read8()))
	u.l.Set(fmt.Sprintf("%02X", u.gb.CPU().L.Read8()))

	u.sp.Set(fmt.Sprintf("%04X", u.gb.CPU().SP.Read16()))
	u.pc.Set(fmt.Sprintf("%04X", u.gb.CPU().PC.Read16()))

	u.win.Content().Refresh()
}
