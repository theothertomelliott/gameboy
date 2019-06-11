package gameboy

/*
Bit7  LCD operation                           | ON            | OFF
Bit6  Window Tile Table address               | 9C00-9FFF     | 9800-9BFF
Bit5  Window display                          | ON            | OFF
Bit4  Tile Pattern Table address              | 8000-8FFF     | 8800-97FF
Bit3  Background Tile Table address           | 9C00-9FFF     | 9800-9BFF
Bit2  Sprite size                             | 8x16          | 8x8
Bit1  Color #0 transparency in the window     | SOLID         | TRANSPARENT
Bit0  Background display
*/
type LCDControl byte

func (l LCDControl) LCDOperation() bool {
	return bitValue(7, byte(l)) != 0
}

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

func (l LCDControl) WindowDisplay() bool {
	return bitValue(5, byte(l)) != 0
}

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

func (l LCDControl) SpriteSize() (byte, byte) {
	if bitValue(2, byte(l)) == 0 {
		return 8, 16
	}
	return 8, 8
}

func (l LCDControl) WindowTransparencyOnColorZero() bool {
	return bitValue(1, byte(l)) != 0
}

func (l LCDControl) BackgroundDisplay() bool {
	return bitValue(0, byte(l)) != 0
}
