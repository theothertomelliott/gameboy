package fyneui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/theothertomelliott/gameboy"
)

func NewUI(gb *gameboy.DMG) *UI {
	data := NewDataTransport(gb)

	ui := &UI{
		gb: gb,

		data: data,
	}

	ui.setupApp()
	ui.setupInput()
	return ui
}

type UI struct {
	gb *gameboy.DMG

	win fyne.Window
	app fyne.App

	data DataTransport
}

func newMonoLabel(s string) *widget.Label {
	l := widget.NewLabel(s)
	l.TextStyle.Monospace = true
	return l
}
