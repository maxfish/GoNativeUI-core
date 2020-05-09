package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type IFont interface {
	FaceName() string
	TextSize(size int, text string) (int, int)
}

type IRenderer interface {
	Init(screen *Gui, dpiScale float32, uiScale float32)
	Draw()
}

type Gui struct {
	screen    IContainer
	mouseData MouseData
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
func (g *Gui) OnMouseCursorMoved(x, y float32) bool {
	//log.Printf("[Gui] Mouse moved %.2f,%.2f\n", x, y)

	// TODO traverse the widgets tree and pass the event
	return false
}

func (g *Gui) OnMouseButtonEvent(buttonIndex ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	//log.Printf("[Gui] Mouse button #%d <%d> modifiers:%d\n", buttonIndex, event, modifiers)

	// TODO traverse the widgets tree and pass the event
	return false
}

func (g *Gui) OnMouseScrolled(scrollX, scrollY float32) bool {
	//log.Printf("[Gui] Mouse wheel scrolled %.2f,%.2f\n", scrollX, scrollY)

	// TODO traverse the widgets tree and pass the event
	return false
}
