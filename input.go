package gameboy

import "github.com/theothertomelliott/gameboy/mmu"

type Key int

const (
	KeyUp Key = iota
	KeyDown
	KeyLeft
	KeyRight
	KeyA
	KeyB
	KeySelect
	KeyStart
)

func NewInput(interrupts *InterruptScheduler) *Input {
	i := &Input{
		interrupts: interrupts,
	}
	i.Reset()
	return i
}

type Input struct {
	states     [8]bool
	interrupts *InterruptScheduler
}

const (
	keyPressed = false
)

func (i *Input) Press(key Key) {
	if i.states[key] != keyPressed {
		i.interrupts.ScheduleInterrupt(InterruptJoypadPress)
	}
	i.states[key] = keyPressed
}

func (i *Input) Release(key Key) {
	i.states[key] = !keyPressed
}

func (i *Input) Reset() {
	for index := range i.states {
		i.states[index] = !keyPressed
	}
}

func (i *Input) Write(mmu *mmu.MMU) {
	joy := mmu.Read8(JOYPAD)
	p14 := bitValue(4, joy)
	p15 := bitValue(5, joy)

	if p15 == 0 {
		joy = setBitValue(0, joy, i.states[KeyA])
		joy = setBitValue(1, joy, i.states[KeyB])
		joy = setBitValue(2, joy, i.states[KeySelect])
		joy = setBitValue(3, joy, i.states[KeyStart])
	}
	if p14 == 0 {
		joy = setBitValue(0, joy, i.states[KeyRight])
		joy = setBitValue(1, joy, i.states[KeyLeft])
		joy = setBitValue(2, joy, i.states[KeyUp])
		joy = setBitValue(3, joy, i.states[KeyDown])
	}
	mmu.Write8(JOYPAD, joy)
}
