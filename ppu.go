package gameboy

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/disintegration/imaging"
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

func (p *PPU) Render() []byte {
	return nil
}

func (p *PPU) RenderWindow() [][]byte {
	return nil
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

func (p *PPU) RenderScreen() *image.RGBA {
	screen := p.RenderBackground()
	screen = p.RenderSprites(screen)
	scrolled := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{160, 144}})
	draw.Draw(scrolled, screen.Bounds(), screen, image.Point{int(p.ScrollX()), int(p.ScrollY())}, draw.Src)
	return scrolled
}

func (p *PPU) RenderSprites(screen *image.RGBA) *image.RGBA {
	// Get tiles from sprite pattern table
	tiles := p.GetTilesForRange(p.LCDControl().TilePatternTableAddress())

	for pos := uint16(0xFE00); pos < 0xFE9F; pos += 4 {
		xPos := p.MMU.Read8(pos)
		x := int(xPos) - 16
		yPos := p.MMU.Read8(pos + 1)
		y := int(yPos) - 8
		tileNumber := p.MMU.Read8(pos + 2)
		flags := p.MMU.Read8(pos + 3)

		//priority := bitValue(7, flags)
		// 		Bit7 Priority
		//  If this bit is set to 0, sprite is
		//  displayed on top of background & window.
		//  If this bit is set to 1, then sprite
		//  will be hidden behind colors 1, 2, and 3
		//  of the background & window. (Sprite only
		//  prevails over color 0 of BG & win.)

		yFlip := bitValue(6, flags)
		xFlip := bitValue(5, flags)

		//paletteNum := bitValue(4, flags)
		// Bit4 Palette number
		// Sprite colors are taken from OBJ1PAL if
		// this bit is set to 1 and from OBJ0PAL
		// otherwise.

		renderedTile := tiles[tileNumber]

		// Skip offscreen sprites
		if x < 0 || y < 0 || x >= 168 || y >= 160 {
			continue
		}

		bounds := image.Rectangle{
			Min: image.Point{
				y,
				x,
			},
			Max: image.Point{
				y + 8,
				x + 8,
			},
		}

		var imgOut image.Image = renderedTile
		if yFlip == 1 {
			imgOut = imaging.FlipV(imgOut)
		}
		if xFlip == 1 {
			imgOut = imaging.FlipH(imgOut)
		}

		// Write the sprite tile into the background
		draw.FloydSteinberg.Draw(screen, bounds, imgOut, image.ZP)
	}
	return screen
}

func (p *PPU) GetTilesForRange(r Range) []*image.RGBA {
	tileData := p.MMU.ReadRange(r)

	var tilesByIndex []*image.RGBA

	for offset := 0; offset < len(tileData); offset += 16 {
		tile := tileData[offset : offset+16]
		renderedTile := renderTile(tile, p.MMU.Read8(BGRDPAL))
		tilesByIndex = append(tilesByIndex, renderedTile)
	}

	return tilesByIndex
}

func (p *PPU) GetBackgroundTiles() []*image.RGBA {
	return p.GetTilesForRange(p.LCDControl().TilePatternTableAddress())
}

func (p *PPU) RenderBackground() *image.RGBA {
	tileMap := p.MMU.ReadRange(p.LCDControl().BackgroundTileTableAddress())
	tiles := p.GetBackgroundTiles()

	// Init 256x256 background
	var out = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{256, 256}})
	if !p.LCDControl().BackgroundDisplay() {
		return out
	}

	for r := 0; r < 32; r++ {
		for c := 0; c < 32; c++ {
			index := r*32 + c
			if index >= len(tileMap) {
				continue
			}
			tileRef := tileMap[index]
			renderedTile := tiles[tileRef]

			bounds := image.Rectangle{
				Min: image.Point{
					8 * c,
					8 * r,
				},
				Max: image.Point{
					8*c + 8,
					8*r + 8,
				},
			}

			// Write the tile into the background
			draw.Draw(out, bounds, renderedTile, image.ZP, draw.Src)
		}
	}

	return out
}

func renderTile(tile []byte, palette byte) *image.RGBA {
	var out = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{8, 8}})
	for line := 0; line < 8; line++ {
		high := tile[line*2+1]
		low := tile[line*2]
		for bit := byte(0); bit < 8; bit++ {
			h := bitValue(7-bit, high)
			l := bitValue(7-bit, low)
			colorValue := l + (h << 1)
			paletteValue := newColor(palette, colorValue)
			out.Set(int(bit), line, paletteValue)
		}
	}
	return out
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

func getTile(number byte, tileMap []byte) []byte {
	start := uint16(number) * 16
	return tileMap[start : start+16]
}

// Detail at: http://bgb.bircd.org/pandocs.htm#videodisplay

// Tile pattern tables & Sprite Pattern Table
//  $8000-8FFF -	sprites, the background, and the window display
//					Its tiles are numbered from 0 to 255.
//  $8800-97FF - The second table can be used for the background and
//				 the window display and its tiles are numbered from
//				 -128 to 127.
// Background Tile Map

// Tile background maps
//  $9800-9BFF
//  $9C00-9FFF

// Sprite attributes reside in the Sprite Attribute Table (OAM - Object Attribute Memory) at $FE00-FE9F

var (
	OAM = Range{
		Start: 0xFE00,
		End:   0xFE9F,
	}
)
