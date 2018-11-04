package gameboy

import (
	"fmt"
	"strings"
	"time"
)

// Built while watching the ultimate Game Boy Talk
// https://www.youtube.com/watch?v=HyzD8pNlpwI

// Reading Game Boy Manual:
// http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf

type CPU struct {
	MMU *MMU

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

	// CB is a placeholder for the prefix
	CB struct{}

	// Cycles remaining to be used for operation
	cycles int

	// if true, no processing will be completed until an interrupt
	isHalted bool

	// if true, the CPU and LCD are halted until a button is pressed
	isStopped bool
}

// NewClock creates time.Ticker with suitable speed
// that can be used with cpu.Run
func NewClock() *time.Ticker {
	return time.NewTicker(time.Second / (4000000))
}

// NewCPU creates a CPU in a zeroed initial state.
func NewCPU(mmu *MMU) *CPU {
	cpu := &CPU{
		MMU: mmu,
		A:   &Register{}, F: &Register{},
		B: &Register{}, C: &Register{},
		D: &Register{}, E: &Register{},
		H: &Register{}, L: &Register{},

		SP: &Address{}, PC: &Address{},
	}
	cpu.AF = &RegisterPair{
		Low: cpu.F, High: cpu.A,
	}
	cpu.BC = &RegisterPair{
		Low: cpu.C, High: cpu.B,
	}
	cpu.DE = &RegisterPair{
		Low: cpu.E, High: cpu.D,
	}
	cpu.HL = &RegisterPair{
		Low: cpu.L, High: cpu.H,
	}

	return cpu
}

// LoadROM places the provided ROM data into RAM
func (c *CPU) LoadROM(data []byte) {
	for index, value := range data {
		c.MMU.Write8(uint16(index), value)
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
	scanTick := time.NewTicker(time.Millisecond)
	defer scanTick.Stop()

	for _ = range clock {
		if c.cycles > 0 {
			c.cycles--
			continue
		}

		// Fake scanlines
		c.fakeScanlines()

		if c.isHalted {
			// If interrupts are disabled (DI) then
			// halt doesn't suspend operation but it does cause
			// the program counter to stop counting for one
			// instruction
			if c.MMU.Read8(IE) == 0x0 {
				c.PC.Inc(1)
				c.cycles = 1
			}
			continue
		}

		c.execute()
	}
}

// Init initializes the device to an appropriate state for loading without a boot ROM
func (c *CPU) Init() {
	c.PC.Write16(0x100)
	c.AF.Write16(0x01)
	c.F.Write8(0xB0)
	c.BC.Write16(0x0013)
	c.DE.Write16(0x00D8)
	c.HL.Write16(0x014D)
	c.SP.Write16(0xFFFE)

	c.MMU.Write8(0xFF05, 0x0, 0x0, 0x0)
	// [$FF05] = $00 ; TIMA
	// [$FF06] = $00 ; TMA
	// [$FF07] = $00 ; TAC
	c.MMU.Write8(0xFF10, 0x80, 0xBF, 0xF3)
	// [$FF10] = $80 ; NR10
	// [$FF11] = $BF ; NR11
	// [$FF12] = $F3 ; NR12
	c.MMU.Write8(0xFF14, 0xBF)
	// [$FF14] = $BF ; NR14
	c.MMU.Write8(0xFF16, 0x3F, 0x00, 0x00, 0xBF)
	// [$FF16] = $3F ; NR21
	// [$FF17] = $00 ; NR22
	// [$FF19] = $BF ; NR24
	c.MMU.Write8(0xFF1A, 0x7F, 0xFF, 0x9F, 0x0, 0xBF, 0x00, 0xFF)
	// [$FF1A] = $7F ; NR30
	// [$FF1B] = $FF ; NR31
	// [$FF1C] = $9F ; NR32
	// [$FF1E] = $BF ; NR33
	// [$FF20] = $FF ; NR41
	c.MMU.Write8(0xFF21, 0x0, 0x0, 0xBF, 0x77, 0xF3, 0xF1)
	// [$FF21] = $00 ; NR42
	// [$FF22] = $00 ; NR43
	// [$FF23] = $BF ; NR30
	// [$FF24] = $77 ; NR50
	// [$FF25] = $F3 ; NR51
	// [$FF26] = $F1-GB, $F0-SGB ; NR52
	c.MMU.Write8(0xFF40, 0x91, 0x0, 0x00, 0x00, 0x0, 0x00, 0x00, 0xFC, 0xFF, 0xFF)
	// [$FF40] = $91 ; LCDC
	// [$FF42] = $00 ; SCY
	// [$FF43] = $00 ; SCX
	// [$FF45] = $00 ; LYC
	// [$FF47] = $FC ; BGP
	// [$FF48] = $FF ; OBP0
	// [$FF49] = $FF ; OBP1
	c.MMU.Write8(0xFF4A, 0x00, 0x00)
	// [$FF4A] = $00 ; WY
	// [$FF4B] = $00 ; WX
	c.MMU.Write8(0xFFFF, 0x00)
	// [$FFFF] = $00 ; IE
}

var scanLineCycleCount int

func (c *CPU) fakeScanlines() {
	if scanLineCycleCount < 50 {
		scanLineCycleCount++
		return
	}
	scanLineCycleCount = 0
	curline := c.MMU.Read8(CURLINE)
	curline++
	if curline > 144 {
		c.MMU.Write8(VBLANK, 0xFF)
	} else {
		c.MMU.Write8(VBLANK, 0x00)
	}
	if curline > 153 {
		curline = 0
	}
	c.MMU.Write8(CURLINE, curline)
}

func (c *CPU) GetOperation() (Opcode, Op) {
	pcBefore := c.PC.Read16()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic at: 0x%X\n", pcBefore)
			panic(r)
		}
	}()

	var table func(*CPU, Opcode) Op
	opcode := Opcode(c.MMU.Read8(c.PC.Read16()))
	switch opcode {
	case 0xCB:
		c.PC.Inc(1)
		opcode = Opcode(c.MMU.Read8(c.PC.Read16()))
		table = cbprefixedOpcodes
	default:
		table = unprefixedOpcodes
	}
	c.PC.Inc(1)
	op := table(c, opcode)
	return opcode, op
}

