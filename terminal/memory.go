package terminal

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (t *TerminalUI) setupMemoryView() {
	t.memoryView = tview.NewTable().
		SetBorders(true).SetFixed(1, 1)
	t.memoryView.SetBorder(true).
		SetTitle("Memory")

	for low := 0; low < 0xFF; low++ {
		t.memoryView.SetCell(
			0,
			low,
			tview.NewTableCell(fmt.Sprintf("0x%02X", low)).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter),
		)
	}

	for high := 0; high < 0xFF; high++ {
		t.memoryView.SetCell(
			high,
			0,
			tview.NewTableCell(fmt.Sprintf("0x%02X", high)).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter),
		)
	}

	for low := 0; low < 0xFF; low++ {
		for high := 0; high < 0xFF; high++ {
			t.memoryView.SetCell(
				high+1,
				low+1,
				tview.NewTableCell(fmt.Sprintf("0x%02X", 0)).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter),
			)
		}
	}
}

func (t *TerminalUI) updateMemory() {
	mmu := t.gb.MMU()
	for low := 0; low < 0xFF; low++ {
		for high := 0; high < 0xFF; high++ {
			cell := t.memoryView.GetCell(high+1, low+1)
			pos := uint16(low) | (uint16(high) << 8)
			text := fmt.Sprintf("0x%02X", mmu.Read8(pos))
			if cell.Text != text {
				cell.SetText(text)
				t.memoryView.Select(high+1, low+1)
				t.memoryView.SetOffset(high, low)
			}
		}
	}
}
