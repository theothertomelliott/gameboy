package fyneui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/theothertomelliott/gameboy"
)

var _ fyne.Widget = &registers{}

type registers struct {
	widget.BaseWidget

	gb *gameboy.DMG

	// Register content
	a binding.Int
	f binding.Int
	b binding.Int
	c binding.Int
	d binding.Int
	e binding.Int
	h binding.Int
	l binding.Int

	sp binding.Int
	pc binding.Int
}

func newRegisters(gb *gameboy.DMG) *registers {
	r := &registers{
		gb: gb,

		a:  binding.NewInt(),
		f:  binding.NewInt(),
		b:  binding.NewInt(),
		c:  binding.NewInt(),
		d:  binding.NewInt(),
		e:  binding.NewInt(),
		h:  binding.NewInt(),
		l:  binding.NewInt(),
		sp: binding.NewInt(),
		pc: binding.NewInt(),
	}
	r.BaseWidget.ExtendBaseWidget(r)
	return r
}

func (r *registers) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewVBox(
		newMonoLabel("Registers"),
		container.NewHBox(
			newMonoLabel("A"),
			newByteLabel(r.a),
			newMonoLabel("F"),
			newByteLabel(r.f),
		),
		container.NewHBox(
			newMonoLabel("B"),
			newByteLabel(r.b),
			newMonoLabel("C"),
			newByteLabel(r.c),
		),
		container.NewHBox(
			newMonoLabel("D"),
			newByteLabel(r.d),
			newMonoLabel("E"),
			newByteLabel(r.e),
		),
		container.NewHBox(
			newMonoLabel("H"),
			newByteLabel(r.h),
			newMonoLabel("L"),
			newByteLabel(r.l),
		),
		container.NewHBox(
			newMonoLabel("SP"),
			newTwoByteLabel(r.sp),
			newMonoLabel("PC"),
			newTwoByteLabel(r.pc),
		),
	))
}

func (r *registers) updateDebugInfo() {
	r.a.Set(int(r.gb.CPU().A.Read8()))
	r.f.Set(int(r.gb.CPU().F.Read8()))

	r.b.Set(int(r.gb.CPU().B.Read8()))
	r.c.Set(int(r.gb.CPU().C.Read8()))

	r.d.Set(int(r.gb.CPU().D.Read8()))
	r.e.Set(int(r.gb.CPU().E.Read8()))

	r.h.Set(int(r.gb.CPU().H.Read8()))
	r.l.Set(int(r.gb.CPU().L.Read8()))

	r.sp.Set(int(r.gb.CPU().SP.Read16()))
	r.pc.Set(int(r.gb.CPU().PC.Read16()))
}

func newTwoByteLabel(b binding.Int) *widget.Label {
	l := widget.NewLabelWithData(binding.IntToStringWithFormat(b, "%04X"))
	l.TextStyle.Monospace = true
	return l
}

func newByteLabel(b binding.Int) *widget.Label {
	l := widget.NewLabelWithData(binding.IntToStringWithFormat(b, "%02X"))
	l.TextStyle.Monospace = true
	return l
}
