package main

import (
	"image/color"

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
		Bounds: pixel.R(0, 0, 1000, 800),
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

func drawMemory(mmu *gameboy.MMU, ppu *gameboy.PPU) {
	memWin.Clear(colornames.Black)
	drawTiles(ppu)
	drawRAM(mmu)
	memWin.Update()
}

func drawRAM(mmu *gameboy.MMU) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 0)
	for x := uint16(0); x < 256; x++ {
		for y := uint16(0); y < 256; y++ {
			value := mmu.Read8(x*256 + y)
			if value == 0 {
				continue
			}

			R := 256 / float64((value&0xE0)>>5) // rgb8 & 1110 0000  >> 5
			G := 256 / float64((value&0x1C)>>2) // rgb8 & 0001 1100  >> 2
			B := 256 / float64(value&0x03)      // rgb8 & 0000 0011
			color := pixel.RGB(R, G, B)
			drawPixel(
				imd,
				pixel.V(float64(x), float64(y)),
				memoryScale(),
				color,
			)
		}
	}
	imd.Draw(memWin)
}

func drawTiles(ppu *gameboy.PPU) {
	imd := imdraw.New(nil)
	tiles := ppu.GetBackgroundTiles()

	drawPixel(imd, pixel.V(800, 0), pixel.V(200, 800), colornames.White)
	for tileIndex, tile := range tiles {
		drawBoundingRect(imd, pixel.V(256, float64(tileIndex*8)), pixel.V(8, 8), memoryScale(), pixel.RGB(1, 0, 0))
		for y, row := range tile {
			for x, value := range row {
				colorVal := 1.0 - (float64(value) / 4)
				color := pixel.RGB(colorVal, colorVal, colorVal)
				color = pixel.RGB(1, 1, 1)
				if value != 0 {
					drawPixel(
						imd,
						pixel.V(float64(256+x), float64(tileIndex*8+(8-y))),
						memoryScale(),
						color,
					)
				}
			}
		}
	}
	imd.Draw(memWin)
}

func memoryScale() pixel.Vec {
	return pixel.V(float64(800/memSizeX), float64(800/memSizeY))
}

func tileScale() pixel.Vec {
	return pixel.V(float64(200/8), float64(200/8))
}

func drawPixel(imd *imdraw.IMDraw, pos, scale pixel.Vec, color color.Color) {
	drawRect(imd, pos, pixel.V(1, 1), scale, color)
}

func drawBoundingRect(imd *imdraw.IMDraw, pos, size, scale pixel.Vec, color color.Color) {
	posScaled := pos.ScaledXY(scale)
	sizeScaled := size.ScaledXY(scale)

	imd.Push(posScaled)
	imd.Push(posScaled.Add(sizeScaled))
	imd.Color = color
	imd.Rectangle(1)
}

func drawRect(imd *imdraw.IMDraw, pos, size, scale pixel.Vec, color color.Color) {
	posScaled := pos.ScaledXY(scale)
	sizeScaled := size.ScaledXY(scale)

	imd.Push(posScaled)
	imd.Push(posScaled.Add(sizeScaled))
	imd.Color = color
	imd.Rectangle(0)
}
