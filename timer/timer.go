package timer

import (
	"github.com/theothertomelliott/gameboy/interrupts"
	"github.com/theothertomelliott/gameboy/ioports"
	"github.com/theothertomelliott/gameboy/mmu"
)

type Timer struct {
	mmu         *mmu.MMU
	interrupts  *interrupts.InterruptScheduler
	clock       int
	divClocksum int
}

func New(mmu *mmu.MMU, interrupts *interrupts.InterruptScheduler) *Timer {
	return &Timer{
		mmu:        mmu,
		interrupts: interrupts,
	}
}

func (t *Timer) Increment(cycles int) {
	cycles = cycles / 4

	//	set divider
	t.divClocksum += cycles
	if t.divClocksum >= 256 {
		t.divClocksum -= 256
		div := t.mmu.Read8(ioports.DIVIDER)
		t.mmu.Write8(ioports.DIVIDER, div+1)
	}

	//	check if timer is on
	if ((t.mmu.Read8(ioports.TIMCONT) >> 2) & 0x1) != 0 {
		//	increase helper counter
		t.clock += cycles * 4

		//	set frequency
		freq := 4096                                 //	Hz
		if (t.mmu.Read8(ioports.TIMCONT) & 3) == 1 { //	mask last 2 bits
			freq = 262144
		} else if (t.mmu.Read8(ioports.TIMCONT) & 3) == 2 { //	mask last 2 bits
			freq = 65536
		} else if (t.mmu.Read8(ioports.TIMCONT) & 3) == 3 { //	mask last 2 bits
			freq = 16384
		}

		//	increment the timer according to the frequency (synched to the processed opcodes)
		for t.clock >= (4194304 / freq) {
			//	increase TIMA
			t.mmu.Write8(ioports.TIMECNT, t.mmu.Read8(ioports.TIMECNT)+1)
			//	check TIMA for overflow
			if t.mmu.Read8(ioports.TIMECNT) == 0x00 {
				//	set timer interrupt request
				t.interrupts.ScheduleInterrupt(interrupts.InterruptTimerOverflow)
				//	reset timer to timer modulo
				t.mmu.Write8(ioports.TIMECNT, t.mmu.Read8(ioports.TIMEMOD))
			}
			t.clock -= (4194304 / freq)
		}
	}
}
