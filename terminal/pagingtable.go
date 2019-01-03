package terminal

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type PagingTable struct {
	*tview.Table

	ui *TerminalUI

	matchCol *int
}

func NewPagingTable(ui *TerminalUI) *PagingTable {
	pt := &PagingTable{
		Table: tview.NewTable(),
		ui:    ui,
	}
	pt.SetInputCapture(pt.pagingFunc)
	return pt
}

func (pt *PagingTable) SetMatchColumn(col int) *PagingTable {
	pt.matchCol = &col
	return pt
}

func (pt *PagingTable) pagingFunc(event *tcell.EventKey) *tcell.EventKey {
	row, col := pt.GetSelection()
	pageSize := 50
	switch event.Rune() {
	case 'b':
		pt.Select(0, 0)
		pt.ScrollToBeginning()
	case 'e':
		pt.Select(pt.GetRowCount(), 0)
		pt.ScrollToEnd()
	case '[':
		if row-pageSize > 0 {
			pt.Select(row-pageSize, col)
		} else {
			pt.Select(0, col)
		}
	case ']':
		if row+pageSize < pt.GetRowCount() {
			pt.Select(row+pageSize, col)
		} else {
			pt.Select(pt.GetRowCount()-1, col)
		}
	case 'g':
		pt.showGotoModal()
	}
	return event
}

func (pt *PagingTable) showGotoModal() {
	modal := NewTextModal().
		SetText("Go to position").
		SetInput("Position", "0x00", "OK").
		SetDoneFunc(func(textValue string) {
			pt.ui.GoToRoot()
			pt.ui.app.SetFocus(pt)

			if pt.matchCol == nil {
				if strings.HasPrefix(textValue, "0x") {
					textValue = strings.Replace(textValue, "0x", "", 1)
				}
				val, err := strconv.ParseInt(textValue, 16, 64)
				if err != nil {
					// TODO: Handle error
					return
				}
				pt.Select(int(val), 0)
				return
			}

			// Find a column matching the input
			for row := 0; row < pt.GetRowCount(); row++ {
				cell := pt.GetCell(row, *pt.matchCol)
				if cell.Text == textValue {
					pt.Select(row, *pt.matchCol)
				}
			}
		})

	pt.ui.app.SetRoot(modal, true)
}
