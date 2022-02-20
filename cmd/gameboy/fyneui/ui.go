package fyneui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/theothertomelliott/gameboy"
)

func NewUI(gb *gameboy.DMG) *UI {
	ui := &UI{
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

	ui.setupApp()
	ui.setupInput()
	return ui
}

type UI struct {
	gb *gameboy.DMG

	win fyne.Window
	app fyne.App

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

func stringBinding(starting string) binding.String {
	b := binding.NewString()
	b.Set("0x00")
	return b
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
