package cpu_test

import (
	"fmt"
	"testing"

	"github.com/theothertomelliott/gameboy/cpu"
	"github.com/theothertomelliott/gameboy/mmu"
)

func TestLD(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)

	src := &cpu.Register{}
	src.Write8(100)
	dst := &cpu.Register{}

	c.LD(dst, src)

	if dst.Read8() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read8())
	}
}

func TestLD16(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)

	src := &cpu.RegisterPair{}
	src.Write16(100)
	dst := &cpu.RegisterPair{}

	c.LD(dst, src)

	if dst.Read16() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read16())
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.a)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.AND(src)

			if got := c.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.a)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.OR(src)

			if got := c.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.a)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.XOR(src)

			if got := c.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			dst := &cpu.Register{}
			dst.Write8(test.dst)
			src := &cpu.Register{}
			src.Write8(test.src)

			c.ADD(dst, src)

			if got := dst.Read8(); got != test.expected {
				t.Errorf("dst: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
			name:  "zero",
			flags: expectedFlags{},
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
			src:      0x1,
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
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := cpu.New(mmu.New(nil), nil)

			dst := &cpu.RegisterPair{}
			dst.Write16(test.dst)
			src := &cpu.RegisterPair{}
			src.Write16(test.src)

			c.ADD(dst, src)

			if got := dst.Read16(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.a)
			src := &cpu.Register{}
			src.Write8(test.in)

			if test.cIn {
				c.F.SetC(test.cIn)
			}

			c.ADC(c.A, src)

			if got := c.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			test.flags.compare(t, c)
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
				h: false,
				c: false,
			},
		},
		{
			name:     "2 minus 1",
			in:       1,
			a:        2,
			expected: 1,
			flags: expectedFlags{
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name:     "half carry",
			in:       0x1,
			a:        0xF0,
			expected: 0xEF,
			flags: expectedFlags{
				n: true,
				h: true,
				c: false,
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
				h: true,
				c: true,
			},
		},
		{
			name:     "borrow and half borrow (not on boundary)",
			a:        0x01,
			in:       0x3,
			expected: 0xFE,
			flags: expectedFlags{
				z: false,
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name:     "carry only",
			in:       0x10,
			a:        0x00,
			expected: 0xF0,
			flags: expectedFlags{
				n: true,
				h: false,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.a)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.SUB(src)

			if got := c.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
				h: false,
				c: false,
			},
		},
		{
			name:     "2 minus 1",
			in:       1,
			a:        2,
			expected: 1,
			flags: expectedFlags{
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name:     "half carry",
			in:       0x1,
			a:        0xF0,
			expected: 0xEF,
			flags: expectedFlags{
				n: true,
				h: true,
				c: false,
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
				h: true,
				c: true,
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
				h: true,
				c: true,
			},
		},
		{
			name:     "carry only",
			in:       0x10,
			a:        0x00,
			expected: 0xF0,
			flags: expectedFlags{
				n: true,
				h: false,
				c: true,
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
				h: true,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.a)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.F.SetC(test.cIn)
			c.SBC(c.A, src)

			if got := c.A.Read8(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
				h: false,
				c: false,
			},
		},
		{
			name: "2 minus 1",
			in:   1,
			a:    2,
			flags: expectedFlags{
				n: true,
				h: false,
				c: false,
			},
		},
		{
			name: "half carry",
			in:   0x1,
			a:    0xF0,
			flags: expectedFlags{
				n: true,
				h: true,
				c: false,
			},
		},
		{
			name: "borrow and half borrow",
			in:   0x1,
			a:    0x00,
			flags: expectedFlags{
				z: false,
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name: "borrow and half borrow (not on boundary)",
			in:   0x3,
			a:    0x01,
			flags: expectedFlags{
				z: false,
				n: true,
				h: true,
				c: true,
			},
		},
		{
			name: "carry only",
			in:   0x10,
			a:    0x00,
			flags: expectedFlags{
				n: true,
				h: false,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.a)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.CP(src)

			if got := c.A.Read8(); got != test.a {
				t.Errorf("A: expected %d, got %d", test.a, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.INC(src)

			if got := src.Read8(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			src := &cpu.RegisterPair{}
			src.Write16(test.in)

			c.INC(src)

			if got := src.Read16(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}
			test.flags.compare(t, c)
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
				h: false,
			},
		},
		{
			name:     "half carry",
			in:       0xF0,
			expected: 0xEF,
			flags: expectedFlags{
				n: true,
				h: true,
			},
		},
		{
			name:     "underflow",
			in:       0x0,
			expected: 0xFF,
			flags: expectedFlags{
				n: true,
				h: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := cpu.New(mmu.New(nil), nil)
			src := &cpu.Register{}
			src.Write8(test.in)

			c.DEC(src)

			if got := src.Read8(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}

			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			src := &cpu.RegisterPair{}
			src.Write16(test.in)

			c.DEC(src)

			if got := src.Read16(); got != test.expected {
				t.Errorf("n: expected %d, got %d", test.expected, got)
			}

			test.flags.compare(t, c)
		})
	}
}

func TestSwap(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)
	in := &cpu.Register{}
	in.Write8(0xAB)

	c.SWAP(in)

	if in.Read8() != 0xBA {
		t.Errorf("input was not swapped, got %d", in.Read8())
	}
}

func TestCPL(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)
	c.A.Write8(0xF0)
	c.CPL()
	if c.A.Read8() != 0x0F {
		t.Errorf("input was not flipped, got %d", c.A.Read8())
	}
}

func TestBit(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)
	c.A.Write8(0x04)
	c.BIT(5, c.A)
	if !c.F.Z() {
		t.Errorf("expected zero on bit 5")
	}
	c.BIT(2, c.A)
	if c.F.Z() {
		t.Errorf("expected one on bit 3")
	}
}

