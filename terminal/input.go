package terminal

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (t *TerminalUI) setupInput() {
	t.app = t.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'p':
			t.gb.TogglePaused()
			t.traceView.SetSelectable(t.gb.IsPaused(), false)
		case 's':
			t.gb.Step()
		case 't':
			t.app.SetFocus(t.traceView)
		case 'd':
			t.app.SetFocus(t.decompileView)
		case 'm':
			t.app.SetFocus(t.memoryView)
		}
		return event
	})
}

func (t *TerminalUI) pagingFunc(table *tview.Table) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		row, col := table.GetSelection()
		pageSize := 50
		switch event.Rune() {
		case '[':
			if row-pageSize > 0 {
				table.Select(row-pageSize, col)
			} else {
				table.Select(0, col)
			}
		case ']':
			if row+pageSize < table.GetRowCount() {
				table.Select(row+pageSize, col)
			} else {
				table.Select(table.GetRowCount()-1, col)
			}
		}
		return event
	}
}
