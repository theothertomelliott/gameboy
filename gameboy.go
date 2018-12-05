package gameboy

import (
	"time"
)

const clockSpeedMHz = 0.8

// DMG provides a means of managing running emulation
type DMG struct {
	cpu    *CPU
	ppu    *PPU
	tracer *Tracer

	paused bool

	done chan struct{}

	Breakpoints []uint16

	err error
}

// NewDMG creates a Game Boy in an uninitialized state
func NewDMG() *DMG {
	tracer := NewTracer()
	mmu := NewMMU(tracer)
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
		var (
			ticks               int
			lastSync            = time.Now()
			cyclesPerSecond     = clockSpeedMHz * 1000 * 1000
			clockSyncsPerSecond = float64(1000)
		)
		for true {
			// Don't continue after error
			if c.err != nil {
				return
			}

			if c.paused {
				// Avoid a busy wait
				time.Sleep(time.Millisecond)
				continue
			}
			err := c.Step()
			if err != nil {
				c.err = err
			}

			// Correct timing
			ticks++
			if float64(ticks) >= cyclesPerSecond/clockSyncsPerSecond {
				elapsed := time.Since(lastSync)
				if elapsed < time.Second/time.Duration(clockSyncsPerSecond) {
					time.Sleep(time.Second/time.Duration(clockSyncsPerSecond) - elapsed)
				}
				ticks = 0
				lastSync = time.Now()
			}

			select {
			case <-c.done:
				return
			default:
			}
		}
	}()
}

// Err returns any error in emulation.
// Emulation will stop running on the first error.
func (c *DMG) Err() error {
	return c.err
}

// TogglePaused with toggle the paused state of the control.
func (c *DMG) TogglePaused() {
	c.paused = !c.paused
}

func (c *DMG) IsPaused() bool {
	return c.paused || c.err != nil
}

// Step will execute the next operation.
// This should usually only be used when paused.
func (c *DMG) Step() error {
	t, err := c.cpu.Step()
	if err != nil {
		return err
	}
	err = c.ppu.Step(t)
	if err != nil {
		return err
	}

	for _, bp := range c.Breakpoints {
		if c.cpu.PC.Read16() == bp {
			c.paused = true
			return nil
		}
	}

	c.tracer.Log()

	return nil
}

// Stop will stop the emulation.
// Once stopped, it cannot be restarted.
func (c *DMG) Stop() {
	close(c.done)
}
