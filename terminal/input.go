package terminal

import (
	"github.com/gdamore/tcell"
)

func (t *TerminalUI) setupInput() {
	t.app = t.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'p':
			t.gb.TogglePaused()
			t.traceView.SetSelectable(t.gb.IsPaused(), false)
		case 's':
			t.gb.Step()
		case 'o':
			if t.gb.IsPaused() {
				t.gb.TogglePaused()
			}
			t.stepOut = true
		case 't':
			t.app.SetFocus(t.traceView)
		case 'd':
			t.app.SetFocus(t.debuggerView)
		case 'm':
			t.app.SetFocus(t.memoryView)
		}
		return event
	})
}
