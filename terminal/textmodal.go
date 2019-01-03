package terminal

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// TextModal is a centered message window used to inform the user or prompt them
// for an immediate decision. It needs to have at least one button (added via
// AddButtons()) or it will never disappear.
//
// See https://github.com/rivo/tview/wiki/Modal for an example.
type TextModal struct {
	*tview.Box

	// The framed embedded in the modal.
	frame *tview.Frame

	// The form embedded in the modal's frame.
	form *tview.Form

	// The message text (original, not word-wrapped).
	text string

	// The text color.
	textColor tcell.Color

	// The optional callback for when the user clicked one of the buttons. It
	// receives the index of the clicked button and the button's label.
	done func(textValue string)
}

// NewTextModal returns a new modal message window.
func NewTextModal() *TextModal {
	m := &TextModal{
		Box:       tview.NewBox(),
		textColor: tview.Styles.PrimaryTextColor,
	}
	m.form = tview.NewForm().
		SetButtonBackgroundColor(tview.Styles.PrimitiveBackgroundColor).
		SetButtonTextColor(tview.Styles.PrimaryTextColor)
	m.form.SetBackgroundColor(tview.Styles.ContrastBackgroundColor).SetBorderPadding(0, 0, 0, 0)
	m.form.SetCancelFunc(func() {
		if m.done != nil {
			m.done("")
		}
	})
	m.frame = tview.NewFrame(m.form).SetBorders(0, 0, 1, 0, 0, 0)
	m.frame.SetBorder(true).
		SetBackgroundColor(tview.Styles.ContrastBackgroundColor).
		SetBorderPadding(1, 1, 1, 1)
	return m
}

// SetTextColor sets the color of the message text.
func (m *TextModal) SetTextColor(color tcell.Color) *TextModal {
	m.textColor = color
	return m
}

// SetDoneFunc sets a handler which is called when one of the buttons was
// pressed. It receives the index of the button as well as its label text. The
// handler is also called when the user presses the Escape key. The index will
// then be negative and the label text an emptry string.
func (m *TextModal) SetDoneFunc(handler func(textValue string)) *TextModal {
	m.done = handler
	return m
}

// SetText sets the message text of the window. The text may contain line
// breaks. Note that words are wrapped, too, based on the final size of the
// window.
func (m *TextModal) SetText(text string) *TextModal {
	m.text = text
	return m
}

// SetInput sets up the input string with a label.
func (m *TextModal) SetInput(textLabel, textDefault, buttonLabel string) *TextModal {
	var textValue string
	m.form.AddInputField(textLabel, textDefault, 0, nil, func(value string) {
		textValue = value
	})
	m.form.AddButton(buttonLabel, func() {
		m.done(textValue)
	})
	return m
}

// Focus is called when this primitive receives focus.
func (m *TextModal) Focus(delegate func(p tview.Primitive)) {
	delegate(m.form)
}

// HasFocus returns whether or not this primitive has focus.
func (m *TextModal) HasFocus() bool {
	return m.form.HasFocus()
}

// Draw draws this primitive onto the screen.
func (m *TextModal) Draw(screen tcell.Screen) {
	// Calculate the width of this modal.
	buttonsWidth := 0

	for i := 0; i < m.form.GetButtonCount(); i++ {
		button := m.form.GetButton(i)
		buttonsWidth += tview.StringWidth(button.GetLabel()) + 4 + 2
	}
	buttonsWidth -= 2
	screenWidth, screenHeight := screen.Size()
	width := screenWidth / 3
	if width < buttonsWidth {
		width = buttonsWidth
	}
	// width is now without the box border.

	// Reset the text and find out how wide it is.
	m.frame.Clear()
	lines := tview.WordWrap(m.text, width)
	for _, line := range lines {
		m.frame.AddText(line, true, tview.AlignCenter, m.textColor)
	}

	// Set the modal's position and size.
	height := len(lines) + 6 + 2
	width += 4
	x := (screenWidth - width) / 2
	y := (screenHeight - height) / 2
	m.SetRect(x, y, width, height)

	// Draw the frame.
	m.frame.SetRect(x, y, width, height)
	m.frame.Draw(screen)
}