// execute handles the next operation
func (c *CPU) execute() {
	pcBefore := c.PC.Read16()
	flagsBefore := flagsToString(c.F)
	_, op := c.GetOperation()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("0x%X: %v\n", pcBefore, op.Description)
			panic(r)
		}
	}()
	if op.Instruction != nil {
		paramsBefore := paramsToString(op.Params...)
		op.Instruction(op.Params...)
		//if opcode != 0 {
		fmt.Printf(
			"0x%X:\t%v\t(%v - %v) -> (%v - %v)\n",
			pcBefore,
			op.Description,
			strings.Join(paramsBefore, ", "),
			flagsBefore,
			strings.Join(paramsToString(op.Params...), ", "),
			flagsToString(c.F),
		)
		//}
		c.cycles = op.Cycles[0] - 1
	}
}

func paramsToString(params ...Param) []string {
	var out []string
	for _, param := range params {
		if s, isString := param.(fmt.Stringer); isString {
			out = append(out, s.String())
			continue
		}
		if n, is8Bit := param.(Value8); is8Bit {
			out = append(out, fmt.Sprintf("0x%X", n.Read8()))
			continue
		}
		if n, is8BitSigned := param.(ValueSigned8); is8BitSigned {
			out = append(out, fmt.Sprintf("%d", n.ReadSigned8()))
			continue
		}
		if n, is16Bit := param.(Value16); is16Bit {
			out = append(out, fmt.Sprintf("0x%X", n.Read16()))
			continue
		}
	}
	return out
}

func flagsToString(f *Register) string {
	return fmt.Sprintf(
		"Z=%v, N=%v, H=%v, C=%v",
		f.Z(),
		f.N(),
		f.H(),
		f.C(),
	)
}
