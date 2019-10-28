package cpu

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/theothertomelliott/gameboy/mmu"
	"github.com/theothertomelliott/gameboy/tracer"
)

// Built while watching the ultimate Game Boy Talk
// https://www.youtube.com/watch?v=HyzD8pNlpwI

// Reading Game Boy Manual:
// http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf

type CPU struct {
	MMU *mmu.MMU

	// Interrupt master enable
	IME bool

	// Registers
	AF *RegisterPair
	BC *RegisterPair
	DE *RegisterPair
	HL *RegisterPair

	A    *Register
	F    *FRegister
	B, C *Register
	D, E *Register
	H, L *Register

	SP *StackPointer
	PC *ProgramCounter

	// CB is a placeholder for the prefix
	CB struct{}

	// if true, no processing will be completed until an interrupt
	isHalted bool

	// if true, the CPU and LCD are halted until a button is pressed
	isStopped bool

	tracer CPUTracer

	countStartTime      time.Time
	opCount             uint64
	OperationsPerSecond uint64
}

func (c *CPU) IsHalted() bool {
	return c.isHalted
}

func (c *CPU) SetHalted(v bool) {
	c.isHalted = v
}

type CPUTracer interface {
	tracer.LogTracer
	RegisterTracer
	AddCPU(pc uint16, description string)
	AddStack(pos, in, out uint16)
}

// New creates a CPU in a zeroed initial state.
func New(mmu *mmu.MMU, tracer CPUTracer) *CPU {
	cpu := &CPU{
		MMU: mmu,

		A: NewRegister("A", tracer), F: &FRegister{Register: *NewRegister("F", tracer)},
		B: NewRegister("B", tracer), C: NewRegister("C", tracer),
		D: NewRegister("D", tracer), E: NewRegister("E", tracer),
		H: NewRegister("H", tracer), L: NewRegister("L", tracer),

		SP: &StackPointer{}, PC: &ProgramCounter{},
		tracer: tracer,
	}
	cpu.AF = NewRegisterPair("AF", cpu.A, cpu.F)
	cpu.BC = NewRegisterPair("BC", cpu.B, cpu.C)
	cpu.DE = NewRegisterPair("DE", cpu.D, cpu.E)
	cpu.HL = NewRegisterPair("HL", cpu.H, cpu.L)

	return cpu
}

type StackPointer struct {
	Address
}

func (s *StackPointer) String() string {
	return "SP"
}

type ProgramCounter struct {
	Address
}

func (s *ProgramCounter) String() string {
	return "PC"
}

type Address struct {
	value uint16
}

func (a *Address) String() string {
	return fmt.Sprintf("0x%X", a.value)
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

func (c *CPU) Logf(message string, args ...interface{}) {
	if c.tracer == nil {
		return
	}
	c.tracer.Logf(message, args...)
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

func (c *CPU) CountSpeed() {
	if c.opCount == 0 {
		c.countStartTime = time.Now()
	}
	c.opCount++

	if c.opCount == 1000 {
		c.OperationsPerSecond = uint64(float64(time.Second/time.Since(c.countStartTime)) * float64(c.opCount))
		c.opCount = 0
	}
}

// Step handles the next operation
func (c *CPU) Step() (int, error) {

	if c.isHalted {
		return 4, nil
	}

	pcBefore := c.PC.Read16()
	description, cycles, err := c.ExecuteOperation()
	if err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("0x%04X", pcBefore))
	}
	if c.tracer != nil {
		c.tracer.AddCPU(pcBefore, description)
	}

	return cycles[0], nil
}

func (c *CPU) ExecuteOperation() (string, []int, error) {
	c.CountSpeed()
	var handler func(*CPU, Opcode) (string, []int, error)
	opcode := Opcode(c.MMU.Read8(c.PC.Read16()))
	switch opcode {
	case 0xCB:
		c.PC.Inc(1)
		opcode = Opcode(c.MMU.Read8(c.PC.Read16()))
		handler = cbprefixedHandler
	default:
		handler = unprefixedHandler
	}
	c.PC.Inc(1)
	return handler(c, opcode)
}
