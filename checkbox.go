package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type Checkbox struct {
	Button
}

func NewCheckbox(text string) *Checkbox {
	b := &Checkbox{}
	widgetInit(b)
	b.text = text
	b.Measure()
	return b
}

func (b *Checkbox) Checked() bool { return b.pressed }

func (b *Checkbox) SetChecked(checked bool) { b.pressed = checked }

func (b *Checkbox) initStyle() {
	t := CurrentGui().Theme()
	b.style = &WidgetStyle{
		Font:             t.TextFont,
		FontSize:         t.TextFontSize,
		TextColor:        t.TextColor,
		BackgroundColor:  utils.TransparentColor,
		Padding:          t.CheckboxPadding,
		ContentAlignment: utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter},
	}
}

func (b *Checkbox) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if event == EventActionPress {
		b.pressed = !b.pressed
		return true
	}
	return false
}
