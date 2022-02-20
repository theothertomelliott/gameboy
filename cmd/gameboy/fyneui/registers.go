package fyneui

import (
	"fmt"

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
	a binding.String
	f binding.String
	b binding.String
	c binding.String
	d binding.String
	e binding.String
	h binding.String
	l binding.String

	sp binding.String
	pc binding.String
}

func newRegisters(gb *gameboy.DMG) *registers {
	r := &registers{
		gb: gb,

		a:  stringBinding("0x00"),
		f:  stringBinding("0x00"),
		b:  stringBinding("0x00"),
		c:  stringBinding("0x00"),
		d:  stringBinding("0x00"),
		e:  stringBinding("0x00"),
		h:  stringBinding("0x00"),
		l:  stringBinding("0x00"),
		sp: stringBinding("0x00"),
		pc: stringBinding("0x00"),
	}
	r.BaseWidget.ExtendBaseWidget(r)
	return r
}

func (r *registers) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewVBox(
		newMonoLabel("Registers"),
		container.NewHBox(
			newMonoLabel("A"),
			newMonoLabelWithData(r.a),
			newMonoLabel("F"),
			newMonoLabelWithData(r.f),
		),
		container.NewHBox(
			newMonoLabel("B"),
			newMonoLabelWithData(r.b),
			newMonoLabel("C"),
			newMonoLabelWithData(r.c),
		),
		container.NewHBox(
			newMonoLabel("D"),
			newMonoLabelWithData(r.d),
			newMonoLabel("E"),
			newMonoLabelWithData(r.e),
		),
		container.NewHBox(
			newMonoLabel("H"),
			newMonoLabelWithData(r.h),
			newMonoLabel("L"),
			newMonoLabelWithData(r.l),
		),
		container.NewHBox(
			newMonoLabel("SP"),
			newMonoLabelWithData(r.sp),
			newMonoLabel("PC"),
			newMonoLabelWithData(r.pc),
		),
	))
}

func (r *registers) updateDebugInfo() {
	r.a.Set(fmt.Sprintf("%02X", r.gb.CPU().A.Read8()))
	r.f.Set(fmt.Sprintf("%02X", r.gb.CPU().F.Read8()))

	r.b.Set(fmt.Sprintf("%02X", r.gb.CPU().B.Read8()))
	r.c.Set(fmt.Sprintf("%02X", r.gb.CPU().C.Read8()))

	r.d.Set(fmt.Sprintf("%02X", r.gb.CPU().D.Read8()))
	r.e.Set(fmt.Sprintf("%02X", r.gb.CPU().E.Read8()))

	r.h.Set(fmt.Sprintf("%02X", r.gb.CPU().H.Read8()))
	r.l.Set(fmt.Sprintf("%02X", r.gb.CPU().L.Read8()))

	r.sp.Set(fmt.Sprintf("%04X", r.gb.CPU().SP.Read16()))
	r.pc.Set(fmt.Sprintf("%04X", r.gb.CPU().PC.Read16()))
}
