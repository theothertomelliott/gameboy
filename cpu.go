package gameboy

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

	A, F *Register
	B, C *Register
	D, E *Register
	H, L *Register

	SP *Address
	PC *Address

	Flags uint8

	D8  *Direct8
	D16 *Direct16

	RAM []byte
}

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

	return cpu
}

type Address struct {
	value uint16
}

func (a *Address) Read() uint16 {
	if a != nil {
		return a.value
	}
	return 0
}

func (a *Address) Write(value uint16) {
	if a != nil {
		a.value = value
	}
}

func (a *Address) Inc(amount uint16) {
	if a != nil {
		a.value += amount
	}
}

// Based on http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html
func (c *CPU) RunCycle() {
	switch c.RAM[c.PC.Read()] {
	case 0xCB:
		c.handleCBPrefixed()
	default:
		c.handleUnprefixed()
	}
}
