package gui

import "github.com/maxfish/gogui/gui/utils"

type IFont interface {
	FaceName() string
	TextSize(size int, text string) (int, int)
}

type IRenderer interface {
	Init(screen *Gui, dpiScale float32, uiScale float32)
	Draw()
}

type Gui struct {
	screen IContainer
}

func NewGui(theme *Theme, w int, h int) *Gui {
	g := &Gui{}
	g.screen = NewContainer()
	g.screen.SetTheme(theme)
	g.screen.SetBounds(utils.Rect{W: w, H: h})
	return g
}

func (g *Gui) Screen() IContainer { return g.screen }
func (g *Gui) Theme() *Theme      { return g.screen.Theme() }

//func (g *Gui) Scale() float32             { return g.scale }
//func (g *Gui) SetScale(scale float32)        { g.scale = scale }

// Mouse handling
func (g *Gui) OnMouseMoved(x, y float32) bool {
	return false
}

func (g *Gui) OnMouseButtonEvent(buttonIndex uint, event EventAction, modifiers ModifierKey) bool {
	return false
}

func (g *Gui) OnMouseScrolled(x, y float32) bool {
	return false
}
