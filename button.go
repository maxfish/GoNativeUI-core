package gui

import "github.com/maxfish/gogui/gui/utils"

type Button struct {
	Label
	pressed bool
}

func NewButton(text string) *Button {
	b := &Button{}
	InitWidget(b.Widget)
	b.fontSize = 25
	b.padding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 2}
	b.bounds = utils.Rect{W: 200, H: b.fontSize}
	b.contentAlignmentH = AlignmentHCenter
	b.contentAlignmentV = AlignmentVCenter
	b.text = text
	return b
}

// Getters
func (b *Button) Pressed() bool { return b.pressed }

func (b *Button) SetTheme(theme *Theme) {
	b.theme = theme
	b.font = theme.ButtonFont
	b.textColor = theme.ButtonTextColor
	b.fontSize = theme.ButtonFontSize
	b.padding = theme.ButtonPadding
}
