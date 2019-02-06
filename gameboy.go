package gameboy

import (
	"fmt"
	"time"
)

const clockSpeedMHz = 0.8

// DMG provides a means of managing running emulation
type DMG struct {
	cpu    *CPU
	ppu    *PPU
	tracer *Tracer
	input  *Input

	interrupts *InterruptScheduler

	paused bool

	done chan struct{}

	Breakpoints map[uint16]struct{}

	err error
}

// NewDMG creates a Game Boy in an uninitialized state
func NewDMG() *DMG {
	tracer := NewTracer()
	mmu := NewMMU(tracer)
	cpu := NewCPU(mmu, tracer)
	interrupts := NewInterruptScheduler(cpu, mmu)
	ppu := NewPPU(mmu, interrupts)

	return &DMG{
		cpu:         cpu,
		ppu:         ppu,
		tracer:      tracer,
		interrupts:  interrupts,
		input:       NewInput(interrupts),
		done:        make(chan struct{}),
		Breakpoints: make(map[uint16]struct{}),
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

func (c *DMG) Input() *Input {
	return c.input
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
				fmt.Println(err)
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

// SetPaused sets the paused state of the control
func (c *DMG) SetPaused(paused bool) {
	c.paused = paused
}

// IsPaused returns the paused state of the control
// If the control has errored out, true is returned
func (c *DMG) IsPaused() bool {
	return c.paused || c.err != nil
}

// Step will execute the next operation.
// This should usually only be used when paused.
func (c *DMG) Step() error {
	c.tracer.Reset()
	t, err := c.cpu.Step()
	if err != nil {
		return err
	}
	c.tracer.Flush()
	err = c.ppu.Step(t)
	if err != nil {
		return err
	}

	for bp := range c.Breakpoints {
		if c.cpu.PC.Read16() == bp {
			c.paused = true
			return nil
		}
	}

	c.interrupts.HandleInterrupts()

	// Write input to memory
	c.input.Write(c.MMU())

	return nil
}

// Stop will stop the emulation.
// Once stopped, it cannot be restarted.
func (c *DMG) Stop() {
	close(c.done)
}
