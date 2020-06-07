package gui

type IPopup interface {
	Widget() IWidget
	Modal() bool
	Visible() bool
	SetVisible(value bool)
}
type Popup struct {
	widget  IWidget
	modal   bool
	visible bool
}

func NewPopup(widget IWidget) *Popup {
	return &Popup{widget: widget}
}

func (p *Popup) Widget() IWidget { return p.widget }
func (p *Popup) Modal() bool     { return p.modal }
func (p *Popup) Visible() bool   { return p.visible }

func (p *Popup) SetVisible(value bool) {
	p.visible = value
}
