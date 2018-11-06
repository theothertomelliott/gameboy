package main

import (
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	"github.com/faiface/pixel/pixelgl"
	"github.com/theothertomelliott/gameboy"
)

func main() {
	f, err := os.Create("profile.out")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

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
