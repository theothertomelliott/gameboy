package timer

import (
	"github.com/theothertomelliott/gameboy/interrupts"
	"github.com/theothertomelliott/gameboy/ioports"
	"github.com/theothertomelliott/gameboy/mmu"
)

type Timer struct {
	mmu        *mmu.MMU
	interrupts *interrupts.InterruptScheduler
	clock      int
}

func New(mmu *mmu.MMU, interrupts *interrupts.InterruptScheduler) *Timer {
	return &Timer{
		mmu:        mmu,
		interrupts: interrupts,
	}
}

func (t *Timer) Increment(cycles int) {
	t.clock += cycles

	var threshold int
	div := t.mmu.Read8(ioports.DIVIDER)
	timerControl := t.mmu.Read8(ioports.TIMCONT)
	timerCount := t.mmu.Read8(ioports.TIMECNT)
	modulo := t.mmu.Read8(ioports.TIMEMOD)

	if (t.clock % 16) == 0 {
		t.mmu.Write8(ioports.DIVIDER, div+1)
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
		t.mmu.Write8(ioports.TIMECNT, timerCount+1)
	}

	if timerCount == 255 {
		t.mmu.Write8(ioports.TIMECNT, modulo)
		t.interrupts.ScheduleInterrupt(interrupts.InterruptTimerOverflow)
	}

}
