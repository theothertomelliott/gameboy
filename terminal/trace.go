package terminal

func (t *TerminalUI) setupTraceView() {
	t.traceView = NewPagingTable(t).
		SetBorders(false).SetSelectable(true, false)
	t.traceView.SetBorder(true).SetTitle("Trace")
}
