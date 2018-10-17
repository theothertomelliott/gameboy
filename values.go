package gameboy

type Value interface {
	Is16Bit() bool
	As16() uint16
	As8() byte
}

type Value8 interface {
	Write(byte)
	Read() byte
}

type In16 interface {
	Read() uint16
}

type Value16 interface {
	In16
	Write(uint16)
}
