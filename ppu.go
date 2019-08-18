package gameboy

import (
	"image"
	"image/color"
)

type PPU struct {
	MMU        *MMU
	interrupts *InterruptScheduler
	modeclock  int
	mode       byte
	line       byte
}

func NewPPU(mmu *MMU, interrupts *InterruptScheduler) *PPU {
	return &PPU{
		MMU:        mmu,
		interrupts: interrupts,
	}
}

func (p *PPU) Step(t int) error {
	defer p.setStatus()
	p.modeclock += t

	switch p.mode {
	// OAM read mode, scanline active
	case 2:
		if p.modeclock >= 80 {
			// Enter scanline mode 3
			p.modeclock = 0
			p.setMode(3)
		}
	// VRAM read mode, scanline active
	// Treat end of mode 3 as end of scanline
	case 3:
		if p.modeclock >= 172 {
			// Enter hblank
			p.modeclock = 0
			p.setMode(0)

			// render a scanline
		}
	// Hblank
	// After the last hblank, push the screen data to canvas
	case 0:
		if p.modeclock >= 204 {
			p.modeclock = 0
			p.incLine()

			if p.line == 143 {
				// Enter vblank
				p.setMode(1)
			} else {
				p.setMode(2)
			}
		}

	// Vblank (10 lines)
	case 1:
		if p.modeclock >= 456 {
			p.modeclock = 0
			p.incLine()

			if p.line > 153 {
				// Restart scanning modes
				p.setMode(2)
				p.setLine(0)
			}
		}
	}

	return nil
}

func (p *PPU) resetLine() {
	p.line = 0
}

func (p *PPU) incLine() {
	p.setLine(p.line + 1)
}

func (p *PPU) setLine(line byte) {
	p.line = line
	if p.MMU.Read8(CURLINE) != p.line {
		// Write the current line to memory
		p.MMU.Write8(CURLINE, p.line)
	}

	// Handle line coincidence
	if p.line == p.MMU.Read8(CMPLINE) {
		curStat := p.MMU.Read8(LCDSTAT)
		if bitValue(6, curStat) != 0 {
			p.interrupts.ScheduleInterrupt(InterruptLCDStatus)
		}
	}
}

func (p *PPU) setMode(newMode byte) {
	if p.mode == newMode {
		return
	}

	curStat := p.MMU.Read8(LCDSTAT)
	if newMode == 0 {
		if bitValue(3, curStat) != 0 {
			p.interrupts.ScheduleInterrupt(InterruptLCDStatus)
		}
	}
	if newMode == 1 {
		if bitValue(4, curStat) != 0 {
			p.interrupts.ScheduleInterrupt(InterruptLCDStatus)
		}
		p.interrupts.ScheduleInterrupt(InterruptVBlank)
	}
	if newMode == 2 {
		if bitValue(5, curStat) != 0 {
			p.interrupts.ScheduleInterrupt(InterruptLCDStatus)
		}
		p.interrupts.ScheduleInterrupt(InterruptVBlank)
	}

	p.mode = newMode
}

func (p *PPU) setStatus() {
	curStat := p.MMU.Read8(LCDSTAT)
	// Set mode stat
	curStat = curStat & 0xF8
	curStat = curStat | p.mode

	// Handle line coincidence
	if p.line == p.MMU.Read8(CMPLINE) {
		curStat = curStat | 0x4
	}

	p.MMU.Write8(LCDSTAT, curStat|0x4)
}

func (p *PPU) LCDControl() LCDControl {
	lcdControl := p.MMU.Read8(LCDCONT)
	return LCDControl(lcdControl)
}

func (p *PPU) ScrollX() byte {
	return p.MMU.Read8(SCROLLX)
}

func (p *PPU) ScrollY() byte {
	return p.MMU.Read8(SCROLLY)
}

func (p *PPU) RenderScreen() image.Image {
	return NewScreen(p)
}

func (p *PPU) GetTilesForRange(r Range) []Tile {
	tileData := p.MMU.ReadRange(r)

	var tilesByIndex []Tile

	for offset := 0; offset < len(tileData); offset += 16 {
		tile := tileData[offset : offset+16]
		tilesByIndex = append(tilesByIndex, NewTile(tile))
	}

	return tilesByIndex
}

func (p *PPU) GetBackgroundTiles() []Tile {
	return p.GetTilesForRange(p.LCDControl().TilePatternTableAddress())
}

func (p *PPU) RenderBackground() image.Image {
	return NewBackground(p)
}

func newColor(palette byte, color byte) color.RGBA {
	// FF47 - BGP - BG Palette Data (R/W) - Non CGB Mode Only
	// This register assigns gray shades to the color numbers of the BG and Window tiles.
	//   Bit 7-6 - Shade for Color Number 3
	//   Bit 5-4 - Shade for Color Number 2
	//   Bit 3-2 - Shade for Color Number 1
	//   Bit 1-0 - Shade for Color Number 0
	switch color {
	case 3:
		return colorForValue((palette & 0xC0) >> 6)
	case 2:
		return colorForValue((palette & 0x30) >> 4)
	case 1:
		return colorForValue((palette & 0xC) >> 2)
	case 0:
		return colorForValue(palette & 0x3)
	}
	return colorForValue(color)
}

func colorForValue(value uint8) color.RGBA {
	adjustedValue := 255 - ((256 / 4) * value)
	return color.RGBA{
		R: adjustedValue, G: adjustedValue, B: adjustedValue, A: 255,
	}
}
