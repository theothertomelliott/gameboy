package main

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/input"
	"github.com/theothertomelliott/gameboy/ppu"
)

func run(
	gb *gameboy.DMG,
) {
	p := gb.PPU()

	setupGraphics()

	for !win.Closed() {
		if ppu.GetLCDControl(p.MMU).LCDOperation() {
			drawGraphics(p.RenderScreen())
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
	keyByIndex = map[input.Key]pixelgl.Button{
		input.KeyA:      pixelgl.KeyZ,
		input.KeyB:      pixelgl.KeyX,
		input.KeyStart:  pixelgl.KeyEnter,
		input.KeySelect: pixelgl.KeySpace,
		input.KeyUp:     pixelgl.KeyUp,
		input.KeyDown:   pixelgl.KeyDown,
		input.KeyLeft:   pixelgl.KeyLeft,
		input.KeyRight:  pixelgl.KeyRight,
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
