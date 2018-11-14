package terminal

import (
	"fmt"
	"sort"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (t *TerminalUI) setupDecompileView() {
	t.decompileView = tview.NewTable().
		SetBorders(false).
		SetSelectable(true, false)

	t.decompileView.
		SetInputCapture(t.pagingFunc(t.decompileView)).
		SetBorder(true).
		SetTitle("Decompile")
}

func (t *TerminalUI) updateDecompilation() {
	t.decompileMtx.Lock()
	defer t.decompileMtx.Unlock()

	var pcs []uint16
	for pc := range t.decompilation {
		pcs = append(pcs, pc)
	}
	sort.Slice(pcs, func(i, j int) bool {
		return pcs[i] < pcs[j]
	})
	for index, pc := range pcs {
		t.decompileView.SetCell(
			index,
			0,
			tview.NewTableCell(fmt.Sprintf("0x%X", pc)).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignRight),
		)
		t.decompileView.SetCell(
			index,
			1,
			tview.NewTableCell(t.decompilation[pc]).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft),
		)
	}

	t.decompileView.ScrollToBeginning()
}
