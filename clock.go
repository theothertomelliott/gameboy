package gameboy

func Step(c *CPU, ppu *PPU) {
	t := c.Step()
	ppu.Step(t)
}