func TestSet(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)
	c.A.Write8(0x04)
	c.SET(1, c.A)
	if got := c.A.Read8(); got != 0x6 {
		t.Errorf("value not as expected, got %d", got)
	}
}

func TestRES(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)
	c.A.Write8(0x04)
	c.RES(2, c.A)
	if got := c.A.Read8(); got != 0 {
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.in)
			c.SLA(c.A)
			if got := c.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.in)
			c.SRA(c.A)
			if got := c.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.in)
			c.SRL(c.A)
			if got := c.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.in)
			c.RLC(c.A)
			if got := c.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.F.SetC(test.carry)
			c.A.Write8(test.in)
			c.RL(c.A)
			if got := c.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, c)
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
			expected: 0xFF,
			flags: expectedFlags{
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := cpu.New(mmu.New(nil), nil)
			c.A.Write8(test.in)
			c.RRC(c.A)
			if got := c.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, c)
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
			c := cpu.New(mmu.New(nil), nil)
			c.F.SetC(test.carry)
			c.A.Write8(test.in)
			c.RR(c.A)
			if got := c.A.Read8(); got != test.expected {
				t.Errorf("expected 0x%X, got 0x%X", test.expected, got)
			}
			test.flags.compare(t, c)
		})
	}
}

func TestDAA(t *testing.T) {
	var tests = []struct {
		name    string
		nBefore bool
		cBefore bool
		hBefore bool

		upperMin byte
		upperMax byte
		lowerMin byte
		lowerMax byte

		correctionExpected byte
		cExpected          bool
	}{
		// Based on table at:
		// http://www.z80.info/z80syntx.htm#DAA
		{
			name:     "no change",
			upperMin: 0x0, upperMax: 0x9,
			lowerMin: 0x0, lowerMax: 0x9,
			correctionExpected: 0x0,
		},
		{
			name:     "add 06",
			upperMin: 0x0, upperMax: 0x8,
			lowerMin: 0xA, lowerMax: 0xF,
			correctionExpected: 0x6,
		},
		{
			name:     "add 06 (h = true)",
			hBefore:  true,
			upperMin: 0x0, upperMax: 0x9,
			lowerMin: 0x0, lowerMax: 0x3,
			correctionExpected: 0x6,
		},
		{
			name:     "add 60",
			upperMin: 0xA, upperMax: 0xF,
			lowerMin: 0x0, lowerMax: 0x9,
			correctionExpected: 0x60,
			cExpected:          true,
		},
		{
			name:     "add 66",
			upperMin: 0x9, upperMax: 0xF,
			lowerMin: 0xA, lowerMax: 0xF,
			correctionExpected: 0x66,
			cExpected:          true,
		},
		{
			name:     "add 66 (h)",
			hBefore:  true,
			upperMin: 0xA, upperMax: 0xF,
			lowerMin: 0x0, lowerMax: 0x3,
			correctionExpected: 0x66,
			cExpected:          true,
		},
		{
			name:     "add 60 (c|h)",
			cBefore:  true,
			upperMin: 0x0, upperMax: 0x2,
			lowerMin: 0x0, lowerMax: 0x9,
			correctionExpected: 0x60,
			cExpected:          true,
		},
		{
			name:     "add 66 (c)",
			cBefore:  true,
			upperMin: 0x0, upperMax: 0x2,
			lowerMin: 0xA, lowerMax: 0xF,
			correctionExpected: 0x66,
			cExpected:          true,
		},
		{
			name:     "add 66 (c|h)",
			cBefore:  true,
			hBefore:  true,
			upperMin: 0x0, upperMax: 0x3,
			lowerMin: 0x0, lowerMax: 0x3,
			correctionExpected: 0x66,
			cExpected:          true,
		},
		{
			name:     "sub 0",
			nBefore:  true,
			upperMin: 0x0, upperMax: 0x9,
			lowerMin: 0x0, lowerMax: 0x9,
		},
		{
			name:     "sub FA",
			nBefore:  true,
			hBefore:  true,
			upperMin: 0x0, upperMax: 0x8,
			lowerMin: 0x6, lowerMax: 0xF,
			correctionExpected: 0xFA,
		},
		{
			name:     "sub A0",
			nBefore:  true,
			cBefore:  true,
			upperMin: 0x7, upperMax: 0xF,
			lowerMin: 0x0, lowerMax: 0x9,
			correctionExpected: 0xA0,
			cExpected:          true,
		},
		{
			name:     "sub A0",
			nBefore:  true,
			cBefore:  true,
			hBefore:  true,
			upperMin: 0x6, upperMax: 0xF,
			lowerMin: 0x6, lowerMax: 0xF,
			correctionExpected: 0x9A,
			cExpected:          true,
		},
	}
	for _, test := range tests {
		for u := test.upperMin; u <= test.upperMax; u++ {
			for l := test.lowerMin; l <= test.lowerMax; l++ {
				valueBefore := u<<4 + l
				valueExpected := valueBefore + test.correctionExpected
				t.Run(fmt.Sprintf("%s: 0x%02X", test.name, valueBefore), func(t *testing.T) {
					c := cpu.New(mmu.New(nil), nil)

					c.F.SetN(test.nBefore)
					c.F.SetC(test.cBefore)
					c.F.SetH(test.hBefore)

					c.A.Write8(valueBefore)

					c.DAA(c.A)

					if got := c.A.Read8(); got != valueExpected {
						t.Errorf("expected 0x%X, got 0x%X", valueExpected, got)
					}
					if c.F.C() != test.cExpected {
						t.Errorf("C, expected %v, got %v", test.cExpected, c.F.C())
					}
				})
			}
		}
	}
}

