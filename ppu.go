package gameboy

type PPU struct {
	MMU *MMU
}

func NewPPU(mmu *MMU) *PPU {
	return &PPU{
		MMU: mmu,
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

func (p *PPU) GetBackgroundTiles() [][][]byte {
	lcdControl := p.MMU.Read8(LCDCONT)

	// Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	tileDataSelect := bitValue(4, lcdControl)

	tileData := p.MMU.ReadRange(PatternTables[tileDataSelect])

	var tilesByIndex [][][]byte

	for offset := 0; offset < len(tileData); offset += 16 {
		tile := tileData[offset : offset+16]
		renderedTile := renderTile(tile, p.MMU.Read8(BGRDPAL))
		tilesByIndex = append(tilesByIndex, renderedTile)
	}

	return tilesByIndex
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
			tileRef := tileMap[r*32+c]
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
	var out [][]byte
	for line := 0; line < 8; line++ {
		var lineData []byte
		high := tile[line*2+1]
		low := tile[line*2]
		for bit := byte(0); bit < 8; bit++ {
			h := bitValue(7-bit, high)
			l := bitValue(7-bit, low)
			colorValue := l + (h << 1)
			paletteValue := color(palette, colorValue)
			lineData = append(lineData, paletteValue)
		}
		out = append(out, lineData)
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
