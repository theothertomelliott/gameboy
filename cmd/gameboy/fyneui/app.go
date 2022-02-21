package fyneui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func (u *UI) setupApp() {
	u.app = app.New()
	u.win = u.app.NewWindow("Game Boy")

	u.win.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu(
			"test menu",
			fyne.NewMenuItem("test menu item", func() {}),
		),
	))

	reg := newRegisters(u.data)
	mem := newMemory(u.data)

	c := container.NewHBox(
		newScreen(u.data, fyne.Size{
			Width:  600,
			Height: 600,
		}),
		container.New(
			layout.NewBorderLayout(reg, nil, nil, nil),
			reg, mem,
		),
	)
	u.win.SetContent(c)
}
