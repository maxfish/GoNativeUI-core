package gui

type Checkbox struct {
	Button
}

func NewCheckbox(text string) *Checkbox {
	b := &Checkbox{}
	widgetInit(b)
	b.text = text
	return b
}

func (b *Checkbox) Checked() bool { return b.pressed }

func (b *Checkbox) SetChecked(checked bool) { b.pressed = checked }

func (b *Checkbox) SetTheme(theme *Theme) {
	b.theme = theme
	b.font = theme.CheckboxFont
	b.textColor = theme.CheckboxTextColor
	b.fontSize = theme.CheckboxFontSize
	b.padding = theme.CheckboxPadding
	b.contentAlignment = theme.CheckboxAlignment
	b.Measure()
}

func (b *Checkbox) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if event == EventActionPress {
		b.pressed = !b.pressed
		return true
	}
	return false
}
