package gameboy_test

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/theothertomelliott/gameboy"
)

func neg(value uint8) byte {
	return 0xFF - value + 1
}

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
			name: "LDH",
			rom: []byte{
				0x3E, 0x11, // LD A 0x11
				0xE0, 0x22, // LD (0xFF22), A
				0x3E, 0x11, // LD A 0x0
				0xF0, 0x22, // LD A, (0xFF22)
			},
			expected: expectation{
				A: 0x11,
				RAM: map[uint16][]byte{
					0xFF22: []byte{0x11},
				},
			},
			cycles: 4,
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
				0xC3, 0xD, 0x0, // JP 0xD
				0xDB,      // Bad opcode (should be skipped over)
				0x06, 0x1, // LD B, 0x1
				0x80,      // ADD B
				0xC9,      // RET
				0x0E, 0x5, // LD C, 0x5
			},
			expected: expectation{
				A: 0x1,
				B: 0x1,
			},
			cycles: 10,
		},
		{
			name: "Call and retc",
			rom: []byte{
				0xCD, 0x7, 0x0, // CALL 0x7
				0xC3, 0x10, 0x0, // JP 0xE
				0xDB,      // Bad opcode (should be skipped over)
				0x06, 0x1, // LD B, 0x1
				0x80,      // ADD B
				0xC8,      // RET Z
				0x06, 0x2, // LD B, 0x2
				0xC0,      // RET NZ
				0x0E, 0x5, // LD C, 0x1
			},
			expected: expectation{
				A: 0x1,
				B: 0x2,
			},
			cycles: 10,
		},
		{
			name: "JP (HL)",
			rom: []byte{
				0x06, 0x1, // LD B, 0x1
				0x21, 0x7, 0x0, // LD HL, 0x7
				0xE9,      // JP (HL)
				0xDB,      // Bad opcode (should be skipped over)
				0x06, 0x2, // LD B, 0x2
			},
			expected: expectation{
				B: 0x2,
				H: 0x0,
				L: 0x7,
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
				0x0, 0x0, 0x0, 0x0, 0x0,

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
		{
			name: "Copy into memory (LDI)",
			rom: []byte{
				// Jump past data
				// (0x0000)
				0xC3, 0xD, 0x0, // JP 0xD
				// Data (0x3)
				0x1, 0x2, 0x3, 0x4, 0x5,
				// Write here (0x8)
				0x0, 0x0, 0x0, 0x0, 0x0,
				// Write here again (0xD)
				0x0, 0x0, 0x0, 0x0, 0x0,

				// Point to relevant places in memory
				// (0x0012)
				0x1, 0x3, 0x0, // LD BC, 0x3
				0x21, 0x8, 0x0, // LD HL, 0x8

				// Load memory into A, write into (HL)
				// (0x0018)
				0xA,  // LD A, (BC)
				0x22, // LDI (HL),A

				0x3, // INC BC

				0x3E, 0x8, // LD A, 0x8
				0xB9,             // CP C
				0xC2, 0x18, 0x00, // JP NZ, 0x18

				// Point to relevant places in memory
				// (0x0012)
				0x1, 0xD, 0x0, // LD BC, 0xD
				0x21, 0x3, 0x0, // LD HL, 0x3

				// Load memory into A, write into (HL)
				// (0x0018)
				0x2A, // LDI A, (HL)
				0x2,  // LD (BC), A

				0x3, // INC BC

				0x3E, 0x12, // LD A, 0x12
				0xB9,         // CP C
				0x20, neg(8), // JR NZ, -8

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
					0xD: []byte{0x1, 0x2, 0x3, 0x4, 0x5},
				},
			},
			cycles: 100,
		},
		{
			name: "Copy into memory (LDD)",
			rom: []byte{
				// Jump past data
				// (0x0000)
				0xC3, 0xD, 0x0, // JP 0xD
				// Data (0x3)
				0x1, 0x2, 0x3, 0x4, 0x5,
				// Write here (0x8)
				0x0, 0x0, 0x0, 0x0, 0x0,
				// Write here again (0xD)
				0x0, 0x0, 0x0, 0x0, 0x0,

				// Point to relevant places in memory
				// (0x0012)
				0x1, 0x3, 0x0, // LD BC, 0x3
				0x21, 0xC, 0x0, // LD HL, 0x12

				// Load memory into A, write into (HL)
				// (0x0018)
				0xA,  // LD A, (BC)
				0x32, // LDD (HL),A

				0x3, // INC BC

				0x3E, 0x8, // LD A, 0x8
				0xB9,             // CP C
				0xC2, 0x18, 0x00, // JP NZ, 0x18

				// Point to relevant places in memory
				// (0x0012)
				0x1, 0xD, 0x0, // LD BC, 0xD
				0x21, 0x7, 0x0, // LD HL, 0x7

				// Load memory into A, write into (HL)
				// (0x0018)
				0x3A, // LDD A, (HL)
				0x2,  // LD (BC), A

				0x3, // INC BC

				0x3E, 0x12, // LD A, 0x12
				0xB9,         // CP C
				0x20, neg(8), // JR NZ, -8

				// Clear registers
				0x1, 0x0, 0x0, // LD BC 0x0
				0x21, 0x0, 0x0, // LD HL 0x0
				0x3E, 0x0, // LD A 0x0

				0x76, // HALT
			},
			expected: expectation{
				F: 0xF0,
				RAM: map[uint16][]byte{
					0x8: []byte{0x5, 0x4, 0x3, 0x2, 0x1},
					0xD: []byte{0x5, 0x4, 0x3, 0x2, 0x1},
				},
			},
			cycles: 100,
		},
		{
			name: "POP AF",
			rom: []byte{
				0xC3, 0x4, 0x0, // 		JP 0x4
				0xDB,             // 	text_failed (bad opcode)
				0x01, 0x00, 0x12, // 	ld   bc,$1200
				0xC5,       // -    	push bc
				0xF1,       //      	pop  af
				0xF5,       //      	push af
				0xD1,       //      	pop  de
				0x79,       //      	ld   a,c
				0xE6, 0xF0, //      	and  $F0
				0xBB,            //     cp   e
				0xC2, 0x3, 0x00, //     jp   nz,test_failed
				0x04,       //      	inc  b
				0x0C,       //      	inc  c
				0x20, 0xF1, //      	jr   nz,-
				0x76, // HALT
			},
			expected: expectation{
				A: 0x3F,
				F: 0x20,
				B: 0x3F,
				C: 0x2D,
				D: 0x3E,
				E: 0x20,
			},
			cycles: 500,
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
			tracer := gameboy.NewTracer()

			mmu := gameboy.NewMMU(tracer)
			mmu.LoadROM(append(test.rom, make([]byte, 0xFF00)...))

			cpu := gameboy.NewCPU(mmu, tracer)
			cpu.SP.Write16(0xFFFE) // Set up stack

			tracer.Logger = func(tm gameboy.TraceMessage) {
				if tm.CPU == nil {
					return
				}
				t.Logf("0x%04X: %v", tm.CPU.PC, tm.CPU.Description)
				t.Logf("BC=%04X, AF=%04X, DE=%04X", cpu.BC.Read16(), cpu.AF.Read16(), cpu.DE.Read16())
				t.Logf("A=0x%02X F=%02X", cpu.A.Read8(), cpu.F.Read8())
			}

			for count := 0; count < test.cycles; count++ {
				_, err := cpu.Step()
				if err != nil {
					t.Fatal(err)
				}
				tracer.Flush()
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

func (e expectation) compareReg(t *testing.T, name string, r gameboy.Value8, expected byte) {
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
