package gui

type IPopup interface {
	Container() IContainer
	Modal() bool
	Visible() bool
	SetVisible(value bool)
}

type Popup struct {
	container IContainer
	modal     bool
	visible   bool
}

func NewPopup(container IContainer) *Popup {
	return &Popup{container: container}
}

func (p *Popup) Container() IContainer { return p.container }
func (p *Popup) Modal() bool           { return p.modal }
func (p *Popup) Visible() bool         { return p.visible }

func (p *Popup) SetVisible(value bool) {
	p.visible = value
}
