package gameboy

import "errors"

// TODO: Implement all interrupts:
// http://imrannazar.com/GameBoy-Emulation-in-JavaScript:-Interrupts

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

func (s *InterruptScheduler) ScheduleInterrupt(i Interrupt) error {
	if !s.cpu.IME {
		return nil
	}

	if s.scheduledInterrupt != nil {
		return errors.New("interrupt scheduled before previous interrupt was processed")
	}

	// Set the VBlank bit in IF to request an interrupt
	ifValue := s.mmu.Read8(IF)
	s.mmu.Write8(IF, setBitValue(i.Bit, ifValue, true))
	s.scheduledInterrupt = &i
	return nil
}

func (s *InterruptScheduler) HandleInterrupts() {
	if !s.cpu.IME || s.scheduledInterrupt == nil {
		return
	}

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

	// Disable interrupts
	s.cpu.DI()

	// Clear interrupt
	s.scheduledInterrupt = nil
}
