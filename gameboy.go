package gameboy

import (
	"time"
)

// DMG provides a means of managing running emulation
type DMG struct {
	cpu    *CPU
	ppu    *PPU
	tracer *Tracer

	paused bool

	done chan struct{}

	Breakpoints []uint16
}

// NewDMG creates a Game Boy in an uninitialized state
func NewDMG() *DMG {
	tracer := NewTracer()
	mmu := NewMMU()
	cpu := NewCPU(mmu, tracer)
	ppu := NewPPU(mmu)

	return &DMG{
		cpu:    cpu,
		ppu:    ppu,
		tracer: tracer,
		done:   make(chan struct{}),
	}
}

func (c *DMG) CPU() *CPU {
	return c.cpu
}

func (c *DMG) PPU() *PPU {
	return c.ppu
}

func (c *DMG) Tracer() *Tracer {
	return c.tracer
}

func (c *DMG) MMU() *MMU {
	return c.cpu.MMU
}

// Start will begin running emulation.
// Emulation will step repeatedly until paused or stopped.
func (c *DMG) Start() {
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
			time.Sleep(time.Nanosecond)
		}
	}()
}

// TogglePaused with toggle the paused state of the control.
func (c *DMG) TogglePaused() {
	c.paused = !c.paused
}

func (c *DMG) IsPaused() bool {
	return c.paused
}

// Step will execute the next operation.
// This should usually only be used when paused.
func (c *DMG) Step() {
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
func (c *DMG) Stop() {
	close(c.done)
}
