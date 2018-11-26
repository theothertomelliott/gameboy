package terminal

import (
	"fmt"
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

	diagnostic := tview.NewFlex().
		AddItem(t.decompileView, 0, 1, true).
		AddItem(t.memoryView, 0, 3, true)

	center := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(diagnostic, 0, 4, true).
		AddItem(t.testOutputView, 5, 1, false)

	root := tview.NewFlex().
		AddItem(t.traceView, 0, 1, true).
		AddItem(center, 0, 4, true).
		AddItem(t.registerView, 20, 1, false)

	t.app.SetRoot(root, true)
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
}

func (t *TerminalUI) updateTestOutput() {
	if err := t.gb.Err(); err != nil {
		t.testOutputView.SetText(fmt.Sprintf("Error: %v", err))
		return
	}
	t.testOutputView.SetText(t.gb.MMU().TestOutput())
}

func (t *TerminalUI) Stop() {
	t.updateTicker.Stop()
	t.app.Stop()
}

func (t *TerminalUI) trace(ev gameboy.TraceMessage) {
	row := int(ev.Count)
	t.traceView.SetCell(
		row,
		0,
		tview.NewTableCell(fmt.Sprintf("0x%X", ev.Event.PC)).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignRight),
	)

	t.traceView.SetCell(
		row,
		1,
		tview.NewTableCell(ev.Event.Description).
			SetTextColor(tcell.ColorWhite))
	t.traceView.Select(row, 0)
	t.traceView.ScrollToEnd()

	t.decompileMtx.Lock()
	t.decompilation[ev.Event.PC] = ev.Event.Description
	t.decompileMtx.Unlock()

	if t.gb.IsPaused() {
		t.update()
	}
}
