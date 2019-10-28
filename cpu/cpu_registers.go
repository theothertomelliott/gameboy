package cpu

import (
	"github.com/theothertomelliott/gameboy/binary"
	"github.com/theothertomelliott/gameboy/values"
)

type (
	Register struct {
		value       byte
		description string
		tracer      RegisterTracer
	}
	FRegister struct {
		Register
	}
	RegisterPair struct {
		Low         values.Value8
		High        values.Value8
		description string
	}

	RegisterTracer interface {
		AddRegister(name string, valueBefore byte, valueAfter byte)
	}
)

func NewRegister(description string, tracer RegisterTracer) *Register {
	return &Register{
		description: description,
		tracer:      tracer,
	}
}

func (r *Register) String() string {
	return r.description
}

func (r *Register) Write8(value byte) {
	if r.tracer != nil {
		r.tracer.AddRegister(r.description, r.value, value)
	}
	r.value = value
}

func (r *Register) Read8() byte {
	if r != nil {
		return r.value
	}
	return 0
}

func NewRegisterPair(description string, high, low values.Value8) *RegisterPair {
	return &RegisterPair{
		High:        high,
		Low:         low,
		description: description,
	}
}

func (r *RegisterPair) String() string {
	return r.description
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

func (r *Register) getBit(pos byte) bool {
	return binary.Bit(pos, r.value) == 1
}

func (r *Register) setBit(v bool, pos byte) {
	r.value = binary.SetBit(pos, r.value, v)
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

func (r *FRegister) Write8(value byte) {
	r.Register.Write8(value & 0xF0)
}
