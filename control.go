package gameboy

import "time"

// Control provides a means of managing running emulation
type Control struct {
	cpu *CPU
	ppu *PPU

	paused bool

	done chan struct{}

	Breakpoints []uint16
}

// NewControl creates a Control from a CPU and PPU.
// It is assumed they share an MMU.
func NewControl(cpu *CPU, ppu *PPU) *Control {
	return &Control{
		cpu:  cpu,
		ppu:  ppu,
		done: make(chan struct{}),
	}
}

// Start will begin running emulation.
// Emulation will step repeatedly until paused or stopped.
func (c *Control) Start() {
	go func() {
		for true {
			if c.paused {
				// Avoid a busy wait
				time.Sleep(time.Millisecond)
				continue
			}
			c.Step()
			select {
			case <-c.done:
				return
			default:
			}
		}
	}()
}

// TogglePaused with toggle the paused state of the control.
func (c *Control) TogglePaused() {
	c.paused = !c.paused
}

func (c *Control) IsPaused() bool {
	return c.paused
}

// Step will execute the next operation.
// This should usually only be used when paused.
func (c *Control) Step() {
	t := c.cpu.Step()
	c.ppu.Step(t)

	for _, bp := range c.Breakpoints {
		if c.cpu.PC.Read16() == bp {
			c.paused = true
			return
		}
	}
}

// Stop will stop the emulation.
// Once stopped, it cannot be restarted.
func (c *Control) Stop() {
	close(c.done)
}
