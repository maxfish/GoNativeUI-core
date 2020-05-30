package gui

type Checkbox struct {
	Button
}

func NewCheckbox(text string, changeCallback ...ButtonChangeCallback) *Checkbox {
	b := &Checkbox{}
	widgetInit(b)
	b.style = CurrentGui().Theme().Checkbox
	b.text = text
	if len(changeCallback) == 1 {
		b.onChangeCallback = changeCallback[0]
	}
	return b
}

func (b *Checkbox) Checked() bool { return b.pressed }

func (b *Checkbox) SetChecked(checked bool) {
	b.pressed = checked
	b.fireChangeEvent(checked)
}

func (b *Checkbox) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if event == EventActionPress {
		b.pressed = !b.pressed
		b.fireChangeEvent(b.pressed)
		return true
	}
	return false
}
