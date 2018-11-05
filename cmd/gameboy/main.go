package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/theothertomelliott/gameboy"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	pixelgl.Run(run)
}

func run() {
	mmu := gameboy.NewMMU()
	cpu := gameboy.NewCPU(mmu)
	ppu := gameboy.NewPPU(mmu)

	setupGraphics()
	setupMemView()

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

	//cpu.Trace = true

	clock := time.NewTicker(time.Microsecond / 4)
	defer clock.Stop()

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

		if !memWin.Closed() {
			drawMemory(
				mmu,
				ppu,
			)
			memWin.UpdateInput()
		}
		win.UpdateInput()

		_ = <-ppu.ShouldDraw()
	}

}
