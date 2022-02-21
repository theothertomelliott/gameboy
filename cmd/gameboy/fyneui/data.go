package fyneui

import (
	"image"
	"sync"

	"fyne.io/fyne/v2/data/binding"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/ppu"
)

func NewDataTransport(gb *gameboy.DMG) DataTransport {
	return &dataTransport{
		gb: gb,

		a:  binding.NewInt(),
		f:  binding.NewInt(),
		b:  binding.NewInt(),
		c:  binding.NewInt(),
		d:  binding.NewInt(),
		e:  binding.NewInt(),
		h:  binding.NewInt(),
		l:  binding.NewInt(),
		sp: binding.NewInt(),
		pc: binding.NewInt(),
	}
}

// DataTransport handles updates to state data.
type DataTransport interface {
	binding.DataItem
	Screen() image.Image

	Update()

	A() binding.Int
	F() binding.Int
	B() binding.Int
	C() binding.Int
	D() binding.Int
	E() binding.Int
	H() binding.Int
	L() binding.Int

	SP() binding.Int
	PC() binding.Int
}

var _ DataTransport = &dataTransport{}

type dataTransport struct {
	gb *gameboy.DMG

	screen image.Image

	listenersMtx sync.Mutex
	listeners    []binding.DataListener

	// Registers
	a binding.Int
	f binding.Int
	b binding.Int
	c binding.Int
	d binding.Int
	e binding.Int
	h binding.Int
	l binding.Int

	sp binding.Int
	pc binding.Int
}

func (dt *dataTransport) Screen() image.Image {
	return dt.screen
}

func (dt *dataTransport) Update() {
	if !ppu.GetLCDControl(dt.gb.MMU()).LCDOperation() {
		return
	}
	dt.screen = dt.gb.PPU().RenderScreen()

	dt.a.Set(int(dt.gb.CPU().A.Read8()))
	dt.f.Set(int(dt.gb.CPU().F.Read8()))

	dt.b.Set(int(dt.gb.CPU().B.Read8()))
	dt.c.Set(int(dt.gb.CPU().C.Read8()))

	dt.d.Set(int(dt.gb.CPU().D.Read8()))
	dt.e.Set(int(dt.gb.CPU().E.Read8()))

	dt.h.Set(int(dt.gb.CPU().H.Read8()))
	dt.l.Set(int(dt.gb.CPU().L.Read8()))

	dt.sp.Set(int(dt.gb.CPU().SP.Read16()))
	dt.pc.Set(int(dt.gb.CPU().PC.Read16()))

	dt.listenersMtx.Lock()
	defer dt.listenersMtx.Unlock()
	for _, ln := range dt.listeners {
		ln.DataChanged()
	}
}

func (dt *dataTransport) AddListener(l binding.DataListener) {
	dt.listenersMtx.Lock()
	defer dt.listenersMtx.Unlock()

	dt.listeners = append(dt.listeners, l)
	l.DataChanged()
}

func (dt *dataTransport) RemoveListener(l binding.DataListener) {
	dt.listenersMtx.Lock()
	defer dt.listenersMtx.Unlock()

	var newListeners []binding.DataListener
	for _, ln := range dt.listeners {
		if ln != l {
			newListeners = append(newListeners, ln)
		}
	}

	dt.listeners = newListeners
}

func (dt *dataTransport) A() binding.Int {
	return dt.a
}

func (dt *dataTransport) F() binding.Int {
	return dt.f
}

func (dt *dataTransport) B() binding.Int {
	return dt.b
}

func (dt *dataTransport) C() binding.Int {
	return dt.c
}

func (dt *dataTransport) D() binding.Int {
	return dt.d
}

func (dt *dataTransport) E() binding.Int {
	return dt.e
}

func (dt *dataTransport) H() binding.Int {
	return dt.h
}

func (dt *dataTransport) L() binding.Int {
	return dt.l
}

func (dt *dataTransport) SP() binding.Int {
	return dt.sp
}

func (dt *dataTransport) PC() binding.Int {
	return dt.pc
}
