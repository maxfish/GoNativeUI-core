package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

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

func (b *Checkbox) initStyle() {
	t := CurrentGui().Theme()
	b.Label.initStyle()
	b.style.Padding = t.CheckboxPadding
	b.style.ContentAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
}

func (b *Checkbox) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if event == EventActionPress {
		b.pressed = !b.pressed
		return true
	}
	return false
}
