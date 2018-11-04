package gameboy

func bitValue(pos byte, value byte) byte {
	mask := (byte(1) << pos)
	return (value & mask) >> pos
}
