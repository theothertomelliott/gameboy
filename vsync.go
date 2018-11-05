package gameboy

import "fmt"

var scanLineCycleCount int

func (c *CPU) fakeScanlines() {
	if scanLineCycleCount < 100 {
		scanLineCycleCount++
		return
	}
	scanLineCycleCount = 0
	curline := c.MMU.Read8(CURLINE)
	curline++
	if curline == 0x90 {
		c.vblankInterrupt()
	}
	if curline >= 153 {
		curline = 0
	}
	c.MMU.Write8(CURLINE, curline)
}

func (c *CPU) vblankInterrupt() {
	// Set the VBlank bit in IF (the request)
	c.MMU.Write8(IF, c.MMU.Read8(IF)|0x1)

	// Expect VSYNC interrupt enabled
	if bitValue(0, c.MMU.Read8(IE)) != 1 {
		return
	}

	fmt.Println("VBLANK Interrupt")

	// CALL VBLANK interrupt vector
	c.CALL(Direct16(IV_VSYNC))

	// Reset the VBlank bit in IF (the request)
	c.MMU.Write8(IF, c.MMU.Read8(IF)&(0xFF-1))
}

const (
	IV_VSYNC = 0x40
)
