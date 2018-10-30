package gameboy

import "testing"

func TestMemoryRead8(t *testing.T) {
	cpu := NewCPU()
	cpu.RAM[10] = 100

	mem := Memory{
		cpu:   cpu,
		index: 10,
	}

	got := mem.Read8()
	if got != 100 {
		t.Errorf("expected 100, got %d", got)
	}
}

func TestMemoryWrite8(t *testing.T) {
	cpu := NewCPU()

	mem := Memory{
		cpu:   cpu,
		index: 10,
	}

	mem.Write8(100)
	got := cpu.RAM[10]
	if got != 100 {
		t.Errorf("expected 100, got %d", got)
	}
}

func TestMemoryRead16(t *testing.T) {
	cpu := NewCPU()
	cpu.RAM[10] = 0x55
	cpu.RAM[11] = 0x66

	mem := Memory{
		cpu:   cpu,
		index: 10,
	}

	got := mem.Read16()
	if got != 0x6655 {
		t.Errorf("expected 0x6655, got 0x%X", got)
	}
}

func TestMemoryWrite16(t *testing.T) {
	cpu := NewCPU()

	mem := Memory{
		cpu:   cpu,
		index: 10,
	}

	mem.Write16(0xABCD)
	got := cpu.RAM[10]
	if got != 0xCD {
		t.Errorf("expected 0xCD, got 0x%X", got)
	}
	got = cpu.RAM[11]
	if got != 0xAB {
		t.Errorf("expected 0xAB, got 0x%X", got)
	}
}

func TestMemoryAt(t *testing.T) {
	var tests = []struct {
		name          string
		in            Param
		expectedIndex uint16
	}{
		{
			name:          "8-bit",
			in:            r8(0xF0),
			expectedIndex: 0xF0,
		},
		{
			name:          "16-bit",
			in:            r16(0x0FFF),
			expectedIndex: 0x0FFF,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCPU()
			mem := c.MemoryAt(test.in)
			if mem.cpu != c {
				t.Error("memory CPU was reference was not expected CPU")
			}
			if mem.index != test.expectedIndex {
				t.Errorf("index not as expected, got 0x%X", mem.index)
			}
		})
	}
}

func TestMemoryAtH(t *testing.T) {
	c := NewCPU()
	mem := c.MemoryAtH(r8(0xAE))
	if mem.cpu != c {
		t.Error("memory CPU was reference was not expected CPU")
	}
	if mem.index != 0xFFAE {
		t.Errorf("index not as expected, got 0x%X", mem.index)
	}
}

func TestDirect8(t *testing.T) {
	cpu := NewCPU()
	cpu.RAM[0x1234] = 0xEF
	cpu.PC.Write16(0x1234)
	got := cpu.D8.Read8()
	if got != 0xEF {
		t.Errorf("value not as expected, got 0x%X", got)
	}
	pc := cpu.PC.Read16()
	if pc != 0x1235 {
		t.Errorf("PC not as expected, got 0x%X", pc)
	}
}

func TestDirect16(t *testing.T) {
	cpu := NewCPU()
	cpu.RAM[0x1234] = 0xEF
	cpu.RAM[0x1235] = 0xAC
	cpu.PC.Write16(0x1234)
	got := cpu.D16.Read16()
	if got != 0xACEF {
		t.Errorf("value not as expected, got 0x%X", got)
	}
	pc := cpu.PC.Read16()
	if pc != 0x1236 {
		t.Errorf("PC not as expected, got 0x%X", pc)
	}
}

type r16 uint16

func (r r16) Read16() uint16 {
	return uint16(r)
}

func (r r16) Write16(uint16) {
	panic("unexpected write")
}

type r8 byte

func (r r8) Read8() byte {
	return byte(r)
}

func (r r8) Write8(byte) {
	panic("unexpected write")
}
