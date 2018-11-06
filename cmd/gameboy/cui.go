package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/theothertomelliott/gameboy"
)

type cui struct {
	gui     *gocui.Gui
	cpu     *gameboy.CPU
	tracer  *gameboy.Tracer
	control *gameboy.Control

	traceBuffer      []gameboy.TraceMessage
	traceBufferIndex int
}

func (c *cui) Close() {
	c.gui.Close()
}

func setupCUI(cpu *gameboy.CPU, tracer *gameboy.Tracer, control *gameboy.Control) (*cui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}

	cui := &cui{
		gui:     g,
		cpu:     cpu,
		tracer:  tracer,
		control: control,

		traceBuffer: make([]gameboy.TraceMessage, 10000),
	}

	g.SetManagerFunc(cui.layout)
	return cui, nil
}

func (c *cui) updateTrace() {
	c.gui.Update(func(g *gocui.Gui) error {
		v, err := g.View("trace")
		if err != nil {
			return nil
		}
		for _, trace := range c.traceBuffer {
			fmt.Fprintf(v, "0x%X: %v\n", trace.Count, trace.Event.Description)
		}
		return nil
	})
}

func (c *cui) updateRegisters() {
	c.gui.Update(func(g *gocui.Gui) error {
		v, err := g.View("registers")
		if err != nil {
			return nil
		}
		v.Clear()

		fmt.Fprintf(v, "A: 0x%X, F: 0x%X\n", c.cpu.A.Read8(), c.cpu.F.Read8())
		fmt.Fprintf(v, "B: 0x%X, C: 0x%X\n", c.cpu.B.Read8(), c.cpu.C.Read8())
		fmt.Fprintf(v, "D: 0x%X, E: 0x%X\n", c.cpu.D.Read8(), c.cpu.E.Read8())
		fmt.Fprintf(v, "H: 0x%X, L: 0x%X\n", c.cpu.H.Read8(), c.cpu.L.Read8())

		fmt.Fprintln(v)

		fmt.Fprintf(v, "PC: 0x%X\n", c.cpu.PC.Read16())
		fmt.Fprintf(v, "SP: 0x%X\n", c.cpu.SP.Read16())

		fmt.Fprintln(v)

		fmt.Fprintf(v, "Ops/Second: %v\n", c.cpu.OperationsPerSecond)

		return nil
	})
}

func startCUI(cui *cui) {
	go func() {
		for trace := range cui.tracer.Event {
			cui.traceBuffer = append(cui.traceBuffer, trace)
			if len(cui.traceBuffer) < 1000 {
				continue
			}
			cui.updateTrace()
			cui.updateRegisters()
			cui.traceBuffer = nil
		}
	}()

	g := cui.gui
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'p', gocui.ModNone, cui.pause); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 's', gocui.ModNone, cui.step); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
	}
}

func (c *cui) pause(g *gocui.Gui, v *gocui.View) error {
	c.control.TogglePaused()
	return nil
}

func (c *cui) step(g *gocui.Gui, v *gocui.View) error {
	c.control.Step()
	c.updateTrace()
	c.updateRegisters()
	c.traceBuffer = nil
	return nil
}

func (c *cui) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("trace", 0, 0, (maxX/3)*2, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = true
	}
	if _, err := g.SetView("registers", (maxX/3)*2+1, 0, maxX-1, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	if v, err := g.SetView("instructions", 0, maxY-5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Press 'p' to pause/resume execution.")
		fmt.Fprintln(v, "Press 's' to step forward.")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
