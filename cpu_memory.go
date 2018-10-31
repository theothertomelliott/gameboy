package gameboy

// Memory references a position on memory.
// The index is fixed on the first read, so MemoryAt may be called to provide a reference
// to the memory addressed by a register, for example, so it may be reused later.
type Memory struct {
	cpu    *CPU
	index  Param
	offset uint16

	cachedPos *uint16 // Cache of position
}

func (m *Memory) pos() uint16 {
	if m.cachedPos != nil {
		return *m.cachedPos
	}

	var pos uint16
	if index8, is8Bit := m.index.(Value8); is8Bit {
		pos = m.offset + uint16(index8.Read8())
	}
	if index16, is16Bit := m.index.(Value16); is16Bit {
		pos = m.offset + uint16(index16.Read16())
	}
	m.cachedPos = &pos
	return pos
}

func (m *Memory) Write8(value byte) {
	m.cpu.RAM[m.pos()] = value
}

func (m *Memory) Read8() byte {
	return m.cpu.RAM[m.pos()]
}

func (m *Memory) Write16(value uint16) {
	pos := m.pos()
	m.cpu.RAM[pos] = byte(value & 0xFF)
	m.cpu.RAM[pos+1] = byte((value & 0xFF00) >> 8)
}

func (m *Memory) Read16() uint16 {
	pos := m.pos()
	return uint16(m.cpu.RAM[pos]) + uint16(m.cpu.RAM[pos+1])<<8
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
		offset: 0xFF00,
		cpu:    c,
	}
}

var _ Value8 = &Direct8{}

type Direct8 struct {
	CPU *CPU
}

func (d Direct8) Read8() byte {
	c := d.CPU
	defer c.PC.Inc(1)
	return c.RAM[c.PC.Read16()]
}

func (d Direct8) Write8(byte) {
	panic("write to direct memory")
}

var _ Value16 = &Direct16{}

type Direct16 struct {
	CPU *CPU
}

func (d Direct16) Read16() uint16 {
	c := d.CPU
	defer c.PC.Inc(2)
	low := c.RAM[c.PC.Read16()]
	high := c.RAM[c.PC.Read16()+1]
	return uint16(high)<<8 | uint16(low)
}

func (d Direct16) Write16(uint16) {
	panic("write to direct memory")
}
