package fyneui

import (
	"time"
)

func (u *UI) Run() {
	go func() {
		for {
			u.win.Content().Refresh()
			u.updateDebugInfo()

			time.Sleep(time.Second / 60)
		}
	}()

	u.win.ShowAndRun()

}
