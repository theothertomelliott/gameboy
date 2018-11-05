package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"

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

	control := NewControl(time.Nanosecond)
	defer control.Close()

	cui, err := setupCUI(cpu, control, tracer)
	if err != nil {
		panic(err)
	}
	defer cui.Close()

	go startCUI(cui)

	pixelgl.Run(func() {
		run(cpu, mmu, ppu, control.C)
	})
}
