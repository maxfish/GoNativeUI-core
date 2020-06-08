package gui

type ButtonChangeCallback func(state bool)

type Button struct {
	Label
	pressed          bool
	onChangeCallback ButtonChangeCallback
}

func NewButton(text string, changeCallback ...ButtonChangeCallback) *Button {
	b := &Button{}
	widgetInit(b)
	b.style = CurrentGui().Theme().Button
	b.text = text
	if len(changeCallback) == 1 {
		b.onChangeCallback = changeCallback[0]
	}
	return b
}

func (b *Button) Pressed() bool { return b.pressed }

func (b *Button) SetOnChangeCallback(f ButtonChangeCallback) {
	b.onChangeCallback = f
}

func (b *Button) setPressed(pressed bool) {
	b.pressed = pressed
	if pressed {
		b.fireChangeEvent(true)
	}
}

func (b *Button) fireChangeEvent(state bool) {
	if b.onChangeCallback != nil {
		b.onChangeCallback(state)
	}
}

func (b *Button) OnMouseEvent(event MouseEvent) IWidget {
	if event.Type == MouseEventButton {
		if event.Button != MouseButtonLeft {
			return nil
		}
		if event.Action == EventActionPress {
			b.setPressed(true)
			return b
		} else if event.Action == EventActionRelease {
			b.setPressed(false)
			return b
		}
	}

	return nil
}
