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

func (p *PPU) RenderScreen() image.Image {
	return NewScreen(p)
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

func (p *PPU) RenderBackground() image.Image {
	return NewBackground(p)
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

var _ image.Image = &Screen{}

func NewBackground(ppu *PPU) *Screen {
	return &Screen{
		BGTileMap:        ppu.MMU.ReadRange(ppu.LCDControl().BackgroundTileTableAddress()),
		BGTiles:          ppu.GetBackgroundTiles(),
		SpriteTiles:      ppu.GetTilesForRange(ppu.LCDControl().TilePatternTableAddress()),
		SpriteData:       ppu.MMU.ReadRange(Range{Start: 0xFE00, End: 0xFE9F}),
		RenderBackground: ppu.LCDControl().BackgroundDisplay(),

		Position: image.Point{
			X: 0,
			Y: 0,
		},

		bounds: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{255, 255},
		},
	}
}

func NewScreen(ppu *PPU) *Screen {
	return &Screen{
		BGTileMap:        ppu.MMU.ReadRange(ppu.LCDControl().BackgroundTileTableAddress()),
		BGTiles:          ppu.GetBackgroundTiles(),
		SpriteTiles:      ppu.GetTilesForRange(ppu.LCDControl().TilePatternTableAddress()),
		SpriteData:       ppu.MMU.ReadRange(Range{Start: 0xFE00, End: 0xFE9F}),
		RenderBackground: ppu.LCDControl().BackgroundDisplay(),

		Position: image.Point{
			X: int(ppu.ScrollX()),
			Y: int(ppu.ScrollY()),
		},

		bounds: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{160, 144},
		},

		RenderSprites: true,
	}
}

type Screen struct {
	BGTileMap        []byte
	BGTiles          []*image.RGBA
	SpriteTiles      []*image.RGBA
	SpriteData       []byte
	RenderBackground bool
	RenderSprites    bool

	Position image.Point
	bounds   image.Rectangle
}

var _ color.Model = &ColorModel{}

type ColorModel struct {
}

func (c *ColorModel) Convert(in color.Color) color.Color {
	return in
}

func (s *Screen) ColorModel() color.Model {
	return &ColorModel{}
}

func (s *Screen) Bounds() image.Rectangle {
	return s.bounds
}

func (s *Screen) atBg(x, y int) color.Color {
	x = x % 255
	y = y % 255

	tileIndex := y/8*32 + x/8
	if tileIndex < 0 || tileIndex >= len(s.BGTileMap) {
		return color.Black
	}
	tileRef := s.BGTileMap[tileIndex]
	renderedTile := s.BGTiles[tileRef]
	return renderedTile.At(x%8, y%8)
}

func (s *Screen) atSprite(pX, pY int) color.Color {
	if !s.RenderSprites {
		return nil
	}
	for pos := 0; pos < len(s.SpriteData); pos += 4 {
		yPos := s.SpriteData[pos]
		y := int(yPos) - 16
		xPos := s.SpriteData[pos+1]
		x := int(xPos) - 8

		x += 8
		y += 8

		if !(x-pX >= 0 && x-pX < 8 && y-pY >= 0 && y-pY < 8) {
			continue
		}

		tileNumber := s.SpriteData[pos+2]

		flags := s.SpriteData[pos+3]

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

		spX := pX % 8
		spY := pY % 8

		if xFlip != 0 {
			spX = 8 - spX
		}
		if yFlip != 0 {
			spY = 8 - spY
		}

		renderedTile := s.SpriteTiles[tileNumber]
		return renderedTile.At(spX, spY)
	}
	return nil
}

func (s *Screen) At(x, y int) color.Color {
	xS := x + s.Position.X
	yS := y + s.Position.Y

	bgPixel := s.atBg(xS, yS)
	spritePixel := s.atSprite(xS, yS)
	if notBlank(spritePixel) {
		return spritePixel
	}
	return bgPixel
}

func notBlank(c color.Color) bool {
	if c == nil {
		return false
	}
	return true
}
