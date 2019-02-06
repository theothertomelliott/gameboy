package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/cmd/gameboy/httpui"
	"github.com/theothertomelliott/gameboy/terminal"
)

func main() {
	var (
		trace       = flag.Bool("trace", false, "If set, a trace UI will be shown in the terminal.")
		breakpoints breakPoints
	)
	flag.Var(&breakpoints, "breakpoint", "A comma-separated list of breakpoints, as 16-bit hex values.")
	flag.Parse()

	gb := gameboy.NewDMG()

	data, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	gb.MMU().LoadCartridge(data)

	if len(flag.Args()) > 1 {
		data, err = ioutil.ReadFile(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		gb.MMU().LoadROM(data)
	} else {
		gb.CPU().Init()
	}

	for _, bp := range breakpoints {
		gb.Breakpoints[bp] = struct{}{}
	}

	var term *terminal.TerminalUI
	if *trace {
		term = terminal.NewTerminalUI(gb)
		defer term.Stop()

		go func() {
			err := term.Run()
			if err != nil {
				panic(err)
			}
		}()
	}

	uiserver := httpui.NewServer(gb)

	gb.Tracer().Logger = func(ev gameboy.TraceMessage) {
		if term != nil {
			term.Trace(ev)
		}
		uiserver.Trace(ev)
	}

	go uiserver.ListenAndServe(8080)

	gb.Start()
	defer gb.Stop()

	pixelgl.Run(func() {
		run(gb)
	})
}

var _ flag.Value = &breakPoints{}

type breakPoints []uint16

func (b *breakPoints) String() string {
	if b == nil {
		return ""
	}
	return fmt.Sprint([]uint16(*b))
}

func (b *breakPoints) Set(value string) error {
	bps := strings.Split(value, ",")
	out := make(breakPoints, 0, len(bps))
	for _, bp := range bps {
		if strings.HasPrefix(bp, "0x") {
			bp = strings.Replace(bp, "0x", "", 1)
		}
		val, err := strconv.ParseInt(bp, 16, 64)
		if err != nil {
			return errors.WithMessage(err, "parsing breakpoint")
		}
		out = append(out, uint16(val))
	}
	*b = out
	return nil
}
