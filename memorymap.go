package gameboy

type LCDControl byte

// LCDOperation ...
// Bit7  LCD operation                           | ON            | OFF
func (l LCDControl) LCDOperation() bool {
	return bitValue(7, byte(l)) != 0
}

// WindowTileTableAddress ...
// Bit6  Window Tile Table address               | 9C00-9FFF     | 9800-9BFF
func (l LCDControl) WindowTileTableAddress() Range {
	if bitValue(6, byte(l)) == 0 {
		return Range{
			Start: 0x9800, End: 0x9BFF,
		}
	}
	return Range{
		Start: 0x9C00, End: 0x9FFF,
	}
}

// WindowDisplay ...
// Bit5  Window display                          | ON            | OFF
func (l LCDControl) WindowDisplay() bool {
	return bitValue(5, byte(l)) != 0
}

func (l LCDControl) BGWindowTileAddressingSigned() bool {
	return bitValue(4, byte(l)) == 0
}

// TilePatternTableAddress ...
// Bit4  BG & Window Tile Data Select            | 8000-8FFF     | 8800-97FF
func (l LCDControl) TilePatternTableAddress() Range {
	if bitValue(4, byte(l)) == 0 {
		return Range{
			Start: 0x8800, End: 0x97FF,
		}
	}
	return Range{
		Start: 0x8000, End: 0x8FFF,
	}
}

// BackgroundTileTableAddress ...
// Bit3  Background Tile Table address           | 9C00-9FFF     | 9800-9BFF
func (l LCDControl) BackgroundTileTableAddress() Range {
	if bitValue(3, byte(l)) == 0 {
		return Range{
			Start: 0x9800, End: 0x9BFF,
		}
	}
	return Range{
		Start: 0x9C00, End: 0x9FFF,
	}
}

// SpriteSize ...
// Bit2  Sprite size                             | 8x16          | 8x8
func (l LCDControl) SpriteSize() (byte, byte) {
	if bitValue(2, byte(l)) == 0 {
		return 8, 16
	}
	return 8, 8
}

// WindowTransparencyOnColorZero ...
// Bit1  Color #0 transparency in the window     | SOLID         | TRANSPARENT
func (l LCDControl) WindowTransparencyOnColorZero() bool {
	return bitValue(1, byte(l)) != 0
}

// BackgroundDisplay ...
// Bit0  Background display
func (l LCDControl) BackgroundDisplay() bool {
	return bitValue(0, byte(l)) != 0
}
