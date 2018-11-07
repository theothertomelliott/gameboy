package main

import (
	"io/ioutil"
	"os"

	"github.com/faiface/pixel/pixelgl"
	"github.com/theothertomelliott/gameboy"
)

func main() {
	gb := gameboy.NewDMG()

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	gb.MMU().LoadCartridge(data)

	if len(os.Args) > 2 {
		data, err = ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		gb.MMU().LoadROM(data)
	} else {
		gb.CPU().Init()
	}

	gb.Breakpoints = []uint16{0x0095, 0xC3C3}

	gb.Start()
	defer gb.Stop()

	term := NewTerminalUI(gb)
	defer term.Stop()

	go func() {
		err := term.Run()
		if err != nil {
			panic(err)
		}
	}()

	pixelgl.Run(func() {
		run(gb)
	})
}
