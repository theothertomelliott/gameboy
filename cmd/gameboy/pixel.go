package main

import (
	"time"

	"github.com/theothertomelliott/gameboy"
)

func run(
	gb *gameboy.DMG,
) {
	ppu := gb.PPU()

	setupGraphics()
	setupMemView()

	for !win.Closed() {
		if ppu.LCDEnabled() {
			bg := ppu.RenderBackground()
			drawGraphics(
				bg,
				ppu.ScrollX(),
				ppu.ScrollY(),
			)
		}

		if !memWin.Closed() {
			drawMemory(
				gb.MMU(),
				gb.PPU(),
			)
			memWin.UpdateInput()
		}
		win.UpdateInput()

		select {
		case <-time.After(time.Second / 60):
		}
	}

}
