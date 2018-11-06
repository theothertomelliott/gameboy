package main

import (
	"io/ioutil"
	"os"

	"github.com/faiface/pixel/pixelgl"
	"github.com/theothertomelliott/gameboy"
)

func main() {
	tracer := gameboy.NewTracer()
	defer tracer.Close()

	mmu := gameboy.NewMMU()
	cpu := gameboy.NewCPU(mmu, tracer)
	ppu := gameboy.NewPPU(mmu)

	data, err := ioutil.ReadFile(os.Args[1])
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

	control := gameboy.NewControl(cpu, ppu)
	control.Start()
	defer control.Stop()

	cui, err := setupCUI(cpu, tracer, control)
	if err != nil {
		panic(err)
	}
	defer cui.Close()

	go startCUI(cui)

	pixelgl.Run(func() {
		run(cpu, mmu, ppu)
	})
}
