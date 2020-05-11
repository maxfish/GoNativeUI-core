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
	l.bounds = utils.Rect{W: 150, H: l.fontSize}
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
	//l.in = theme.LabelFontSize
}

func (l *Label) SetText(text string) {
	l.text = text
	l.contentSizeValid = false
	l.Layout()
}

func (l *Label) Layout() {
	l.computeContentSize()
	widgetLayout(l)
}

func (l *Label) computeContentSize() {
	if l.contentSizeValid {
		return
	}
	l.contentWidth, l.contentHeight = l.theme.LabelFont.TextSize(l.fontSize, l.text)
	l.contentSizeValid = true
}
