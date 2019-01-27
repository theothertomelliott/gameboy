package terminal

import (
	"fmt"
	"sort"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/theothertomelliott/gameboy"
)

type StackTable struct {
	*tview.Table

	gb *gameboy.DMG

	stackEvents map[uint16]gameboy.StackEvent

	pushCommand   map[uint16]uint16
	pushOperation map[uint16]int64
}

func NewStackTable(gb *gameboy.DMG) *StackTable {
	st := &StackTable{
		gb:            gb,
		stackEvents:   make(map[uint16]gameboy.StackEvent),
		pushCommand:   make(map[uint16]uint16),
		pushOperation: make(map[uint16]int64),
	}
	st.Table = tview.NewTable().
		SetBorders(false).
		SetFixed(1, 1)
	st.Table.SetBorder(true).SetTitle("Stack")
	return st
}

func (s *StackTable) Update() {
	cpu := s.gb.CPU()
	//mmu := s.gb.MMU()

	spPos := cpu.SP.Read16()

	s.Clear()

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

	var positions []uint16
	for pos := range s.stackEvents {
		positions = append(positions, pos)
	}
	sort.Slice(positions, func(i int, j int) bool {
		return positions[i] > positions[j]
	})

	var row int
	var wroteSP bool
	for _, pos := range positions {
		if pos < spPos && !wroteSP {
			s.SetCell(
				row+1,
				0,
				tview.NewTableCell("SP").
					SetTextColor(tcell.ColorYellow).
					SetAlign(tview.AlignCenter),
			)
			row++
			wroteSP = true
			continue
		}

		ev := s.stackEvents[pos]
		s.SetCell(
			row+1,
			0,
			tview.NewTableCell(fmt.Sprintf("0x%04X", ev.Pos)).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter),
		)
		s.SetCell(
			row+1,
			1,
			tview.NewTableCell(fmt.Sprintf("0x%04X", ev.ValueIn)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter),
		)

		pushText := fmt.Sprintf("0x%04X", s.pushCommand[pos])
		opText := fmt.Sprintf("0x%04X", s.pushOperation[pos])
		s.SetCell(
			row+1,
			2,
			tview.NewTableCell(pushText).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter),
		)
		s.SetCell(
			row+1,
			3,
			tview.NewTableCell(opText).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter),
		)
		row++
	}
	s.ScrollToEnd()
}

func (s *StackTable) Trace(ev gameboy.TraceMessage) {
	for _, stackOp := range ev.Stack {
		if stackOp.ValueIn != 0 {
			s.stackEvents[stackOp.Pos] = stackOp
			s.pushCommand[stackOp.Pos] = ev.CPU.PC
			s.pushOperation[stackOp.Pos] = ev.Count
		}
	}
}
