package main

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/theothertomelliott/gameboy"
)

func run(
	gb *gameboy.DMG,
) {
	ppu := gb.PPU()

	setupGraphics()

	for !win.Closed() {
		if ppu.LCDControl().LCDOperation() {
			drawGraphics(ppu.RenderScreen())
		}

		win.UpdateInput()

		handleKeys(gb)

		select {
		case <-time.After(time.Second / 60):
		}
	}

}

// Store key press state (Press and Release)
var (
	keyByIndex = map[gameboy.Key]pixelgl.Button{
		gameboy.KeyA:      pixelgl.KeyZ,
		gameboy.KeyB:      pixelgl.KeyX,
		gameboy.KeyStart:  pixelgl.KeyEnter,
		gameboy.KeySelect: pixelgl.KeySpace,
		gameboy.KeyUp:     pixelgl.KeyUp,
		gameboy.KeyDown:   pixelgl.KeyDown,
		gameboy.KeyLeft:   pixelgl.KeyLeft,
		gameboy.KeyRight:  pixelgl.KeyRight,
	}
	keysDown [16]*time.Ticker
)

func handleKeys(gb *gameboy.DMG) {
	for index, key := range keyByIndex {
		if !win.Pressed(key) {
			gb.Input().Release(index)
		} else if win.Pressed(key) {
			gb.Input().Press(index)
		}
	}
}
