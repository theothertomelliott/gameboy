package interrupts

import (
	"github.com/theothertomelliott/gameboy/binary"
	"github.com/theothertomelliott/gameboy/cpu"
	"github.com/theothertomelliott/gameboy/ioports"
	"github.com/theothertomelliott/gameboy/mmu"
)

type Interrupt struct {
	Bit byte
	ISR uint16
}

var (
	InterruptVBlank = Interrupt{
		Bit: 0,
		ISR: 0x0040,
	}
	InterruptLCDStatus = Interrupt{
		Bit: 1,
		ISR: 0x0048,
	}
	InterruptTimerOverflow = Interrupt{
		Bit: 2,
		ISR: 0x0050,
	}
	InterruptSerialLink = Interrupt{
		Bit: 3,
		ISR: 0x0058,
	}
	InterruptJoypadPress = Interrupt{
		Bit: 4,
		ISR: 0x0060,
	}
	allInterrupts = []Interrupt{
		InterruptVBlank,
		InterruptLCDStatus,
		InterruptTimerOverflow,
		InterruptSerialLink,
		InterruptJoypadPress,
	}
)

func New(cpu *cpu.CPU, mmu *mmu.MMU) *InterruptScheduler {
	return &InterruptScheduler{
		cpu: cpu,
		mmu: mmu,
	}
}

type InterruptScheduler struct {
	cpu *cpu.CPU
	mmu *mmu.MMU
}

func (s *InterruptScheduler) ScheduleInterrupt(i Interrupt) {
	// Set the appropriate bit in IF to request an interrupt
	ifValue := s.mmu.Read8(ioports.IF)
	s.mmu.Write8(ioports.IF, binary.SetBit(i.Bit, ifValue, true))
}

func (s *InterruptScheduler) HandleInterrupts() {
	ieValue := s.mmu.Read8(ioports.IE)
	ifValue := s.mmu.Read8(ioports.IF)

	if ieValue&ifValue == 0 {
		return
	}

	for _, i := range allInterrupts {
		if binary.Bit(i.Bit, ieValue) != 0 && binary.Bit(i.Bit, ifValue) != 0 {
			// When halt is enabled, don't service the interrupt
			// Just un-halt the CPU
			if !s.cpu.IME {
				s.cpu.SetHalted(false)
				return
			}

			// Reset the bit in IF (the request)
			s.mmu.Write8(ioports.IF, binary.SetBit(i.Bit, ifValue, false))

			// Disable interrupts
			s.cpu.DI()

			// CALL interrupt vector
			s.cpu.CALL(cpu.Direct16(i.ISR))
			return
		}
	}
}
