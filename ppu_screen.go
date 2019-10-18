package gameboy

import (
	"image"
	"image/color"
)

var _ image.Image = &Screen{}

// https://gbdev.gg8.se/wiki/articles/Video_Display#VRAM_Sprite_Attribute_Table_.28OAM.29

func NewBackground(mmu *MMU) *Screen {
	return &Screen{
		BGTileMap:        mmu.ReadRange(GetLCDControl(mmu).BackgroundTileTableAddress()),
		BGTiles:          GetBackgroundTiles(mmu),
		SpriteTiles:      GetTilesForRange(mmu, GetLCDControl(mmu).TilePatternTableAddress()),
		SpriteData:       mmu.ReadRange(Range{Start: 0xFE00, End: 0xFE9F}),
		RenderBackground: GetLCDControl(mmu).BackgroundDisplay(),
		RenderWindow:     false,

		BGRDPAL: mmu.Read8(BGRDPAL),
		OBJ0PAL: mmu.Read8(OBJ0PAL),
		OBJ1PAL: mmu.Read8(OBJ1PAL),

		// Off screen
		WindowPos: image.Point{
			X: 167,
			Y: 144,
		},
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

func NewScreen(mmu *MMU) *Screen {
	lcdControl := GetLCDControl(mmu)
	return &Screen{
		BGTileMap:        mmu.ReadRange(lcdControl.BackgroundTileTableAddress()),
		BGTiles:          GetBackgroundTiles(mmu),
		WindowTiles:      GetWindowTiles(mmu),
		SpriteTiles:      GetTilesForRange(mmu, lcdControl.TilePatternTableAddress()),
		SpriteData:       mmu.ReadRange(Range{Start: 0xFE00, End: 0xFE9F}),
		RenderBackground: lcdControl.BackgroundDisplay(),
		RenderWindow:     lcdControl.WindowDisplay(),

		BGRDPAL: mmu.Read8(BGRDPAL),
		OBJ0PAL: mmu.Read8(OBJ0PAL),
		OBJ1PAL: mmu.Read8(OBJ1PAL),

		WindowPos: GetWindowPos(mmu),
		Position:  GetScroll(mmu),

		bounds: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{160, 144},
		},

		RenderSprites: true,
	}
}

type Screen struct {
	BGTileMap        []byte
	WindowTiles      []Tile
	BGTiles          []Tile
	SpriteTiles      []Tile
	SpriteData       []byte
	RenderBackground bool
	RenderSprites    bool
	RenderWindow     bool

	BGRDPAL byte
	OBJ0PAL byte
	OBJ1PAL byte

	WindowPos image.Point
	Position  image.Point
	bounds    image.Rectangle
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

func (s *Screen) atBg(x, y int) byte {
	x = x % 255
	y = y % 255

	tileIndex := y/8*32 + x/8
	if tileIndex < 0 || tileIndex >= len(s.BGTileMap) {
		return 0
	}
	tileRef := s.BGTileMap[tileIndex]
	renderedTile := s.BGTiles[tileRef]
	return valueInPalette(s.BGRDPAL, renderedTile.At(x%8, y%8))
}

func (s *Screen) atWindow(x, y int) byte {
	tileIndex := y/8*32 + x/8
	if tileIndex < 0 || tileIndex >= len(s.BGTileMap) {
		return 0
	}
	tileRef := s.BGTileMap[tileIndex]
	renderedTile := s.WindowTiles[tileRef]
	return valueInPalette(s.BGRDPAL, renderedTile.At(x%8, y%8))
}

func (s *Screen) atSprite(pX, pY int, bg byte) byte {
	if !s.RenderSprites {
		return bg
	}
	for pos := 0; pos < len(s.SpriteData); pos += 4 {
		yPos := s.SpriteData[pos]
		y := int(yPos) - 8
		xPos := s.SpriteData[pos+1]
		x := int(xPos)

		if !(x-pX >= 0 &&
			x-pX < 8 &&
			y-pY >= 0 &&
			y-pY < 8) {
			continue
		}

		tileNumber := s.SpriteData[pos+2]

		flags := byte(0)
		if pos+3 < len(s.SpriteData) {
			flags = s.SpriteData[pos+3]
		}

		priority := bitValue(7, flags)
		// 		Bit7 Priority
		//  If this bit is set to 0, sprite is
		//  displayed on top of background & window.
		//  If this bit is set to 1, then sprite
		//  will be hidden behind colors 1, 2, and 3
		//  of the background & window. (Sprite only
		//  prevails over color 0 of BG & win.)

		yFlip := bitValue(6, flags)
		xFlip := bitValue(5, flags)

		// Bit4 Palette number
		paletteNum := bitValue(4, flags)
		// Sprite colors are taken from OBJ1PAL if
		// this bit is set to 1 and from OBJ0PAL
		// otherwise.
		paletteValue := s.OBJ0PAL
		if paletteNum == 1 {
			paletteValue = s.OBJ1PAL
		}

		spX := (pX - 1) % 8
		spY := (pY - 1) % 8

		if xFlip != 0 {
			spX = 8 - spX
		}
		if yFlip != 0 {
			spY = 8 - spY
		}

		renderedTile := s.SpriteTiles[tileNumber]
		val := valueInPalette(paletteValue, renderedTile.At(spX, spY))

		if priority == 0 || bg == 0 {
			return val
		}
	}
	return bg
}

func (s *Screen) valueAt(x, y int) byte {
	xS := x + s.Position.X
	yS := y + s.Position.Y

	var pixel byte

	if s.RenderWindow &&
		s.WindowPos.X >= 0 && s.WindowPos.X <= 166 &&
		s.WindowPos.Y >= 0 && s.WindowPos.Y <= 143 {
		pixel = s.atWindow(xS, yS)
	} else {
		pixel = s.atBg(xS, yS)
	}
	pixel = s.atSprite(xS+1, yS+1, pixel)
	return pixel
}

func (s *Screen) At(x, y int) color.Color {
	return colorForValue(s.valueAt(x, y))
}
