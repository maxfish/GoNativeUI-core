package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type LabelText struct {
	font      IFont
	fontSize  int
	textColor utils.Color
	fillColor utils.Color
	text      string
}

// Getters
func (l *LabelText) Text() string           { return l.text }
func (l *LabelText) Font() IFont            { return l.font }
func (l *LabelText) FontSize() int          { return l.fontSize }
func (l *LabelText) TextColor() utils.Color { return l.textColor }
func (l *LabelText) FillColor() utils.Color { return l.fillColor }

type Label struct {
	Widget
	LabelText
}

func NewLabel(text string) *Label {
	l := &Label{}
	widgetInit(l)
	l.text = text
	l.fontSize = 25
	return l
}

func (l *Label) SetTheme(theme *Theme) {
	l.theme = theme
	l.font = theme.LabelFont
	l.textColor = theme.LabelTextColor
	l.fillColor = theme.LabelFillColor
	l.fontSize = theme.LabelFontSize
	l.padding = theme.LabelPadding
	l.contentAlignment = theme.LabelAlignment
	l.Measure()
}

func (l *Label) SetText(text string) {
	l.text = text
	l.Measure()
}

func (l *Label) Measure() {
	l.computeContentSize()
	l.measuredWidth = l.contentWidth + l.padding.Left + l.padding.Right
	l.measuredHeight = l.contentHeight + l.padding.Top + l.padding.Bottom
	l.measuredFlex = l.flex
}

func (l *Label) computeContentSize() {
	textSize := l.theme.LabelFont.TextSize(l.fontSize, l.text)
	l.contentWidth = textSize.W()
	l.contentHeight = textSize.H()
}
