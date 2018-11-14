package terminal

import (
	"github.com/rivo/tview"
)

func (t *TerminalUI) setupTraceView() {
	t.traceView = tview.NewTable().
		SetBorders(false)
	t.traceView.SetInputCapture(t.pagingFunc(t.traceView)).SetBorder(true).SetTitle("Trace")
}
