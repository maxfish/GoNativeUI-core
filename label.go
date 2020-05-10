package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Label struct {
	Widget

	text      string
	font      IFont
	fontSize  int
	textColor utils.ColorF
}

func NewLabel(text string) *Label {
	l := &Label{}
	l.visible = true
	l.enabled = true
	l.text = text
	l.fontSize = 25
	l.bounds = utils.Rect{W: 150, H: l.fontSize}
	return l
}

// Getters
func (l *Label) Text() string            { return l.text }
func (l *Label) Font() IFont             { return l.font }
func (l *Label) FontSize() int           { return l.fontSize }
func (l *Label) TextColor() utils.ColorF { return l.textColor }

func (l *Label) SizeToContent() {
	w, h := l.theme.LabelFont.TextSize(l.fontSize, l.text)
	l.bounds.W = w + l.padding.Left + l.padding.Right
	l.bounds.H = h + l.padding.Top + l.padding.Bottom
}

func (l *Label) SetTheme(theme *Theme) {
	l.theme = theme
	l.font = theme.LabelFont
	l.textColor = theme.LabelTextColor
	l.fontSize = theme.LabelFontSize
	l.padding = theme.LabelPadding
	//l.in = theme.LabelFontSize
}
