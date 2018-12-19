package gameboy

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

func NewInput() *Input {
	i := &Input{}
	i.Reset()
	return i
}

type Input struct {
	states [8]bool // false == pressed
}

func (i *Input) Press(key Key) {
	i.states[key] = false
}

func (i *Input) Release(key Key) {
	i.states[key] = true
}

func (i *Input) Reset() {
	for index := range i.states {
		i.states[index] = true
	}
}

func (i *Input) Write(mmu *MMU) {
	joy := mmu.Read8(JOYPAD)
	if bitValue(4, joy) == 1 {
		joy = setBitValue(0, joy, i.states[KeyA])
		joy = setBitValue(1, joy, i.states[KeyB])
		joy = setBitValue(2, joy, i.states[KeySelect])
		joy = setBitValue(3, joy, i.states[KeyStart])
	}
	if bitValue(5, joy) == 1 {
		joy = setBitValue(0, joy, i.states[KeyRight])
		joy = setBitValue(1, joy, i.states[KeyLeft])
		joy = setBitValue(2, joy, i.states[KeyUp])
		joy = setBitValue(3, joy, i.states[KeyDown])
	}
	mmu.Write8(JOYPAD, joy)
}
