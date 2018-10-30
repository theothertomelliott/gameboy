package gameboy

type Memory struct {
	cpu   *CPU
	index uint16
}

func (m *Memory) Write8(value byte) {
	m.cpu.RAM[m.index] = value
}

func (m *Memory) Read8() byte {
	return m.cpu.RAM[m.index]
}

func (m *Memory) Write16(value uint16) {
	m.cpu.RAM[m.index] = byte(value & 0xFF)
	m.cpu.RAM[m.index+1] = byte((value & 0xFF00) >> 8)
}

func (m *Memory) Read16() uint16 {
	return uint16(m.cpu.RAM[m.index]) + uint16(m.cpu.RAM[m.index+1])<<8
}

func (c *CPU) MemoryAt(pos Param) *Memory {
	if pos16, ok := pos.(Value16); ok {
		return &Memory{
			index: pos16.Read16(),
			cpu:   c,
		}
	}
	if pos8, ok := pos.(Value8); ok {
		return &Memory{
			index: uint16(pos8.Read8()),
			cpu:   c,
		}
	}
	return nil
}

func (c *CPU) MemoryAtH(pos Param) *Memory {
	pos8 := pos.(Value8)
	return &Memory{
		index: 0xFF00 | uint16(pos8.Read8()),
		cpu:   c,
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
