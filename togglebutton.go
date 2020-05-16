package gui

type ToggleButton struct {
	Button
	buttonGroup []*ToggleButton
}

func NewToggleButton(text string) *ToggleButton {
	b := &ToggleButton{}
	widgetInit(b)
	b.text = text
	return b
}
func (b *ToggleButton) SetButtonGroup(buttonGroup []*ToggleButton) {
	b.buttonGroup = buttonGroup
}

func (b *ToggleButton) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
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
		return true
	}
	return false
}
