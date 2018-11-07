package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/theothertomelliott/gameboy"
)

type TerminalUI struct {
	gb *gameboy.DMG

	updateTicker *time.Ticker

	app          *tview.Application
	traceView    *tview.Table
	registerView *tview.Table

	rCellA *tview.TableCell
	rCellF *tview.TableCell
	rCellB *tview.TableCell
	rCellC *tview.TableCell
	rCellD *tview.TableCell
	rCellE *tview.TableCell
	rCellH *tview.TableCell
	rCellL *tview.TableCell
}

func NewTerminalUI(gb *gameboy.DMG) *TerminalUI {
	app := tview.NewApplication()

	t := &TerminalUI{
		gb: gb,

		app: app,
	}

	t.setupTraceView()
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
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Decompile"), 0, 4, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(t.registerView, 20, 1, false)

	t.app.SetRoot(root, true)
}

func (t *TerminalUI) setupTraceView() {
	t.traceView = tview.NewTable().
		SetBorders(false)
	t.traceView.SetBorder(true).SetTitle("Trace")
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
}

func (t *TerminalUI) Run() error {
	t.updateTicker = time.NewTicker(time.Second / 30)
	go func() {
		for range t.updateTicker.C {
			cpu := t.gb.CPU()
			t.rCellA.SetText(fmt.Sprintf("0x%02X", cpu.A.Read8()))
			t.rCellF.SetText(fmt.Sprintf("0x%02X", cpu.F.Read8()))
			t.rCellB.SetText(fmt.Sprintf("0x%02X", cpu.B.Read8()))
			t.rCellC.SetText(fmt.Sprintf("0x%02X", cpu.C.Read8()))
			t.rCellD.SetText(fmt.Sprintf("0x%02X", cpu.D.Read8()))
			t.rCellE.SetText(fmt.Sprintf("0x%02X", cpu.E.Read8()))
			t.rCellH.SetText(fmt.Sprintf("0x%02X", cpu.H.Read8()))
			t.rCellL.SetText(fmt.Sprintf("0x%02X", cpu.L.Read8()))

			t.app.Draw()
		}
	}()
	return t.app.Run()
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
}
