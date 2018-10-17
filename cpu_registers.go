package gameboy

type (
	Register struct {
		value byte
	}
	RegisterPair struct {
		Low  *Register
		High *Register
	}
)

func (r *Register) Write(value byte) {
	r.value = value
}

func (r *Register) Read() byte {
	if r != nil {
		return r.value
	}
	return 0
}

func (r *RegisterPair) Write(value uint16) {
	r.Low.Write(byte(value & 0xFF))
	r.High.Write(byte(value>>8) & 0xFF)
}

func (r *RegisterPair) Read() uint16 {
	if r != nil {
		return uint16(r.Low.Read()) & (uint16(r.High.Read()) << 8)
	}
	return 0
}
