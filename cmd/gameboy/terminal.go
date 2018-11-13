package main

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/theothertomelliott/gameboy"
)

type TerminalUI struct {
	gb *gameboy.DMG

	updateTicker *time.Ticker

	app *tview.Application

	traceView      *tview.Table
	registerView   *tview.Table
	memoryView     *tview.Table
	testOutputView *tview.TextView
	decompileView  *tview.Table

	rCellA  *tview.TableCell
	rCellF  *tview.TableCell
	rCellB  *tview.TableCell
	rCellC  *tview.TableCell
	rCellD  *tview.TableCell
	rCellE  *tview.TableCell
	rCellH  *tview.TableCell
	rCellL  *tview.TableCell
	rCellSP *tview.TableCell
	rCellPC *tview.TableCell

	decompilation map[uint16]string
	decompileMtx  sync.Mutex
}

func NewTerminalUI(gb *gameboy.DMG) *TerminalUI {
	app := tview.NewApplication()
	t := &TerminalUI{
		gb: gb,

		app: app,

		decompilation: make(map[uint16]string),
	}

	t.setupTraceView()
	t.setupDecompileView()
	t.setupTestOutputView()
	t.setupMemoryView()
	t.setupRegisterView()
	t.setupRoot()

	t.setupInput()

	gb.Tracer().Logger = t.trace

	return t
}

func (t *TerminalUI) setupRoot() {
	root := tview.NewFlex().
		AddItem(t.traceView, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(t.decompileView, 0, 4, true).
			AddItem(t.testOutputView, 5, 1, false), 0, 2, true).
		AddItem(t.registerView, 20, 1, false)

	t.app.SetRoot(root, true)
}

func (t *TerminalUI) setupTraceView() {
	t.traceView = tview.NewTable().
		SetBorders(false)
	t.traceView.SetBorder(true).SetTitle("Trace")
}

func (t *TerminalUI) setupDecompileView() {
	t.decompileView = tview.NewTable().
		SetBorders(false).SetSelectable(true, false)
	t.decompileView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		row, col := t.decompileView.GetSelection()
		pageSize := 50
		switch event.Rune() {
		case '[':
			if row-pageSize > 0 {
				t.decompileView.Select(row-pageSize, col)
			} else {
				t.decompileView.Select(0, col)
			}
		case ']':
			if row+pageSize < t.decompileView.GetRowCount() {
				t.decompileView.Select(row+pageSize, col)
			} else {
				t.decompileView.Select(t.decompileView.GetRowCount()-1, col)
			}
		}
		return event
	}).SetBorder(true).SetTitle("Decompile")
}

func (t *TerminalUI) setupTestOutputView() {
	t.testOutputView = tview.NewTextView()
	t.testOutputView.SetBorder(true).SetTitle("Test ROM Output")
}

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

func (t *TerminalUI) setupInput() {
	t.app = t.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'p':
			t.gb.TogglePaused()
		case 's':
			t.gb.Step()
		}
		return event
	})
}

func (t *TerminalUI) setupRegisterView() {
	t.registerView = tview.NewTable().
		SetBorders(false)
	t.registerView.SetBorder(true).SetTitle("Registers")

	t.rCellA = tview.NewTableCell("0x00")
	t.rCellF = tview.NewTableCell("0x00")
	t.rCellB = tview.NewTableCell("0x00")
	t.rCellC = tview.NewTableCell("0x00")
	t.rCellD = tview.NewTableCell("0x00")
	t.rCellE = tview.NewTableCell("0x00")
	t.rCellH = tview.NewTableCell("0x00")
	t.rCellL = tview.NewTableCell("0x00")
	t.rCellSP = tview.NewTableCell("0x00")
	t.rCellPC = tview.NewTableCell("0x00")

	r := 0
	t.registerView.SetCell(r, 0, tview.NewTableCell("A"))
	t.registerView.SetCell(r, 1, t.rCellA)
	t.registerView.SetCell(r, 2, tview.NewTableCell("F"))
	t.registerView.SetCell(r, 3, t.rCellF)

	r++
	t.registerView.SetCell(r, 0, tview.NewTableCell("B"))
	t.registerView.SetCell(r, 1, t.rCellB)
	t.registerView.SetCell(r, 2, tview.NewTableCell("C"))
	t.registerView.SetCell(r, 3, t.rCellC)

	r++
	t.registerView.SetCell(r, 0, tview.NewTableCell("D"))
	t.registerView.SetCell(r, 1, t.rCellD)
	t.registerView.SetCell(r, 2, tview.NewTableCell("E"))
	t.registerView.SetCell(r, 3, t.rCellE)

	r++
	t.registerView.SetCell(r, 0, tview.NewTableCell("H"))
	t.registerView.SetCell(r, 1, t.rCellH)
	t.registerView.SetCell(r, 2, tview.NewTableCell("L"))
	t.registerView.SetCell(r, 3, t.rCellL)

	r++
	t.registerView.SetCell(r, 0, tview.NewTableCell("PC"))
	t.registerView.SetCell(r, 1, t.rCellPC)
	r++
	t.registerView.SetCell(r, 0, tview.NewTableCell("SP"))
	t.registerView.SetCell(r, 1, t.rCellSP)
}

func (t *TerminalUI) Run() error {
	t.updateTicker = time.NewTicker(time.Second / 30)
	go func() {
		for range t.updateTicker.C {
			t.testOutputView.SetText(t.gb.MMU().TestOutput())
			t.updateRegisters()
			t.updateMemory()
			t.updateDecompilation()
			t.app.Draw()
		}
	}()
	return t.app.Run()
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

func (t *TerminalUI) updateRegisters() {
	cpu := t.gb.CPU()
	t.rCellA.SetText(fmt.Sprintf("0x%02X", cpu.A.Read8()))
	t.rCellF.SetText(fmt.Sprintf("0x%02X", cpu.F.Read8()))
	t.rCellB.SetText(fmt.Sprintf("0x%02X", cpu.B.Read8()))
	t.rCellC.SetText(fmt.Sprintf("0x%02X", cpu.C.Read8()))
	t.rCellD.SetText(fmt.Sprintf("0x%02X", cpu.D.Read8()))
	t.rCellE.SetText(fmt.Sprintf("0x%02X", cpu.E.Read8()))
	t.rCellH.SetText(fmt.Sprintf("0x%02X", cpu.H.Read8()))
	t.rCellL.SetText(fmt.Sprintf("0x%02X", cpu.L.Read8()))

	t.rCellPC.SetText(fmt.Sprintf("0x%04X", cpu.PC.Read16()))
	t.rCellSP.SetText(fmt.Sprintf("0x%04X", cpu.SP.Read16()))
}

func (t *TerminalUI) Stop() {
	t.updateTicker.Stop()
	t.app.Stop()
}

func (t *TerminalUI) trace(ev gameboy.TraceMessage) {
	t.traceView.SetCell(
		int(ev.Count),
		0,
		tview.NewTableCell(fmt.Sprintf("0x%X", ev.Event.PC)).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignRight),
	)

	t.traceView.SetCell(
		int(ev.Count),
		1,
		tview.NewTableCell(ev.Event.Description).
			SetTextColor(tcell.ColorWhite))

	t.traceView.ScrollToEnd()

	t.decompileMtx.Lock()
	defer t.decompileMtx.Unlock()
	t.decompilation[ev.Event.PC] = ev.Event.Description
}
