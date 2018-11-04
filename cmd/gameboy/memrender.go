package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/theothertomelliott/gameboy"
	"golang.org/x/image/colornames"
)

const (
	memSizeX, memSizeY = 256, 256
)

var (
	memWin *pixelgl.Window
)

func setupMemView() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game Boy RAM",
		Bounds: pixel.R(0, 0, 800, 800),
		VSync:  true,
	}
	var err error
	memWin, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	// Workaround for Mojave
	// https://github.com/faiface/pixel/issues/140
	memWin.SetPos(win.GetPos().Add(pixel.V(0, 1)))
}

func drawMemory(mmu *gameboy.MMU) {
	memWin.Clear(colornames.Black)
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 0)
	width, height := float64(memWin.Bounds().W()/memSizeX), float64(memWin.Bounds().H()/memSizeY)
	for x := uint16(0); x < 256; x++ {
		for y := uint16(0); y < 256; y++ {
			value := mmu.Read8(x*256 + y)
			if value == 0 {
				continue
			}
			imd.Push(pixel.V(width*float64(x), height*float64(y)))
			imd.Push(pixel.V(width*float64(x)+width, height*float64(y)+height))
			R := 256 / float64((value&0xE0)>>5) // rgb8 & 1110 0000  >> 5
			G := 256 / float64((value&0x1C)>>2) // rgb8 & 0001 1100  >> 2
			B := 256 / float64(value&0x03)      // rgb8 & 0000 0011

			imd.Color = pixel.RGB(R, G, B)
			imd.Rectangle(0)
		}
	}
	imd.Draw(memWin)
	memWin.Update()
}
