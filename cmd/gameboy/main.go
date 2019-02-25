package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
	"github.com/theothertomelliott/gameboy"
	"github.com/theothertomelliott/gameboy/cmd/gameboy/httpui"
)

func main() {
	var (
		breakpoints breakPoints
	)
	flag.Var(&breakpoints, "breakpoint", "A comma-separated list of breakpoints, as 16-bit hex values.")
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s cartridge [bootROM]:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	gb := gameboy.NewDMG()

	cartridgeFile := flag.Arg(0)
	if cartridgeFile == "" {
		fmt.Fprintf(os.Stderr, "A cartridge ROM file is required\n\n")
		flag.Usage()
		return
	}
	bootROMFile := flag.Arg(1)

	data, err := ioutil.ReadFile(cartridgeFile)
	if err != nil {
		log.Fatal(err)
	}
	gb.MMU().LoadCartridge(data)

	if bootROMFile != "" {
		data, err = ioutil.ReadFile(bootROMFile)
		if err != nil {
			log.Fatal(err)
		}
		gb.MMU().LoadROM(data)
	} else {
		gb.CPU().Init()
	}

	for _, bp := range breakpoints {
		gb.Breakpoints[bp] = struct{}{}
	}

	uiserver := httpui.NewServer(gb)

	gb.Tracer().Logger = func(ev gameboy.TraceMessage) {
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
