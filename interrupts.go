package gameboy

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

func NewInterruptScheduler(cpu *CPU, mmu *MMU) *InterruptScheduler {
	return &InterruptScheduler{
		cpu: cpu,
		mmu: mmu,
	}
}

type InterruptScheduler struct {
	cpu *CPU
	mmu *MMU
}

func (s *InterruptScheduler) ScheduleInterrupt(i Interrupt) {
	// Set the appropriate bit in IF to request an interrupt
	ifValue := s.mmu.Read8(IF)
	s.mmu.Write8(IF, setBitValue(i.Bit, ifValue, true))
}

func (s *InterruptScheduler) HandleInterrupts() {
	ieValue := s.mmu.Read8(IE)
	ifValue := s.mmu.Read8(IF)

	if ieValue&ifValue == 0 {
		return
	}

	for _, i := range allInterrupts {
		if bitValue(i.Bit, ieValue) != 0 && bitValue(i.Bit, ifValue) != 0 {
			// When halt is enabled, don't service the interrupt
			// Just un-halt the CPU
			if !s.cpu.IME {
				s.cpu.isHalted = false
				return
			}

			// Reset the bit in IF (the request)
			s.mmu.Write8(IF, setBitValue(i.Bit, ifValue, false))

			// Disable interrupts
			s.cpu.DI()

			// CALL interrupt vector
			s.cpu.CALL(Direct16(i.ISR))
			return
		}
	}
}
