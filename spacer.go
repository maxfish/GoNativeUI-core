package gui

type Spacer struct {
	Widget
}

func NewSpacer() *Spacer {
	s := &Spacer{}
	widgetInit(s)
	s.flex = 1
	return s
}
