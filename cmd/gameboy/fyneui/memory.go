package fyneui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &memory{}

type memory struct {
	widget.BaseWidget

	data DataTransport
}

func newMemory(data DataTransport) *memory {
	m := &memory{
		data: data,
	}
	m.BaseWidget.ExtendBaseWidget(m)
	return m
}

var _ fyne.WidgetRenderer = &memoryRenderer{}

type memoryRenderer struct {
	fyne.WidgetRenderer
	table *widget.Table
}

func (m *memory) CreateRenderer() fyne.WidgetRenderer {
	t := widget.NewTable(
		func() (int, int) {
			return len(m.data.Memory()), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("00")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*widget.Label).SetText(fmt.Sprintf("%04X", i.Row))
			} else if i.Col == 1 {
				o.(*widget.Label).SetText(fmt.Sprintf("%02X", m.data.Memory()[i.Row]))
			}
		},
	)
	t.SetColumnWidth(0, 50)
	t.SetColumnWidth(1, 40)

	l := widget.NewLabel("Memory")

	mr := &memoryRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(
			container.New(
				layout.NewBorderLayout(l, nil, nil, nil),
				l, t,
			),
		),
	}
	mr.table = t
	m.data.AddListener(mr)
	return mr
}

func (mr *memoryRenderer) DataChanged() {
	mr.table.Refresh()
}
