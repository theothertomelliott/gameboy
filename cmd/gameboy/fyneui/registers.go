package fyneui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func (u *UI) registerState() fyne.CanvasObject {

	return container.NewVBox(
		newMonoLabel("Registers"),
		container.NewHBox(
			newMonoLabel("A"),
			newMonoLabelWithData(u.a),
			newMonoLabel("F"),
			newMonoLabelWithData(u.f),
		),
		container.NewHBox(
			newMonoLabel("B"),
			newMonoLabelWithData(u.b),
			newMonoLabel("C"),
			newMonoLabelWithData(u.c),
		),
		container.NewHBox(
			newMonoLabel("D"),
			newMonoLabelWithData(u.d),
			newMonoLabel("E"),
			newMonoLabelWithData(u.e),
		),
		container.NewHBox(
			newMonoLabel("H"),
			newMonoLabelWithData(u.h),
			newMonoLabel("L"),
			newMonoLabelWithData(u.l),
		),
		container.NewHBox(
			newMonoLabel("SP"),
			newMonoLabelWithData(u.sp),
			newMonoLabel("PC"),
			newMonoLabelWithData(u.pc),
		),
	)
}

func (u *UI) updateDebugInfo() {
	u.a.Set(fmt.Sprintf("%02X", u.gb.CPU().A.Read8()))
	u.f.Set(fmt.Sprintf("%02X", u.gb.CPU().F.Read8()))

	u.b.Set(fmt.Sprintf("%02X", u.gb.CPU().B.Read8()))
	u.c.Set(fmt.Sprintf("%02X", u.gb.CPU().C.Read8()))

	u.d.Set(fmt.Sprintf("%02X", u.gb.CPU().D.Read8()))
	u.e.Set(fmt.Sprintf("%02X", u.gb.CPU().E.Read8()))

	u.h.Set(fmt.Sprintf("%02X", u.gb.CPU().H.Read8()))
	u.l.Set(fmt.Sprintf("%02X", u.gb.CPU().L.Read8()))

	u.sp.Set(fmt.Sprintf("%04X", u.gb.CPU().SP.Read16()))
	u.pc.Set(fmt.Sprintf("%04X", u.gb.CPU().PC.Read16()))

	u.win.Content().Refresh()
}