func TestLDHL(t *testing.T) {
	var tests = []struct {
		name     string
		a        uint16
		b        int8
		expected uint16
		flags    expectedFlags
	}{
		{
			name: "zero",
		},
		{
			name:     "1 plus 1",
			a:        1,
			b:        1,
			expected: 2,
		},
		{
			name:     "carry and half carry",
			a:        0xFFFF,
			b:        1,
			expected: 0,
			flags: expectedFlags{
				h: true,
				c: true,
			},
		},
		{
			name:     "carry and half carry (not on boundary",
			a:        0xFFFE,
			b:        3,
			expected: 1,
			flags: expectedFlags{
				h: true,
				c: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := mmu.New(nil)
			m.Write16(test.expected, 0xAEBF)

			c := cpu.New(m, nil)
			a := &cpu.RegisterPair{}
			a.Write16(test.a)
			b := cpu.DirectSigned8(test.b)

			c.LDHL(a, b)

			if got := c.HL.Read16(); got != test.expected {
				t.Errorf("HL: expected 0x%04X, got 0x%04X", test.expected, got)
			}
			test.flags.compare(t, c)
		})
	}
}

func TestCCF(t *testing.T) {
	cpu := cpu.New(mmu.New(nil), nil)

	cpu.F.SetZ(true)
	cpu.F.SetN(true)
	cpu.F.SetH(true)
	cpu.F.SetC(true)

	cpu.CCF()

	expectedFlags{
		z: true,
	}.compare(t, cpu)
}

func TestSCF(t *testing.T) {
	c := cpu.New(mmu.New(nil), nil)

	c.F.SetZ(true)
	c.F.SetN(true)
	c.F.SetH(true)
	c.F.SetC(false)

	c.SCF()

	expectedFlags{
		z: true,
		c: true,
	}.compare(t, c)
}

type expectedFlags struct {
	z, n, h, c bool
}

func (e expectedFlags) compare(t *testing.T, cpu *cpu.CPU) {
	t.Helper()
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
