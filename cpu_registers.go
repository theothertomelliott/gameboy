package gameboy

type (
	Register struct {
		value byte
	}
	RegisterPair struct {
		Low  Value8
		High Value8
	}
)

func (r *Register) Write8(value byte) {
	r.value = value
}

func (r *Register) Read8() byte {
	if r != nil {
		return r.value
	}
	return 0
}

func (r *RegisterPair) Write16(value uint16) {
	if r.Low == nil {
		r.Low = &Register{}
	}
	if r.High == nil {
		r.High = &Register{}
	}
	r.Low.Write8(byte(value & 0xFF))
	r.High.Write8(byte(value>>8) & 0xFF)
}

func (r *RegisterPair) Read16() uint16 {
	if r != nil {
		return uint16(r.Low.Read8()) | (uint16(r.High.Read8()) << 8)
	}
	return 0
}

func (r *Register) getBit(pos uint) bool {
	return r.value&(1<<pos) > 0
}

func (r *Register) setBit(v bool, pos uint) {
	if v {
		r.value |= (1 << pos)
	} else {
		r.value &^= (1 << pos)
	}
}

func (r *Register) Z() bool {
	return r.getBit(7)
}

func (r *Register) SetZ(v bool) {
	r.setBit(v, 7)
}

func (r *Register) N() bool {
	return r.getBit(6)
}

func (r *Register) SetN(v bool) {
	r.setBit(v, 6)
}

func (r *Register) H() bool {
	return r.getBit(5)
}

func (r *Register) SetH(v bool) {
	r.setBit(v, 5)
}

func (r *Register) C() bool {
	return r.getBit(4)
}

func (r *Register) SetC(v bool) {
	r.setBit(v, 4)
}
