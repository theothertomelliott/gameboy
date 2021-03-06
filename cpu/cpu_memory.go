package cpu

import (
	"fmt"

	"github.com/theothertomelliott/gameboy/ioports"
	"github.com/theothertomelliott/gameboy/values"
)

// Memory references a position on memory.
// The index is fixed on the first read, so MemoryAt may be called to provide a reference
// to the memory addressed by a register, for example, so it may be reused later.
type Memory struct {
	cpu    *CPU
	index  Param
	offset uint16

	cachedPos *uint16 // Cache of position
}

func (m *Memory) GetIndex() Param {
	return m.index
}

func (m *Memory) String() string {
	if m.offset != 0 {
		return fmt.Sprintf("(0x%04X+%v)", m.offset, m.index)
	}
	return fmt.Sprintf("(%v)", m.index)
}

func (m *Memory) pos() uint16 {
	if m.cachedPos != nil {
		return *m.cachedPos
	}

	var pos uint16
	if index8, is8Bit := m.index.(values.Value8); is8Bit {
		pos = m.offset + uint16(index8.Read8())
	}
	if index16, is16Bit := m.index.(values.Value16); is16Bit {
		pos = m.offset + uint16(index16.Read16())
	}
	m.cachedPos = &pos
	return pos
}

func (m *Memory) Write8(value byte) {
	pos := m.pos()
	switch pos {
	case ioports.DIVIDER:
		value = 0
	}
	m.cpu.MMU.Write8(pos, value)
}

func (m *Memory) Read8() byte {
	return m.cpu.MMU.Read8(m.pos())
}

func (m *Memory) Write16(value uint16) {
	m.cpu.MMU.Write16(m.pos(), value)
}

func (m *Memory) Read16() uint16 {
	return m.cpu.MMU.Read16(m.pos())
}

func (c *CPU) MemoryAt(pos Param) *Memory {
	return &Memory{
		index: pos,
		cpu:   c,
	}
}

func (c *CPU) MemoryAtH(pos Param) *Memory {
	return &Memory{
		index:  pos,
		cpu:    c,
		offset: 0xFF00,
	}
}

var _ values.Value8 = Direct8(0)
var _ values.Value16 = Direct16(0)
var _ values.ValueSigned8 = DirectSigned8(0)

type Direct8 byte

func (b Direct8) String() string {
	return fmt.Sprintf("0x%X", byte(b))
}

func (b Direct8) Read8() byte {
	return byte(b)
}

func (b Direct8) Write8(byte) {
	panic("write to direct memory")
}

type DirectSigned8 int8

func (b DirectSigned8) String() string {
	return fmt.Sprintf("%d", int8(b))
}

func (b DirectSigned8) ReadSigned8() int8 {
	return int8(b)
}

type Direct16 uint16

func (b Direct16) String() string {
	return fmt.Sprintf("0x%X", uint16(b))
}

func (b Direct16) Read16() uint16 {
	return uint16(b)
}

func (b Direct16) Write16(uint16) {
	panic("write to direct memory")
}

func (c *CPU) D8() values.Value8 {
	v := c.MMU.Read8(c.PC.Read16())
	c.PC.Inc(1)
	return Direct8(v)
}

func (c *CPU) R8() values.ValueSigned8 {
	v := c.MMU.Read8(c.PC.Read16())
	c.PC.Inc(1)
	return DirectSigned8(v)
}

func (c *CPU) D16() values.Value16 {
	v := c.MMU.Read16(c.PC.Read16())
	c.PC.Inc(2)
	return Direct16(v)
}

func (c *CPU) A8() values.Value16 {
	v := c.MMU.Read16(c.PC.Read16())
	c.PC.Inc(1)
	return Direct16(0xFF00 | v)
}

func (c *CPU) A16() values.Value16 {
	v := c.MMU.Read16(c.PC.Read16())
	c.PC.Inc(2)
	return Direct16(v)
}
