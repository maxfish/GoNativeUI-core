package gui

type Button struct {
	Label
	pressed bool
}

func NewButton(text string) *Button {
	b := &Button{}
	widgetInit(b)
	b.text = text
	return b
}

func (b *Button) Pressed() bool { return b.pressed }

func (b *Button) initStyle() {
	b.Label.initStyle()
	t := CurrentGui().Theme()
	b.style.BackgroundColor = t.ButtonColor
	b.style.Padding = t.ButtonPadding
}

func (b *Button) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if event == EventActionPress {
		b.pressed = true
		return true
	} else if event == EventActionRelease {
		b.pressed = false
		return true
	}
	return false
}
