package gameboy

type Value8 interface {
	Write8(byte)
	Read8() byte
}

type ValueSigned8 interface {
	ReadSigned8() int8
}

type Value16 interface {
	Read16() uint16
	Write16(uint16)
}
