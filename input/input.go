package input

import (
	"github.com/theothertomelliott/gameboy/binary"
	"github.com/theothertomelliott/gameboy/interrupts"
	"github.com/theothertomelliott/gameboy/ioports"
	"github.com/theothertomelliott/gameboy/mmu"
)

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

func New(interrupts *interrupts.InterruptScheduler) *Input {
	i := &Input{
		interrupts: interrupts,
	}
	i.Reset()
	return i
}

type Input struct {
	states     [8]bool
	interrupts *interrupts.InterruptScheduler
}

const (
	keyPressed = false
)

func (i *Input) Press(key Key) {
	if i.states[key] != keyPressed {
		i.interrupts.ScheduleInterrupt(interrupts.InterruptJoypadPress)
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
	joy := mmu.Read8(ioports.JOYPAD)
	p14 := binary.Bit(4, joy)
	p15 := binary.Bit(5, joy)

	if p15 == 0 {
		joy = binary.SetBit(0, joy, i.states[KeyA])
		joy = binary.SetBit(1, joy, i.states[KeyB])
		joy = binary.SetBit(2, joy, i.states[KeySelect])
		joy = binary.SetBit(3, joy, i.states[KeyStart])
	}
	if p14 == 0 {
		joy = binary.SetBit(0, joy, i.states[KeyRight])
		joy = binary.SetBit(1, joy, i.states[KeyLeft])
		joy = binary.SetBit(2, joy, i.states[KeyUp])
		joy = binary.SetBit(3, joy, i.states[KeyDown])
	}
	mmu.Write8(ioports.JOYPAD, joy)
}
