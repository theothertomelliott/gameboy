package terminal

import (
	"fmt"

	"github.com/rivo/tview"
)

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
