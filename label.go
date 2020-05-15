package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Label struct {
	Widget
	font      IFont
	fontSize  int
	textColor utils.ColorF

	text string
}

func NewLabel(text string) *Label {
	l := &Label{}
	widgetInit(l)
	l.text = text
	l.fontSize = 25
	return l
}

// Getters
func (l *Label) Text() string            { return l.text }
func (l *Label) Font() IFont             { return l.font }
func (l *Label) FontSize() int           { return l.fontSize }
func (l *Label) TextColor() utils.ColorF { return l.textColor }

func (l *Label) SetTheme(theme *Theme) {
	l.theme = theme
	l.font = theme.LabelFont
	l.textColor = theme.LabelTextColor
	l.fontSize = theme.LabelFontSize
	l.padding = theme.LabelPadding
	l.Measure()
}

func (l *Label) SetText(text string) {
	l.text = text
	l.Measure()
}

func (l *Label) Measure() {
	l.computeContentSize()
	l.measuredWidth = l.contentWidth
	l.measuredHeight = l.contentHeight
	l.measuredFlex = l.flex
}

func (l *Label) Layout() {
	// TODO: Layout content
}

func (l *Label) computeContentSize() {
	textSize := l.theme.LabelFont.TextSize(l.fontSize, l.text)
	l.contentWidth = textSize.W()
	l.contentHeight = textSize.H()
}
