package terminal

import (
	"fmt"
	"strings"
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
	debuggerView   *tview.Table
	stackView      *StackTable

	rootView tview.Primitive

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
	latestPC      uint16
	decompileMtx  sync.Mutex

	stepOut bool
}

func NewTerminalUI(gb *gameboy.DMG) *TerminalUI {
	app := tview.NewApplication()
	t := &TerminalUI{
		gb: gb,

		app: app,

		decompilation: make(map[uint16]string),
	}

	t.stackView = NewStackTable(gb)
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

	output := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.registerView, 0, 1, true).
		AddItem(t.stackView, 0, 2, true).
		AddItem(t.testOutputView, 0, 1, false)

	t.rootView = tview.NewFlex().
		AddItem(t.traceView, 0, 3, true).
		AddItem(t.debuggerView, 0, 1, true).
		AddItem(output, 0, 1, true)

	t.GoToRoot()
}

func (t *TerminalUI) GoToRoot() {
	t.app.SetRoot(t.rootView, true)
}

func (t *TerminalUI) setupTestOutputView() {
	t.testOutputView = tview.NewTextView()
	t.testOutputView.SetBorder(true).SetTitle("Test ROM Output")
}

func (t *TerminalUI) Run() error {
	t.updateTicker = time.NewTicker(time.Second / 10)
	go func() {
		for range t.updateTicker.C {
			if !t.gb.IsPaused() {
				t.update()
			}
			t.app.Draw()
		}
	}()
	return t.app.Run()
}

func (t *TerminalUI) update() {
	t.updateTestOutput()
	t.updateRegisters()
	t.updateMemory()
	t.updateDecompilation()
	t.stackView.Update()
}

func (t *TerminalUI) updateTestOutput() {
	if err := t.gb.Err(); err != nil {
		t.testOutputView.SetText(fmt.Sprintf("Error: %v", err))
		return
	}
	testOutput := t.gb.MMU().TestOutput()
	t.testOutputView.SetText(testOutput)
	t.testOutputView.ScrollToEnd()
}

func (t *TerminalUI) Stop() {
	t.updateTicker.Stop()
	t.app.Stop()
}

func (t *TerminalUI) trace(ev gameboy.TraceMessage) {
	if ev.CPU == nil {
		return
	}

	t.stackView.Trace(ev)

	if ev.MMU != nil && ev.MMU.Pos == 0xFF02 && ev.MMU.ValuesAfter[0] == 0x81 {
		t.gb.SetPaused(true)
	}

	row := int(ev.Count)
	t.traceView.SetCell(
		row,
		0,
		tview.NewTableCell(fmt.Sprintf("0x%X", row)).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignRight),
	)

	t.traceView.SetCell(
		row,
		1,
		tview.NewTableCell(fmt.Sprintf("0x%X", ev.CPU.PC)).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignRight),
	)

	t.traceView.SetCell(
		row,
		2,
		tview.NewTableCell(ev.CPU.Description).
			SetTextColor(tcell.ColorWhite),
	)

	var col = 3
	for _, log := range ev.Log {
		t.traceView.SetCell(
			row,
			col,
			tview.NewTableCell(log.Text).
				SetTextColor(tcell.ColorWhite),
		)
		col++
	}
	for _, stack := range ev.Stack {
		t.traceView.SetCell(
			row,
			col,
			tview.NewTableCell(fmt.Sprint(stack)).
				SetTextColor(tcell.ColorWhite),
		)
		col++
	}
	if ev.MMU != nil {
		t.traceView.SetCell(
			row,
			col,
			tview.NewTableCell(fmt.Sprint(ev.MMU)).
				SetTextColor(tcell.ColorWhite),
		)
		col++
	}
	for _, reg := range ev.Registers {
		t.traceView.SetCell(
			row,
			col,
			tview.NewTableCell(fmt.Sprint(reg)).
				SetTextColor(tcell.ColorWhite),
		)
		col++
	}

	t.traceView.Select(row, 0)
	t.traceView.ScrollToEnd()

	t.decompileMtx.Lock()
	t.decompilation[ev.CPU.PC] = ev.CPU.Description
	t.decompileMtx.Unlock()

	t.latestPC = ev.CPU.PC

	if t.gb.IsPaused() {
		t.update()
	}

	if t.stepOut {
		if strings.HasPrefix(ev.CPU.Description, "RET") ||
			strings.HasPrefix(ev.CPU.Description, "RST") {
			t.gb.SetPaused(true)
			t.stepOut = false
		}
	}
}
