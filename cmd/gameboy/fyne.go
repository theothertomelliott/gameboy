package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/input"
	"github.com/theothertomelliott/gameboy/ppu"
)

func run(
	gb *gameboy.DMG,
) {
	p := gb.PPU()

	setupGraphics()

	if deskCanvas, ok := win.Canvas().(desktop.Canvas); ok {
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
				drawGraphics(p.RenderScreen())
			}

			select {
			case <-time.After(time.Second / 60):
			}
		}
	}()

	win.ShowAndRun()

}

// Store key press state (Press and Release)
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
