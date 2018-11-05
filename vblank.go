package gameboy

import "log"

func (c *CPU) vblankInterrupt() {
	// Expect VSYNC interrupt enabled
	if bitValue(0, c.MMU.Read8(IE)) != 1 || bitValue(0, c.MMU.Read8(IF)) != 1 {
		return
	}

	log.Print("VBLANK Interrupt")

	// CALL VBLANK interrupt vector
	c.CALL(Direct16(IV_VSYNC))

	// Reset the VBlank bit in IF (the request)
	c.MMU.Write8(IF, c.MMU.Read8(IF)&(0xFF-1))
}

const (
	IV_VSYNC = 0x40
)
