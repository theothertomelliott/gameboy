package gameboy

func bitValue(pos byte, value byte) byte {
	mask := (byte(1) << pos)
	return (value & mask) >> pos
}

func setBitValue(pos byte, value byte, bitValue bool) byte {
	if bitValue {
		value |= (1 << pos)
	} else {
		value &^= (1 << pos)
	}
	return value
}
