package terminal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (t *TerminalUI) setupMemoryView() {
	t.memoryView = NewPagingTable(t).
		SetGotoFunc(func(textValue string) (int, int) {
			if strings.HasPrefix(textValue, "0x") {
				textValue = strings.Replace(textValue, "0x", "", 1)
			}
			val, err := strconv.ParseInt(textValue, 16, 64)
			if err != nil {
				// TODO: Handle error
				return 0, 0
			}
			row, col := int(val&0xFFF0)>>4, int(val&0xF)
			return row + 1, col + 1
		}).
		SetBorders(false).
		SetFixed(1, 1).
		SetSelectable(true, true)

	t.memoryView.
		SetBorder(true).
		SetTitle("Memory")

	for pos := 0; pos <= 0xF; pos++ {
		t.memoryView.SetCell(
			0,
			pos+1,
			tview.NewTableCell(fmt.Sprintf("0x%02X", pos)).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter),
		)
	}

	for pos := 0; pos <= 0xFFFF; pos++ {
		row := (pos / 16) + 1
		if pos%16 == 0 {
			t.memoryView.SetCell(
				row,
				0,
				tview.NewTableCell(fmt.Sprintf("0x%04X", pos)).
					SetTextColor(tcell.ColorYellow).
					SetAlign(tview.AlignCenter),
			)
		}

		t.memoryView.SetCell(
			row,
			(pos%16)+1,
			tview.NewTableCell(fmt.Sprintf("0x%02X", 0)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter),
		)
	}
}

func (t *TerminalUI) updateMemory() {
	mmu := t.gb.MMU()
	for pos := 0; pos <= 0xFFFF; pos++ {
		cell := t.memoryView.GetCell((pos/16)+1, (pos%16)+1)
		value := mmu.Read8(uint16(pos))
		cell.SetText(
			fmt.Sprintf("0x%02X", value),
		)
	}
}
