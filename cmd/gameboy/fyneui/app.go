package fyneui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	raster := u.gbScreen()

	c := container.NewHBox(raster, u.registerState())
	u.win.SetContent(c)
}
