package gameboy

type Timer struct {
	mmu        *MMU
	interrupts *InterruptScheduler
	clock      int
}

func NewTimer(mmu *MMU, interrupts *InterruptScheduler) *Timer {
	return &Timer{
		mmu:        mmu,
		interrupts: interrupts,
	}
}

func (t *Timer) Increment(cycles int) {
	t.clock += cycles

	var threshold int
	div := t.mmu.Read8(DIVIDER)
	timerControl := t.mmu.Read8(TIMCONT)
	timerCount := t.mmu.Read8(TIMECNT)
	modulo := t.mmu.Read8(TIMEMOD)

	if (t.clock % 16) == 0 {
		t.mmu.Write8(DIVIDER, div+1)
	}

	if timerControl&4 == 0 {
		return
	}

	switch timerControl & 3 {
	case 0:
		threshold = 64
	case 1:
		threshold = 1
	case 2:
		threshold = 4
	case 3:
		threshold = 16
	}

	if t.clock > threshold {
		t.clock = 0
		t.mmu.Write8(TIMECNT, timerCount+1)
	}

	if timerCount == 255 {
		t.mmu.Write8(TIMECNT, modulo)
		t.interrupts.ScheduleInterrupt(InterruptTimerOverflow)
	}

}
