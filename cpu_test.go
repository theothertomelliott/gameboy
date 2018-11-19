package gameboy_test

import (
	"fmt"
	"log"
	"runtime/debug"
	"testing"

	"github.com/theothertomelliott/gameboy"
)

// TestPrograms runs some simple programs and verifies the state of the CPU after
func TestPrograms(t *testing.T) {
	var tests = []struct {
		name     string
		rom      []byte
		expected expectation
		cycles   int
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
				B: 0x22,
				C: 0x11,
			},
			cycles: 1,
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
			cycles: 3,
		},
		{
			name: "Push and Pop",
			rom: []byte{
				0x1, 0x1, 0x2, // LD BC, 0x0201
				0xC5,          // PUSH BC
				0x1, 0x3, 0x4, // LD BC, 0x0403
				0xC5,          // PUSH BC
				0x1, 0x0, 0x0, // LD BC, 0x0
				0xC1, // POP BC
				0xC1, // POP BC
			},
			expected: expectation{
				B: 0x2,
				C: 0x1,
			},
			cycles: 10,
		},
		{
			name: "Call and ret",
			rom: []byte{
				0xCD, 0x7, 0x0, // CALL 0x7
				0xC3, 0xB, 0x0, // JP 0xD
				0xDB,      // Bad opcode (should be skipped over)
				0x06, 0x1, // LD B, 0x1
				0x80,      // ADD B
				0xC9,      // RET
				0x0E, 0x5, // LD C, 0x1
			},
			expected: expectation{
				A: 0x1,
				B: 0x1,
				C: 0x5,
			},
			cycles: 10,
		},
		{
			name: "Copy into memory",
			rom: []byte{
				// Jump past data
				// (0x0000)
				0xC3, 0xD, 0x0, // JP 0xD
				// Data (0x3)
				0x1, 0x2, 0x3, 0x4, 0x5,
				// Write here (0x8)
				0x1, 0x2, 0x3, 0x4, 0x5,

				// Point to relevant places in memory
				// (0x000D)
				0x1, 0x3, 0x0, // LD BC, 0x3
				0x21, 0x8, 0x0, // LD HL, 0x8

				// Load memory into A, write into (HL)
				// (0x0013)
				0xA,  // LD A, (BC)
				0x77, // LD (HL),A

				0x23, // INC HL
				0x3,  // INC BC

				0x3E, 0x8, // LD A, 0x8
				0xB9,             // CP C
				0xC2, 0x13, 0x00, // JP NZ, 0x13

				// Clear registers
				0x1, 0x0, 0x0, // LD BC 0x0
				0x21, 0x0, 0x0, // LD HL 0x0
				0x3E, 0x0, // LD A 0x0

				0x76, // HALT
			},
			expected: expectation{
				F: 0xF0,
				RAM: map[uint16][]byte{
					0x8: []byte{0x1, 0x2, 0x3, 0x4, 0x5},
				},
			},
			cycles: 100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Panic: %v", r)
					debug.PrintStack()
				}
			}()

			mmu := gameboy.NewMMU()
			mmu.LoadROM(append(test.rom, make([]byte, 0xFF00)...))

			tracer := gameboy.NewTracer()
			tracer.Logger = func(t gameboy.TraceMessage) {
				log.Print(t.Event.Description)
			}

			cpu := gameboy.NewCPU(mmu, tracer)
			cpu.SP.Write16(0xFFFE) // Set up stack

			for count := 0; count < test.cycles; count++ {
				_, err := cpu.Step()
				if err != nil {
					t.Fatal(err)
				}
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

	RAM map[uint16][]byte
}

func (e expectation) compare(t *testing.T, cpu *gameboy.CPU) {
	t.Helper()

	e.compareReg(t, "A", cpu.A, e.A)
	e.compareReg(t, "F", cpu.F, e.F)
	e.compareReg(t, "B", cpu.B, e.B)
	e.compareReg(t, "C", cpu.C, e.C)
	e.compareReg(t, "D", cpu.D, e.D)
	e.compareReg(t, "E", cpu.E, e.E)
	e.compareReg(t, "H", cpu.H, e.H)
	e.compareReg(t, "L", cpu.L, e.L)

	// Check memory
	for index, ram := range e.RAM {
		for offset, value := range ram {
			if value != cpu.MMU.RAM[index+uint16(offset)] {
				t.Errorf("RAM at 0x%x did not match expected value, got %v", index, sprintRAM(cpu, int(index), len(ram)))
				break
			}
		}
	}
}

func (e expectation) compareReg(t *testing.T, name string, r *gameboy.Register, expected byte) {
	t.Helper()
	if got := r.Read8(); got != expected {
		t.Errorf("%s: expected 0x%X, got 0x%X", name, expected, got)
	}
}

func sprintRAM(cpu *gameboy.CPU, index, length int) []string {
	var out []string
	data := cpu.MMU.RAM[index : index+length]
	for _, d := range data {
		out = append(out, fmt.Sprintf("0x%X", d))
	}
	return out
}
