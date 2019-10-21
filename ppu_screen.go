package gameboy

import (
	"image"
	"image/color"
)

var _ image.Image = &Screen{}

// https://gbdev.gg8.se/wiki/articles/Video_Display#VRAM_Sprite_Attribute_Table_.28OAM.29

func NewBackground(mmu *MMU) *Screen {
	return &Screen{
		BGTileMap:                    mmu.ReadRange(GetLCDControl(mmu).BackgroundTileTableAddress()),
		BGTiles:                      GetBackgroundTiles(mmu),
		SpriteTiles:                  GetTilesForRange(mmu, Range{Start: 0x8000, End: 0x8FFF}),
		SpriteData:                   mmu.ReadRange(Range{Start: 0xFE00, End: 0xFE9F}),
		RenderBackground:             GetLCDControl(mmu).BackgroundDisplay(),
		BGWindowTileAddressingSigned: GetLCDControl(mmu).BGWindowTileAddressingSigned(),
		RenderWindow:                 false,

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

func NewWindow(mmu *MMU) *Screen {
	return &Screen{
		BGTileMap:                    mmu.ReadRange(GetLCDControl(mmu).BackgroundTileTableAddress()),
		BGTiles:                      GetBackgroundTiles(mmu),
		WindowTileMap:                mmu.ReadRange(GetLCDControl(mmu).WindowTileTableAddress()),
		SpriteTiles:                  GetTilesForRange(mmu, Range{Start: 0x8000, End: 0x8FFF}),
		SpriteData:                   mmu.ReadRange(Range{Start: 0xFE00, End: 0xFE9F}),
		RenderBackground:             GetLCDControl(mmu).BackgroundDisplay(),
		BGWindowTileAddressingSigned: GetLCDControl(mmu).BGWindowTileAddressingSigned(),
		RenderWindow:                 true,

		BGRDPAL: mmu.Read8(BGRDPAL),
		OBJ0PAL: mmu.Read8(OBJ0PAL),
		OBJ1PAL: mmu.Read8(OBJ1PAL),

		// Put the window on screen
		WindowPos: image.Point{
			X: 7,
			Y: 0,
		},
		Position: image.Point{
			X: 0,
			Y: 0,
		},

		bounds: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{160, 144},
		},
	}
}

func NewScreen(mmu *MMU) *Screen {
	lcdControl := GetLCDControl(mmu)
	return &Screen{
		BGTileMap:                    mmu.ReadRange(lcdControl.BackgroundTileTableAddress()),
		BGTiles:                      GetBackgroundTiles(mmu),
		WindowTileMap:                mmu.ReadRange(lcdControl.WindowTileTableAddress()),
		SpriteTiles:                  GetTilesForRange(mmu, Range{Start: 0x8000, End: 0x8FFF}),
		SpriteData:                   mmu.ReadRange(Range{Start: 0xFE00, End: 0xFE9F}),
		RenderBackground:             lcdControl.BackgroundDisplay(),
		RenderWindow:                 lcdControl.WindowDisplay(),
		BGWindowTileAddressingSigned: lcdControl.BGWindowTileAddressingSigned(),

		BGRDPAL: mmu.Read8(BGRDPAL),
		OBJ0PAL: mmu.Read8(OBJ0PAL),
		OBJ1PAL: mmu.Read8(OBJ1PAL),

		WindowPos:  GetWindowPos(mmu),
		Position:   GetScroll(mmu),
		SpriteSize: lcdControl.SpriteSize(),

		bounds: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{160, 144},
		},

		RenderSprites: true,
	}
}

type Screen struct {
	BGTileMap                    []byte
	WindowTileMap                []byte
	BGTiles                      []Tile
	SpriteTiles                  []Tile
	SpriteData                   []byte
	RenderBackground             bool
	RenderSprites                bool
	RenderWindow                 bool
	BGWindowTileAddressingSigned bool

	BGRDPAL byte
	OBJ0PAL byte
	OBJ1PAL byte

	SpriteSize image.Point
	WindowPos  image.Point
	Position   image.Point
	bounds     image.Rectangle
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
	if s.BGWindowTileAddressingSigned {
		tileRef = (tileRef + 128) % 255
	}
	renderedTile := s.BGTiles[tileRef]
	return valueInPalette(s.BGRDPAL, renderedTile.At(x%8, y%8))
}

func (s *Screen) atWindow(x, y int) byte {
	tileIndex := y/8*32 + x/8
	if tileIndex < 0 || tileIndex >= len(s.WindowTileMap) {
		return 0
	}
	tileRef := s.WindowTileMap[tileIndex]
	if s.BGWindowTileAddressingSigned {
		tileRef = (tileRef + 128) % 255
	}
	renderedTile := s.BGTiles[tileRef]
	return valueInPalette(s.BGRDPAL, renderedTile.At(x%8, y%8))
}

func (s *Screen) atSprite(pX, pY int, bg byte) byte {
	if !s.RenderSprites {
		return bg
	}
	var value = bg
	for pos := 0; pos < len(s.SpriteData); pos += 4 {
		y := int(s.SpriteData[pos]) - 16
		x := int(s.SpriteData[pos+1]) - 8

		if !(x <= pX &&
			pX-x < s.SpriteSize.X &&
			y <= pY &&
			pY-y < s.SpriteSize.Y) {
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
		if priority == 1 && bg != 0 {
			continue
		}

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

		spX := (pX - x)
		tiledY := (pY - y)
		spY := tiledY % 8

		if xFlip != 0 {
			spX = 8 - spX
		}
		if yFlip != 0 {
			spY = 8 - spY
		}
		var renderedTile Tile
		if tiledY >= 8 {
			renderedTile = s.SpriteTiles[tileNumber+1]
		} else {
			renderedTile = s.SpriteTiles[tileNumber]
		}
		val := renderedTile.At(spX, spY)
		// Transparent
		if val == 0 {
			continue
		}
		value = valueInPalette(paletteValue, val)
	}
	return value
}

func (s *Screen) valueAt(x, y int) byte {
	var pixel byte

	pixel = s.atBg(x+s.Position.X, y+s.Position.Y)
	if s.RenderWindow {
		wX := x - (s.WindowPos.X - 7)
		wY := y - s.WindowPos.Y
		if wX >= 0 && wX < 166 &&
			wY >= 0 && wY < 144 {
			pixel = s.atWindow(wX, wY)
		}
	}
	pixel = s.atSprite(x, y, pixel)
	return pixel
}

func (s *Screen) At(x, y int) color.Color {
	return colorForValue(s.valueAt(x, y))
}
