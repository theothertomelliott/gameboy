package gameboy

import (
	"fmt"
	"runtime"
	"testing"
)

// Iterate through all opcodes and verify they can be called with empty values

func TestUnprefixedOpcodes(t *testing.T) {
	executeAllOpcodes(t, unprefixedOpcodes)
}

func TestCBPrefixedOpcodes(t *testing.T) {
	executeAllOpcodes(t, cbprefixedOpcodes)
}

func executeAllOpcodes(t *testing.T, opcodeMap func(c *CPU) map[Opcode]Op) {
	cpu := NewCPU()
	for code, def := range opcodeMap(cpu) {
		t.Run(fmt.Sprintf("0x%X", code), func(t *testing.T) {
			executeOp(t, cpu, def)
		})
	}
}

func executeOp(t *testing.T, cpu *CPU, op Op) {
	cpu.PC.Write16(0)
	defer func() {
		// Offset PC and stack to handle decrementing
		cpu.SP.Write16(500)
		cpu.PC.Write16(500)
		// Reset RAM and registers
		cpu.RAM = make([]byte, 0x10000)
		cpu.AF.Write16(0)
		cpu.BC.Write16(0)
		cpu.DE.Write16(0)
		cpu.HL.Write16(0)

		if r := recover(); r != nil {
			t.Errorf("panic: %v", r)
			trace := make([]byte, 1024)
			_ = runtime.Stack(trace, true)
			t.Log(string(trace))
		}
	}()
	op.Instruction(op.Params...)
}
