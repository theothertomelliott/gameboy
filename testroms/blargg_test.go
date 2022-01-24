package testroms

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/theothertomelliott/gameboy"
)

func TestBlarggCPU(t *testing.T) {
	var roms = map[string]string{
		"01-special":            "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/01-special.gb",
		"02-interrupts":         "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/02-interrupts.gb",
		"03-op sp,hl":           "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/03-op%20sp%2Chl.gb",
		"04-op r,imm":           "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/04-op%20r%2Cimm.gb",
		"05-op rp":              "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/05-op%20rp.gb",
		"06-ld r,r.gb":          "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/06-ld%20r%2Cr.gb",
		"07-jr,jp,call,ret,rst": "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/07-jr%2Cjp%2Ccall%2Cret%2Crst.gb",
		"08-misc instrs":        "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/08-misc%20instrs.gb",
		"09-op r,r":             "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/09-op%20r%2Cr.gb",
		"10-bit ops":            "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/10-bit%20ops.gb",
		"11-op a,(hl)":          "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/individual/11-op%20a%2C(hl).gb",
		"combined":              "https://github.com/retrio/gb-test-roms/raw/master/cpu_instrs/cpu_instrs.gb",
	}
	runTestROMs(t, roms)
}

func TestBlarggInstructionTiming(t *testing.T) {
	runTestROM(t, "https://github.com/retrio/gb-test-roms/raw/master/instr_timing/instr_timing.gb")
}

func TestBlarggInterruptTime(t *testing.T) {
	runTestROM(t, "https://github.com/retrio/gb-test-roms/raw/master/interrupt_time/interrupt_time.gb")
}

func TestBlarggMemTiming1(t *testing.T) {
	var roms = map[string]string{
		"01-read_timing":   "https://github.com/retrio/gb-test-roms/raw/master/mem_timing/individual/01-read_timing.gb",
		"02-write_timing":  "https://github.com/retrio/gb-test-roms/raw/master/mem_timing/individual/02-write_timing.gb",
		"03-modify_timing": "https://github.com/retrio/gb-test-roms/raw/master/mem_timing/individual/03-modify_timing.gb",
		"combined":         "https://github.com/retrio/gb-test-roms/raw/master/mem_timing/mem_timing.gb",
	}
	runTestROMs(t, roms)
}

func TestBlarggMemTiming2(t *testing.T) {
	var roms = map[string]string{
		"01-read_timing":   "https://github.com/retrio/gb-test-roms/raw/master/mem_timing-2/rom_singles/01-read_timing.gb",
		"02-write_timing":  "https://github.com/retrio/gb-test-roms/raw/master/mem_timing-2/rom_singles/02-write_timing.gb",
		"03-modify_timing": "https://github.com/retrio/gb-test-roms/raw/master/mem_timing-2/rom_singles/03-modify_timing.gb",
		"combined":         "https://github.com/retrio/gb-test-roms/raw/master/mem_timing-2/mem_timing.gb",
	}
	runTestROMs(t, roms)
}

func TestBlarggOAMBug(t *testing.T) {
	var roms = map[string]string{
		"1-lcd_sync":        "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/1-lcd_sync.gb",
		"2-causes":          "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/2-causes.gb",
		"3-non_causes":      "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/3-non_causes.gb",
		"4-scanline_timing": "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/4-scanline_timing.gb",
		"5-timing_bug":      "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/5-timing_bug.gb",
		"6-timing_no_bug":   "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/6-timing_no_bug.gb",
		"7-timing_effect":   "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/7-timing_effect.gb",
		"8-instr_effect":    "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/rom_singles/8-instr_effect.gb",
		"combined":          "https://github.com/retrio/gb-test-roms/raw/master/oam_bug/oam_bug.gb",
	}
	runTestROMs(t, roms)
}

func TestBlarggDMGSound(t *testing.T) {
	var roms = map[string]string{
		"01-registers":             "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/01-registers.gb",
		"02-len ctr":               "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/02-len%20ctr.gb",
		"03-trigger":               "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/03-trigger.gb",
		"04-sweep":                 "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/04-sweep.gb",
		"05-sweep details":         "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/05-sweep%20details.gb",
		"06-overflow on trigger":   "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/06-overflow%20on%20trigger.gb",
		"07-len sweep period sync": "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/07-len%20sweep%20period%20sync.gb",
		"08-len ctr during power":  "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/08-len%20ctr%20during%20power.gb",
		"09-wave read while on":    "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/09-wave%20read%20while%20on.gb",
		"10-wave trigger while on": "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/10-wave%20trigger%20while%20on.gb",
		"11-regs after power":      "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/11-regs%20after%20power.gb",
		"12-wave write while on":   "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/rom_singles/12-wave%20write%20while%20on.gb",
		"combined":                 "https://github.com/retrio/gb-test-roms/raw/master/dmg_sound/dmg_sound.gb",
	}
	runTestROMs(t, roms)
}

func runTestROMs(t *testing.T, roms map[string]string) {
	// Ensure tests are run in alphabetical order
	var names []string
	for name := range roms {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			runTestROM(t, roms[name])
		})
	}
}

func runTestROM(t *testing.T, romFile string) {
	if testing.Short() {
		t.Skip()
	}

	resp, err := http.Get(romFile)
	if err != nil {
		t.Fatalf("Could not load ROM: %v", err)
	}
	romData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Could not read ROM data: %v", err)
	}

	err = runTestRom(t, romData)
	if err != nil {
		t.Error(err)
	}
}

func discardOuptut() func() {
	var stdout, stderr = os.Stdout, os.Stderr
	os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)
	os.Stderr = os.NewFile(uintptr(syscall.Stdin), os.DevNull)
	log.SetOutput(ioutil.Discard)
	return func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(stdout)
	}
}

func runTestRom(t *testing.T, romData []byte) error {
	screenshotFile := fmt.Sprintf("screenshots/%v.png", strings.ReplaceAll(t.Name(), "/", ""))
	_ = os.Remove(screenshotFile)

	// Don't log output from DMG
	cleanup := discardOuptut()
	defer cleanup()

	gb := gameboy.NewDMGWithNoRateLimit()

	err := gb.MMU().LoadCartridge(romData)
	if err != nil {
		return err
	}
	gb.CPU().Init()

	var cycles int

	timeout := time.After(10 * time.Second)
	for {
		err := gb.Step()
		if err != nil {
			return err
		}
		cycles++
		if cycles%1000 == 0 {
			fmt.Println(gb.MMU().TestOutput())
			to := strings.ToLower(gb.MMU().TestOutput())
			if strings.Contains(to, "failed") {
				return fmt.Errorf("%v", to)
			}
			if strings.Contains(to, "passed") {
				return nil
			}
		}

		select {
		case <-timeout:
			screen := gb.PPU().RenderScreen()
			f, err := os.Create(screenshotFile)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			// Encode to `PNG` with `DefaultCompression` level
			// then save to file
			err = png.Encode(f, screen)
			if err != nil {
				panic(err)
			}
			return fmt.Errorf("Timeout after %d cycles:\n%v", cycles, gb.MMU().TestOutput())
		default:
		}
	}
}
