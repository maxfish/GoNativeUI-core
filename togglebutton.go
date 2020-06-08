package gui

type ToggleButton struct {
	Button
	buttonGroup []*ToggleButton
}

func NewToggleButton(text string, changeCallback ...ButtonChangeCallback) *ToggleButton {
	b := &ToggleButton{}
	widgetInit(b)
	b.style = CurrentGui().Theme().Button
	b.text = text
	if len(changeCallback) == 1 {
		b.onChangeCallback = changeCallback[0]
	}
	return b
}

func (b *ToggleButton) SetButtonGroup(buttonGroup []*ToggleButton) {
	b.buttonGroup = buttonGroup
}

func (b *ToggleButton) SetPressed(pressed bool) {
	if b.buttonGroup == nil {
		b.pressed = pressed
		b.fireChangeEvent(pressed)
		return
	}

	if pressed && !b.pressed {
		for _, btn := range b.buttonGroup {
			if btn == b {
				if b.pressed != true {
					b.pressed = true
					b.fireChangeEvent(true)
				}
			} else {
				if btn.pressed != false {
					btn.pressed = false
					btn.fireChangeEvent(false)
				}
			}
		}
	}
}

func (b *ToggleButton) OnMouseEvent(event MouseEvent) IWidget {
	if event.Type == MouseEventButton {
		if event.Button != MouseButtonLeft {
			return nil
		}
		if event.Action == EventActionPress {
			b.SetPressed(!b.pressed)
			return b
		}
	}

	return nil
}
