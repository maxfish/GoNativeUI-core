package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type LabelText struct {
	text string
}

// Getters
func (l *LabelText) Text() string { return l.text }

type Label struct {
	Widget
	LabelText
}

func NewLabel(text string) *Label {
	l := &Label{}
	widgetInit(l)
	l.text = text
	return l
}

func (l *Label) initStyle() {
	t := CurrentGui().Theme()
	l.style = &WidgetStyle{
		Font:             t.TextFont,
		FontSize:         t.TextFontSize,
		TextColor:        t.TextColor,
		BackgroundColor:  utils.TransparentColor,
		ContentAlignment: utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter},
	}
}

func (l *Label) SetText(text string) {
	l.text = text
	l.Measure()
}

func (l *Label) Measure() {
	l.computeContentSize()
	l.measuredWidth = l.contentWidth + l.style.Padding.Left + l.style.Padding.Right
	l.measuredHeight = l.contentHeight + l.style.Padding.Top + l.style.Padding.Bottom
	l.measuredFlex = l.flex
}

func (l *Label) computeContentSize() {
	textSize := l.style.Font.TextSize(l.style.FontSize, l.text)
	l.contentWidth = textSize.W()
	l.contentHeight = textSize.H()
}
