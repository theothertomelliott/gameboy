package terminal

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/theothertomelliott/gameboy"
)

type StackTable struct {
	*tview.Table

	gb *gameboy.DMG

	pushCommand map[uint16]uint16
}

func NewStackTable(gb *gameboy.DMG) *StackTable {
	st := &StackTable{
		gb:          gb,
		pushCommand: make(map[uint16]uint16),
	}
	st.Table = tview.NewTable().
		SetBorders(false).
		SetFixed(1, 1)
	st.Table.SetBorder(true).SetTitle("Stack")
	return st
}

func (s *StackTable) Update() {
	cpu := s.gb.CPU()
	mmu := s.gb.MMU()

	psPos := cpu.SP.Read16()

	s.SetCell(
		0,
		0,
		tview.NewTableCell("SP").
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter),
	)
	s.SetCell(
		0,
		1,
		tview.NewTableCell("Value").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter),
	)
	s.SetCell(
		0,
		2,
		tview.NewTableCell("Pushed At").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter),
	)

	var row int
	for pos := uint16(0xFFFE); pos >= psPos; pos-- {
		value := mmu.Read8(pos)
		s.SetCell(
			row+1,
			0,
			tview.NewTableCell(fmt.Sprintf("0x%04X", pos)).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter),
		)
		s.SetCell(
			row+1,
			1,
			tview.NewTableCell(fmt.Sprintf("0x%02X", value)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter),
		)
		pushText := ""
		if row%2 == 0 {
			pushText = fmt.Sprintf("0x%04X", s.pushCommand[pos])
		}
		s.SetCell(
			row+1,
			2,
			tview.NewTableCell(pushText).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter),
		)
		row++
	}
}

func (s *StackTable) Trace(ev gameboy.TraceMessage) {
	for _, stackOp := range ev.Stack {
		if stackOp.ValueIn != 0 {
			s.pushCommand[stackOp.Pos] = ev.CPU.PC
		}
	}
}
