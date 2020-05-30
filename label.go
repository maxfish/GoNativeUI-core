package gui

type Label struct {
	Widget
	text string
}

func NewLabel(text string) *Label {
	l := &Label{}
	widgetInit(l)
	l.style = CurrentGui().Theme().Label
	l.text = text
	return l
}

func (l *Label) Text() string { return l.text }
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
