package fyneui

import (
	"time"
)

func (u *UI) Run() {
	go func() {
		for {
			u.screenContent.Set(u.gb)
			u.registers.updateDebugInfo()

			time.Sleep(time.Second / 60)
		}
	}()

	u.win.ShowAndRun()

}
