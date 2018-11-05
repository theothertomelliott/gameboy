package main

import "time"

type control struct {
	C    chan struct{}
	step chan struct{}

	Paused bool

	timer *time.Ticker
}

func NewControl(duration time.Duration) *control {
	c := &control{
		C:     make(chan struct{}),
		step:  make(chan struct{}),
		timer: time.NewTicker(duration),
	}
	go func() {
		for range c.timer.C {
			if c.Paused {
				continue
			}
			c.C <- struct{}{}
		}
	}()
	go func() {
		for range c.step {
			c.C <- struct{}{}
		}
	}()
	return c
}

func (c *control) Step() {
	c.step <- struct{}{}
}

func (c *control) Close() {
	c.timer.Stop()
	close(c.step)
	close(c.C)
}
