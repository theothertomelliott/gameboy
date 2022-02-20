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

		registers: newRegisters(gb),
	}

	ui.setupApp()
	ui.setupInput()
	return ui
}

type UI struct {
	gb *gameboy.DMG

	win fyne.Window
	app fyne.App

	registers *registers
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
