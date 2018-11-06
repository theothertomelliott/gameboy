package main

import (
	"time"

	"github.com/theothertomelliott/gameboy"
)

func run(
	cpu *gameboy.CPU,
	mmu *gameboy.MMU,
	ppu *gameboy.PPU,
) {
	setupGraphics()
	//setupMemView()

	for !win.Closed() {
		if ppu.LCDEnabled() {
			bg := ppu.RenderBackground()
			drawGraphics(
				bg,
				ppu.ScrollX(),
				ppu.ScrollY(),
			)
		}

		// if !memWin.Closed() {
		// 	drawMemory(
		// 		mmu,
		// 		ppu,
		// 	)
		// 	memWin.UpdateInput()
		// }
		win.UpdateInput()

		select {
		case <-time.After(time.Second / 60):
		}
	}

}
