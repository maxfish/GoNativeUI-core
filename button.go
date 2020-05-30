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

func (b *Button) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if event == EventActionPress {
		b.setPressed(true)
		return true
	} else if event == EventActionRelease {
		b.setPressed(false)
		return true
	}
	return false
}
