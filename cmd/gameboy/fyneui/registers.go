package fyneui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &registers{}

type registers struct {
	widget.BaseWidget

	data DataTransport
}

func newRegisters(data DataTransport) *registers {
	r := &registers{
		data: data,
	}
	r.BaseWidget.ExtendBaseWidget(r)
	return r
}

func (r *registers) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewVBox(
		newMonoLabel("Registers"),
		container.NewHBox(
			newMonoLabel("A"),
			newByteLabel(r.data.A()),
			newMonoLabel("F"),
			newByteLabel(r.data.F()),
		),
		container.NewHBox(
			newMonoLabel("B"),
			newByteLabel(r.data.B()),
			newMonoLabel("C"),
			newByteLabel(r.data.C()),
		),
		container.NewHBox(
			newMonoLabel("D"),
			newByteLabel(r.data.D()),
			newMonoLabel("E"),
			newByteLabel(r.data.E()),
		),
		container.NewHBox(
			newMonoLabel("H"),
			newByteLabel(r.data.H()),
			newMonoLabel("L"),
			newByteLabel(r.data.L()),
		),
		container.NewHBox(
			newMonoLabel("SP"),
			newTwoByteLabel(r.data.SP()),
			newMonoLabel("PC"),
			newTwoByteLabel(r.data.PC()),
		),
	))
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
