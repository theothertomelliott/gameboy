package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/theothertomelliott/gameboy"
)

var romDirs = []string{
	"cpu_instrs/individual",
	"instr_timing",
	"interrupt_time",
	"mem_timing",
	"mem_timing-2",
}

func main() {
	basePath := os.Args[1]

	var (
		totalRoms  int
		failedRoms int
	)
	for _, dir := range romDirs {
		romDirPath := path.Join(basePath, dir)
		fmt.Println(romDirPath)
		files, err := ioutil.ReadDir(romDirPath)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".gb") {
				totalRoms++
				pathToRom := path.Join(romDirPath, file.Name())
				err := runTestRom(pathToRom)
				if err != nil {
					log.Printf("%v", err)
					failedRoms++
					continue
				}
			}
		}

	}
	fmt.Printf("Ran %v ROMs, %v passed, %v failed.\n", totalRoms, (totalRoms - failedRoms), failedRoms)
}

func runTestRom(pathToRom string) error {
	fmt.Println(pathToRom)

	romData, err := ioutil.ReadFile(pathToRom)
	if err != nil {
		return err
	}

	gb := gameboy.NewDMGWithNoRateLimit()

	gb.MMU().LoadCartridge(romData)
	gb.CPU().Init()

	defer fmt.Println()
	var cycles int

	timeout := time.After(30 * time.Second)
	for {
		err := gb.Step()
		if err != nil {
			return err
		}
		cycles++
		if cycles%1000 == 0 {
			to := strings.ToLower(gb.MMU().TestOutput())
			if strings.Contains(to, "f") {
				return fmt.Errorf("%v", to)
			}
			if strings.Contains(to, "passed") {
				return nil
			}
		}

		select {
		case <-timeout:
			return fmt.Errorf("Timeout:\n%v", gb.MMU().TestOutput())
		default:
		}
	}
}
