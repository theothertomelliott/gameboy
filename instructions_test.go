package gameboy_test

import (
	"testing"

	"github.com/theothertomelliott/gameboy"
)

func TestLD(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.Register{}
	src.Write(100)
	dst := &gameboy.Register{}

	cpu.LD(dst, src)

	if dst.Read() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read())
	}
}

func TestLD16(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.RegisterPair{}
	src.Write(100)
	dst := &gameboy.RegisterPair{}

	cpu.LD(dst, src)

	if dst.Read() != 100 {
		t.Errorf("dst: expected %d, got %d", 100, dst.Read())
	}
}

func TestLDI(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.Register{}
	src.Write(10)
	dst := &gameboy.Register{}

	cpu.LDI(dst, src)

	if dst.Read() != 11 {
		t.Errorf("dst: expected %d, got %d", 11, dst.Read())
	}
}

func TestLDD(t *testing.T) {
	cpu := gameboy.NewCPU()

	src := &gameboy.Register{}
	src.Write(10)
	dst := &gameboy.Register{}

	cpu.LDD(dst, src)

	if dst.Read() != 9 {
		t.Errorf("dst: expected %d, got %d", 9, dst.Read())
	}
}

func TestAND(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		expected   byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
			h:    true,
		},
		{
			name:     "a pair of 1s is ANDed",
			in:       1,
			a:        1,
			expected: 1,
			h:        true,
		},
		{
			name:     "not OR",
			in:       1,
			a:        2,
			expected: 0,
			z:        true,
			h:        true,
		},
		{
			name:     "All bits set",
			in:       0xFF,
			a:        0xFF,
			expected: 0xFF,
			z:        false,
			h:        true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			cpu.AND(src)

			if got := cpu.A.Read(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}

func TestOR(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		expected   byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
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
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			cpu.OR(src)

			if got := cpu.A.Read(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}

func TestXOR(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		expected   byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
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
			z:        true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			cpu.XOR(src)

			if got := cpu.A.Read(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}

func TestADD(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		expected   byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
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
			h:        true,
		},
		{
			name:     "carry and half carry",
			in:       0xFF,
			a:        1,
			expected: 0,
			z:        true,
			h:        true,
			c:        true,
		},
		{
			name:     "carry and half carry (not on boundary",
			in:       0xFE,
			a:        3,
			expected: 1,
			h:        true,
			c:        true,
		},
		{
			name:     "carry only",
			in:       0xF0,
			a:        0x10,
			expected: 0,
			z:        true,
			c:        true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			cpu.ADD(src)

			if got := cpu.A.Read(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}

func TestADC(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		cIn        bool
		expected   byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
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
			h:        true,
		},
		{
			name:     "carry and half carry",
			in:       0xFF,
			a:        1,
			expected: 0,
			z:        true,
			h:        true,
			c:        true,
		},
		{
			name:     "carry only",
			in:       0xF0,
			a:        0x10,
			expected: 0,
			z:        true,
			c:        true,
		},
		{
			name:     "adds the carry",
			in:       0xFF,
			a:        1,
			cIn:      true,
			expected: 1,
			h:        true,
			c:        true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			if test.cIn {
				cpu.F.SetC(test.cIn)
			}

			cpu.ADC(src)

			if got := cpu.A.Read(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}

func TestSUB(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		expected   byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
			n:    true,
			h:    true,
			c:    true,
		},
		{
			name:     "2 minus 1",
			in:       1,
			a:        2,
			expected: 1,
			n:        true,
			h:        true,
			c:        true,
		},
		{
			name:     "half carry",
			in:       0x1,
			a:        0xF0,
			expected: 0xEF,
			n:        true,
			h:        false,
			c:        true,
		},
		{
			name:     "borrow and half borrow",
			in:       0x1,
			a:        0x00,
			expected: 0xFF,
			z:        false,
			n:        true,
			h:        false,
			c:        false,
		},
		{
			name:     "borrow and half borrow (not on boundary)",
			in:       0x3,
			a:        0x01,
			expected: 0xFE,
			z:        false,
			n:        true,
			h:        false,
			c:        false,
		},
		{
			name:     "carry only",
			in:       0x10,
			a:        0x00,
			expected: 0xF0,
			n:        true,
			h:        true,
			c:        false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			cpu.SUB(src)

			if got := cpu.A.Read(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}

func TestSBC(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		cIn        bool
		expected   byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
			n:    true,
			h:    true,
			c:    true,
		},
		{
			name:     "2 minus 1",
			in:       1,
			a:        2,
			expected: 1,
			n:        true,
			h:        true,
			c:        true,
		},
		{
			name:     "half carry",
			in:       0x1,
			a:        0xF0,
			expected: 0xEF,
			n:        true,
			h:        false,
			c:        true,
		},
		{
			name:     "borrow and half borrow",
			in:       0x1,
			a:        0x00,
			expected: 0xFF,
			z:        false,
			n:        true,
			h:        false,
			c:        false,
		},
		{
			name:     "borrow and half borrow (not on boundary)",
			in:       0x3,
			a:        0x01,
			expected: 0xFE,
			z:        false,
			n:        true,
			h:        false,
			c:        false,
		},
		{
			name:     "carry only",
			in:       0x10,
			a:        0x00,
			expected: 0xF0,
			n:        true,
			h:        true,
			c:        false,
		},
		{
			name:     "includes the carry",
			cIn:      true,
			in:       0x2,
			a:        0x01,
			expected: 0xFE,
			z:        false,
			n:        true,
			h:        false,
			c:        false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			cpu.F.SetC(test.cIn)
			cpu.SBC(src)

			if got := cpu.A.Read(); got != test.expected {
				t.Errorf("A: expected %d, got %d", test.expected, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}

func TestCP(t *testing.T) {
	var tests = []struct {
		name       string
		in         byte
		a          byte
		z, n, h, c bool
	}{
		{
			name: "zero",
			z:    true,
			n:    true,
			h:    true,
			c:    true,
		},
		{
			name: "2 minus 1",
			in:   1,
			a:    2,
			n:    true,
			h:    true,
			c:    true,
		},
		{
			name: "half carry",
			in:   0x1,
			a:    0xF0,
			n:    true,
			h:    false,
			c:    true,
		},
		{
			name: "borrow and half borrow",
			in:   0x1,
			a:    0x00,
			z:    false,
			n:    true,
			h:    false,
			c:    false,
		},
		{
			name: "borrow and half borrow (not on boundary)",
			in:   0x3,
			a:    0x01,
			z:    false,
			n:    true,
			h:    false,
			c:    false,
		},
		{
			name: "carry only",
			in:   0x10,
			a:    0x00,
			n:    true,
			h:    true,
			c:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.A.Write(test.a)
			src := &gameboy.Register{}
			src.Write(test.in)

			cpu.CP(src)

			if got := cpu.A.Read(); got != test.a {
				t.Errorf("A: expected %d, got %d", test.a, got)
			}

			if got := cpu.F.Z(); got != test.z {
				t.Errorf("Z Flag: expected %v, got %v", test.z, got)
			}
			if got := cpu.F.N(); got != test.n {
				t.Errorf("N Flag: expected %v, got %v", test.n, got)
			}
			if got := cpu.F.H(); got != test.h {
				t.Errorf("H Flag: expected %v, got %v", test.h, got)
			}
			if got := cpu.F.C(); got != test.c {
				t.Errorf("C Flag: expected %v, got %v", test.c, got)
			}
		})
	}
}
