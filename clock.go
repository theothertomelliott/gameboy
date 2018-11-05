package gameboy

func Step(c *CPU, ppu *PPU) {
	if c.isHalted {
		// If interrupts are disabled (DI) then
		// halt doesn't suspend operation but it does cause
		// the program counter to stop counting for one
		// instruction
		if c.MMU.Read8(IE) == 0x0 {
			c.PC.Inc(1)
			c.cycles = 1
		}
		return
	}

	t := c.Step()
	ppu.Step(t)
}
