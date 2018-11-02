package gameboy

import (
	"fmt"
	"time"
)

// Built while watching the ultimate Game Boy Talk
// https://www.youtube.com/watch?v=HyzD8pNlpwI

// Reading Game Boy Manual:
// http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf

type CPU struct {

	// Registers
	AF *RegisterPair
	BC *RegisterPair
	DE *RegisterPair
	HL *RegisterPair

	A    *Register
	F    *Register
	B, C *Register
	D, E *Register
	H, L *Register

	SP *Address
	PC *Address

	D8  *Direct8
	D16 *Direct16

	RAM []byte

	// CB is a placeholder for the prefix
	CB struct{}

	// Cycles remaining to be used for operation
	cycles int

	// if true, no processing will be completed until an interrupt
	isHalted bool

	// if true, the CPU and LCD are halted until a button is pressed
	isStopped bool
}

const (
	// Interrupt flag locations
	// See: http://bgb.bircd.org/pandocs.htm#interrupts
	IE = 0xFFFF // Interrupt Enable (R/W)
	IF = 0xFF0F // Interrupt Flag (R/W)
)

func NewCPU() *CPU {
	cpu := &CPU{
		A: &Register{}, F: &Register{},
		B: &Register{}, C: &Register{},
		D: &Register{}, E: &Register{},
		H: &Register{}, L: &Register{},

		SP: &Address{}, PC: &Address{},
	}
	cpu.AF = &RegisterPair{
		Low: cpu.A, High: cpu.F,
	}
	cpu.BC = &RegisterPair{
		Low: cpu.B, High: cpu.C,
	}
	cpu.DE = &RegisterPair{
		Low: cpu.D, High: cpu.E,
	}
	cpu.HL = &RegisterPair{
		Low: cpu.H, High: cpu.L,
	}

	cpu.D8 = &Direct8{
		CPU: cpu,
	}
	cpu.D16 = &Direct16{
		CPU: cpu,
	}

	cpu.RAM = make([]byte, 0x10000)

	return cpu
}

// LoadROM places the provided ROM data into RAM
func (c *CPU) LoadROM(data []byte) {
	for index, value := range data {
		c.RAM[index] = value
	}
}

type Address struct {
	value uint16
}

func (a *Address) Read16() uint16 {
	if a != nil {
		return a.value
	}
	return 0
}

func (a *Address) Write16(value uint16) {
	if a != nil {
		a.value = value
	}
}

func (a *Address) Inc(amount int8) {
	if a != nil {
		v := int32(a.value) + int32(amount)
		a.value = uint16(v)
	}
}

func (c *CPU) Run(clock <-chan time.Time) {
	for _ = range clock {
		if c.isHalted {
			// If interrupts are disabled (DI) then
			// halt doesn't suspend operation but it does cause
			// the program counter to stop counting for one
			// instruction
			continue
		}
		if c.cycles > 0 {
			c.cycles--
			continue
		}

		c.Cycle()
	}
}

// Based on http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html
func (c *CPU) Cycle() {
	var table map[Opcode]Op
	opcode := Opcode(c.RAM[c.PC.Read16()])
	var isCB bool
	switch opcode {
	case 0xCB:
		c.PC.Inc(1)
		opcode = Opcode(c.RAM[c.PC.Read16()])
		table = cbprefixedOpcodes(c)
		isCB = true
	default:
		table = unprefixedOpcodes(c)
	}
	op := table[opcode]

	defer func() {
		if r := recover(); r != nil {
			if isCB {
				fmt.Print("0xCB ")
			}
			fmt.Printf("%#x\n", opcode)
			panic(r)
		}
	}()
	c.PC.Inc(1)
	if op.Instruction != nil {
		op.Instruction(op.Params...)
	}
	c.cycles = op.Cycles[0] - 1
}
