package gui

type ButtonType int

const (
	ButtonTypeNormal ButtonType = iota
	ButtonTypeToggle
	ButtonTypeCheckbox
)

type Button struct {
	Label
	buttonType  ButtonType
	buttonGroup []*Button
	pressed     bool
}

func NewButton(text string) *Button {
	b := &Button{}
	widgetInit(b)
	b.buttonType = ButtonTypeNormal
	b.text = text
	return b
}

func NewToggleButton(text string) *Button {
	b := NewButton(text)
	b.buttonType = ButtonTypeToggle
	return b
}

func NewCheckbox(text string) *Button {
	b := NewButton(text)
	b.buttonType = ButtonTypeCheckbox
	return b
}

func (b *Button) Pressed() bool { return b.pressed }

func (b *Button) SetTheme(theme *Theme) {
	b.theme = theme

	if b.buttonType == ButtonTypeCheckbox {
		b.font = theme.CheckboxFont
		b.textColor = theme.CheckboxTextColor
		b.fontSize = theme.CheckboxFontSize
		b.padding = theme.CheckboxPadding
		b.contentAlignment = theme.CheckboxAlignment
	} else {
		b.font = theme.ButtonFont
		b.textColor = theme.ButtonTextColor
		b.fontSize = theme.ButtonFontSize
		b.padding = theme.ButtonPadding
		b.contentAlignment = theme.ButtonAlignment
	}

	b.Measure()
}

func (b *Button) SetButtonGroup(buttonGroup []*Button) {
	b.buttonGroup = buttonGroup
}

func (b *Button) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if b.buttonType == ButtonTypeNormal {
		switch event {
		case EventActionPress:
			b.pressed = true
		case EventActionRelease:
			b.pressed = false
		}
		return true
	}

	// Toggle buttons
	if event == EventActionPress {
		if b.buttonGroup == nil {
			b.pressed = !b.pressed
		} else {
			for _, btn := range b.buttonGroup {
				if btn == b {
					b.pressed = true
				} else {
					btn.pressed = false
				}
			}
		}
	}
	return true
}
