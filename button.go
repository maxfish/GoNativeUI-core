package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type Button struct {
	Label
	pressed bool
}

func NewButton(text string) *Button {
	b := &Button{}
	widgetInit(b)
	b.text = text
	b.Measure()
	return b
}

func (b *Button) Pressed() bool { return b.pressed }

func (b *Button) initStyle() {
	t := CurrentGui().Theme()
	b.style = &WidgetStyle{
		Font:             t.TextFont,
		FontSize:         t.TextFontSize,
		TextColor:        t.TextColor,
		BackgroundColor:  t.ButtonColor,
		Padding:          t.ButtonPadding,
		ContentAlignment: utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter},
	}
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
