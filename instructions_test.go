package gameboy_test

import (
	"testing"

	"github.com/theothertomelliott/gameboy"
)

func TestLD(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())

	src := &gameboy.Register{}
	src.Write8(100)
	dst := &gameboy.Register{}

	cpu.LD(dst, src)

	if dst.Read8() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read8())
	}
}

func TestLD16(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())

	src := &gameboy.RegisterPair{}
	src.Write16(100)
	dst := &gameboy.RegisterPair{}

	cpu.LD(dst, src)

	if dst.Read16() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read16())
	}
}

func TestLDI(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())

	src := &gameboy.Register{}
	src.Write8(10)
	dst := &gameboy.Register{}

	cpu.LDI(dst, src)

	if dst.Read8() != 11 {
		t.Errorf("dst: expected %d, got %d", 11, dst.Read8())
	}
}

func TestLDD(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())

	src := &gameboy.Register{}
	src.Write8(10)
	dst := &gameboy.Register{}

	cpu.LDD(dst, src)

	if dst.Read8() != 9 {
		t.Errorf("dst: expected %d, got %d", 9, dst.Read8())
	}
}

func TestAND(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		a        byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
				h: true,
			},
		},
		{
			name:     "a pair of 1s is ANDed",
			in:       1,
			a:        1,
			expected: 1,
			flags: expectedFlags{
				h: true,
			},
		},
		{
			name:     "not OR",
			in:       1,
			a:        2,
			expected: 0,
			flags: expectedFlags{
				z: true,
				h: true,
			},
		},
		{
			name:     "All bits set",
			in:       0xFF,
			a:        0xFF,
			expected: 0xFF,
			flags: expectedFlags{
				z: false,
				h: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.a)
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.AND(src)

			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestOR(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		a        byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "a pair of 1s is retained",
			in:       1,
			a:        1,
			expected: 1,
		},
		{
			name:     "is OR",
			in:       1,
			a:        2,
			expected: 3,
		},
		{
			name:     "All bits set",
			in:       0xFF,
			a:        0xFF,
			expected: 0xFF,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.a)
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.OR(src)

			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestXOR(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		a        byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "is XOR",
			in:       1,
			a:        2,
			expected: 3,
		},
		{
			name:     "All bits match",
			in:       0xFF,
			a:        0xFF,
			expected: 0,
			flags: expectedFlags{
				z: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.a)
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.XOR(src)

			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestADD(t *testing.T) {
	var tests = []struct {
		name     string
		dst      byte
		src      byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "1 plus 1",
			src:      1,
			dst:      1,
			expected: 2,
		},
		{
			name:     "half carry",
			src:      0xF,
			dst:      1,
			expected: 0x10,
			flags: expectedFlags{
				h: true,
			},
		},
		{
			name:     "carry and half carry",
			dst:      0xFF,
			src:      1,
			expected: 0,
			flags: expectedFlags{
				z: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "carry and half carry (not on boundary",
			dst:      0xFE,
			src:      3,
			expected: 1,
			flags: expectedFlags{
				h: true,
				c: true,
			},
		},
		{
			name:     "carry only",
			dst:      0xF0,
			src:      0x10,
			expected: 0,
			flags: expectedFlags{
				z: true,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			dst := &gameboy.Register{}
			dst.Write8(test.dst)
			src := &gameboy.Register{}
			src.Write8(test.src)

			cpu.ADD(dst, src)

			if got := dst.Read8(); got != test.expected {
				t.Errorf("dst: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestADD16(t *testing.T) {
	var tests = []struct {
		name     string
		dst      uint16
		src      uint16
		expected uint16
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "1 plus 1",
			dst:      1,
			src:      1,
			expected: 2,
		},
		{
			name:     "half carry",
			dst:      0xFFF,
			src:      1,
			expected: 0x1000,
			flags: expectedFlags{
				h: true,
			},
		},
		{
			name:     "carry + half carruy",
			dst:      0xFFFF,
			src:      0x1,
			expected: 0x0000,
			flags: expectedFlags{
				z: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "carry",
			dst:      0xF000,
			src:      0x1000,
			expected: 0x0000,
			flags: expectedFlags{
				z: true,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())

			dst := &gameboy.RegisterPair{}
			dst.Write16(test.dst)
			src := &gameboy.RegisterPair{}
			src.Write16(test.src)

			cpu.ADD(dst, src)

			if got := dst.Read16(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestADDSP(t *testing.T) {
	var tests = []struct {
		name     string
		dst      uint16
		src      byte
		expected uint16
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "1 plus 1",
			dst:      1,
			src:      1,
			expected: 2,
		},
		{
			name:     "half carry",
			dst:      0xFFF,
			src:      1,
			expected: 0x1000,
			flags: expectedFlags{
				h: true,
			},
		},
		{
			name:     "carry + half carruy",
			dst:      0xFFFF,
			src:      0x1,
			expected: 0x0000,
			flags: expectedFlags{
				z: true,
				h: true,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())

			dst := &gameboy.RegisterPair{}
			dst.Write16(test.dst)
			src := &gameboy.Register{}
			src.Write8(test.src)

			cpu.ADD(dst, src)

			if got := dst.Read16(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			test.flags.compare(t, cpu)
		})
	}
}

func TestADC(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		a        byte
		cIn      bool
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "1 plus 1",
			in:       1,
			a:        1,
			expected: 2,
		},
		{
			name:     "half carry",
			in:       0xF,
			a:        1,
			expected: 0x10,
			flags: expectedFlags{
				h: true,
			},
		},
		{
			name:     "carry and half carry",
			in:       0xFF,
			a:        1,
			expected: 0,
			flags: expectedFlags{
				z: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "carry only",
			in:       0xF0,
			a:        0x10,
			expected: 0,
			flags: expectedFlags{
				z: true,
				c: true,
			},
		},
		{
			name:     "adds the carry",
			in:       0xFF,
			a:        1,
			cIn:      true,
			expected: 1,
			flags: expectedFlags{
				h: true,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.a)
			src := &gameboy.Register{}
			src.Write8(test.in)

			if test.cIn {
				cpu.F.SetC(test.cIn)
			}

			cpu.ADC(src)

			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			test.flags.compare(t, cpu)
		})
	}
}

func TestSUB(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		a        byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "2 minus 1",
			in:       1,
			a:        2,
			expected: 1,
			flags: expectedFlags{
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "half carry",
			in:       0x1,
			a:        0xF0,
			expected: 0xEF,
			flags: expectedFlags{
				n: true,
				h: false,
				c: true,
			},
		},
		{
			name:     "borrow and half borrow",
			in:       0x1,
			a:        0x00,
			expected: 0xFF,
			flags: expectedFlags{
				z: false,
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name:     "borrow and half borrow (not on boundary)",
			in:       0x3,
			a:        0x01,
			expected: 0xFE,
			flags: expectedFlags{
				z: false,
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name:     "carry only",
			in:       0x10,
			a:        0x00,
			expected: 0xF0,
			flags: expectedFlags{
				n: true,
				h: true,
				c: false,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.a)
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.SUB(src)

			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestSBC(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		a        byte
		cIn      bool
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "2 minus 1",
			in:       1,
			a:        2,
			expected: 1,
			flags: expectedFlags{
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "half carry",
			in:       0x1,
			a:        0xF0,
			expected: 0xEF,
			flags: expectedFlags{
				n: true,
				h: false,
				c: true,
			},
		},
		{
			name:     "borrow and half borrow",
			in:       0x1,
			a:        0x00,
			expected: 0xFF,
			flags: expectedFlags{
				z: false,
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name:     "borrow and half borrow (not on boundary)",
			in:       0x3,
			a:        0x01,
			expected: 0xFE,
			flags: expectedFlags{
				z: false,
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name:     "carry only",
			in:       0x10,
			a:        0x00,
			expected: 0xF0,
			flags: expectedFlags{
				n: true,
				h: true,
				c: false,
			},
		},
		{
			name:     "includes the carry",
			cIn:      true,
			in:       0x2,
			a:        0x01,
			expected: 0xFE,
			flags: expectedFlags{
				z: false,
				n: true,
				h: false,
				c: false,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.a)
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.F.SetC(test.cIn)
			cpu.SBC(src)

			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestCP(t *testing.T) {
	var tests = []struct {
		name  string
		in    byte
		a     byte
		flags expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name: "2 minus 1",
			in:   1,
			a:    2,
			flags: expectedFlags{
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name: "half carry",
			in:   0x1,
			a:    0xF0,
			flags: expectedFlags{
				n: true,
				h: false,
				c: true,
			},
		},
		{
			name: "borrow and half borrow",
			in:   0x1,
			a:    0x00,
			flags: expectedFlags{
				z: false,
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name: "borrow and half borrow (not on boundary)",
			in:   0x3,
			a:    0x01,
			flags: expectedFlags{
				z: false,
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name: "carry only",
			in:   0x10,
			a:    0x00,
			flags: expectedFlags{
				n: true,
				h: true,
				c: false,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.a)
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.CP(src)

			if got := cpu.A.Read8(); got != test.a {
				t.Errorf("A: expected %d, got %d", test.a, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestINC(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		expected byte
		flags    expectedFlags
	}{
		{
			name:     "zero",
			expected: 1,
		},
		{
			name:     "1 plus 1",
			in:       1,
			expected: 2,
		},
		{
			name:     "half carry",
			in:       0xF,
			expected: 0x10,
			flags: expectedFlags{
				h: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.INC(src)

			if got := src.Read8(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestINC16(t *testing.T) {
	var tests = []struct {
		name     string
		in       uint16
		expected uint16
		flags    expectedFlags
	}{
		{
			name:     "zero",
			expected: 1,
		},
		{
			name:     "1 plus 1",
			in:       1,
			expected: 2,
		},
		{
			name:     "no half carry",
			in:       0x0F,
			expected: 0x0010,
		},
		{
			name:     "no carry",
			in:       0xFF,
			expected: 0x0100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			src := &gameboy.RegisterPair{}
			src.Write16(test.in)

			cpu.INC(src)

			if got := src.Read16(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestDEC(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		expected byte
		flags    expectedFlags
	}{
		{
			name:     "1 - 1",
			in:       1,
			expected: 0,
			flags: expectedFlags{
				z: true,
				n: true,
				h: true,
			},
		},
		{
			name:     "half carry",
			in:       0xF0,
			expected: 0xEF,
			flags: expectedFlags{
				h: false,
				n: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			src := &gameboy.Register{}
			src.Write8(test.in)

			cpu.DEC(src)

			if got := src.Read8(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}

			test.flags.compare(t, cpu)
		})
	}
}

func TestDEC16(t *testing.T) {
	var tests = []struct {
		name     string
		in       uint16
		expected uint16
		flags    expectedFlags
	}{
		{
			name:     "1 - 1",
			in:       1,
			expected: 0,
		},
		{
			name:     "no half borrow",
			in:       0xF0,
			expected: 0xEF,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			src := &gameboy.RegisterPair{}
			src.Write16(test.in)

			cpu.DEC(src)

			if got := src.Read16(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}

			test.flags.compare(t, cpu)
		})
	}
}

func TestSwap(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())
	in := &gameboy.Register{}
	in.Write8(0xAB)

	cpu.SWAP(in)

	if in.Read8() != 0xBA {
		t.Errorf("input was not swapped, got %d", in.Read8())
	}
}

func TestCPL(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())
	cpu.A.Write8(0xF0)
	cpu.CPL()
	if cpu.A.Read8() != 0x0F {
		t.Errorf("input was not flipped, got %d", cpu.A.Read8())
	}
}

func TestBit(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())
	cpu.A.Write8(0x04)
	cpu.BIT(5, cpu.A)
	if cpu.F.Z() {
		t.Errorf("expected zero on bit 5")
	}
	cpu.BIT(2, cpu.A)
	if !cpu.F.Z() {
		t.Errorf("expected one on bit 3")
	}
}

func TestSet(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())
	cpu.A.Write8(0x04)
	cpu.SET(1, cpu.A)
	if got := cpu.A.Read8(); got != 0x6 {
		t.Errorf("value not as expected, got %d", got)
	}
}

func TestRES(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())
	cpu.A.Write8(0x04)
	cpu.RES(2, cpu.A)
	if got := cpu.A.Read8(); got != 0 {
		t.Errorf("value not as expected, got 0x%X", got)
	}
}

func TestSLA(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "no carry",
			in:       0x1,
			expected: 0x2,
		},
		{
			name:     "carry",
			in:       0xFF,
			expected: 0xFE,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.in)
			cpu.SLA(cpu.A)
			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestSRA(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "no carry",
			in:       0x2,
			expected: 0x1,
		},
		{
			name:     "carry",
			in:       0x01,
			expected: 0x00,
			flags: expectedFlags{
				z: true,
				c: true,
			},
		},
		{
			name:     "MSB unchanged",
			in:       0xFF,
			expected: 0xFF,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.in)
			cpu.SRA(cpu.A)
			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestSRL(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "no carry",
			in:       0x2,
			expected: 0x1,
		},
		{
			name:     "carry",
			in:       0x01,
			expected: 0x00,
			flags: expectedFlags{
				z: true,
				c: true,
			},
		},
		{
			name:     "MSB set to zero",
			in:       0xFF,
			expected: 0x7F,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.in)
			cpu.SRL(cpu.A)
			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestRLC(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "rotates left",
			in:       0x1,
			expected: 0x2,
		},
		{
			name:     "full rotation and carry bit",
			in:       0xFF,
			expected: 0xFF,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.in)
			cpu.RLC(cpu.A)
			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestRL(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		carry    bool
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "rotates left",
			in:       0x1,
			expected: 0x2,
		},
		{
			name:     "full rotation through carry bit",
			in:       0xFF,
			expected: 0xFE,
			flags: expectedFlags{
				c: true,
			},
		},
		{
			name:     "full rotation through carry bit (set)",
			in:       0xFF,
			carry:    true,
			expected: 0xFF,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.F.SetC(test.carry)
			cpu.A.Write8(test.in)
			cpu.RL(cpu.A)
			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestRRC(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "rotates right",
			in:       0x2,
			expected: 0x1,
		},
		{
			name:     "full rotation and carry bit",
			in:       0xFF,
			expected: 0x7F,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.A.Write8(test.in)
			cpu.RRC(cpu.A)
			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestRR(t *testing.T) {
	var tests = []struct {
		name     string
		in       byte
		carry    bool
		expected byte
		flags    expectedFlags
	}{
		{
			name: "zero",
			flags: expectedFlags{
				z: true,
			},
		},
		{
			name:     "rotates right",
			in:       0x2,
			expected: 0x1,
		},
		{
			name:     "full rotation through carry bit",
			in:       0xFF,
			expected: 0x7F,
			flags: expectedFlags{
				c: true,
			},
		},
		{
			name:     "full rotation through carry bit (set)",
			in:       0xFF,
			carry:    true,
			expected: 0xFF,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU(gameboy.NewMMU())
			cpu.F.SetC(test.carry)
			cpu.A.Write8(test.in)
			cpu.RR(cpu.A)
			if got := cpu.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, cpu)
		})
	}
}

func TestDAA(t *testing.T) {
	cpu := gameboy.NewCPU(gameboy.NewMMU())
	cpu.A.Write8(55)
	cpu.DAA()
	if got := cpu.A.Read8(); got != 0x55 {
		t.Errorf("value not as expected, got 0x%X", got)
	}
}

type expectedFlags struct {
	z, n, h, c bool
}

func (e expectedFlags) compare(t *testing.T, cpu *gameboy.CPU) {
	if got := cpu.F.Z(); got != e.z {
		t.Errorf("Z Flag: expected %v, got %v", e.z, got)
	}
	if got := cpu.F.N(); got != e.n {
		t.Errorf("N Flag: expected %v, got %v", e.n, got)
	}
	if got := cpu.F.H(); got != e.h {
		t.Errorf("H Flag: expected %v, got %v", e.h, got)
	}
	if got := cpu.F.C(); got != e.c {
		t.Errorf("C Flag: expected %v, got %v", e.c, got)
	}
}
