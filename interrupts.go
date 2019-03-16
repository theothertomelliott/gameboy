package gameboy

import "fmt"

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
)

func NewInterruptScheduler(cpu *CPU, mmu *MMU) *InterruptScheduler {
	return &InterruptScheduler{
		cpu: cpu,
		mmu: mmu,
	}
}

type InterruptScheduler struct {
	scheduledInterrupt *Interrupt
	cpu                *CPU
	mmu                *MMU
}

func (s *InterruptScheduler) ScheduleInterrupt(i Interrupt) {
	if !s.cpu.IME {
		return
	}

	if s.scheduledInterrupt != nil {
		fmt.Println("interrupt scheduled before previous interrupt was processed")
	}

	// Set the VBlank bit in IF to request an interrupt
	ifValue := s.mmu.Read8(IF)
	s.mmu.Write8(IF, setBitValue(i.Bit, ifValue, true))
	s.scheduledInterrupt = &i
}

func (s *InterruptScheduler) HandleInterrupts() {
	if !s.cpu.IME || s.scheduledInterrupt == nil {
		return
	}

	// Clear interrupt
	defer func() {
		s.scheduledInterrupt = nil
	}()

	i := s.scheduledInterrupt
	ieValue := s.mmu.Read8(IE)
	ifValue := s.mmu.Read8(IF)
	if bitValue(i.Bit, ieValue&ifValue) != 1 {
		return
	}

	// Reset the bit in IF (the request)
	s.mmu.Write8(IF, setBitValue(i.Bit, ifValue, false))

	// CALL interrupt vector
	s.cpu.CALL(Direct16(i.ISR))

	if s.cpu.isHalted {
		s.cpu.isHalted = false
	}

	// Disable interrupts
	s.cpu.DI()
}
