package fyneui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/theothertomelliott/gameboy/input"
)

func (u *UI) setupInput() {
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
