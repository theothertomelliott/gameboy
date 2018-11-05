package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/theothertomelliott/gameboy"
)

type cui struct {
	gui     *gocui.Gui
	cpu     *gameboy.CPU
	tracer  *gameboy.Tracer
	control *control

	opCount      uint64
	opsPerSecond uint64
}

func (c *cui) Close() {
	c.gui.Close()
}

func setupCUI(cpu *gameboy.CPU, control *control, tracer *gameboy.Tracer) (*cui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}

	cui := &cui{
		gui:     g,
		cpu:     cpu,
		tracer:  tracer,
		control: control,
	}

	g.SetManagerFunc(cui.layout)
	return cui, nil
}

func startCUI(cui *cui) {
	go func() {
		for range time.Tick(time.Second) {
			cui.opsPerSecond = cui.opCount
			cui.opCount = 0
		}
	}()
	go func() {
		for trace := range cui.tracer.Event {
			cui.opCount++
			trace := trace
			cui.gui.Update(func(g *gocui.Gui) error {
				v, err := g.View("trace")
				if err != nil {
					return nil
				}
				fmt.Fprintf(v, "0x%X: %v\n", trace.Count, trace.Event.Operation.Description)
				return nil
			})
			cui.gui.Update(func(g *gocui.Gui) error {
				v, err := g.View("registers")
				if err != nil {
					return nil
				}
				v.Clear()

				fmt.Fprintf(v, "A: 0x%X, F: 0x%X\n", cui.cpu.A.Read8(), cui.cpu.F.Read8())
				fmt.Fprintf(v, "B: 0x%X, C: 0x%X\n", cui.cpu.B.Read8(), cui.cpu.C.Read8())
				fmt.Fprintf(v, "D: 0x%X, E: 0x%X\n", cui.cpu.D.Read8(), cui.cpu.E.Read8())
				fmt.Fprintf(v, "H: 0x%X, L: 0x%X\n", cui.cpu.H.Read8(), cui.cpu.L.Read8())

				fmt.Fprintln(v)

				fmt.Fprintf(v, "PC: 0x%X\n", cui.cpu.PC.Read16())
				fmt.Fprintf(v, "SP: 0x%X\n", cui.cpu.SP.Read16())

				fmt.Fprintln(v)

				fmt.Fprintf(v, "Ops/Second: %d\n", cui.opsPerSecond)

				return nil
			})
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
	c.control.Paused = !c.control.Paused
	return nil
}

func (c *cui) step(g *gocui.Gui, v *gocui.View) error {
	c.control.Step()
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
