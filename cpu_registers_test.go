package gameboy_test

import (
	"testing"

	"github.com/theothertomelliott/gameboy"
)

func TestRegister8(t *testing.T) {
	src := &gameboy.Register{}
	src.Write8(100)
	if src.Read8() != 100 {
		t.Errorf("expected %d, got %d", 100, src.Read8())
	}
}

func TestRegister16(t *testing.T) {
	src := &gameboy.RegisterPair{}
	src.Write16(100)
	t.Logf("%#v: %v, %v", src, src.Low, src.High)
	if src.Read16() != 100 {
		t.Errorf("expected %d, got %d", 100, src.Read16())
	}
}

func TestRegisterFlagRead(t *testing.T) {
	var tests = []struct {
		name       string
		value      byte
		z, n, h, c bool
	}{
		{
			name: "zero",
		},
		{
			name:  "all",
			value: 0xFF,
			z:     true,
			n:     true,
			h:     true,
			c:     true,
		},
		{
			name:  "Z",
			value: 128,
			z:     true,
		},
		{
			name:  "N",
			value: 64,
			n:     true,
		},
		{
			name:  "H",
			value: 32,
			h:     true,
		},
		{
			name:  "C",
			value: 16,
			c:     true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			src := &gameboy.Register{}
			src.Write8(test.value)
			if src.Z() != test.z {
				t.Errorf("Z: expected %v, got %v", test.z, src.Z())
			}
			if src.N() != test.n {
				t.Errorf("N: expected %v, got %v", test.n, src.N())
			}
			if src.H() != test.h {
				t.Errorf("H: expected %v, got %v", test.h, src.H())
			}
			if src.C() != test.c {
				t.Errorf("C: expected %v, got %v", test.c, src.C())
			}
		})
	}
}

func TestRegisterFlagSet(t *testing.T) {
	var tests = []struct {
		name       string
		value      byte
		z, n, h, c bool
	}{
		{
			name: "none",
		},
		{
			name: "all",
			z:    true,
			n:    true,
			h:    true,
			c:    true,
		},
		{
			name: "Z",
			z:    true,
		},
		{
			name: "N",
			n:    true,
		},
		{
			name: "H",
			h:    true,
		},
		{
			name: "C",
			c:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			src := &gameboy.Register{}
			src.Write8(test.value)

			src.SetZ(test.z)
			src.SetN(test.n)
			src.SetH(test.h)
			src.SetC(test.c)

			if src.Z() != test.z {
				t.Errorf("Z: expected %v, got %v", test.z, src.Z())
			}
			if src.N() != test.n {
				t.Errorf("N: expected %v, got %v", test.n, src.N())
			}
			if src.H() != test.h {
				t.Errorf("H: expected %v, got %v", test.h, src.H())
			}
			if src.C() != test.c {
				t.Errorf("C: expected %v, got %v", test.c, src.C())
			}
		})
	}
}
