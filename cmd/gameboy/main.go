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

	if len(os.Args) > 2 {
		data, err = ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		cpu.LoadROM(data)
	} else {
		cpu.Init()
	}

	data, err = ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	mmu.LoadCartridge(data)

	clock := gameboy.NewClock()
	videoClock := time.NewTicker(time.Second / 60)
	defer videoClock.Stop()
	defer clock.Stop()

	go cpu.Run(clock.C)

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
		_ = <-videoClock.C
	}

	clock.Stop()
	videoClock.Stop()

}
