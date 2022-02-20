package fyneui

import (
	"time"

	"github.com/theothertomelliott/gameboy/ppu"
)

func (u *UI) Run() {
	go func() {
		p := u.gb.PPU()

		for {
			if ppu.GetLCDControl(p.MMU).LCDOperation() {
				u.drawGraphics(p.RenderScreen())
			}
			u.updateDebugInfo()

			time.Sleep(time.Second / 60)
		}
	}()

	u.win.ShowAndRun()

}
