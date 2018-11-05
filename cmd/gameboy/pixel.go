package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/theothertomelliott/gameboy"
)

func run(
	cpu *gameboy.CPU,
	mmu *gameboy.MMU,
	ppu *gameboy.PPU,
	clock chan struct{},
) {
	setupGraphics()
	//setupMemView()

	var (
		data []byte
		err  error
	)

	data, err = ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	mmu.LoadCartridge(data)

	if len(os.Args) > 2 {
		data, err = ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		mmu.LoadROM(data)
	} else {
		cpu.Init()
	}

	go func() {
		for true {
			gameboy.Step(cpu, ppu)
		}
	}()

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
