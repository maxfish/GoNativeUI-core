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

func (b *Checkbox) OnMouseEvent(event MouseEvent) IWidget {
	if event.Type == MouseEventButton {
		if event.Button != MouseButtonLeft {
			return nil
		}
		if event.Action == EventActionPress {
			b.pressed = !b.pressed
			b.fireChangeEvent(b.pressed)
			return b
		}
	}

	return nil
}
