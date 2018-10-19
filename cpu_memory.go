package gameboy

type Memory8 struct {
	value *byte
}

func (m Memory8) Write(value byte) {
	*m.value = value
}

func (m Memory8) Read() byte {
	if m.value != nil {
		return *m.value
	}
	return 0
}

type Memory16 struct {
	value *uint16
}

func (m Memory16) Write(value uint16) {
	*m.value = value
}

func (m Memory16) Read() uint16 {
	if m.value != nil {
		return *m.value
	}
	return 0
}

func (c *CPU) MemoryAt(pos Param) Value8 {
	if pos16, ok := pos.(In16); ok {
		return Memory8{
			value: &c.RAM[pos16.Read()],
		}
	}
	if pos8, ok := pos.(Value8); ok {
		return Memory8{
			value: &c.RAM[pos8.Read()],
		}
	}
	return nil
}

func (c *CPU) MemoryAtH(pos Param) Value8 {
	pos8 := pos.(Value8)
	return Memory8{
		value: &c.RAM[0xFF00|uint16(pos8.Read())],
	}
}

func (c *CPU) MemoryAt16(pos In16) Value16 {
	low := c.RAM[pos.Read()]
	high := c.RAM[pos.Read()+1]
	value := uint16(low) | (uint16(high) << 8)
	return Memory16{
		value: &value,
	}
}

type Direct8 struct {
	CPU *CPU
}

func (d Direct8) Read() byte {
	c := d.CPU
	defer func() {
		c.PC.Inc(1)
	}()
	return c.RAM[c.PC.Read()]
}

type Direct16 struct {
	CPU *CPU
}

func (d Direct16) Read() uint16 {
	c := d.CPU
	low := c.RAM[c.PC.Read()]
	c.PC.Inc(1)
	high := c.RAM[c.PC.Read()]
	c.PC.Inc(1)
	return uint16(high) << 8 & uint16(low)
}
