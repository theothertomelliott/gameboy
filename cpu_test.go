package gameboy_test

import (
	"testing"
	"time"

	"github.com/theothertomelliott/gameboy"
)

// TestPrograms runs some simple programs and verifies the state of the CPU after
func TestPrograms(t *testing.T) {
	var tests = []struct {
		name       string
		rom        []byte
		expected   expectation
		operations int
	}{
		{
			name: "empty",
		},
		{
			name: "LD",
			rom: []byte{
				0x01, 0x11, 0x22, // LD BC, 0x2211 (12)
			},
			expected: expectation{
				B: 0x11,
				C: 0x22,
			},
			operations: 12,
		},
		{
			name: "0x2 + 0x1",
			rom: []byte{
				0x3E, 0x2, // LD A, 0x2 (8)
				0x06, 0x1, // LD B, 0x1 (16)
				0x80, // ADD A, B (8)
			},
			expected: expectation{
				A: 0x3,
				B: 0x1,
			},
			operations: 8 + 16 + 8,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpu := gameboy.NewCPU()
			cpu.LoadROM(test.rom)

			var clock = make(chan time.Time)
			go func() {
				for i := 0; i < test.operations; i++ {
					clock <- time.Now()
				}
				close(clock)
			}()

			var exited = make(chan struct{})
			go func() {
				cpu.Run(clock)
				exited <- struct{}{}
			}()

			select {
			case <-exited:
			case <-time.After(time.Second):
				t.Errorf("Timed out waiting for CPU to exit")
			}
			test.expected.compare(t, cpu)
		})
	}
}

type expectation struct {
	A    byte
	F    byte
	B, C byte
	D, E byte
	H, L byte

	RAM map[uint16]byte
}

func (e expectation) compare(t *testing.T, cpu *gameboy.CPU) {
	e.compareReg(t, "A", cpu.A, e.A)
	e.compareReg(t, "F", cpu.F, e.F)
	e.compareReg(t, "B", cpu.B, e.B)
	e.compareReg(t, "C", cpu.C, e.C)
	e.compareReg(t, "D", cpu.D, e.D)
	e.compareReg(t, "E", cpu.E, e.E)
	e.compareReg(t, "H", cpu.H, e.H)
	e.compareReg(t, "L", cpu.L, e.L)
}

func (e expectation) compareReg(t *testing.T, name string, r *gameboy.Register, expected byte) {
	if got := r.Read8(); got != expected {
		t.Errorf("%s: expected 0x%X, got 0x%X", name, expected, got)
	}
}
