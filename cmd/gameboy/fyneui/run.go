package fyneui

import (
	"time"
)

func (u *UI) Run() {
	go func() {
		for {
			u.data.Update()

			time.Sleep(time.Second / 60)
		}
	}()

	u.win.ShowAndRun()

}
