package gameboy

type Value8 interface {
	Write8(byte)
	Read8() byte
}

type Value16 interface {
	Read16() uint16
	Write16(uint16)
}
