package gameboy

type PPU struct {
	MMU        *MMU
	interrupts *InterruptScheduler
	modeclock  int
	mode       byte
	line       byte
	drawChan   chan struct{}
}

func NewPPU(mmu *MMU, interrupts *InterruptScheduler) *PPU {
	return &PPU{
		MMU:        mmu,
		interrupts: interrupts,
		drawChan:   make(chan struct{}),
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
	curStat = curStat & 0x3
	p.MMU.Write8(LCDSTAT, curStat|p.mode)

	// Handle line coincidence
	if p.line == p.MMU.Read8(CMPLINE) {
		curStat := p.MMU.Read8(LCDSTAT)
		p.MMU.Write8(LCDSTAT, curStat|0x4)
	}
}

func (p *PPU) Render() []byte {
	return nil
}

func (p *PPU) RenderWindow() [][]byte {
	return nil
}

func (p *PPU) LCDEnabled() bool {
	lcdControl := p.MMU.Read8(LCDCONT)
	enabled := bitValue(7, lcdControl)
	return enabled > 0
}

func (p *PPU) BackgroundEnabled() bool {
	lcdControl := p.MMU.Read8(LCDCONT)
	enabled := bitValue(0, lcdControl)
	return enabled > 0
}

func (p *PPU) ScrollX() byte {
	return p.MMU.Read8(SCROLLX)
}

func (p *PPU) ScrollY() byte {
	return p.MMU.Read8(SCROLLY)
}

func (p *PPU) RenderScreen() [][]byte {
	screen := p.RenderBackground()
	screen = p.RenderSprites(screen)
	return screen
}

func (p *PPU) RenderSprites(screen [][]byte) [][]byte {
	// Get tiles from sprite pattern table
	tiles := p.GetTilesByIndex(1)

	for pos := uint16(0xFE00); pos < 0xFE9F; pos += 4 {
		xPos := p.MMU.Read8(pos)
		x := int(xPos) - 16
		yPos := p.MMU.Read8(pos + 1)
		y := int(yPos) - 8
		tileNumber := p.MMU.Read8(pos + 2)
		//flags := p.MMU.Read8(pos + 3)

		renderedTile := tiles[tileNumber]

		// Skip offscreen sprites
		if x < 0 || y < 0 || x >= 168 || y >= 160 {
			continue
		}

		// Write the sprite to the screen
		for tr, rowValues := range renderedTile {
			for tc, value := range rowValues {
				screen[x+tr][y+tc] = value
			}
		}

	}
	return screen
}

func (p *PPU) GetTilesByIndex(tileDataSelect byte) [][][]byte {
	tileData := p.MMU.ReadRange(PatternTables[tileDataSelect])

	var tilesByIndex [][][]byte

	for offset := 0; offset < len(tileData); offset += 16 {
		tile := tileData[offset : offset+16]
		renderedTile := renderTile(tile, p.MMU.Read8(BGRDPAL))
		tilesByIndex = append(tilesByIndex, renderedTile)
	}

	return tilesByIndex
}

func (p *PPU) GetBackgroundTiles() [][][]byte {
	lcdControl := p.MMU.Read8(LCDCONT)

	// Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	tileDataSelect := bitValue(4, lcdControl)
	return p.GetTilesByIndex(tileDataSelect)
}

func (p *PPU) RenderBackground() [][]byte {
	lcdControl := p.MMU.Read8(LCDCONT)

	// Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	tileMapSelect := bitValue(3, lcdControl)
	tileMap := p.MMU.ReadRange(TileBackgroundMaps[tileMapSelect])
	tiles := p.GetBackgroundTiles()

	// Init 256x256 background
	var out = make([][]byte, 256)
	for i := range out {
		out[i] = make([]byte, 256)
	}
	if !p.BackgroundEnabled() {
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

			var total byte
			for _, rowValues := range renderedTile {
				for _, value := range rowValues {
					total += value
				}
			}

			if total == 0 {
				continue
			}

			// Write the tile into the background
			for tr, rowValues := range renderedTile {
				for tc, value := range rowValues {
					out[8*r+tr][8*c+tc] = value
				}
			}
		}
	}

	return out
}

func renderTile(tile []byte, palette byte) [][]byte {
	var out [][]byte = make([][]byte, 8)
	for line := 0; line < 8; line++ {
		var lineData = make([]byte, 8)
		high := tile[line*2+1]
		low := tile[line*2]
		for bit := byte(0); bit < 8; bit++ {
			h := bitValue(7-bit, high)
			l := bitValue(7-bit, low)
			colorValue := l + (h << 1)
			paletteValue := color(palette, colorValue)
			lineData[bit] = paletteValue
		}
		out[line] = lineData
	}
	return out
}

func color(palette byte, color byte) byte {
	// FF47 - BGP - BG Palette Data (R/W) - Non CGB Mode Only
	// This register assigns gray shades to the color numbers of the BG and Window tiles.
	//   Bit 7-6 - Shade for Color Number 3
	//   Bit 5-4 - Shade for Color Number 2
	//   Bit 3-2 - Shade for Color Number 1
	//   Bit 1-0 - Shade for Color Number 0
	switch color {
	case 3:
		return (palette & 0xC0) >> 6
	case 2:
		return (palette & 0x30) >> 4
	case 1:
		return (palette & 0xC) >> 2
	case 0:
		return palette & 0x3
	}
	return color
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

// Bit 7 - LCD Display Enable             (0=Off, 1=On)
// Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
// Bit 5 - Window Display Enable          (0=Off, 1=On)
// Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
// Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
// Bit 2 - OBJ (Sprite) Size              (0=8x8, 1=8x16)
// Bit 1 - OBJ (Sprite) Display Enable    (0=Off, 1=On)
// Bit 0 - BG Display (for CGB see below) (0=Off, 1=On)

var (
	// Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	PatternTables = []Range{
		Range{
			Start: 0x8800,
			End:   0x97FF,
		},
		Range{
			Start: 0x8000,
			End:   0x8FFF,
		},
	}
	// Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	TileBackgroundMaps = []Range{
		Range{
			Start: 0x9800,
			End:   0x9BFF,
		},
		Range{
			Start: 0x9C00,
			End:   0x9FFF,
		},
	}
	OAM = Range{
		Start: 0xFE00,
		End:   0xFE9F,
	}
)
