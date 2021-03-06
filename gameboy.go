package gameboy

import (
	"fmt"
	"time"

	"github.com/theothertomelliott/gameboy/cpu"
	"github.com/theothertomelliott/gameboy/input"
	"github.com/theothertomelliott/gameboy/interrupts"
	"github.com/theothertomelliott/gameboy/mmu"
	"github.com/theothertomelliott/gameboy/ppu"
	"github.com/theothertomelliott/gameboy/timer"
	"github.com/theothertomelliott/gameboy/tracer"
)

// DMG provides a means of managing running emulation
type DMG struct {
	cpu         *cpu.CPU
	ppu         *ppu.PPU
	tracer      *tracer.Tracer
	input       *input.Input
	timer       *timer.Timer
	rateLimiter *RateLimiter

	interrupts *interrupts.InterruptScheduler

	paused bool

	done chan struct{}

	// TODO: This is not thread safe
	// Need to unexport and create an appropriate interface
	Breakpoints map[uint16]struct{}

	err error
}

// NewDMG creates a Game Boy in an uninitialized state
func NewDMG() *DMG {
	dmg := NewDMGWithNoRateLimit()
	dmg.rateLimiter = NewDefaultRateLimiter()
	return dmg
}

// NewDMGWithNoRateLimit creates a Game Boy in an uninitialized state
// with no rate limit to control speed
func NewDMGWithNoRateLimit() *DMG {
	tracer := tracer.New()
	m := mmu.New(tracer)
	cpu := cpu.New(m, tracer)
	interrupts := interrupts.New(cpu, m)
	ppu := ppu.New(m, interrupts)
	timer := timer.New(m, interrupts)

	return &DMG{
		cpu:         cpu,
		ppu:         ppu,
		tracer:      tracer,
		interrupts:  interrupts,
		input:       input.New(interrupts),
		done:        make(chan struct{}),
		Breakpoints: make(map[uint16]struct{}),
		timer:       timer,
	}
}

// Reset resets the emulator, retaining the current loaded cartridge
func (c *DMG) Reset() {
	c.cpu.MMU.ResetCartridge()
	c.cpu.Init()
}

func (c *DMG) CPU() *cpu.CPU {
	return c.cpu
}

func (c *DMG) PPU() *ppu.PPU {
	return c.ppu
}

func (c *DMG) Tracer() *tracer.Tracer {
	return c.tracer
}

func (c *DMG) MMU() *mmu.MMU {
	return c.cpu.MMU
}

func (c *DMG) Input() *input.Input {
	return c.input
}

// Start will begin running emulation.
// Emulation will step repeatedly until paused or stopped.
func (c *DMG) Start() {
	go func() {
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

	c.timer.Increment(t)
	if c.rateLimiter != nil {
		c.rateLimiter.Increment(t)
	}

	for bp := range c.Breakpoints {
		if c.cpu.PC.Read16() == bp {
			c.paused = true
			return nil
		}
	}

	// Write input to memory
	c.input.Write(c.MMU())

	c.interrupts.HandleInterrupts()

	return nil
}

// Stop will stop the emulation.
// Once stopped, it cannot be restarted.
func (c *DMG) Stop() {
	close(c.done)
}
