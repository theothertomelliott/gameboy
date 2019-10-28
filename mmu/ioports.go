package mmu

const (
	// Interrupt flag locations
	// See: http://bgb.bircd.org/pandocs.htm#interrupts

	IE = 0xFFFF // Interrupt Enable (R/W)
	IF = 0xFF0F // Interrupt Flag (R/W)

	// See: https://fms.komkon.org/GameBoy/Tech/Software.html
	JOYPAD  = 0xFF00 // [RW] Joypad port
	SIODATA = 0xFF01 // [RW] Serial I/O Data
	SIOCONT = 0xFF02 // [RW] Serial I/O Control

	DIVIDER = 0xFF04 // [RW] Divider [meaning unknown]

	TIMECNT = 0xFF05 // [RW] Timer Counter
	TIMEMOD = 0xFF06 // [RW] Timer Modulo
	TIMCONT = 0xFF07 // [RW] Timer Control

	SNDREG10 = 0xFF10 // [RW] Sweep [Sound Mode #1]
	SNDREG11 = 0xFF11 // [RW] Sound Length/Pattern Duty [Sound Mode #1]
	SNDREG12 = 0xFF12 // [RW] Control [Sound Mode #1]
	SNDREG13 = 0xFF13 // [W] Frequency Low [Sound Mode #1]

	SNDREG14 = 0xFF14 // [RW] Frequency High [Sound Mode #1]
	SNDREG21 = 0xFF16 // [RW] Sound Length/Pattern Duty [Sound Mode #2]
	SNDREG22 = 0xFF17 // [RW] Control [Sound Mode #2]
	SNDREG23 = 0xFF18 // [W] Frequency Low [Sound Mode #2]
	SNGREG24 = 0xFF19 // [RW] Frequency High [Sound Mode #2]
	SNDREG30 = 0xFF1A // [RW] Control [Sound Mode #3]
	SNDREG31 = 0xFF1B // [RW] Sound Length [Sound Mode #3]
	SNDREG32 = 0xFF1C // [RW] Output Level [Sound Mode #3]
	SNDREG33 = 0xFF1D // [W] Frequency Low [Sound Mode #3]
	SNDREG34 = 0xFF1E // [RW] Frequency High [Sound Mode #3]
	SNDREG41 = 0xFF20 // [RW] Sound Length/Pattern Duty [Sound Mode #4]
	SNDREG42 = 0xFF21 // [RW] Control [Sound Mode #4]
	SNDREG43 = 0xFF22 // [RW] Polynomial Counter [Sound Mode #4]
	SNDREG44 = 0xFF23 // SNDREG44 [RW] Frequency High [Sound Mode #4]
	SNDREG50 = 0xFF24 // [RW] Channel and Volume Control
	SNDREG51 = 0xFF25 // [RW] Sound Output Terminal Selector
	SNDREG52 = 0xFF26 // SNDREG52 [RW] Sound ON/OFF

	LCDCONT = 0xFF40 //  [RW] LCD Control
	LCDSTAT = 0xFF41 // [RW] LCD Status
	SCROLLY = 0xFF42 // [RW] Background Vertical Scrolling
	SCROLLX = 0xFF43 // [RW] Background Horizontal Scrolling
	CURLINE = 0xFF44 // [RW] Current Scanline
	CMPLINE = 0xFF45 // [RW] Scanline Comparison

	BGRDPAL = 0xFF47 // [W] Background Palette
	OBJ0PAL = 0xFF48 // [W] Sprite Palette #0
	OBJ1PAL = 0xFF49 // [W] Sprite Palette #1

	WNDPOSY = 0xFF4A // [RW] Window Y Position
	WNDPOSX = 0xFF4B // [RW] Window X Position

	DMACONT = 0xFF46 // [W] DMA Transfer Control
)
